package resp

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Pattern struct {
	Code int
	Msg  string
	Data any
}

func Suc(ctx *fiber.Ctx, data any) error {
	ctx.Response().Header.SetContentType(fiber.MIMEApplicationJSON)
	return ctx.JSON(Pattern{
		Code: 1,
		Msg:  "成功",
		Data: data,
	})
}

func Err(ctx *fiber.Ctx, msg ...any) error {
	ctx.Response().Header.SetContentType(fiber.MIMEApplicationJSON)
	return ctx.JSON(Pattern{
		Code: -1,
		Msg:  strings.TrimRight(fmt.Sprintln(msg...), "\n"),
	})
}
