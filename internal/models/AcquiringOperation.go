package models

import "time"

type AcquiringOper struct {
	ID                  int64      `json:"id"`                  // bigint, первичный ключ
	Term                int        `json:"term"`                // int, обязательное поле
	Code                string     `json:"code"`                // nvarchar(50), обязательное поле
	Merchant            *string    `json:"merchant"`            // nvarchar(10), NULL
	OrderID             *string    `json:"orderId"`             // nvarchar(50), NULL
	SessionID           *string    `json:"sessionId"`           // nvarchar(50), NULL
	OrderIDEncrypted    *string    `json:"orderIdEncrypted"`    // nvarchar(50), NULL
	PAN                 *string    `json:"pan"`                 // nvarchar(50), NULL
	AmountCur           *float64   `json:"amountCur"`           // float, NULL
	Amount              *float64   `json:"amount"`              // float, NULL
	Note                *string    `json:"note"`                // nvarchar(50), NULL
	Articul             *string    `json:"articul"`             // nvarchar(50), NULL
	TryCount            int        `json:"tryCount"`            // int, обязательное поле
	Status              int        `json:"status"`              // int, обязательное поле
	Currency            *string    `json:"currency"`            // nvarchar(10), NULL
	Date                time.Time  `json:"date"`                // date, обязательное поле
	InputDate           *string    `json:"inputDate"`           // nvarchar(50), NULL
	OutputDate          *string    `json:"outputDate"`          // nvarchar(50), NULL
	Phone               *string    `json:"phone"`               // nvarchar(20), NULL
	TranDateTime        *time.Time `json:"tranDateTime"`        // datetime, NULL
	ResponseCode        *string    `json:"responseCode"`        // nvarchar(10), NULL
	ResponseDescription *string    `json:"responseDescription"` // nvarchar(250), NULL
	CardHolderName      *string    `json:"cardHolderName"`      // nvarchar(50), NULL
	Brand               *string    `json:"brand"`               // nvarchar(20), NULL
	ApprovalCode        *string    `json:"approvalCode"`        // nvarchar(20), NULL
	OrderDescription    *string    `json:"orderDescription"`    // nvarchar(550), NULL
	AcqFee              *string    `json:"acqFee"`              // nvarchar(10), NULL
	CardID              *string    `json:"cardId"`              // nvarchar(50), NULL
	UrlRequest          *string    `json:"urlRequest"`          // nvarchar(500), NULL
	ApproveURL          *string    `json:"approveUrl"`          // nvarchar(500), NULL
	DeclineURL          *string    `json:"declineUrl"`          // nvarchar(500), NULL
	CancelURL           *string    `json:"cancelUrl"`           // nvarchar(500), NULL
	AgentID             *string    `json:"agentId"`             // nvarchar(50), NULL
	ACSURL              *string    `json:"acsUrl"`              // nvarchar(max), NULL
	PaReq               *string    `json:"paReq"`               // nvarchar(max), NULL
	MD                  *string    `json:"md"`                  // nvarchar(max), NULL
	Month1              *string    `json:"month1"`              // nvarchar(100), NULL
	Year1               *string    `json:"year1"`               // nvarchar(100), NULL
	CVC1                *string    `json:"cvc1"`                // nvarchar(100), NULL
	Respond             *string    `json:"respond"`             // nvarchar(max), NULL
	Request             *string    `json:"request"`             // nvarchar(max), NULL
	StatusCallback      int        `json:"statusCallback"`      // int, обязательное поле
	SendCallback        *string    `json:"sendCallback"`        // nvarchar(50), NULL
}
