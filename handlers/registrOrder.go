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

// RegisterOrder регистрация заказа и получения ссылки для редиректа пользователя на форму оплаты
func (h *Handlers) RegisterOrder(c *gin.Context) ([]byte, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URL") + os.Getenv("ORDER_REGISTRATION_URL")

	method := "POST"

	// Получаем параметры из запроса
	amount := c.PostForm("amount")
	currency := c.PostForm("currency")
	language := c.PostForm("language")
	orderNumber := c.PostForm("orderNumber")
	returnUrl := c.PostForm("returnUrl")
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	clientId := c.PostForm("clientId")

	// Формируем тело запроса с полученными параметрами
	payload := strings.NewReader(fmt.Sprintf(
		"amount=%s&currency=%s&language=%s&orderNumber=%s&returnUrl=%s&userName=%s&password=%s&clientId=%s",
		amount, currency, language, orderNumber, returnUrl, userName, password, clientId,
	))

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Создаем HTTP-запрос
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return []byte(""), fmt.Errorf("error creating request: %w", err)
	}

	// Устанавливаем заголовки
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Добавляем авторизацию (Basic Auth)
	req.SetBasicAuth(userName, password)

	// Отправляем запрос
	res, err := client.Do(req)
	if err != nil {
		return []byte(""), fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte(""), fmt.Errorf("Error reading response: %w", err)
	}

	// Возвращаем тело ответа как строку
	return body, nil
}

//func (h *Handlers) RegistrOrder(c *gin.Context) (string, error) {
//	url := "https://vtb.rbsuat.com/payment/rest/register.do"
//	method := "POST"
//
//	// Формируем тело запроса
//	payload := strings.NewReader("amount=100&currency=643&language=en&orderNumber=503&returnUrl=https%3A%2F%2Fmybestmerchantreturnurl.com&userName=Test-Api-api&password=cnngn-D0&clientId=259753456")
//
//	// Создаем HTTP-клиент
//	client := &http.Client{}
//
//	// Создаем HTTP-запрос
//	req, err := http.NewRequest(method, url, payload)
//	if err != nil {
//		fmt.Println("Error creating request:", err)
//		return "", fmt.Errorf("error creating request: %w", err)
//	}
//
//	// Устанавливаем заголовки
//	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//
//	// Добавляем авторизацию (Basic Auth)
//	req.SetBasicAuth("Test-Api-api", "cnngn-D0")
//
//	// Отправляем запрос
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println("Error sending request:", err)
//		return "", fmt.Errorf("error sending request: %w", err)
//	}
//	defer res.Body.Close()
//
//	// Читаем ответ
//	body, err := io.ReadAll(res.Body)
//	if err != nil {
//
//		return "", fmt.Errorf("Error reading response: %w", err)
//	}
//	return string(body), nil
//}
