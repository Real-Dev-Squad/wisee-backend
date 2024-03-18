# Start from the latest golang base image
FROM golang:latest

# Declare build arguments
ARG ENV
ARG GIN_MODE
ARG JWT_SECRET
ARG JWT_VALIDITY_IN_HOURS
ARG JWT_ISSUER
ARG DOMAIN
ARG AUTH_REDIRECT_URL
ARG DB_URL
ARG TEST_DB_URL
ARG DB_MAX_OPEN_CONNECTIONS
ARG GOOGLE_CLIENT_ID
ARG GOOGLE_CLIENT_SECRET
ARG GOOGLE_REDIRECT_URL

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

# Set the environment variables
ENV ENV=$ENV
ENV GIN_MODE=$GIN_MODE

ENV JWT_SECRET=$JWT_SECRET
ENV JWT_VALIDITY_IN_HOURS=$JWT_VALIDITY_IN_HOURS
ENV JWT_ISSUER=$JWT_ISSUER

ENV DOMAIN=$DOMAIN
ENV AUTH_REDIRECT_URL=$AUTH_REDIRECT_URL

ENV DB_URL=$DB_URL
ENV TEST_DB_URL=$TEST_DB_URL
ENV DB_MAX_OPEN_CONNECTIONS=$DB_MAX_OPEN_CONNECTIONS

ENV GOOGLE_CLIENT_ID=$GOOGLE_CLIENT_ID
ENV GOOGLE_CLIENT_SECRET=$GOOGLE_CLIENT_SECRET
ENV GOOGLE_REDIRECT_URL=$GOOGLE_REDIRECT_URL

# Build the Go app
RUN make build

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./bin/wisee"]
