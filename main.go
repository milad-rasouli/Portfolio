package main

import (
	"log"

	"github.com/Milad75Rasouli/portfolio/internal/config"
	"github.com/Milad75Rasouli/portfolio/internal/handler"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()
	log.Printf("Config:%+v", cfg)

	var (
		logger *zap.Logger
		err    error
	)
	if cfg.Debug == true {
		logger, err = zap.NewDevelopment()
	} else {

		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync()

	app := fiber.New()
	{
		logger := logger.Named("http")
		h := handler.Home{
			Logger: logger.Named("home"),
		}

		b := handler.Blog{
			Logger: logger.Named("blog"),
		}

		home := app.Group("/")
		blog := app.Group("/blog")

		h.Register(home)
		b.Register(blog)
	}

	app.Listen(":5000")
}
