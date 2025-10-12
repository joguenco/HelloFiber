package util

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConfigureLogger(app *fiber.App) {
	app.Use(logger.New())

	file, err := os.OpenFile("./server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// file.Close() //is not deferred here to keep the file open for logging
	app.Use(logger.New(logger.Config{
		Output: file,
	}))
}
