package main

import (
	"log"
	"net/http"
)

const addr = "0.0.0.0:8043"

func main() {
	router := NewRouter()

	router.Add("/", func(w http.ResponseWriter, r *http.Request) {
		HelloTemplate().Render(r.Context(), w)
	})

	router.Add("/cv", func(w http.ResponseWriter, r *http.Request) {
		cv, err := loadData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		CVTemplate(cv).Render(r.Context(), w)
	})

	log.Println("starting server on " + addr)
	log.Fatalf("error starting server: %v", http.ListenAndServe(addr, router))
}
