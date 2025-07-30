package app

import (
	"context"
	"corporate/config"
	"corporate/internal/adapter/handler"
	"corporate/internal/adapter/repository"
	"corporate/internal/adapter/storage"
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
	heroSectionRepo := repository.NewHeroSectionRepository(db.DB)
	clientSectionRepo := repository.NewClientSectionRepository(db.DB)
	aboutCompanyRepo := repository.NewAboutCompanyRepository(db.DB)
	aboutCompanyKeynoteRepo := repository.NewAboutCompanyKeynoteRepository(db.DB)
	faqRepo := repository.NewFaqSectionRepository(db.DB)
	ourTeamRepo := repository.NewOurTeamRepository(db.DB)
	serviceSectionRepo := repository.NewServiceSectionRepository(db.DB)
	appointmentRepo := repository.NewAppointmentRepository(db.DB)
	portofolioRepo := repository.NewPortofolioSectionRepository(db.DB)
	portofolioDetailRepo := repository.NewPortofolioDetailRepository(db.DB)

	userService := service.NewUserService(userRepo, cfg, jwt)
	heroSectionService := service.NewHeroSectionService(heroSectionRepo)
	clientSectionService := service.NewClientSectionService(clientSectionRepo)
	aboutCompanyService := service.NewAboutCompanyService(aboutCompanyRepo)
	aboutCompanyKeynoteService := service.NewAboutCompanyKeynoteService(aboutCompanyKeynoteRepo, aboutCompanyRepo)
	faqService := service.NewFaqSectionService(faqRepo)
	ourTeamService := service.NewOurTeamService(ourTeamRepo)
	serviceSectionService := service.NewServiceSectionService(serviceSectionRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo)
	portofolioService := service.NewPortofolioSectionService(portofolioRepo)
	portofolioDetailService := service.NewPortofolioDetailService(portofolioDetailRepo, portofolioRepo)
	storageAdapter := storage.NewSupabase(cfg)

	e := echo.New()
	e.Use(middleware.CORS())
	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(customValidator.Validator, customValidator.Translator)
	e.Validator = customValidator
	e.GET("/api/check", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	handler.NewUserHandler(e, userService)
	handler.NewUploadImage(e, storageAdapter, cfg)
	handler.NewHeroSectionHandler(e, cfg, heroSectionService)
	handler.NewClientSectionHandler(e, clientSectionService, cfg)
	handler.NewAboutCompanyHandler(e, aboutCompanyService, cfg)
	handler.NewAboutCompanyKeynoteHandler(e, aboutCompanyKeynoteService, cfg)
	handler.NewFaqSectionHandler(e, faqService, cfg)
	handler.NewOurTeamHandler(e, ourTeamService, cfg)
	handler.NewServiceSectionHandler(e, serviceSectionService, cfg)
	handler.NewAppointmentHandler(e, appointmentService, cfg)
	handler.NewPortofolioSectionHandler(e, portofolioService, cfg)
	handler.NewPortofolioDetailHandler(e, portofolioDetailService, cfg)

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
