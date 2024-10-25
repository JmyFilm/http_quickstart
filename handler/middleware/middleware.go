package middleware

import (
	"edit-your-project-name/handler/resp"
	"edit-your-project-name/slog"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	if e := new(fiber.Error); errors.As(err, &e) {
		if e.Code >= fiber.StatusInternalServerError {
			slog.Err("HTTP Response Status", e.Code, ctx.IP(), ctx.OriginalURL(), e.Message)
		}
		return resp.Err(ctx, e.Code, e.Message)
	}

	slog.ErrWithStack("ServerError", ctx.IP(), ctx.OriginalURL(), err)
	return resp.Err(ctx, "X_X SERVER ERROR")
}

func ToHTTPS(ctx *fiber.Ctx) error {
	if ctx.Protocol() == "http" {
		return ctx.Redirect("https://"+ctx.Hostname()+ctx.OriginalURL(), fiber.StatusMovedPermanently)
	}
	return ctx.Next()
}

func Limit(max, exp int) func(ctx *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: time.Duration(exp) * time.Second,
		LimitReached: func(ctx *fiber.Ctx) error {
			return resp.Err(ctx, fiber.StatusTooManyRequests, "Too Many Requests")
		},
	})
}
