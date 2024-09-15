package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItems(t *testing.T) {
	req, _ := http.NewRequest("GET", "/items", nil)
	res := httptest.NewRecorder()
	getItems(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", res.Code)
	}
}
func TestCreateItem(t *testing.T) {
	newItem := Item{Name: "Test Item", Price: 100}
	jsonItem, _ := json.Marshal(newItem)
	req, _ := http.NewRequest("POST", "/items/create", bytes.NewBuffer(jsonItem))
	res := httptest.NewRecorder()
	createItem(res, req)
	if res.Code != http.StatusCreated {
		t.Errorf("Expected status Created, got %v", res.Code)
	}
}
