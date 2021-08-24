package handlers

import (
	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/db"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/services/member"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const memberHandlers = "Member Handlers"

func GetMembers(c *fiber.Ctx) (err error){
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.ErrorDetail(memberHandlers, "error load config: %v", err)
		return
	}

	// Connect Database
	conn := db.ConnectDB(cfg.DB)
	defer conn.Close()

	// New Services
	service := member.NewService(conn)

	
	members, err := service.GetMembers(c.Context())

	if err != nil {
		log.Error("Hello")
	}

	c.Send([]byte(members[0].Username))

	return
}