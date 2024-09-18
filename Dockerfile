FROM golang:22.5
WORKDIR /short_url
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 8080
CMD ["/docker-short-url"]