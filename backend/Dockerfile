# Backend Dockerfile
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY app/go.mod app/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY app/ .

# Build the application
RUN go build -o main .
RUN go build -o migrate ./database/goose

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./database/goose/migrate
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/database ./database
COPY entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

EXPOSE 7070

# Run the executable
CMD ["./entrypoint.sh"]
