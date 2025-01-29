package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skewertoreversepolarity/vtbApiValidation/tree/main/internal/models"
	"net/http"
	"time"
)

func (h *Handlers) GetOrderStatusExtendedHandler(c *gin.Context) {
	// Вызываем функцию GetOrderStatusExtended, которая отправляет запрос к VTB API
	response, err := h.GetOrderStatusExtended(c)
	if err != nil {
		h.Logger.Error("ошибка при получении статуса ордера", err, err.Error())
		// Если произошла ошибка при запросе, возвращаем ошибку в ответ
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_ = bytes.NewReader(response)

	var responseBody models.GetOrderStatusExtended
	if err := json.Unmarshal(response, &responseBody); err != nil {
		h.Logger.Error("Ошибка при парсинге ответа", "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	responseBody.Date = time.UnixMilli(responseBody.Date.UnixMilli())
	responseBody.AuthDateTime = time.UnixMilli(responseBody.AuthDateTime.UnixMilli())
	if err := h.DB.Create(&responseBody); err != nil {
		h.Logger.Error("failed to save order status", "err", err.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error.Error(),
		})
		return
	}
	//if err := json.NewDecoder(r).Decode(&responseBody); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	fmt.Println(responseBody)

	// Возвращаем успешный ответ с результатом запроса
	c.JSON(http.StatusOK, gin.H{
		"response": responseBody,
	})
}
