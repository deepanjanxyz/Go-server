package routes

import "github.com/gofiber/fiber/v2"

func StatusRoute(app *fiber.App) {
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "online",
			"message": "Server OK",
		})
	})
}
