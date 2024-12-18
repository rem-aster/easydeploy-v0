FROM golang:1.22.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - && apt-get install -y nodejs && \
    npm install -g npm && \
    npm install && \
    npx tailwindcss -i styles.css -o public/static/css/tw.css
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/public ./public
EXPOSE 8080
CMD ["./main"]
