package main

import (
	"github.com/boro8eyy/go-restapi-struct/configs"
	"github.com/boro8eyy/go-restapi-struct/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

    routes.NameRoute(app)

    app.Listen(":8080")
}