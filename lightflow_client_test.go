package lightflow

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoRequest(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer testtoken" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer ts.Close()

	client := NewAPIClient(ts.URL, "testtoken")

	// Test successful request
	resp, err := client.DoRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %v", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != `{"message": "success"}` {
		t.Fatalf("expected body to be `{\"message\": \"success\"}`, got %v", string(body))
	}

	// Test unauthorized request
	client.AuthToken = "wrongtoken"
	resp, err = client.DoRequest("GET", "/test", nil)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
	if resp != nil && resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %v", resp.StatusCode)
	}
}
