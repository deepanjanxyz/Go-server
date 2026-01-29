package routes

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func loadFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func FiltersRoute(app *fiber.App) {
	app.Get("/filters/:type", func(c *fiber.Ctx) error {

		filterType := c.Params("type")

		fileMap := map[string]string{
			"global":   "data/filters/global.txt",
			"ads":      "data/filters/ads.txt",
			"tracking": "data/filters/tracking.txt",
			"malware":  "data/filters/malware.txt",
		}

		filePath, ok := fileMap[filterType]

		if !ok {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid filter type",
			})
		}

		list, err := loadFile(filePath)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Unable to load filter file",
			})
		}

		return c.JSON(fiber.Map{
			"type":    filterType,
			"count":   len(list),
			"domains": list,
		})
	})
}
