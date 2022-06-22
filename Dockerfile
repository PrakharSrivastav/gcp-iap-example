FROM golang:1.18 as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o application

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.16.0
RUN apk add --no-cache ca-certificates sed
# Copy the binary to the production image from the builder stage.
COPY  --from=builder /app/application /application
# Run the web service on container startup.
CMD ["/application"]
