// server.go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blevesearch/bleve/v2"
)

type searchRequest struct {
	Query string `json:"query"`
}

type searchResponse struct {
	Files []string `json:"files"`
}

func Server(index bleve.Index) {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		var req searchRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		searchResult, err := Search(req.Query, index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		files := make([]string, len(searchResult.Hits))
		for i, hit := range searchResult.Hits {
			files[i] = hit.ID
		}

		resp := searchResponse{
			Files: files,
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.Handle("/", http.FileServer(http.Dir("./ui")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}