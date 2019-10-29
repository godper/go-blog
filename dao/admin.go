package dao

import (
	"blog/model"
)

// GetAdminByID 用ID获取用户
func (d *Dao) GetAdminByID(ID interface{}) (model.Admin, error) {
	var Admin model.Admin
	result := d.DB.First(&Admin, ID)
	return Admin, result.Error
}

//GetAdminByAdminname 用Adminname获取用户
func (d *Dao) GetAdminByAdminname(Adminname string) (model.Admin, error) {
	var Admin model.Admin
	result := d.DB.Model(&model.Admin{}).Where("Adminname = ?", Adminname).First(&Admin)
	return Admin, result.Error
}

//ExistAdminByAdminName 用AdminName判断用户是否存在
func (d *Dao) ExistAdminByAdminName(AdminName string) bool {
	count := 0
	d.DB.Model(&model.Admin{}).Where("Adminname = ?", AdminName).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

//CreateAdmin 创建新用户
func (d *Dao) CreateAdmin(Admin *model.Admin) error {

	err := d.DB.Create(&Admin).Error

	return err
}

//UpdateAdmin 更新用户
func (d *Dao) UpdateAdmin(currentAdmin *model.Admin, modAdmin *model.Admin) error {
	err := d.DB.Model(currentAdmin).Updates(modAdmin).Error
	return err
}

// GetAdminsByOffset gets a list of admins based on paging constraints
func (d *Dao) GetAdminsByOffset(offset int, limit int) ([]*model.Admin, error) {
	var admins []*model.Admin

	err := d.DB.Select("id, adminname, status, nickname, created_at, avatar").
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&admins).
		Error

	if err != nil {
		return nil, err
	}
	return admins, nil
}

// GetAdminCount gets the total number of admins based on the constraints
func (d *Dao) GetAdminCount() (int, error) {
	var count int
	if err := d.DB.Model(&model.Admin{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteAdmin delete a single admin
func (d *Dao) DeleteAdmin(id uint) error {
	if err := d.DB.Where("id = ?", id).Delete(model.Admin{}).Error; err != nil {
		return err
	}
	return nil
}
