package rest_api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"gandiwa/internal/core"
	"gandiwa/internal/presentation/restApi/handler"
	"gandiwa/internal/presentation/restApi/rest_api_common"
	validator3 "gandiwa/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

type Server interface {
	Start() error
	Shutdown() error
	CallServerApp() *fiber.App
}

type serverContainer struct {
	app     *fiber.App
	appPort string
}

func NewFiberServer(useCases core.UseCases) Server {
	sc := serverContainer{}

	validatorEngine, err := validator3.NewEngine()
	if err != nil {
		log.Fatalf("failed to set up validator engine. err:%s", err.Error())
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: rest_api_common.ResponseErr,
	})

	app.Use(helmet.New())
	app.Use(recover2.New())

	appHandler := handler.NewHandler(app, useCases, validatorEngine)
	appHandler.SetUpHandler()

	sc.app = app

	return &sc
}

func (s *serverContainer) Start() error {
	if err := s.app.Listen(fmt.Sprintf(":%s", s.appPort)); err != nil {
		// ErrServerClosed is expected behaviour when exiting app
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server is closed caused by: %s", err.Error())
		}

		return err
	}

	return nil
}

func (s *serverContainer) Shutdown() error {
	if err := s.app.Shutdown(); err != nil {
		return err
	}

	log.Println("http Server is stopped")

	return nil
}

func (s *serverContainer) CallServerApp() *fiber.App {
	return s.app
}
