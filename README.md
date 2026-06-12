# Goodreads Collage Generator 📚

A simple web app built in **Go** that takes your Goodreads CSV data export and turns it into a clean, modern graphic image showing your reading stats and currently-reading list.

## 🚀 How to Run It

You can run this application locally either by using Go directly or via Docker.

### Method 1: Using Docker (Recommended & Easiest)
You don't need Go installed on your machine. Just run the pre-built Docker image directly from Docker Hub:

```bash
docker run -d -p 8080:8080 --name reading-collage vikonad/book-collage:latest
```

### Method 2: Running with Go locally
1. Clone the repository:
```bash
git clone [https://github.com/vikonad/book-collage.git](https://github.com/vikonad/book-collage.git)
cd book-collage
```
2. Start the local server:
```bash
go run .
```
After running either method, open http://localhost:8080 in your browser, upload your Goodreads CSV, and download your image!

## 🛠️ Tech Stack

- Go (using the native net/http package for the server)
- gg (github.com/fogleman/gg) for rendering the image and text
