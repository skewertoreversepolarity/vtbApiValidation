package main

import "vtbapivalidation/internal/models"

func main() {
	orderRegistration := models.Requset{
		Method:   "POST",
		Endpoint: "https://vtb.rbsuat.com/payment/rest/register.do",
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Body: map[string]string{
			"amount":       "100",
			"currency":     "643",
			"language":     "en",
			"returnUrl":    "https://mybestmerchantreturnurl.com",
			"userName":     "Test-Api-order",
			"userPassword": "cnngn-D0",
			"clientId":     "259753456",
		},
	}

}
