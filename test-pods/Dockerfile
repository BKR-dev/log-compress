# BASE
FROM golang:1.21.0
# Application directory
WORKDIR /app
# dependencies
COPY go.mod go.sum ./
# resolve deps
RUN go mod download
# get sources
COPY *.go ./
# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /log-spam