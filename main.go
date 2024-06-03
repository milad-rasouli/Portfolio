package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/cipher"
	"github.com/Milad75Rasouli/portfolio/internal/config"
	"github.com/Milad75Rasouli/portfolio/internal/handler"
	"github.com/Milad75Rasouli/portfolio/internal/jwt"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	sqlitedb "github.com/Milad75Rasouli/portfolio/internal/store/sqliteDB"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

func main() {
	var (
		logger *zap.Logger
		err    error
		db     store.Store
	)

	cfg := config.New()

	sqlite := sqlitedb.SqliteInit{Folder: "data"}
	db, cancelDB, err := sqlite.Init(false, cfg.Database, logger)
	if err != nil {
		logger.Fatal("sqlite init error", zap.Error(err))
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
	logger.Info("Configuration", zap.Any("config.toml", cfg))

	if err != nil {
		logger.Fatal("zap error", zap.Error(err))
	}
	defer logger.Sync()

	app := fiber.New(fiber.Config{
		// Immutable: true,
		AppName: "Milad Rasouli Portfolio",
		Views:   engine,
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
		m := handler.NewMetricsMiddleware(logger.Named("metrics"))

		cp := handler.ControlPanel{
			Logger: logger.Named("control-panel"),
			DB:     db,
		}

		home := app.Group("/", m.Middleware)
		aboutMe := app.Group("/about-me", m.Middleware)
		blog := app.Group("/blog", a.LimitToAuthMiddleWare, m.Middleware)
		contact := app.Group("/contact", m.Middleware)
		auth := app.Group("/user", m.Middleware)
		// controlPanel := app.Group("/admin", a.LimitToAdminMiddleWare, m.Middleware)
		controlPanel := app.Group("/admin", m.Middleware) // TODO: uncomment out the above line and remove this line
		app.Get("/health", handler.GetHealth)

		h.Register(home)
		am.Register(aboutMe)
		b.Register(blog)
		c.Register(contact)
		a.Register(auth)
		cp.Register(controlPanel)
	}

	app.Static("/static", "./frontend/static")
	go func() {
		http.Handle("GET /metrics", promhttp.Handler())
		http.ListenAndServe(":5000", nil)
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- app.Listen(":5001")
	}()

	select {
	case err := <-serverErrors:
		logger.Fatal("listening and serving error", zap.Error(err))

	case <-stop:
		logger.Info("main : Start shutdown")
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err := app.ShutdownWithContext(ctx); err != nil {
			logger.Fatal("graceful shutdown did not complete in the timeout", zap.Any("timeout", timeout), zap.Error(err))
		}
	}

}
