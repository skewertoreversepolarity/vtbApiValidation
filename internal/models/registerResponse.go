package models

type RegisterResponse struct {
	OrderID string `json:"orderId"`
	FormURL string `json:"formUrl"`
}
