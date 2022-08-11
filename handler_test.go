package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVersionsSuccessfully(t *testing.T) {
	db := mockDatabase{storage: map[string]string{
		KBios:  "1.2.3",
		KBand:  "1.2.3",
		KBhua:  "1.2.3",
		YLCios: "1.2.3",
		YLCand: "1.2.3",
		YLChua: "1.2.3",
	}}
	handler := Handler{db: &db}
	app := fiber.New()
	app.Get("/versions", handler.GetVersion)
	req := httptest.NewRequest(http.MethodGet, "http://localhost/versions", nil)
	req.Header.Set("content-type", "application/json")
	response, _ := app.Test(req, -1)
	if response.StatusCode != fiber.StatusOK {
		t.Errorf("Response code expected 200 but got %v", response.Status)
	}
	payload := Service{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Expected body to be valid : %v", err)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.Errorf("Expected payload to be valid: %v", err)
	}
	if payload.KullanBirak.Ios != "1.2.3" {
		t.Errorf("Expected Kullan Bırak Ios version '1.2.3' but got %v", payload.KullanBirak.Ios)
	}
	if payload.KullanBirak.Android != "1.2.3" {
		t.Errorf("Expected Kullan Bırak Android version '1.2.3' but got %v", payload.KullanBirak.Android)
	}
	if payload.KullanBirak.Huawei != "1.2.3" {
		t.Errorf("Expected Kullan Bırak Huawei version '1.2.3' but got %v", payload.KullanBirak.Huawei)
	}
	if payload.Yolcu360.Ios != "1.2.3" {
		t.Errorf("Expected Yolcu360 Ios version '1.2.3' but got %v", payload.Yolcu360.Ios)
	}
	if payload.Yolcu360.Android != "1.2.3" {
		t.Errorf("Expected Yolcu360 Android version '1.2.3' but got %v", payload.Yolcu360.Android)
	}
	if payload.Yolcu360.Huawei != "1.2.3" {
		t.Errorf("Expected Yolcu360 Huawei version '1.2.3' but got %v", payload.Yolcu360.Huawei)
	}
}
