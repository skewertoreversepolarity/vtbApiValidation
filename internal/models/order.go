package models

import (
	"gorm.io/datatypes"
	"time"
)

type Order struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	OrderNumber  string         `gorm:"unique;not null" json:"orderNumber"`
	OrderID      string         `gorm:"unique;not null" json:"orderId"`
	Amount       int            `gorm:"not null" json:"amount"`
	Currency     string         `gorm:"not null" json:"currency"`
	OrderStatus  string         `gorm:"default:'pending'" json:"orderStatus"` // pending, paid, failed
	RequestData  datatypes.JSON `gorm:"type:jsonb" json:"requestData"`        // JSON-тело запроса
	ResponseData datatypes.JSON `gorm:"type:jsonb" json:"responseData"`       // JSON-ответ от VTB API
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
}
