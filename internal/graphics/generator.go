package graphics

import (
	"io"

	"github.com/fogleman/gg"
	"github.com/vikonad/book-collage/internal/parser"
)

func GenerateCollage(w io.Writer, books []parser.Book) error {
	const S = 1200
	dc := gg.NewContext(S, S)

	dc.SetRGB(0.12, 0.13, 0.16)
	dc.Clear()

	dc.SetRGB(1, 1, 1)
	_ = dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 24)

	yPosition := 100.0
	for i, book := range books {
		if i >= 10 {
			break
		}
		dc.DrawString(book.Title+" - "+book.Author, 100, yPosition)
		yPosition += 60
	}

	return dc.EncodePNG(w)
}
