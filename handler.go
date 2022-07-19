package main

import (
	"github.com/gofiber/fiber/v2"
	"regexp"
)

var semver = regexp.MustCompile(`^([1-9]\d?)+\.\d+\.\d+$`)

type Handler struct {
	db Database
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
	iosVersion, err := h.db.Get(c.UserContext(), "ios")
	if err != nil {
		return err
	}
	andVersion, err := h.db.Get(c.UserContext(), "and")
	if err != nil {
		return err
	}
	huaVersion, err := h.db.Get(c.UserContext(), "hua")
	if err != nil {
		return err
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
	err = h.db.Set(c.UserContext(), "ios", versions.Ios)
	if err != nil {
		return err
	}
	err = h.db.Set(c.UserContext(), "and", versions.Android)
	if err != nil {
		return err
	}
	err = h.db.Set(c.UserContext(), "hua", versions.Huawei)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(versions)
}
