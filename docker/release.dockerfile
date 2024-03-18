# Start from the latest golang base image
FROM golang:latest

# Install make
RUN apt-get update && apt-get install -y make

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.* ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Set the environment variable
# this is required for prod deployment
ENV ENV="production"
# this is require for gin mode to run in release mode
ENV GIN_MODE="release"

ENV JWT_SECRET="secret"
ENV JWT_VALIDITY_IN_HOURS=24
ENV JWT_ISSUER="wisee-backend"

ENV DOMAIN="localhost"
ENV AUTH_REDIRECT_URL="http://localhost:3000/dashboard"

ENV DB_URL="postgresql://postgres:postgres@host.docker.internal:5432/wisee_core?sslmode=disable"
ENV TEST_DB_URL="postgresql://postgres:postgres@host.docker.internal:5432/wisee_core_test?sslmode=disable"
ENV DB_MAX_OPEN_CONNECTIONS=10

ENV GOOGLE_CLIENT_ID="google-client-id"
ENV GOOGLE_CLIENT_SECRET="google-client-secret"
ENV GOOGLE_REDIRECT_URL="http://localhost:8080/v1/auth/google/callback"


# Build the Go app
RUN make build

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./bin/wisee"]
