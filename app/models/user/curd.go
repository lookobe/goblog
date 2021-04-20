package user

import (
	"goblog/pkg/types"
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

// Create 创建用户，通过User.ID 来判断是否创建成功
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (User,error) {
	var user User
	if err := model.DB.Where("email = ?",email).First(&user).Error; err != nil {
		return user,err
	}
	return user, nil
}

// Get 通过ID 获取用户
func Get(idstr string) (User,error) {
	var user User
	id := types.StringToInt(idstr)
	if err := model.DB.First(&user,id).Error; err != nil {
		return user,err
	}
	return user,nil
}