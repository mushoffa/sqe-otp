package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sqe-otp/config"
	"sqe-otp/infrastructure/postgres"
	"sqe-otp/infrastructure/rest"
	"sqe-otp/presentation/controller"
	"sqe-otp/presentation/repository"
	"sqe-otp/usecase"
)

func main() {
	shutdown, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM /*syscall.SIGTSTP*/)
	defer stop()

	cfg := config.Get()
	db := postgres.New(cfg.Database)
	server := rest.NewServer(cfg.HttpServer)

	r := repository.NewOtpRepository(db)
	u := usecase.NewOtpUsecase(r)
	c := controller.NewOtpController(u)

	server.AddRoutes(
		config.BASE_API,
		c.Routes(),
	)
	go server.Start()

	<-shutdown.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db.Shutdown()
	server.Shutdown(ctx)
}
