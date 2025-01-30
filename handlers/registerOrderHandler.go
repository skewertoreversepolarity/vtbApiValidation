package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/skewertoreversepolarity/vtbApiValidation/tree/main/internal/models"
	"gorm.io/datatypes"
	"net/http"
)

func (h *Handlers) RegisterOrderHandler(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Вызываем функцию RegistrOrder, которая отправляет запрос к VTB API
	response, err := h.RegisterOrder(requestData)
	if err != nil {
		// Если произошла ошибка при запросе, возвращаем ошибку в ответ
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var responseBody models.RegisterOrderResponse
	if err := json.Unmarshal(response, &responseBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	order := models.Order{
		OrderNumber: requestData["orderNumber"].(string),
		OrderID:     responseBody.OrderID,
		Amount:      int(requestData["amount"].(float64)),
		Currency:    requestData["currency"].(string),
	}

	jsonData, err := json.Marshal(order)
	if err != nil {
		c.JSON
	}
	//r := bytes.NewReader(response)
	//
	//var responseBody models.RegisterResponse
	//
	//if err := json.NewDecoder(r).Decode(&responseBody); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//
	//order := models.RegisterResponse{
	//	OrderID: responseBody.OrderID,
	//	FormURL: responseBody.FormURL,
	//}
	//
	//if err := h.DB.Create(&order).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error": "Failed to save order to DB " + err.Error(),
	//	})
	//	return
	//}
	//// Возвращаем успешный ответ с результатом запроса
	//c.JSON(http.StatusOK, gin.H{
	//	"orderId": responseBody.OrderID,
	//	"formUrl": responseBody.FormURL,
	//})
}
