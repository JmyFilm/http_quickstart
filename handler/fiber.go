package handler

import (
	"edit-your-project-name/conf"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"time"
)

func InitHandler() {
	app := fiber.New(fiber.Config{
		AppName:      conf.Fiber.AppName,
		JSONDecoder:  sonic.Unmarshal,
		JSONEncoder:  sonic.Marshal,
		ErrorHandler: ErrorHandler,
	})
	app.Use(recover.New(recover.Config{
		StackTraceHandler: func(_ *fiber.Ctx, e any) {
			conf.ErrWithStack(e)
		},
	}), cors.New())

	app.All("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok - " + time.Now().Format(time.DateTime))
	})

	if err := app.Listen(conf.Fiber.Addr); err != nil {
		conf.Fatal("Fiber ERROR", err)
	}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	if e := new(fiber.Error); errors.As(err, &e) {
		if e.Code >= fiber.StatusInternalServerError {
			conf.Err(e.Code, c.IP(), c.OriginalURL(), e.Message)
		}
		return c.Status(e.Code).SendString(fmt.Sprintf("%d %s", e.Code, e.Message))
	}

	conf.ErrWithStack("ServerError", c.IP(), c.OriginalURL(), err)
	return c.Status(fiber.StatusInternalServerError).SendString("X_X SERVER ERROR")
}
