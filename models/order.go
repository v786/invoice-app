// models/book.go

package models

import (
	"time"
)

type Order struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	CreatedBy    string    `json:"createdBy"`
	OrderDetails string    `json:"orderDetails"`
	Amount       int       `json:"amount"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Mid          uint      `json:"mid"`
}
