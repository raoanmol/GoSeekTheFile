package main

import (
	"log"
	"os"

	"github.com/blevesearch/bleve/v2"
)

func main() {
	indexPath := "index.bleve"
	directory := "./documents"

	var index bleve.Index
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexPath, mapping)
		if err != nil {
			log.Fatal(err)
		}
		defer index.Close()

		err = IndexFiles(directory, index)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Indexing completed.")
	} else {
		index, err = bleve.Open(indexPath)
		if err != nil {
			log.Fatal(err)
		}
		defer index.Close()
	}

	Server(index)
}