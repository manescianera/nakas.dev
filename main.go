package main

import (
	"net/http"
)

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

	http.ListenAndServe(":8080", router)
}
