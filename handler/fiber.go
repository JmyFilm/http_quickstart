package handler

import (
	"PROJECTNAME/conf"
	"PROJECTNAME/handler/middleware"
	"PROJECTNAME/xlog"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"time"
)

func Init() {
	app := fiber.New(fiber.Config{
		JSONDecoder:                  sonic.Unmarshal,
		JSONEncoder:                  sonic.Marshal,
		ErrorHandler:                 middleware.ErrorHandler,
		DisablePreParseMultipartForm: true,
		ReadTimeout:                  time.Second * 10,
		Concurrency:                  conf.Fiber.MaxConns * 1024,
	})
	app.Use(recover.New(recover.Config{
		StackTraceHandler: func(_ fiber.Ctx, e any) {
			xlog.ErrWithStack("Panic Recover", e)
		},
	}), cors.New(), middleware.Limit(conf.Fiber.LimitMax, conf.Fiber.LimitMax, conf.Fiber.LimitKey, conf.App.DebugMode))
	if conf.Fiber.ListenOption == "HTTPS" {
		app.Use(middleware.ToHTTPS)
	}
	if conf.Fiber.RequestStdOut {
		app.Use(logger.New())
	}

	handler(app)

	switch conf.Fiber.ListenOption {
	case "HTTP":
		if err := app.Listen(conf.Fiber.HTTPListenAddr, fiber.ListenConfig{
			//EnablePrefork:         true,
			GracefulContext:       xlog.ShutdownCtx,
			DisableStartupMessage: true,
			BeforeServeFunc: func(app *fiber.App) error {
				xlog.Info("HTTP Server Run At", conf.Fiber.HTTPListenAddr)
				return nil
			},
		}); err != nil {
			xlog.Fatal("API HTTP Listen ERROR", err)
		}
	case "HTTPS":
		go func() {
			if err := app.Listen(conf.Fiber.HTTPListenAddr, fiber.ListenConfig{
				//EnablePrefork:         true,
				GracefulContext:       xlog.ShutdownCtx,
				DisableStartupMessage: true,
			}); err != nil {
				xlog.Fatal("API HTTP Listen ERROR", err)
			}
		}()
		if err := app.Listen(conf.Fiber.HTTPSListenAddr, fiber.ListenConfig{
			//EnablePrefork:         true,
			GracefulContext:       xlog.ShutdownCtx,
			DisableStartupMessage: true,
			CertFile:              conf.Fiber.TLSCertFile,
			CertKeyFile:           conf.Fiber.TLSKeyFile,
			BeforeServeFunc: func(app *fiber.App) error {
				xlog.Info("HTTPS Server Run At", conf.Fiber.HTTPSListenAddr)
				return nil
			},
		}); err != nil {
			xlog.Fatal("API HTTPS Listen ERROR", err)
		}
	default:
		xlog.Fatal("Config ListenOption expect HTTP or HTTPS, but:", conf.Fiber.ListenOption)
	}
}
