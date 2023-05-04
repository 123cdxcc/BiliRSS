package model

import (
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Mid      int64  `gorm:"unique" json:"mid"`
	Pic      string `json:"pic"`
	ShareUrl string `json:"share_url"`
}
