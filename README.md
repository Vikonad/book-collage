# Goodreads Collage Generator 📚

A simple web app built in **Go** that takes your Goodreads CSV data export and turns it into a clean, modern graphic image showing your reading stats and currently-reading list.

## 🚀 How to Run It

1. Clone the repository:
```bash
   git clone [https://github.com/vikonad/book-collage.git](https://github.com/vikonad/book-collage.git)
   cd book-collage
```
2. Start the local server:
```bash
   go run cmd/book-server/main.go
```
3. Open http://localhost:8080 in your browser, upload your Goodreads CSV, and download your image!


## 🛠️ Tech Stack

- Go (using the native net/http package for the server)
- gg (github.com/fogleman/gg) for rendering the image and text
