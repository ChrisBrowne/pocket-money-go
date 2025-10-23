package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"pocketmoney/internal/server"
	"testing"
)

func TestApi(t *testing.T) {
	server := httptest.NewServer(server.AppHandler())
	defer server.Close()

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expexted status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expected := "Pocket money, golang edition, mooo"
	if string(body) != expected {
		t.Errorf("Expected response body %q, got %q", expected, string(body))
	}

}
