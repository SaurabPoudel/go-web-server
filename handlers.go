package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (cfg *apiConfig) adminHandleMetrices(w http.ResponseWriter, r *http.Request) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, `
		<html>

<body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
</body>

</html>
	`,
		cfg.fileserverHits)
}

func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request) {
	cfg.mu.Lock()
	cfg.fileserverHits = 0
	cfg.mu.Unlock()
	w.WriteHeader(http.StatusOK)
}

func handleValidateChirp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var chirpReq ChirpRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&chirpReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Something went wrong"})
		return
	}
	if len(chirpReq.Body) > 40 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Chirp is too long"})
		return
	}

	w.WriteHeader(http.StatusOK)
	words := []string{"kerfuffle", "sharbert", "fornax"}
	for _, word := range words {
		chirpReq.Body = strings.Replace(chirpReq.Body, word, "****", 1)
	}
	json.NewEncoder(w).Encode(map[string]string{"valid": "true", "cleaned_body": chirpReq.Body})
}
