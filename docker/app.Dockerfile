FROM golang:1.21-alpine

# Install make
RUN apk add --no-cache make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Build the binary
RUN make build

ENTRYPOINT [ "./bin/wisee" ]
