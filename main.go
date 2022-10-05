package main

import (
	"fiber-repo-pattern/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.SetupConnection()
	config.Initalize(app)

	app.Listen(":3000")
}
