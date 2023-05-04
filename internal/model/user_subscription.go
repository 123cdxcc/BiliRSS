package model

import "gorm.io/gorm"

type UserSubscription struct {
	gorm.Model
	UID uint `json:"uid"`
	SID uint `json:"sid"`
}
