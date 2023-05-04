package service

import (
	"bilibiliRSS/internal/db"
	"bilibiliRSS/internal/model"
	"gorm.io/gorm"
)

type userSubscriptionDao struct {
}

var UserSubscriptionDao = userSubscriptionDao{}

func (dao *userSubscriptionDao) db() *gorm.DB {
	return db.DB
}
func (dao *userSubscriptionDao) GetAllByUserID(uid uint) (list []*model.UserSubscription, err error) {
	err = dao.db().Where("UID = ?", uid).Find(&list).Error
	return
}
