# Build
FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY . .
RUN go mod download

RUN go build -o ./web-page-analyzer


# RUN
FROM debian:12
# FROM gcr.io/distroless/base-debian12

RUN apt-get update && \
    apt-get install -y \
    wget \
    curl \
    gnupg2 \
    ca-certificates \
    fonts-liberation \
    libappindicator3-1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libatk1.0-0 \
    libcups2 \
    libx11-xcb1 \
    libxcomposite1 \
    libxdamage1 \
    libxrandr2 \
    xdg-utils \
    --no-install-recommends && \
    # Add Google Chrome repository and install it
    wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" | tee -a /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && \
    apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder ./build/web-page-analyzer ./web-page-analyzer

COPY templates/ ./templates/
COPY config.json ./config.json

CMD ["/app/web-page-analyzer"]

EXPOSE 8080
