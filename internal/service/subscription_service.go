package service

import (
	"bilibiliRSS/internal/db"
	"bilibiliRSS/internal/model"
	"gorm.io/gorm"
)

type subscriptionService struct {
}

var SubscriptionService = subscriptionService{}

func (service *subscriptionService) db() *gorm.DB {
	return db.DB
}
func (service *subscriptionService) GetAll() (list []*model.Subscription, err error) {
	err = service.db().Find(&list).Error
	return
}
func (service *subscriptionService) GetByIds(ids []uint) (list []*model.Subscription, err error) {
	err = service.db().Where("ID in ?", ids).Find(&list).Error
	return
}
func (service *subscriptionService) GetByUserID(uid uint) (list []*model.Subscription, err error) {
	us := make([]*model.UserSubscription, 0)
	err = service.db().Where("uid = ?", uid).Find(&us).Error
	if err != nil {
		return
	}
	ids := make([]uint, 0, len(us))
	for i, v := range us {
		ids[i] = v.SID
	}
	list, err = service.GetByIds(ids)
	if err != nil {
		return
	}
	return
}
func (service subscriptionService) GetByMid(mid int64) (subscription *model.Subscription, err error) {
	err = service.db().Where("mid=?", mid).Find(&subscription).Error
	return
}
func (service *subscriptionService) Insert(s *model.Subscription) error {
	return service.db().Create(&s).Error
}
func (service *subscriptionService) Update(s *model.Subscription) error {
	return service.db().Save(&s).Error
}
