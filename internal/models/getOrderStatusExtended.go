package models

// GetOrderStatusExtended представляет структуру ответа
type GetOrderStatusExtended struct {
	ID                    int64             `gorm:"primaryKay" json:"id"`
	ErrorCode             string            `gorm:"not null" json:"errorCode"`
	ErrorMessage          string            `gorm:"not null" json:"errorMessage"`
	OrderNumber           string            `gorm:"unique;not null" json:"orderNumber"`
	OrderStatus           int               `gorm:"not null" json:"orderStatus"`
	ActionCode            int               `gorm:"not null" json:"actionCode"`
	ActionCodeDescription string            `gorm:"not null" json:"actionCodeDescription"`
	Amount                int               `gorm:"not null" json:"amount"`
	Currency              string            `gorm:"not null" json:"currency"`
	Date                  int64             `gorm:"not null" json:"date"`
	OrderDescription      string            `gorm:"not null" json:"orderDescription"`
	IP                    string            `gorm:"not null" json:"ip"`
	MerchantOrderParams   []interface{}     `gorm:"not null" json:"merchantOrderParams"`
	TransactionAttributes []interface{}     `gorm:"not null" json:"transactionAttributes"`
	Attributes            []Attribute       `gorm:"not null" json:"attributes"`
	CardAuthInfo          CardAuthInfo      `gorm:"not null" json:"cardAuthInfo"`
	AuthDateTime          int64             `gorm:"not null" json:"authDateTime"`
	TerminalID            string            `gorm:"not null" json:"terminalId"`
	AuthRefNum            string            `gorm:"not null" json:"authRefNum"`
	PaymentAmountInfo     PaymentAmountInfo `gorm:"not null" json:"paymentAmountInfo"`
	BankInfo              BankInfo          `gorm:"not null" json:"bankInfo"`
}

// Attribute представляет элемент массива attributes
type Attribute struct {
	Name  string `gorm:"" json:"name"`
	Value string `gorm:"" json:"value"`
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
	PaymentState    string `gorm:"" json:"paymentState"`
	ApprovedAmount  int    `gorm:"" json:"approvedAmount"`
	DepositedAmount int    `gorm:"" json:"depositedAmount"`
	RefundedAmount  int    `gorm:"" json:"refundedAmount"`
}

// BankInfo содержит информацию о банке
type BankInfo struct {
	BankName        string `gorm:"" json:"bankName"`
	BankCountryCode string `gorm:"" json:"bankCountryCode"`
	BankCountryName string `gorm:"" json:"bankCountryName"`
}
