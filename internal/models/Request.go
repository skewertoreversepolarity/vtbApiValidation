package models

type Requset struct {
	Method   string            `json:"method"`
	Endpoint string            `json:"endpoint"`
	Headers  map[string]string `json:"headers"`
	Body     map[string]string `json:"body"`
}
