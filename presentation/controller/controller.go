package controller

import (
	"sqe-otp/infrastructure/rest"
	"sqe-otp/usecase"

	"github.com/gofiber/fiber/v2"
)

type OtpController interface {
	Routes() func(fiber.Router)
}

type otp struct {
	rest.BaseHandler
	u usecase.OtpUsecase
}

func NewOtpController(u usecase.OtpUsecase) OtpController {
	return &otp{
		u: u,
	}
}

func (c *otp) Routes() func(fiber.Router) {
	return func(router fiber.Router) {
		router.Route("/otp", func(api fiber.Router) {
			api.Post("/request", c.createOtp())
			api.Post("/validate", c.validateOtp())
		})
	}
}
