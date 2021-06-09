# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the server binary.
RUN go build -v -o server

# Build the migrate binary.
RUN go build ./bin/migrate -v -o migrate

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binaries to the production image from the builder stage.
COPY --from=builder /app/server /app/server
COPY --from=builder /app/migrate /app/migrate

# Run the migrations 

# First run the migrations and after it completes run the server on container startup.
# Note: It would be better if migration could be run in the build step because it would
# ensure that the it is only run once. Otherwise, there are chances that multiple instances
# of cloud run may run it at the same time & cause failure.
CMD /app/migrate; /app/server