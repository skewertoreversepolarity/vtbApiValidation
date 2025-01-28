package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func (h *Handlers) GetOrderStatusExtended(c *gin.Context) ([]byte, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URL") + os.Getenv("GET_ORDER_STATUS_EXTENDED")
	method := "POST"
	//получаем параметры из запроса
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	language := c.PostForm("language")
	orderNumber := c.PostForm("orderNumber")

	payload := strings.NewReader(fmt.Sprintf(
		"userName=%s&password=%s&language=%s&orderId=%s",
		userName, password, language, orderNumber,
	))

	//формируем тоело запроса с получением параметров
	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(userName, password)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	//читаем ответ
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении ответа%s", err)
	}

	return body, nil

}
