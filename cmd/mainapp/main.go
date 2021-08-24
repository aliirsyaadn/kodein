package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/db"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/internal/router"
)

const mainTag = "Main"

func main(){
	godotenv.Load()

	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.ErrorDetail(mainTag, "error load config: %v", err)
		return
	}

	// Fiber App
	app := fiber.New()

	// Connect Database
	conn := db.ConnectDB(cfg.DB)
	defer conn.Close()

	// Initiate Middleware
	app.Use(logger.New())

	// SetUp Router
	router.SetUpRouter(app)

	log.InfoDetail(mainTag, "app started at :%s", cfg.APP.Port)
	log.FatalDetail(mainTag, "Aborting...", app.Listen(":"+cfg.APP.Port))
}