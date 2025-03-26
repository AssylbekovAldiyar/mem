package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthCheck проверяет работу сервиса
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// GenerateMeme – заглушка для генерации мема
func GenerateMeme(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Meme generated!",
		"url":     "https://example.com/meme.jpg",
	}
	json.NewEncoder(w).Encode(response)
}
