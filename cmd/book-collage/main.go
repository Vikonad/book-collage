package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/vikonad/book-collage/internal/graphics"
	"github.com/vikonad/book-collage/internal/parser"
)

var templates = template.Must(template.ParseFiles("web/templates/index.html"))

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/generate", handleGenerate)

	fmt.Println("Server starting on http://localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("📥 Received an upload request...")

	file, _, err := r.FormFile("goodreads_csv")
	if err != nil {
		log.Println("❌ Error retrieving file:", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Println("📊 Parsing CSV...")
	books, err := parser.ParseCSV(file)
	if err != nil {
		log.Println("❌ CSV Parsing Failed:", err)
		http.Error(w, "Error parsing CSV: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("📚 Books parsed successfully! Total found:", len(books))
	if len(books) > 0 {
		log.Printf("First book found: %+v\n", books[0])
	} else {
		log.Println("⚠️ WARNING: No books were added to the slice. Is the CSV layout correct?")
	}

	log.Println("🎨 Starting image generation...")
	var buf bytes.Buffer
	err = graphics.GenerateCollage(&buf, books)
	if err != nil {
		log.Println("❌ Image Generation Failed:", err)
		http.Error(w, "Error generating image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("🚀 Image generated perfectly! Sending to browser...")
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))

	if _, err := buf.WriteTo(w); err != nil {
		log.Println("❌ Error writing image to response:", err)
	}
}
