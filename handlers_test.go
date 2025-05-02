package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lxnewayfarer/shortlinks/handlers"
	"github.com/lxnewayfarer/shortlinks/storage"
)
func TestJSONResponse(t *testing.T) {
    rr := httptest.NewRecorder()
    data := map[string]string{"key": "value"}
    
    handlers.JSONResponse(rr, http.StatusOK, data)
    
    if rr.Code != http.StatusOK {
        t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
    }
    
    if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
        t.Errorf("expected content-type %s, got %s", "application/json", ct)
    }
    
    var response map[string]string
    if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
        t.Fatalf("could not unmarshal response: %v", err)
    }
    
    if response["key"] != "value" {
        t.Errorf("expected response key 'value', got '%s'", response["key"])
    }
}

func TestPing(t *testing.T) {
	rdb, _ := storage.InitMockRedis()

    handler := handlers.Ping(rdb)
    
    req := httptest.NewRequest("GET", "/ping", nil)
    rr := httptest.NewRecorder()
    
    handler.ServeHTTP(rr, req)
    
    if rr.Code != http.StatusOK {
        t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
    }
    
    var response map[string]string
    if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
        t.Fatalf("could not unmarshal response: %v", err)
    }
    
    if response["response"] != "pong" {
        t.Errorf("expected response 'pong', got '%s'", response["response"])
    }
}