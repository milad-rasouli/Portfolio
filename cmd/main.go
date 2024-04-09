package main

import (
	"github.com/Milad75Rasouli/portfolio/internal/handler"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func main() {
	app := fiber.New()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

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
