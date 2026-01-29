package main

import (
    "os"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":  true,
            "message": "Server is running 🚀",
        })
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "10000"
    }

    app.Listen(":" + port)
}
