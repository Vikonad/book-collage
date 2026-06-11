package graphics

import (
	"fmt"
	"io"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/vikonad/book-collage/internal/parser"
)

func truncateText(text string, maxChars int) string {
	runes := []rune(text)

	if len(runes) > maxChars {
		return string(runes[:maxChars]) + "..."
	}
	return text
}

func GenerateCollage(w io.Writer, books []parser.Book) error {
	const S = 1200
	dc := gg.NewContext(S, S)
	dc.SetRGB(0.11, 0.13, 0.17)
	dc.Clear()

	logoImg, err := gg.LoadImage("logo.png")
	if err != nil {
		return fmt.Errorf("CRITICAL ERROR: failed to load logo: %w", err)
	}
	logoImg = resize.Resize(400, 0, logoImg, resize.Lanczos3)
	dc.DrawImage(logoImg, 700, 50)

	if err := dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 48); err != nil {
		return fmt.Errorf("could not load font: %w", err)
	}
	dc.SetRGB(0.31, 0.53, 0.96)
	dc.DrawString("My Literary Journey", 100, 95)

	if err := dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 20); err != nil {
		return fmt.Errorf("could not load font: %w", err)
	}
	dc.SetRGB(0.55, 0.60, 0.68)
	dc.DrawString("Top rated books exported from Goodreads", 100, 130)

	//
	// STATS
	//

	if err := dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 30); err != nil {
		return fmt.Errorf("could not load font: %w", err)
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawString("To-Read", 190, 210)

	if err := dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 30); err != nil {
		return fmt.Errorf("could not load font: %w", err)
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawString("Currently-Reading", 468, 210)

	if err := dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 30); err != nil {
		return fmt.Errorf("could not load font: %w", err)
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawString("Read", 918, 210)

	dc.SetRGB(0.22, 0.26, 0.33)
	dc.SetLineWidth(2)
	dc.DrawLine(100, 170, 1100, 170)

	dc.DrawLine(400, 0, 400, 1200)
	dc.DrawLine(800, 0, 800, 1200)

	dc.Stroke()

	yPosition := 500.0

	for i, book := range books {
		if i >= 12 {
			break
		}

		cleanTitle := truncateText(book.Title, 75)

		cleanAuthor := truncateText(book.Author, 40)

		if err := dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 26); err != nil {
			return err
		}
		dc.SetRGB(1, 1, 1)
		dc.DrawString(cleanTitle, 130, yPosition)

		if err := dc.LoadFontFace("RobotoSlab-VariableFont_wght.ttf", 18); err != nil {
			return err
		}
		dc.SetRGB(0.93, 0.60, 0.23)
		dc.DrawString("by "+cleanAuthor, 150, yPosition+30)

		yPosition += 75
	}

	return dc.EncodePNG(w)
}
