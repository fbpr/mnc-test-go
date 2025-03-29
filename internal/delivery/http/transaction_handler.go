package http

import (
	"github.com/fbpr/mnc-test-go/internal/domain"
	"github.com/fbpr/mnc-test-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	paymentUseCase *usecase.TransactionUseCase
}

func NewTransactionHandler(paymentUseCase usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		paymentUseCase: &paymentUseCase,
	}
}

func (h *TransactionHandler) Pay(ctx *fiber.Ctx) error {
	transactionID := ctx.Params("id")
	if transactionID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "transaction ID is required",
		})
	}

	var paymentRequest domain.PaymentRequest
	if err := ctx.BodyParser(&paymentRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid request format",
		})
	}

	if paymentRequest.CustomerID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "customer ID is required",
		})
	}

	response, err := h.paymentUseCase.ProcessPayment(transactionID, paymentRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "payment successful",
		"data":    response,
	})
}
