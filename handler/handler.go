package handler

import (
	"PROJECTNAME/xlog"
	"github.com/gofiber/fiber/v3"
	"time"
)

func handler(app *fiber.App) {
	app.All("/", func(c fiber.Ctx) error {
		return c.Send([]byte(xlog.AppName()))
	})
	app.All("/monitoring/heartbeat", func(c fiber.Ctx) error {
		return c.Send([]byte("ok"))
	})
	app.Get("/time", func(c fiber.Ctx) error {
		return c.Send([]byte(time.Now().Format(time.DateTime)))
	})

}
