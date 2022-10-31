package main

import (
	"github.com/fiber/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.SetupDatabase()

	app.Listen(":3000")
}
