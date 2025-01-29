package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skewertoreversepolarity/vtbApiValidation/tree/main/handlers"
	"github.com/skewertoreversepolarity/vtbApiValidation/tree/main/internal"
	"github.com/skewertoreversepolarity/vtbApiValidation/tree/main/internal/models"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := internal.GormConnect()
	if err != nil {
		logger.Error("cannot connect to database", "err", err.Error())
		return
	}

	if err := db.AutoMigrate(&models.GetOrderStatusExtended{}); err != nil {
		logger.Error("cannot auto migrate order", "err", err.Error())
		return
	}

	h := handlers.NewHandlers(db, logger)

	router := gin.New()

	// Группа маршрутов для API VTB
	vtbPaymentHandlers := router.Group("/vtb.rbsuat.com/payment")
	{
		vtbPaymentHandlers.POST("/rest/register.do", h.RegisterOrderHandler) // Обработчик POST-запроса
		vtbPaymentHandlers.POST("/rest/getOrderStatusExtended.do", h.GetOrderStatusExtendedHandler)
	}

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		logger.Error("server failed to start", "err", err.Error())
	}
}
