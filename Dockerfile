# Stage 1: Build environment
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency manifests first to utilize Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a highly optimized, statically-linked production binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o book-collage ./cmd/book-collage

# Stage 2: Clean runtime environment
FROM alpine:3.19
RUN apk add --no-cache ca-certificates

WORKDIR /app

# 1. Copy the compiled app from the builder stage
COPY --from=builder /app/book-collage .

# 2. Copy your design assets (fonts, images)
COPY --from=builder /app/assets ./assets

# 3. FIX: Copy the frontend templates folder into the container!
COPY --from=builder /app/web ./web

EXPOSE 8080

CMD ["./book-server"]
