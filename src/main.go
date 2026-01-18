package main

import (
	"HelloFiber/src/database"
	"HelloFiber/src/handlers"
	"HelloFiber/src/util"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}
	PORT := os.Getenv("PORT")

	database.ConnectPostgres()

	app := fiber.New()

	util.ConfigureLogger(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Application": "Hello Fiber!"})
	})

	app.Get("/ping", handlers.GetPing)
	app.Get("/version", handlers.GetVersion)

	log.Printf("Server is running on port %s\n", PORT)
	app.Listen(":" + PORT)
}
