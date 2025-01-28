package models

type StatusResponse struct {
	ErrorCode             string `json:"errorCode"`
	ErrorMessage          string `json:"errorMessage"`
	OrderID               string `json:"orderID"`
	OrderStatus           int    `json:"orderStatus"`
	ActionCode            int    `json:"actionCode"`
	ActionCodeDescription string `json:"actionCodeDescription"`
	Amount                int    `json:"amount"`
	Currency              string `json:"currency"`
	MaskedPan             string `json:"maskedPan,omitempty"`
	CardholderName        string `json:"cardholderName,omitempty"`
	BankName              string `json:"bankName,omitempty"`
	ApprovedAmount        int    `json:"approvedAmount,omitempty"`
	Date                  int64  `json:"date"`
	OrderDescription      string `json:"orderDescription,omitempty"`
	Ip                    string `json:"ip,omitempty"`
	MerchantOrderParams   []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"merchantOrderParams,omitempty"`
	TransactionAttributes []interface{} `json:"transactionAttributes,omitempty"`
	Attributes            []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"attributes,omitempty"`
	CardAuthInfo struct {
		MaskedPan      string `json:"maskedPan"`
		Expiration     string `json:"expiration"`
		CardholderName string `json:"cardholderName"`
		ApprovalCode   string `json:"approvalCode"`
		Pan            string `json:"pan"`
	} `json:"cardAuthInfo,omitempty"`
	AuthDateTime      int64  `json:"authDateTime,omitempty"`
	TerminalId        string `json:"terminalId,omitempty"`
	AuthRefNum        string `json:"authRefNum,omitempty"`
	PaymentAmountInfo struct {
		PaymentState    string `json:"paymentState"`
		ApprovedAmount  int    `json:"approvedAmount"`
		DepositedAmount int    `json:"depositedAmount"`
		RefundedAmount  int    `json:"refundedAmount"`
	} `json:"paymentAmountInfo,omitempty"`
	BankInfo struct {
		BankName        string `json:"bankName"`
		BankCountryCode string `json:"bankCountryCode"`
		BankCountryName string `json:"bankCountryName"`
	} `json:"bankInfo,omitempty"`
}
