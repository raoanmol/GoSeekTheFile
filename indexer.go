package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/blevesearch/bleve/v2"
)

func IndexFiles(directory string, index bleve.Index) error {
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			text, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}	
			
			index.Index(path, string(text))
		}

		return nil
	})
}

func Search(query string, index bleve.Index) (*bleve.SearchResult, error) {
	searchRequest := bleve.NewSearchRequest(bleve.NewMatchQuery(query))
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}