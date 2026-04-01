package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	r := setupRouter()

	body := []byte(`{"values":[1,2,3]}`)
	req := httptest.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("expected 200, got %d", w.Code)
	}
}