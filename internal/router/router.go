package router

import (
	"github.com/fbpr/mnc-test-go/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app                *fiber.App
	authHandler        *http.AuthHandler
	transactionHandler *http.TransactionHandler
}

func NewRouter(app *fiber.App, authHandler *http.AuthHandler, transactionHandler *http.TransactionHandler) *Router {
	return &Router{
		app:                app,
		authHandler:        authHandler,
		transactionHandler: transactionHandler,
	}
}

func (r *Router) Routes() {
	api := r.app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/login", r.authHandler.Login)
	auth.Post("/logout", r.authHandler.Logout)

	transaction := api.Group("/transactions")
	transaction.Post("/:id/pay", r.transactionHandler.Pay)
}
