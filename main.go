package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define a route that returns a JSON response
	app.Get("/health", func(c *fiber.Ctx) error {
		response := map[string]string{
			"status": "Ok",
		}
		return c.JSON(response)
	})

	// Start the server on port 3000
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
