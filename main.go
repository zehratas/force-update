package main

import (
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	app := fiber.New()

	handler := Handler{redis: rdb}
	app.Get("/versions", handler.GetVersion)
	app.Post("/versions", handler.SetVersions)

	_ = app.Listen(":3000")
}
