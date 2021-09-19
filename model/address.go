package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street  *string `gorm:"type:varchar(100)"`
	City    *string `gorm:"type:varchar(100)"`
	Country *string `gorm:"type:varchar(100)"`
	UserID  uint
}
