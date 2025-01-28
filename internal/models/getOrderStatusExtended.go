package models

import "time"

// GetOrderStatusExtended представляет структуру ответа
type GetOrderStatusExtended struct {
	ID                    int64             `gorm:"primaryKay" json:"id"`
	ErrorCode             string            `gorm:"not null" json:"errorCode"`
	ErrorMessage          string            `gorm:"not null" json:"errorMessage"`
	OrderNumber           string            `gorm:"unique;not null" json:"orderNumber"`
	OrderStatus           int               `gorm:"not null" json:"orderStatus"`
	ActionCode            int               `gorm:"not null" json:"actionCode"`
	ActionCodeDescription string            `gorm:"" json:"actionCodeDescription"`
	Amount                int               `gorm:"not null" json:"amount"`
	Currency              string            `gorm:"not null" json:"currency"`
	Date                  time.Time         `gorm:"not null" json:"date"`
	OrderDescription      string            `gorm:"" json:"orderDescription"`
	IP                    string            `gorm:"" json:"ip"`
	MerchantOrderParams   []interface{}     `gorm:"" json:"merchantOrderParams"`
	TransactionAttributes []interface{}     `gorm:"" json:"transactionAttributes"`
	Attributes            []Attribute       `gorm:"" json:"attributes"`
	CardAuthInfo          CardAuthInfo      `gorm:"not null" json:"cardAuthInfo"`
	AuthDateTime          time.Time         `gorm:"not null" json:"authDateTime"`
	TerminalID            string            `gorm:"not null" json:"terminalId"`
	AuthRefNum            string            `gorm:"" json:"authRefNum"`
	PaymentAmountInfo     PaymentAmountInfo `gorm:"not null" json:"paymentAmountInfo"`
	BankInfo              BankInfo          `gorm:"not null" json:"bankInfo"`
}

// Attribute представляет элемент массива attributes
type Attribute struct {
	Name  string `gorm:"not null" json:"name"`
	Value string `gorm:"not null" json:"value"`
}

// CardAuthInfo содержит информацию о карте
type CardAuthInfo struct {
	MaskedPan      string `gorm:"" json:"maskedPan"`
	Expiration     string `gorm:"" json:"expiration"`
	CardholderName string `gorm:"" json:"cardholderName"`
	ApprovalCode   string `gorm:"" json:"approvalCode"`
	Pan            string `gorm:"" json:"pan"`
}

// PaymentAmountInfo содержит информацию об оплате
type PaymentAmountInfo struct {
	PaymentState    string `gorm:"not null" json:"paymentState"`
	ApprovedAmount  int    `gorm:"" json:"approvedAmount"`
	DepositedAmount int    `gorm:"" json:"depositedAmount"`
	RefundedAmount  int    `gorm:"" json:"refundedAmount"`
}

// BankInfo содержит информацию о банке
type BankInfo struct {
	BankName        string `gorm:"" json:"bankName"`
	BankCountryCode string `gorm:"not null" json:"bankCountryCode"`
	BankCountryName string `gorm:"not null" json:"bankCountryName"`
}
