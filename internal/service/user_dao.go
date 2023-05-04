package service

import (
	"bilibiliRSS/internal/db"
	"bilibiliRSS/internal/model"
	"gorm.io/gorm"
)

type userService struct {
}

var UserService = userService{}

func (service userService) db() *gorm.DB {
	return db.DB
}
func (service *userService) GetUserByUsernameAndPassword(username, password string) (user *model.User, err error) {
	err = service.db().Where("username = ? AND password = ?", username, password).First(&user).Error
	return
}
func (service *userService) Insert(username, password string) error {
	user := &model.User{
		Username: username,
		Password: password,
	}
	return service.db().Create(&user).Error
}
