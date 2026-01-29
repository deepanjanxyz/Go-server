package routes

import "github.com/gofiber/fiber/v2"

func StatsRoute(app *fiber.App) {
	app.Get("/stats", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"blocked_ads":      0,
			"blocked_trackers": 0,
			"data_saved_mb":    0,
		})
	})
}
