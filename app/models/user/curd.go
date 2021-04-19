package user

import (
	"goblog/pkg/model"
	"goblog/pkg/logger"
)

// Create 创建用户，通过User.ID 来判断是否创建成功
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}