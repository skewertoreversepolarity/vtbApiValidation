package models

type RegisterOrderResponse struct {
	OrderID string `json:"orderId"`
	FormURL string `json:"formUrl"`
}
