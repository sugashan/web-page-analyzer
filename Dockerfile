# Build
FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY . .
RUN go mod download /.

RUN go build -o ./web-page-analyzer


# RUN
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder ./web-page-analyzer

CMD ["/app/web-page-analyzer"]

EXPOSE 8080
