package main

import (
	"context"
	"log"

	"github.com/Milad75Rasouli/portfolio/internal/config"
	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/handler"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/zap"
)

func main() {
	var (
		logger *zap.Logger
		err    error
	)

	cfg := config.New()
	log.Printf("Config:%+v", cfg)

	// var db *sqlitex.Pool
	db, err := db.New(cfg.Database)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)
	conn := db.Get(context.TODO())
	_ = conn.Prep(`CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT,
		image_path TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		modified_at DATETIME 
	);`)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Put(conn)

	engine := html.New("frontend/views/pages/", ".html")
	if cfg.Debug == true {
		logger, err = zap.NewDevelopment()
		engine.Reload(true)
	} else {
		logger, err = zap.NewProduction()
		engine.Reload(false)
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync()
	app := fiber.New(fiber.Config{
		Immutable: true,
		Views:     engine,
	})
	{
		logger := logger.Named("http")
		h := handler.Home{
			Logger: logger.Named("home"),
		}

		b := handler.Blog{
			Logger: logger.Named("blog"),
		}

		a := handler.Auth{
			Logger: logger.Named("auth"),
		}

		home := app.Group("/")
		blog := app.Group("/blog")
		auth := app.Group("/user")

		h.Register(home)
		b.Register(blog)
		a.Register(auth)
	}

	app.Static("/static", "./frontend/static")
	log.Fatalln(app.Listen(":5001"))
}
