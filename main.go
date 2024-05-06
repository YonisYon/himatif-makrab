package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Health check endpoint
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Fiber OK!",
		})
	})

	// Start the Fiber app
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
