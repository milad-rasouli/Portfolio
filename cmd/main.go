package main

import (
	"log"

	"github.com/Milad75Rasouli/portfolio/internal/api"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	endpoint := ":5000"
	a, err := api.New(endpoint, logger.Named("api"))
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()
	app.Get("/", a.HomeHandler)
	app.Get("/post", a.PostListHandler)
	app.Get("/post/:postID", a.PostHandler)
	log.Fatal(app.Listen(endpoint))
}
