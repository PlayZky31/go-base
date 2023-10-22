package rest_api_common

import (
	customError2 "gandiwa/pkg/customError"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`

	Error    string            `json:"error,omitempty"`
	ErrField map[string]string `json:"errorField,omitempty"`
}

func ResponseOK(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(Response{
		Status:  200,
		Message: "success",
		Data:    data,
	})
}

func ResponseOKWithMsg(ctx *fiber.Ctx, msg string, data any) error {
	return ctx.JSON(Response{
		Status:  200,
		Message: msg,
		Data:    data,
	})
}

func ResponseErr(ctx *fiber.Ctx, err error) error {
	resp := Response{
		Status:  500,
		Message: customError2.GeneralErrMessage,
		Error:   err.Error(),
		Data:    struct{}{},
	}

	if val, ok := err.(*fiber.Error); ok {
		resp.Status = val.Code
		resp.Message = val.Message
		resp.Error = val.Error()
		return ctx.JSON(resp)
	}

	if val, ok := err.(*customError2.CustomError); ok {
		resp.Status = val.Code
		resp.Message = val.Message
		resp.Error = val.Error()
		if val.ErrField != nil {
			resp.ErrField = val.ErrField
		}
		return ctx.JSON(resp)
	}

	return ctx.JSON(resp)
}
