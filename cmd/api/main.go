package main

import (
	"fmt"
	"path/filepath"

	"github.com/fbpr/mnc-test-go/config"
	handler "github.com/fbpr/mnc-test-go/internal/delivery/http"
	"github.com/fbpr/mnc-test-go/internal/repository/persistent"
	"github.com/fbpr/mnc-test-go/internal/router"
	"github.com/fbpr/mnc-test-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	dataDir := filepath.Join(".", "data")

	customerRepository := persistent.NewCustomerRepositoryJSON(dataDir)
	transactionRepository := persistent.NewTransactionRepositoryJSON(dataDir)
	historyRepository := persistent.NewHistoryRepositoryJSON(dataDir)

	authUseCase := usecase.NewAuthUseCase(customerRepository, historyRepository)
	authHandler := handler.NewAuthHandler(*authUseCase)

	transactionUseCase := usecase.NewTransactionUseCase(transactionRepository, customerRepository, historyRepository)
	transactionHandler := handler.NewTransactionHandler(*transactionUseCase)

	router := router.NewRouter(app, authHandler, transactionHandler)
	router.Routes()

	app.Listen(fmt.Sprintf(":%s", cfg.HttpPort))
}
