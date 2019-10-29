package service

import (
	"blog/cache"
	"blog/helpers"
	"blog/http/request"
	"blog/model"
	"encoding/json"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AdminValid 验证表单
func (s *Service) AdminValid(req *request.AdminRegister) error {
	if req.PasswordConfirm != req.Password {
		return errors.New("两次输入的密码不相同")
	}

	if s.Dao.ExistAdminByAdminName(req.Adminname) {
		return errors.New("用户名已经注册")
	}
	return nil
}

// AdminRegister 用户注册
func (s *Service) AdminRegister(req *request.AdminRegister) (model.Admin, error) {
	admin := model.Admin{
		Adminname: req.Adminname,
		Status:    1,
	}

	// 表单验证
	if err := s.AdminValid(req); err != nil {
		return admin, err
	}

	// 加密密码
	if err := setAdminPassword(&admin, req.Password); err != nil {
		return admin, errors.New("密码加密失败")
	}

	// 创建用户
	if err := s.Dao.CreateAdmin(&admin); err != nil {
		return admin, errors.New("注册失败")
	}
	//清理缓存
	s.Cache.DeleteLike("blog_admins_by_offset*")
	s.Cache.Delete("blog_admin_count")

	return admin, nil
}

// SetPassword 设置密码
func setAdminPassword(admin *model.Admin, password string) error {
	admin.Salt = helpers.SaltMaker(SaltLen)
	bytes, err := bcrypt.GenerateFromPassword([]byte(admin.Salt+password), PassWordCost)
	if err != nil {
		return err
	}
	admin.PasswordDigest = string(bytes)
	return nil
}

// AdminLogin 用户登录函数
func (s *Service) AdminLogin(req *request.AdminLogin) (model.Admin, error) {

	admin, err := s.Dao.GetAdminByAdminname(req.Adminname)
	if err != nil {
		return admin, errors.New("账号或密码错误")
	}

	if checkAdminPassword(&admin, req.Password) == false {
		return admin, errors.New("账号或密码错误")
	}
	return admin, nil
}

// CheckPassword 校验密码
func checkAdminPassword(admin *model.Admin, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordDigest), []byte(admin.Salt+password))
	return err == nil
}

//AdminModify 用户信息修改
func (s *Service) AdminModify(currentAdmin *model.Admin, req *request.AdminInfoModify) error {
	admin := model.Admin{
		Nickname: req.Nickname,
	}

	if req.Password != "" {
		if req.Password != req.PasswordConfirm {
			return errors.New("密码不一致")
		}
		if err := setAdminPassword(&admin, req.Password); err != nil {
			return errors.New("密码加密失败")
		}
	}

	if err := s.Dao.UpdateAdmin(currentAdmin, &admin); err != nil {
		return errors.New("修改失败")
	}
	//清理缓存
	s.Cache.DeleteLike("blog_admins_by_offset*")
	return nil
}

//getAdminsByOffset 管理员列表
func (s *Service) getAdminsByOffset(PageNum int, PageSize int) ([]*model.Admin, error) {
	var admins []*model.Admin
	offset, limit := s.Page.OffsetLimit(PageNum, PageSize)

	//获取Redis缓存
	cachekey := cache.GenAdminsByOffset(offset, limit)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &admins)
		return admins, nil
	}

	//数据库查询数据
	admins, err = s.Dao.GetAdminsByOffset(offset, limit)
	if err != nil {
		return nil, err
	}

	//Redis缓存数据
	s.Cache.Set(cachekey, admins, 5*time.Hour)
	return admins, nil
}

//AdminCount 获取管理员数量
func (s *Service) getAdminCount() (int, error) {
	var count int
	//获取Redis缓存
	cachekey := cache.GenAdminCountKey()
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &count)
		return count, nil
	}
	//数据库查询数据
	if count, err = s.Dao.GetAdminCount(); err != nil {
		return 0, nil
	}
	//Redis缓存数据
	s.Cache.Set(cachekey, count, 5*time.Hour)
	return count, nil
}

//GetAdminsWithTotal 获取管理员列表服务
func (s *Service) GetAdminsWithTotal(PageNum int, PageSize int) (map[string]interface{}, error) {

	admins, err := s.getAdminsByOffset(PageNum, PageSize)
	if err != nil {
		return nil, err
	}
	counts, err := s.getAdminCount()
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"items": admins,
		"total": counts,
	}
	return res, nil
}

//AdminDelete 删除管理员
func (s *Service) AdminDelete(ID uint) error {
	//清理缓存
	s.Cache.DeleteLike("blog_admins_by_offset*")
	s.Cache.Delete("blog_admin_count")

	return s.Dao.DeleteAdmin(ID)
}
