package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/skewertoreversepolarity/vtbApiValidation/tree/main/internal/models"
	"io"
	"net/http"
)

func SendRequest(r models.Requset) ([]byte, error) {
	body := new(bytes.Buffer)
	for key, value := range r.Headers {
		body.WriteString(key + "=" + value + "&")
	}
	body.Truncate(body.Len() - 1) // remove the trailing '&'

	// create HTTP request
	req, err := http.NewRequest(r.Method, r.Endpoint, body)
	if err != nil {
		return nil, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	// Send the req
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return nil, err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, bodyBytes, "", "  ")
	if err != nil {
		fmt.Println("Response", string(bodyBytes))
		return nil, err
	}

	fmt.Println(prettyJSON.String())

	return io.ReadAll(resp.Body)
}
