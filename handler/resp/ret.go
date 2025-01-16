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
	c.Response().Header.SetContentType(fiber.MIMEApplicationJSON)
	return c.JSON(Pattern{
		Code: 1,
		Msg:  "成功",
		Data: data,
	})
}

func Err(c fiber.Ctx, msg ...any) error {
	c.Response().Header.SetContentType(fiber.MIMEApplicationJSON)
	return c.JSON(Pattern{
		Code: -1,
		Msg:  strings.TrimRight(fmt.Sprintln(msg...), "\n"),
	})
}
