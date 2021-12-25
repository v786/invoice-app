// models/book.go

package models

import (
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	UserName    string  `json:"userName"`
	TinNumber   string  `json:"tinNumber"`
	PhoneNumber string  `json:"phoneNumber"`
	Email       string  `json:"email"`
	Orders      []Order `gorm:"foreignKey:Mid"`
}
