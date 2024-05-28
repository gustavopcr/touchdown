package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gustavopcr/touchdown/internal"
	"github.com/gustavopcr/touchdown/types"
)

func TestHelloHandler(t *testing.T) {
	internal.StartTouchdown() // inicializar logica que calcula as pontuações
	requestPayload := types.Score{
		Score: "3x15",
	}

	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/verify", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleVerify)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected response: got %v want %v", status, http.StatusOK)
	}

	expectedResponse := types.Combination{
		Combination: 4,
	}
	var response types.Combination
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if response != expectedResponse {
		t.Errorf("unexpected response: got %v want %v", response, expectedResponse)
	}
}
