package router

import (
	"github.com/aliirsyaadn/kodein/handlers"
	"github.com/aliirsyaadn/kodein/model"
	"github.com/aliirsyaadn/kodein/services/member"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App, model *model.Queries){
	api := app.Group("/api")

	// Member
	memberService := member.NewService(model)
	handlers.NewMemberHandler(memberService).Register(api)
}