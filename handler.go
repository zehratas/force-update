package main

import (
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"regexp"
)

var semver = regexp.MustCompile(`^([1-9]\d?)+\.\d+\.\d+$`)

type Handler struct {
	redis *redis.Client
}

type Versions struct {
	Ios     string `json:"ios"`
	Android string `json:"and"`
	Huawei  string `json:"hua"`
}

type ApiError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func (h *Handler) GetVersion(c *fiber.Ctx) error {
	iosVersion, err := h.redis.Get(c.UserContext(), "ios").Result()
	if err != nil {
		panic(err)
	}
	andVersion, err := h.redis.Get(c.UserContext(), "and").Result()
	if err != nil {
		panic(err)
	}
	huaVersion, err := h.redis.Get(c.UserContext(), "hua").Result()
	if err != nil {
		panic(err)
	}

	versions := Versions{
		Ios:     iosVersion,
		Android: andVersion,
		Huawei:  huaVersion,
	}
	return c.Status(fiber.StatusOK).JSON(versions)
}

func (h *Handler) SetVersions(c *fiber.Ctx) error {
	versions := Versions{}
	err := c.BodyParser(&versions)
	if err != nil {
		return err
	}
	if !semver.MatchString(versions.Ios) {
		return c.Status(fiber.StatusBadRequest).JSON(ApiError{Code: 201, Error: "Ios version number is invalid"})
	}
	if !semver.MatchString(versions.Android) {
		return c.Status(fiber.StatusBadRequest).JSON(ApiError{Code: 202, Error: "Android version number is invalid"})
	}
	if !semver.MatchString(versions.Huawei) {
		return c.Status(fiber.StatusBadRequest).JSON(ApiError{Code: 203, Error: "Huawei version number is invalid"})
	}
	err = h.redis.Set(c.UserContext(), "ios", versions.Ios, 0).Err()
	if err != nil {
		panic(err)
	}
	err = h.redis.Set(c.UserContext(), "and", versions.Android, 0).Err()
	if err != nil {
		panic(err)
	}
	err = h.redis.Set(c.UserContext(), "hua", versions.Huawei, 0).Err()
	if err != nil {
		panic(err)
	}
	return c.Status(fiber.StatusOK).JSON(versions)
}
