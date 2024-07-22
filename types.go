package main

import "sync"

type apiConfig struct {
	fileserverHits int
	mu             sync.Mutex
}

type ChirpRequest struct {
	Body string `json:"body"`
}
