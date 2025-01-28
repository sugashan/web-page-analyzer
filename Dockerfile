FROM golang:1.23
WORKDIR /app
COPY . .
RUN go build -o web-page-analyzer
CMD ["./web-page-analyzer"]
EXPOSE 8080
