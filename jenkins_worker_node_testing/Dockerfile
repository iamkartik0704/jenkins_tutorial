FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY testfile.go .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o app testfile.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]