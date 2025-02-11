package middleware

import (
	"PROJECTNAME/handler/resp"
	"PROJECTNAME/xlog"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"time"
)

func ErrorHandler(c fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	if e := new(fiber.Error); errors.As(err, &e) {
		if e.Code >= fiber.StatusInternalServerError {
			xlog.Err(xlog.PS("HTTP Response Code", e.Code, c.IP(), c.OriginalURL()), e.Message)
		}
		return resp.Err(c, e.Code, e.Message)
	}

	xlog.Err(xlog.PS("Server Error", c.IP(), c.OriginalURL()), err)
	return resp.Err(c, "X_X SERVER ERROR")
}

func ToHTTPS(c fiber.Ctx) error {
	if c.Protocol() == "http" {
		return c.Redirect().Status(301).To("https://" + c.Hostname() + c.OriginalURL())
	}
	return c.Next()
}

func Limit(max, exp int, key string, needSkip bool) fiber.Handler {
	if needSkip {
		return func(c fiber.Ctx) error {
			return c.Next()
		}
	}

	if max == 0 || exp == 0 {
		return func(c fiber.Ctx) error {
			return c.Next()
		}
	}

	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: time.Duration(exp) * time.Second,
		KeyGenerator: func(c fiber.Ctx) string {
			if key != "" {
				if key == fiber.HeaderXForwardedFor {
					if len(c.IPs()) > 0 {
						return c.IPs()[0]
					}
				} else {
					return key
				}
			}
			return c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			return resp.Err(c, fiber.StatusTooManyRequests, "Too Many Requests")
		},
	})
}
