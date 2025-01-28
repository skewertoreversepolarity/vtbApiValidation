package models

import "time"

type AcquiringOperation struct {
	ID               uint   `gorm:"primaryKey"`
	Term             int    `gorm:"not null"`
	Code             string `gorm:"varchar(50); not null"`
	Merchant         string `gorm:"varchar(10); not null"`
	OrderID          string `gorm:"unique;not null"` // orderId из ответа API
	SessionId        string `gorm:"varchar(50);"`
	OrderIDEncrypted string `gorm:"varchar(50);"`
	
	FormURL     string    `gorm:"not null"`       // formUrl из ответа API
	Amount      string    `gorm:"not null"`       // Сумма заказа
	Currency    string    `gorm:"not null"`       // Валюта
	OrderNumber string    `gorm:"not null"`       // Уникальный номер заказа
	CreatedAt   time.Time `gorm:"autoCreateTime"` // Время создания записи
}
