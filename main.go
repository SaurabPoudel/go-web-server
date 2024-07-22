package main

import (
	"net/http"
)

func main() {
	apiConfig := &apiConfig{
		fileserverHits: 0,
	}
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("."))
	mux.Handle("/app/", apiConfig.middlewareMatricsInc(http.StripPrefix("/app", fileserver)))
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("GET /admin/metrics", apiConfig.adminHandleMetrices)
	mux.HandleFunc("GET /api/reset", apiConfig.handleReset)
	mux.Handle("POST /api/validate_chirp", http.HandlerFunc(handleValidateChirp))
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
