package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go-server/internal/routes"
)

func main() {
	app := fiber.New()

	routes.StatusRoute(app)
	routes.FiltersRoute(app)
	routes.StatsRoute(app)

	log.Println("🚀 Server running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
