package handlers

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
)

type Version struct {
	Name            string `json:"message"`
	Author          string `json:"author"`
	Version         string `json:"version"`
	VersionOS       string `json:"versionOS"`
	VersionLanguage string `json:"versionLanguage"`
}

func GetVersion(c *fiber.Ctx) error {
	return c.JSON(
		Version{
			Name:            "Hello Fiber",
			Author:          "Jorge Luis",
			Version:         "1.0.0",
			VersionOS:       runtime.GOOS + " " + runtime.GOARCH,
			VersionLanguage: runtime.Version(),
		})
}
