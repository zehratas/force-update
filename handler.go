package main

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	db Database
}

type Versions struct {
	Ios     string `json:"ios"`
	Android string `json:"and"`
	Huawei  string `json:"hua"`
}

type Service struct {
	KullanBirak Versions `json:"kullanbirak"`
	Yolcu360    Versions `json:"yolcu360"`
}

func (h *Handler) GetVersion(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Service{
		KullanBirak: Versions{
			Ios:     h.db.Get(KBios),
			Android: h.db.Get(KBand),
			Huawei:  h.db.Get(KBhua),
		},
		Yolcu360: Versions{
			Ios:     h.db.Get(YLCios),
			Android: h.db.Get(YLCand),
			Huawei:  h.db.Get(YLChua),
		},
	})
}
