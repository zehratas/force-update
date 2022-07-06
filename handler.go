package main

import "github.com/gofiber/fiber/v2"

type Handler struct {
}
type Versions struct {
	Ios     string `json:"ios"`
	Android string `json:"and"`
	Huawei  string `json:"hua"`
}

func (h *Handler) GetVersion(c *fiber.Ctx) error {
	versions := Versions{
		Ios:     "1.0.0",
		Android: "1.0.0",
		Huawei:  "1.0.0",
	}
	return c.Status(fiber.StatusOK).JSON(versions)
}
