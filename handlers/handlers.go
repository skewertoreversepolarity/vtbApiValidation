package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skewertoreversepolarity/vtbApiValidation/tree/main/internal/models"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

type RegisterOrderResponse struct {
	OrderID string `json:"orderId"`
	FormUrl string `json:"formUrl"`
}
type Handlers struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

func NewHandlers(db *gorm.DB, logger *slog.Logger) *Handlers {
	return &Handlers{
		DB:     db,
		Logger: logger,
	}
}

func (h *Handlers) RegistrOrderHandler(c *gin.Context) {
	// Вызываем функцию RegistrOrder, которая отправляет запрос к VTB API
	response, err := h.RegisterOrder(c)
	if err != nil {
		// Если произошла ошибка при запросе, возвращаем ошибку в ответ
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := bytes.NewReader(response)

	var responseBody RegisterOrderResponse

	if err := json.NewDecoder(r).Decode(&responseBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Возвращаем успешный ответ с результатом запроса
	c.JSON(http.StatusOK, gin.H{
		"orderId": responseBody.OrderID,
		"formUrl": responseBody.FormUrl,
	})
}

func (h *Handlers) GetOrderStatusExtendedHandler(c *gin.Context) {
	// Вызываем функцию GetOrderStatusExtended, которая отправляет запрос к VTB API
	response, err := h.GetOrderStatusExtended(c)
	if err != nil {
		// Если произошла ошибка при запросе, возвращаем ошибку в ответ
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := bytes.NewReader(response)

	var responseBody models.GetOrderStatusExtended

	if err := json.NewDecoder(r).Decode(&responseBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(responseBody)

	// Возвращаем успешный ответ с результатом запроса
	c.JSON(http.StatusOK, gin.H{
		"response": responseBody,
	})
}
