FROM golang:alpine as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o webapp

FROM alpine
COPY --from=builder /app/webapp /app/webapp
EXPOSE 80
CMD ["/app/webapp"]