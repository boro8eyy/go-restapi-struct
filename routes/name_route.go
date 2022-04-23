package routes

import (
	"github.com/boro8eyy/go-restapi-struct/controllers"
	"github.com/gofiber/fiber/v2"
)

func NameRoute(app *fiber.App){
	app.Get("api/names", controllers.GetNames)
}