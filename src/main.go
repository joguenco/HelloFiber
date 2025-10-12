package main

import (
	"HelloFiber/src/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	util.ConfigureLogger(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber!")
	})

	app.Listen(":3000")
}
