# Build
FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY . .
RUN go mod download

RUN go build -o ./web-page-analyzer


# RUN
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder ./build/web-page-analyzer ./web-page-analyzer

COPY templates/ ./templates/
COPY config.json ./config.json

CMD ["/app/web-page-analyzer"]

EXPOSE 8080
