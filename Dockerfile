# Dockerfile

FROM golang:1.22.2
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-server
EXPOSE 80
CMD ["/go-server"]