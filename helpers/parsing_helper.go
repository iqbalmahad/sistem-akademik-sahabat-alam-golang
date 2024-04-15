package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseID(c *fiber.Ctx) (uint, error) {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
