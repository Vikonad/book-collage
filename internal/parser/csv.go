package parser

import (
	"encoding/csv"
	"io"
)

type Book struct {
	Title    string
	Author   string
	MyRating string
	Status   string
}

func ParseCSV(r io.Reader) ([]Book, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books []Book
	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) > 1 {
			books = append(books, Book{
				Title:    row[1],
				Author:   row[2],
				MyRating: row[7],
				Status:   row[17],
			})
		}
	}

	return books, nil
}
