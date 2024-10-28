package handler

import (
	"edit-your-project-name/config"
	"edit-your-project-name/handler/middleware"
	"edit-your-project-name/handler/resp"
	"edit-your-project-name/slog"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"time"
)

func handler(app *fiber.App) {
	app.All("/health", func(c *fiber.Ctx) error {
		return resp.Suc(c, "ok - "+time.Now().Format(time.DateTime))
	})
}

// ====^

func InitHandler() {
	app := fiber.New(fiber.Config{
		JSONDecoder:           sonic.Unmarshal,
		JSONEncoder:           sonic.Marshal,
		ErrorHandler:          middleware.ErrorHandler,
		DisableStartupMessage: true,
		ReadTimeout:           time.Second * 10,
	})
	app.Use(recover.New(recover.Config{
		StackTraceHandler: func(_ *fiber.Ctx, e any) {
			slog.ErrWithStack("Panic Recover", e)
		},
	}), cors.New(), middleware.Limit(config.Fiber.LimitMax, config.Fiber.LimitExp))
	if config.Fiber.ListenOption == "HTTPS" {
		app.Use(middleware.ToHTTPS)
	}
	if config.Log.RequestLogStdout {
		app.Use(logger.New())
	}

	handler(app)

	switch config.Fiber.ListenOption {
	case "HTTP":
		if err := app.Listen(config.Fiber.HTTPListenAddr); err != nil {
			slog.Fatal("API HTTP Listen ERROR", err)
		}
	case "HTTPS":
		go func() {
			if err := app.Listen(config.Fiber.HTTPListenAddr); err != nil {
				slog.Fatal("API HTTP Listen ERROR", err)
			}
		}()
		if err := app.ListenTLS(config.Fiber.HTTPSListenAddr, config.Fiber.TLSCertFile, config.Fiber.TLSKeyFile); err != nil {
			slog.Fatal("API HTTPS Listen ERROR", err)
		}
	default:
		slog.Fatal("Config Fiber.ListenOption expect HTTP or HTTPS, but:", config.Fiber.ListenOption)
	}
}
