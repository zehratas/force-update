package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIosVersionNumberInvalid(t *testing.T) {
	handler := Handler{}
	app := fiber.New()
	app.Post("/versions", handler.SetVersions)
	req := httptest.NewRequest(http.MethodPost, "http://localhost/versions", getInvalidIosPayload())
	req.Header.Set("content-type", "application/json")
	response, _ := app.Test(req, -1)
	if response.StatusCode != fiber.StatusBadRequest {
		t.Errorf("Response code expected 400 but got %v", response.Status)
	}
	payload := ApiError{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Expected body to be valid : %v", err)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.Errorf("Expected payload to be valid: %v", err)
	}
	if payload.Code != 201 {
		t.Errorf("Expected payload code 201 but got %v", payload.Code)
	}
	if payload.Error != "Ios version number is invalid" {
		t.Errorf("Expected payload message 'Ios version number is invalid' but got '%v'", payload.Error)
	}
}

func TestAndVersionNumberInvalid(t *testing.T) {
	handler := Handler{}
	app := fiber.New()
	app.Post("/versions", handler.SetVersions)
	req := httptest.NewRequest(http.MethodPost, "http://localhost/versions", getInvalidAndPayload())
	req.Header.Set("content-type", "application/json")
	response, _ := app.Test(req, -1)
	if response.StatusCode != fiber.StatusBadRequest {
		t.Errorf("Response code expected 400 but got %v", response.Status)
	}
	payload := ApiError{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Expected body to be valid : %v", err)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.Errorf("Expected payload to be valid: %v", err)
	}
	if payload.Code != 202 {
		t.Errorf("Expected payload code 202 but got %v", payload.Code)
	}
	if payload.Error != "Android version number is invalid" {
		t.Errorf("Expected payload message 'Android version number is invalid' but got '%v'", payload.Error)
	}
}

func TestHuaVersionNumberInvalid(t *testing.T) {
	handler := Handler{}
	app := fiber.New()
	app.Post("/versions", handler.SetVersions)
	req := httptest.NewRequest(http.MethodPost, "http://localhost/versions", getInvalidHuaPayload())
	req.Header.Set("content-type", "application/json")
	response, _ := app.Test(req, -1)
	if response.StatusCode != fiber.StatusBadRequest {
		t.Errorf("Response code expected 400 but got %v", response.Status)
	}
	payload := ApiError{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Expected body to be valid : %v", err)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.Errorf("Expected payload to be valid: %v", err)
	}
	if payload.Code != 203 {
		t.Errorf("Expected payload code 203 but got %v", payload.Code)
	}
	if payload.Error != "Huawei version number is invalid" {
		t.Errorf("Expected payload message 'Huawei version number is invalid' but got '%v'", payload.Error)
	}
}

func TestUpdateVersionsSuccessfully(t *testing.T) {
	db := mockDatabase{storage: map[string]string{}}
	handler := Handler{db: &db}
	app := fiber.New()
	app.Post("/versions", handler.SetVersions)
	req := httptest.NewRequest(http.MethodPost, "http://localhost/versions", getValidPayload())
	req.Header.Set("content-type", "application/json")
	response, _ := app.Test(req, -1)
	if response.StatusCode != fiber.StatusOK {
		t.Errorf("Response code expected 200 but got %v", response.Status)
	}
	payload := Versions{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Expected body to be valid : %v", err)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.Errorf("Expected payload to be valid: %v", err)
	}
	if payload.Ios != "1.2.3" {
		t.Errorf("Expected Ios version '1.2.3' but got %v", payload.Ios)
	}
	if payload.Android != "1.2.3" {
		t.Errorf("Expected Android version '1.2.3' but got %v", payload.Android)
	}
	if payload.Huawei != "1.2.2" {
		t.Errorf("Expected Huawei version '1.2.2' but got %v", payload.Huawei)
	}
	if db.storage["ios"] != "1.2.3" {
		t.Errorf("Expected Ios version '1.2.3' but got %v", db.storage["ios"])
	}
	if db.storage["and"] != "1.2.3" {
		t.Errorf("Expected Android version '1.2.3' but got %v", db.storage["and"])
	}
	if db.storage["hua"] != "1.2.2" {
		t.Errorf("Expected Huawei version '1.2.2' but got %v", db.storage["hua"])
	}
}

func TestGetVersionsSuccessfully(t *testing.T) {
	db := mockDatabase{storage: map[string]string{
		"ios": "1.2.3",
		"and": "1.2.3",
		"hua": "1.2.2",
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
	payload := Versions{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Expected body to be valid : %v", err)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.Errorf("Expected payload to be valid: %v", err)
	}
	if payload.Ios != "1.2.3" {
		t.Errorf("Expected Ios version '1.2.3' but got %v", payload.Ios)
	}
	if payload.Android != "1.2.3" {
		t.Errorf("Expected Android version '1.2.3' but got %v", payload.Android)
	}
	if payload.Huawei != "1.2.2" {
		t.Errorf("Expected Huawei version '1.2.2' but got %v", payload.Huawei)
	}
}

func getInvalidIosPayload() io.Reader {
	versions := Versions{Ios: "1.a.3", Android: "1.2.3", Huawei: "109.2.2"}
	data, _ := json.Marshal(versions)
	return bytes.NewReader(data)
}

func getInvalidAndPayload() io.Reader {
	versions := Versions{Ios: "1.2.3", Android: "a.2.3", Huawei: "109.2.2"}
	data, _ := json.Marshal(versions)
	return bytes.NewReader(data)
}

func getInvalidHuaPayload() io.Reader {
	versions := Versions{Ios: "1.2.3", Android: "1.2.3", Huawei: "b.2.2"}
	data, _ := json.Marshal(versions)
	return bytes.NewReader(data)
}

func getValidPayload() io.Reader {
	versions := Versions{Ios: "1.2.3", Android: "1.2.3", Huawei: "1.2.2"}
	data, _ := json.Marshal(versions)
	return bytes.NewReader(data)
}
