package check

import (
	"PROJECTNAME/xlog"
	"github.com/gofiber/fiber/v3"
)

func Query(c fiber.Ctx, keys ...string) error {
	for _, key := range keys {
		if v, ok := c.Queries()[key]; !ok || v == "" {
			return xlog.NE(key, "is nil")
		}
	}

	return nil
}
