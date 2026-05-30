package main

import (
	"fmt"
	"log"

	"github.com/vikonad/book-collage/internal/parser"
)

func main() {
	books, err := parser.ParseCSV("goodreads_library_export.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Printf("Title: %s | Rating: %s\n", book.Title, book.MyRating)
	}
}
