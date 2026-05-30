package graphics

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/vikonad/book-collage/internal/parser"
)

func GenerateCollage(outputPath string, books []parser.Book) error {
	const S = 1200
	dc := gg.NewContext(S, S)

	dc.SetRGB(0.12, 0.13, 0.16)
	dc.Clear()

	dc.SetRGB(1, 1, 1)

	err := dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 24)
	if err != nil {
		print("Warning: could not load font file, using fallback\n")
	}

	yPosition := 100.0
	for i, book := range books {
		fmt.Println(book)
		if i >= 5 {
			break
		}
		dc.DrawString(book.Title+" - "+book.Author, 100, yPosition)
		yPosition += 60
	}

	return dc.SavePNG(outputPath)
}
