package rest

import (
	"github.com/gofiber/fiber/v2"
)

type BaseHandler struct {
}

func (c *BaseHandler) Success(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (c *BaseHandler) BadRequest(ctx *fiber.Ctx, error any) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(error)
}

func (c *BaseHandler) InternalServerError(ctx *fiber.Ctx, error any) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(error)
}
