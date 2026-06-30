package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := "0.0.0.0:" + port 

	router := NewRouter()

	router.Add("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		HelloTemplate().Render(r.Context(), w)
	})

	router.Add("/cv", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cv, err := loadData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		CVTemplate(cv).Render(r.Context(), w)
	})

	router.AddStatic("/fonts/", fontsHandler())

	log.Println("starting server on " + addr)
	log.Fatalf("error starting server: %v", http.ListenAndServe(addr, router))
}
