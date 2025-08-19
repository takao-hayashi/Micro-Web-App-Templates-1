package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var store = map[string]string{} // code -> url

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var body struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.URL == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"invalid body"}`))
			return
		}
		code := randCode(5)
		store[code] = body.URL
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"code":  code,
			"short": "/" + code,
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[1:]
		if code == "" {
			w.Write([]byte("POST /create {\"url\":\"https://example.com\"} -> returns /{code}\n"))
			return
		}
		if url, ok := store[code]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		http.NotFound(w, r)
	})

	log.Println("URL Shortener running -> http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func randCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}