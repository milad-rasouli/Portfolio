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
		logger *zap.Logger
		err    error
		db     store.Store
	)

	cfg := config.New()
	log.Printf("Config:%+v", cfg)

	sqlite := sqlitedb.SqliteInit{Folder: "data"}
	db, cancelDB, err := sqlite.Init(false, cfg.Database, logger)
	if err != nil {
		log.Fatal(err)
	}
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
		log.Fatal(err)
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
			Logger:    logger.Named("home"),
			HomeStore: db,
		}
		am := handler.AboutMe{
			Logger:       logger.Named("aboutMe"),
			AboutMeStore: db,
		}
		b := handler.Blog{
			Logger: logger.Named("blog"),
		}
		c := handler.Contact{
			Logger:       logger.Named("contact"),
			ContactStore: db,
		}
		a := handler.Auth{
			AdminEmail:   cfg.AdminEmail,
			Logger:       logger.Named("auth"),
			UserStore:    db,
			UserPassword: userPassword,
			JWTToken:     jwtToken,
		}

		cp := handler.ControlPanel{
			Logger: logger.Named("control-panel"),
			DB:     db,
		}

		home := app.Group("/")
		aboutMe := app.Group("about-me")
		blog := app.Group("/blog", a.LimitToAuthMiddleWare)
		contact := app.Group("/contact")
		auth := app.Group("/user")
		controlPanel := app.Group("/admin", a.LimitToAdminMiddleWare) //TODO: add an auth middleware for this path with only admin access

		h.Register(home)
		am.Register(aboutMe)
		b.Register(blog)
		c.Register(contact)
		a.Register(auth)
		cp.Register(controlPanel)
	}

	app.Static("/static", "./frontend/static")

	log.Fatalln(app.Listen(":5003"))
}
