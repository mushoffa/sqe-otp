package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c *otp) validateOtp() fiber.Handler {
	request := new(validate_otp_request)
	return func(ctx *fiber.Ctx) error {
		// if err := c.ValidateBody(ctx, request); err != nil {
		// 	return c.BadRequest(ctx, err.Error())
		// }

		if err := ctx.BodyParser(request); err != nil {
			return c.BadRequest(ctx, err)
		}

		if err := c.u.ValidateOtp(ctx.UserContext(), request.UserID, request.Otp); err != nil {
			return c.InternalServerError(ctx, err)
		}

		// TODO: Create success response data structure
		return c.Success(ctx, map[string]any{
			"user_id": request.UserID,
			"message": "OTP validated succesfully.",
		})
	}
}

type validate_otp_request struct {
	UserID string `json:"user_id" validate:"required"`
	Otp    string `json:"otp" validate:"required"`
}
