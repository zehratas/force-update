package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	handler := Handler{}
	app.Get("/versions", handler.GetVersion)

	_ = app.Listen(":3000")
}
