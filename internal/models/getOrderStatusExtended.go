package models

import (
	"gorm.io/datatypes"
	"time"
)

// GetOrderStatusExtended представляет структуру ответа
type GetOrderStatusExtended struct {
	ID                    int64             `gorm:"primaryKey" json:"id"`
	ErrorCode             string            `gorm:"not null" json:"errorCode"`
	ErrorMessage          string            `gorm:"not null" json:"errorMessage"`
	OrderNumber           string            `gorm:"unique;not null" json:"orderNumber"`
	OrderStatus           int               `gorm:"not null" json:"orderStatus"`
	ActionCode            int               `gorm:"not null" json:"actionCode"`
	ActionCodeDescription string            `json:"actionCodeDescription"`
	Amount                int               `gorm:"not null" json:"amount"`
	Currency              string            `gorm:"not null" json:"currency"`
	Date                  time.Time         `gorm:"not null" json:"date"`
	OrderDescription      string            `json:"orderDescription"`
	IP                    string            `json:"ip"`
	MerchantOrderParams   datatypes.JSON    `gorm:"type:jsonb" json:"merchantOrderParams"`
	TransactionAttributes datatypes.JSON    `gorm:"type:jsonb" json:"transactionAttributes"`
	Attributes            []Attribute       `gorm:"foreignKey:OrderID;references:ID" json:"attributes"`
	CardAuthInfo          CardAuthInfo      `gorm:"embedded" json:"cardAuthInfo"`
	AuthDateTime          time.Time         `gorm:"not null" json:"authDateTime"`
	TerminalID            string            `gorm:"not null" json:"terminalId"`
	AuthRefNum            string            `json:"authRefNum"`
	PaymentAmountInfo     PaymentAmountInfo `gorm:"embedded" json:"paymentAmountInfo"`
	BankInfo              BankInfo          `gorm:"embedded" json:"bankInfo"`
}

// Attribute представляет элемент массива attributes
type Attribute struct {
	ID      int64  `gorm:"primaryKey" json:"id"`
	OrderID int64  `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orderID"`
	Name    string `gorm:"not null" json:"name"`
	Value   string `gorm:"not null" json:"value"`
}

// CardAuthInfo содержит информацию о карте
type CardAuthInfo struct {
	MaskedPan      string `json:"maskedPan"`
	Expiration     string `json:"expiration"`
	CardholderName string `json:"cardholderName"`
	ApprovalCode   string `json:"approvalCode"`
	Pan            string `json:"pan"`
}

// PaymentAmountInfo содержит информацию об оплате
type PaymentAmountInfo struct {
	PaymentState    string `gorm:"not null" json:"paymentState"`
	ApprovedAmount  int    `json:"approvedAmount"`
	DepositedAmount int    `json:"depositedAmount"`
	RefundedAmount  int    `json:"refundedAmount"`
}

// BankInfo содержит информацию о банке
type BankInfo struct {
	BankName        string `json:"bankName"`
	BankCountryCode string `gorm:"not null" json:"bankCountryCode"`
	BankCountryName string `gorm:"not null" json:"bankCountryName"`
}
