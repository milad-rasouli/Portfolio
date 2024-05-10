package main

import (
	"log"

	"github.com/Milad75Rasouli/portfolio/internal/cipher"
	"github.com/Milad75Rasouli/portfolio/internal/config"
	"github.com/Milad75Rasouli/portfolio/internal/handler"
	"github.com/Milad75Rasouli/portfolio/internal/jwt"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	sqlitedb "github.com/Milad75Rasouli/portfolio/internal/store/sqliteDB"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/zap"
)

func main() {
	var (
		logger    *zap.Logger
		err       error
		userStore store.Store
	)

	cfg := config.New()
	log.Printf("Config:%+v", cfg)

	sqlite := sqlitedb.SqliteInit{Folder: "data"}
	userStore, cancelDB, err := sqlite.Init(false, cfg.Database, logger)
	defer cancelDB()

	userPassword := cipher.NewUserPassword(cfg.Cipher)

	jwtToken := jwt.New(cfg.JWT)

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

	// app.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token",
	// 	CookieName:     "_csrf",	// app.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token",
	// 	CookieName:     "_csrf",
	// 	CookieSameSite: "Strict",
	// }))

	// 	CookieSameSite: "Strict",
	// }))

	{
		logger := logger.Named("http")
		h := handler.Home{
			Logger: logger.Named("home"),
		}

		b := handler.Blog{
			Logger: logger.Named("blog"),
		}

		a := handler.Auth{
			Logger:       logger.Named("auth"),
			UserStore:    userStore,
			UserPassword: userPassword,
			JWTToken:     jwtToken,
		}

		home := app.Group("/")
		blog := app.Group("/blog", a.LimitToAuthMiddleWare)
		auth := app.Group("/user")

		h.Register(home)
		b.Register(blog)
		a.Register(auth)
	}

	app.Static("/static", "./frontend/static")
	log.Fatalln(app.Listen(":5001"))
}
