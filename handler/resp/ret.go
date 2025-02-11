package resp

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"strings"
)

type Pattern struct {
	Code int
	Msg  string
	Data any
}

func Suc(c fiber.Ctx, data any) error {
	return c.JSON(Pattern{
		Code: 1,
		Data: data,
	})
}

func Err(c fiber.Ctx, msg ...any) error {
	return c.JSON(Pattern{
		Code: -1,
		Msg:  strings.TrimRight(fmt.Sprintln(msg...), "\n"),
	})
}

func Raw(c fiber.Ctx, body []byte) error {
	return c.Send(body)
}
