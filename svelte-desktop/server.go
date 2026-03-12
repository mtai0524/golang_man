package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/dist/*
var frontendAssets embed.FS

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func main() {
	// Root of the embedded filesystem
	content, err := fs.Sub(frontendAssets, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files
	fs := http.FileServer(http.FS(content))
	http.Handle("/", fs)

	// API Endpoints
	http.HandleFunc("/api/sum", handleSum)
	http.HandleFunc("/api/greet", handleGreet)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	app := NewApp()
	result := app.Sum(req.A, req.B)

	json.NewEncoder(w).Encode(SumResponse{Result: result})
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	app := NewApp()
	greeting := app.Greet(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"greeting": greeting})
}
