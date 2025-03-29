package http

import (
	"github.com/fbpr/mnc-test-go/internal/domain"
	"github.com/fbpr/mnc-test-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	var request domain.LoginRequest
	var response domain.LoginResponse

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid request format",
		})
	}

	if request.Email == "" || request.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": request.Password,
		})
	}

	response, err := h.authUseCase.Login(request.Email, request.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "login successful",
		"data":    response,
	})
}

func (h *AuthHandler) Logout(ctx *fiber.Ctx) error {
	var request domain.LogoutRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid request format",
		})
	}

	if request.Email == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "email is required",
		})
	}

	err := h.authUseCase.Logout(request.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "logout successful",
	})
}
