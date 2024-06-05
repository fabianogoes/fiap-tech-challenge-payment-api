package main

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-payment/domain/usecases"
	"github.com/fabianogoes/fiap-payment/frameworks/repository"
	"log/slog"
	"os"

	"github.com/fabianogoes/fiap-payment/domain/entities"

	"github.com/fabianogoes/fiap-payment/frameworks/rest"
)

func init() {
	fmt.Println("Initializing...")

	var logHandler *slog.JSONHandler

	config, _ := entities.NewConfig()
	if config.Environment == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func main() {
	fmt.Println("Starting web server...")

	ctx := context.Background()
	var err error

	config, err := entities.NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := repository.InitDB(ctx, config)
	if err != nil {
		panic(err)
	}

	paymentRepository := repository.NewPaymentRepository(db)
	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	paymentHandler := rest.NewPaymentHandler(paymentUseCase)

	router, err := rest.NewRouter(
		paymentHandler,
	)
	if err != nil {
		panic(err)
	}

	err = router.Run(config.AppPort)
	if err != nil {
		panic(err)
	}
}
