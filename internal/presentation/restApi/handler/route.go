package handler

import (
	"gandiwa/internal/core"
	"gandiwa/internal/presentation/restApi/rest_api_common"
	"gandiwa/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	app             *fiber.App
	useCases        core.UseCases
	validatorEngine *validator.ValidationEngine
}

func (h Handler) SetUpHandler() {
	h.app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Hello World :)")
	})

	v1 := h.app.Group("/api/v1")
	v1.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return rest_api_common.ResponseOKWithMsg(ctx, "i'm ok, are u okay? ", struct{}{})
	})
}

func NewHandler(app *fiber.App, useCase core.UseCases, validatorEngine *validator.ValidationEngine) *Handler {
	return &Handler{
		app:             app,
		useCases:        useCase,
		validatorEngine: validatorEngine,
	}
}
