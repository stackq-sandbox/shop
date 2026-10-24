package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCartTotal(t *testing.T) {
	rec := httptest.NewRecorder()
	routes().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/cart/total", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "$40.48" {
		t.Fatalf("body = %q, want $40.48", got)
	}
}
