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
	req := httptest.NewRequest(http.MethodPost, "http://localhost/versions", getPayload())
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

func getPayload() io.Reader {
	versions := Versions{Ios: "1.a.3", Android: "1.2.3", Huawei: "101.2.2"}
	data, _ := json.Marshal(versions)
	return bytes.NewReader(data)
}
