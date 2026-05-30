package parser

import (
	"encoding/csv"
	"os"
)

type Book struct {
	Title    string
	Author   string
	MyRating string
}

func ParseCSV(filePath string) ([]Book, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books []Book

	for i, row := range records {
		if i == 0 {
			continue
		}

		if len(row) >= 3 {
			book := Book{
				Title:    row[1],
				Author:   row[2],
				MyRating: row[7],
			}

			books = append(books, book)
		}
	}

	return books, nil
}
