package dao

import (
	"blog/model"
)

// GetUserByID 用ID获取用户
func (d *Dao) GetUserByID(ID interface{}) (model.User, error) {
	var user model.User
	result := d.DB.First(&user, ID)
	return user, result.Error
}

//GetUserByUsername 用Username获取用户
func (d *Dao) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := d.DB.Model(&model.User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}

//ExistUserByUserName 用UserName判断用户是否存在
func (d *Dao) ExistUserByUserName(userName string) bool {
	count := 0
	d.DB.Model(&model.User{}).Where("username = ?", userName).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

//CreateUser 创建新用户
func (d *Dao) CreateUser(user *model.User) error {

	err := d.DB.Create(&user).Error

	return err
}

//UpdateUser 更新用户
func (d *Dao) UpdateUser(currentUser *model.User, modUser *model.User) error {
	err := d.DB.Model(currentUser).Updates(modUser).Error
	return err
}

// GetUsersByOffset gets a list of users based on paging constraints
func (d *Dao) GetUsersByOffset(offset int, limit int) ([]*model.User, error) {
	var users []*model.User

	err := d.DB.Select("id, username, status, nickname, created_at, avatar").
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&users).
		Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserCount gets the total number of users based on the constraints
func (d *Dao) GetUserCount() (int, error) {
	var count int
	if err := d.DB.Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteUser delete a single user
func (d *Dao) DeleteUser(id uint) error {
	if err := d.DB.Where("id = ?", id).Delete(model.User{}).Error; err != nil {
		return err
	}
	return nil
}
