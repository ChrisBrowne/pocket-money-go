package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"pocketmoney/internal/server"
	"strings"
	"testing"
)

func TestApi(t *testing.T) {
	var pocketMoneyManagerCommandChannel = make(chan server.PocketMoneyCommand)
	go server.PocketMoneyManager(pocketMoneyManagerCommandChannel)

	server := httptest.NewServer(server.AppHandler(pocketMoneyManagerCommandChannel))
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

	if strings.Contains(string(body), "Elizabeth") == false {
		t.Errorf("Expected response body to contain 'elizabeth', got %q", string(body))
	}
}
