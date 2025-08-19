package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/pretty", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		var v any
		if err := json.Unmarshal(raw, &v); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad json"))
			return
		}
		b, _ := json.MarshalIndent(v, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})
	log.Println("JSON Pretty -> POST http://localhost:8081/pretty")
	log.Fatal(http.ListenAndServe(":8081", nil))
}