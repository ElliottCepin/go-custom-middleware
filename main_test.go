package main

import (
	"net/http/httptest"
	"bytes"
	"encoding/json"
	"strconv"
	"testing"
	"net/http"
)

func TestCalculate(t *testing.T) {
	calculation := Calculation{8, 9, "multiply"}
	value := 8 * 9
	r, err := json.Marshal(calculation)
	
	if (err != nil) {
		t.Error("Client error")
	}

	reader := bytes.NewReader(r)
	req := httptest.NewRequest("POST", "/calculate", reader)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	serveCalculate(rec, req)

	if (rec.Body.String() != strconv.Itoa(value)) {
		t.Errorf("Got %q, expected %q", rec.Body.String(), strconv.Itoa(value))
	}

}

func TestMiddleware(t *testing.T) {
	
	calculation := Calculation{8, 9, "multiply"}
	r, err := json.Marshal(calculation)
	
	a := strconv.Itoa(uuid)
	if (err != nil) {
		t.Error("Client error")
	}

	reader := bytes.NewReader(r)
	req := httptest.NewRequest("POST", "/calculate", reader)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	updateHeader(http.HandlerFunc(serveCalculate))(rec, req)
	if (rec.Header().Get("X-Request-ID") != a) {
		t.Errorf("Got %q, expected %q", rec.Header().Get("X-Request-ID"), a)
	}
}
