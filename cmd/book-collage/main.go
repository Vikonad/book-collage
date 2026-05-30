package main

import (
	"log"

	"github.com/vikonad/book-collage/internal/graphics"
	"github.com/vikonad/book-collage/internal/parser"
)

func main() {
	// Inside your main() function:
	books, err := parser.ParseCSV("goodreads_library_export.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Pass the 'books' slice we got from the parser into the graphics package!
	err = graphics.GenerateCollage("my_collage.png", books)
	if err != nil {
		log.Fatal(err)
	}
}
