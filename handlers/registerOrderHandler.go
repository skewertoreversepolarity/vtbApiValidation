package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/skewertoreversepolarity/vtbApiValidation/tree/main/internal/models"
	"net/http"
)

func (h *Handlers) RegisterOrderHandler(c *gin.Context) {
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

	var responseBody models.RegisterResponse

	if err := json.NewDecoder(r).Decode(&responseBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	order := models.RegisterResponse{
		OrderID: responseBody.OrderID,
		FormURL: responseBody.FormURL,
	}

	if err := h.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save order to DB " + err.Error(),
		})
		return
	}
	// Возвращаем успешный ответ с результатом запроса
	c.JSON(http.StatusOK, gin.H{
		"orderId": responseBody.OrderID,
		"formUrl": responseBody.FormURL,
	})
}
