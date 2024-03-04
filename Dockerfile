# syntax=docker/dockerfile:1

FROM golang:1.19.4

# Set destination for COPY
WORKDIR /app

# Copy the source code with go.mod & go.sum
COPY ./src/ ./

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 80

# Run
CMD ["/docker-gs-ping"]