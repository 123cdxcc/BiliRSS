package dao

import (
	db2 "bilibiliRSS/internal/db"
	"bilibiliRSS/internal/model"
	"gorm.io/gorm"
)

type subscriptionDao struct {
}

var SubscriptionDao = subscriptionDao{}

func (dao *subscriptionDao) db() *gorm.DB {
	return db2.DB
}
func (dao *subscriptionDao) GetAll() (list []*model.Subscription, err error) {
	err = dao.db().Find(&list).Error
	return
}
func (dao subscriptionDao) GetByMid(mid int64) (subscription *model.Subscription, err error) {
	err = dao.db().Where("mid=?", mid).Find(&subscription).Error
	return
}
func (dao *subscriptionDao) Insert(s *model.Subscription) error {
	return dao.db().Create(&s).Error
}
func (dao *subscriptionDao) Update(s *model.Subscription) error {
	return dao.db().Save(&s).Error
}
