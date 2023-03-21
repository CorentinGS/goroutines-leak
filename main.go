package main

func main() {

	// Create the app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Listen to port 1815
	app.Listen(":1815")
}
