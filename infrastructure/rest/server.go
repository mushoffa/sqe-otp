package rest

import (
	"context"
	"errors"
	"net/http"

	c "sqe-otp/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type routerFn func(fiber.Router)
type mountFn func() (string, *fiber.App)

type HttpServer interface {
	AddConfig(...configFn)
	AddRoutes(string, ...routerFn)
	MountRoutes(...mountFn)
	AddMiddleware(...fiber.Handler)
	Start() error
	Shutdown(context.Context) error
}

type server struct {
	config
	instance *fiber.App
}

func NewServer(cfg c.HttpServer) HttpServer {
	config := defaultConfig()
	config.port = cfg.Port

	app := fiber.New()
	s := &server{
		config:   config,
		instance: app,
	}

	s.AddMiddleware(
		logger.New(
			logger.Config{
				Format:     "${time} | ${status} | ${latency} | ${method} | ${path} | ${body}\n",
				TimeFormat: "2006-01-02 15:04:05",
			},
		),
	)

	return s
}

func (s *server) AddConfig(configs ...configFn) {
	for _, config := range configs {
		config(&s.config)
	}
}

func (s *server) AddRoutes(api string, fn ...routerFn) {
	for _, f := range fn {
		s.instance.Route(api, f)
	}
}

func (s *server) MountRoutes(fn ...mountFn) {
	for _, f := range fn {
		s.instance.Mount(f())
	}
}

func (s *server) AddMiddleware(fn ...fiber.Handler) {
	for _, f := range fn {
		s.instance.Use(f)
	}
}

// Start starts [server] to run.
func (s *server) Start() error {
	if err := s.instance.Listen(s.config.address()); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// Shutdown shutdown the [server].
func (s *server) Shutdown(ctx context.Context) error {
	if err := s.instance.ShutdownWithContext(ctx); err != nil {
		return err
	}

	return nil
}
