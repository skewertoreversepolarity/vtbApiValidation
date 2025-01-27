package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

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
	response, err := h.RegistrOrder(c)
	if err != nil {
		// Если произошла ошибка при запросе, возвращаем ошибку в ответ
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Возвращаем успешный ответ с результатом запроса
	c.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}
