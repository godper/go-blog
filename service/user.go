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

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	//SaltLen 盐长度
	SaltLen = 8
)

// Valid 验证表单
func (s *Service) Valid(req *request.UserRegister) error {
	if req.PasswordConfirm != req.Password {
		return errors.New("两次输入的密码不相同")
	}

	if s.Dao.ExistUserByUserName(req.Username) {
		return errors.New("用户名已经注册")
	}
	return nil
}

// Register 用户注册
func (s *Service) Register(req *request.UserRegister) (model.User, error) {
	user := model.User{
		Username: req.Username,
		Status:   1,
	}

	// 表单验证
	if err := s.Valid(req); err != nil {
		return user, err
	}

	// 加密密码
	if err := setPassword(&user, req.Password); err != nil {
		return user, errors.New("密码加密失败")
	}

	// 创建用户
	if err := s.Dao.CreateUser(&user); err != nil {
		return user, errors.New("注册失败")
	}
	//清理缓存
	s.Cache.DeleteLike("blog_users_by_offset*")
	s.Cache.Delete("blog_user_count")
	return user, nil
}

// SetPassword 设置密码
func setPassword(user *model.User, password string) error {
	user.Salt = helpers.SaltMaker(SaltLen)
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Salt+password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// Login 用户登录函数
func (s *Service) Login(req *request.UserLogin) (model.User, error) {

	user, err := s.Dao.GetUserByUsername(req.Username)
	if err != nil {
		return user, errors.New("账号或密码错误")
	}

	if checkPassword(&user, req.Password) == false {
		return user, errors.New("账号或密码错误")
	}
	return user, nil
}

// CheckPassword 校验密码
func checkPassword(user *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(user.Salt+password))
	return err == nil
}

//UserModify 用户信息修改
func (s *Service) UserModify(currentUseer *model.User, req *request.UserInfoModify) error {
	user := model.User{
		Nickname: req.Nickname,
	}

	if req.Password != "" {
		if req.Password != req.PasswordConfirm {
			return errors.New("密码不一致")
		}
		if err := setPassword(&user, req.Password); err != nil {
			return errors.New("密码加密失败")
		}
	}

	if err := s.Dao.UpdateUser(currentUseer, &user); err != nil {
		return errors.New("修改失败")
	}
	//清理缓存
	s.Cache.DeleteLike("blog_users_by_offset*")
	return nil
}

//getUsersByOffset 用户列表
func (s *Service) getUsersByOffset(PageNum int, PageSize int) ([]*model.User, error) {
	var users []*model.User
	offset, limit := s.Page.OffsetLimit(PageNum, PageSize)

	//获取Redis缓存
	cachekey := cache.GenUsersByOffset(offset, limit)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &users)
		return users, nil
	}

	//数据库查询数据
	users, err = s.Dao.GetUsersByOffset(offset, limit)
	if err != nil {
		return nil, err
	}

	//Redis缓存数据
	s.Cache.Set(cachekey, users, 5*time.Hour)
	return users, nil
}

//UserCount 获取用户数量
func (s *Service) getUserCount() (int, error) {
	var count int
	//获取Redis缓存
	cachekey := cache.GenUserCountKey()
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &count)
		return count, nil
	}
	//数据库查询数据
	if count, err = s.Dao.GetUserCount(); err != nil {
		return 0, nil
	}
	//Redis缓存数据
	s.Cache.Set(cachekey, count, 5*time.Hour)
	return count, nil
}

//GetUsersWithTotal 获取管理员列表服务
func (s *Service) GetUsersWithTotal(PageNum int, PageSize int) (map[string]interface{}, error) {

	users, err := s.getUsersByOffset(PageNum, PageSize)
	if err != nil {
		return nil, err
	}
	counts, err := s.getUserCount()
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"items": users,
		"total": counts,
	}
	return res, nil
}

//UserDelete 删除文章
func (s *Service) UserDelete(ID uint) error {
	//清理缓存
	s.Cache.DeleteLike("blog_users_by_offset*")
	s.Cache.Delete("blog_user_count")

	return s.Dao.DeleteUser(ID)
}
