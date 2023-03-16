package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Message(s bool, m string) map[string]interface{} {
	return map[string]interface{}{"status": s, "message": m}
}

func Response(c *fiber.Ctx, data map[string]interface{}) {
	c.Accepts("application/json")
	json.NewEncoder(c).Encode(data)
}
