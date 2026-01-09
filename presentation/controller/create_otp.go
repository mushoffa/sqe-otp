package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c *otp) createOtp() fiber.Handler {
	request := new(create_otp_request)
	return func(ctx *fiber.Ctx) error {
		// if err := c.ValidateBody(ctx, request); err != nil {
		// 	return c.BadRequest(ctx, err.Error())
		// }
		if err := ctx.BodyParser(request); err != nil {
			return c.BadRequest(ctx, err)
		}

		otp, err := c.u.GenerateOtp(ctx.UserContext(), request.UserID)
		if err != nil {
			return c.BadRequest(ctx, err.Error())
		}

		return c.Success(ctx, map[string]any{
			"user_id": request.UserID,
			"otp":     otp,
		})
	}
}

type create_otp_request struct {
	UserID string `json:"user_id" validate:"required"`
}
