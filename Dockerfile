FROM golang:1.22.6 AS build
WORKDIR /app

# Copy the source code.
COPY . .
# Installs Go dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/notifier
FROM alpine:3.20.2
ENV VERSION=0.1.0
COPY --from=build /app/notifier /usr/local/bin/notifier
RUN chmod +x /usr/local/bin/notifier

ENTRYPOINT ["/usr/local/bin/notifier"]