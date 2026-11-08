# Stage 1: Build binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download dependencies if any
COPY go.mod ./
RUN go mod download

# Copy source code and build statically linked executable
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o go-portfolio .

# Stage 2: Minimal runtime image
FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/go-portfolio .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 5000

ENV PORT=5000

USER guest

ENTRYPOINT ["./go-portfolio"]

