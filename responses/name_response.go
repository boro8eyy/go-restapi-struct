package responses

import "github.com/gofiber/fiber/v2"

type NameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload    *fiber.Map `json:"payload"`
}