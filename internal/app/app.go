package app

import (
	"context"
	"corporate/config"
	"corporate/internal/adapter/handler"
	"corporate/internal/adapter/repository"
	"corporate/internal/core/service"
	"corporate/utils/auth"
	"corporate/utils/validator"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	en "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatalf("Error connection to database: %v", err)
		return
	}

	jwt := auth.NewJwt(cfg)
	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo, cfg, jwt)

	e := echo.New()
	e.Use(middleware.CORS())
	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(customValidator.Validator, customValidator.Translator)
	e.Validator = customValidator
	e.GET("/api/check", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	handler.NewUserHandler(e, userService)

	// starting server

	go func() {
		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv("APP_PORT")
		}
		err := e.Start(":" + cfg.App.AppPort)
		if err != nil {
			log.Fatal("error starting server: ", err)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	// block until a signal is received
	<-quit
	log.Println("server shutdown of 5 second")

	// gracefully shutdown the server, waiting max 5 seconds for current operations to complete

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	e.Shutdown(ctx)
}
