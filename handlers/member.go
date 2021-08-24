package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetMembers(c *fiber.Ctx) (err error){
	c.Send([]byte("Members"))

	return
}