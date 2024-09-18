FROM golang:1.22.5
WORKDIR /short_url
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /app
EXPOSE 8080
CMD ["/app"]