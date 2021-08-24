package router

import (
	"github.com/aliirsyaadn/kodein/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App){
	member := app.Group("/member")

	member.Get("/", handlers.GetMembers)
}