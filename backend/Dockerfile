FROM golang:alpine as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o backend

FROM alpine
COPY --from=builder /app/backend /app/backend
CMD ["/app/backend"]