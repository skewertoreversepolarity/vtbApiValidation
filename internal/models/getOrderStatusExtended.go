package models

// GetOrderStatusExtended представляет структуру ответа
type GetOrderStatusExtended struct {
	ErrorCode             string            `json:"errorCode"`
	ErrorMessage          string            `json:"errorMessage"`
	OrderNumber           string            `json:"orderNumber"`
	OrderStatus           int               `json:"orderStatus"`
	ActionCode            int               `json:"actionCode"`
	ActionCodeDescription string            `json:"actionCodeDescription"`
	Amount                int               `json:"amount"`
	Currency              string            `json:"currency"`
	Date                  int64             `json:"date"`
	OrderDescription      string            `json:"orderDescription"`
	IP                    string            `json:"ip"`
	MerchantOrderParams   []interface{}     `json:"merchantOrderParams"`
	TransactionAttributes []interface{}     `json:"transactionAttributes"`
	Attributes            []Attribute       `json:"attributes"`
	CardAuthInfo          CardAuthInfo      `json:"cardAuthInfo"`
	AuthDateTime          int64             `json:"authDateTime"`
	TerminalID            string            `json:"terminalId"`
	AuthRefNum            string            `json:"authRefNum"`
	PaymentAmountInfo     PaymentAmountInfo `json:"paymentAmountInfo"`
	BankInfo              BankInfo          `json:"bankInfo"`
}

// Attribute представляет элемент массива attributes
type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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
	PaymentState    string `json:"paymentState"`
	ApprovedAmount  int    `json:"approvedAmount"`
	DepositedAmount int    `json:"depositedAmount"`
	RefundedAmount  int    `json:"refundedAmount"`
}

// BankInfo содержит информацию о банке
type BankInfo struct {
	BankName        string `json:"bankName"`
	BankCountryCode string `json:"bankCountryCode"`
	BankCountryName string `json:"bankCountryName"`
}
