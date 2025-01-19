# Build stage
FROM golang:1.24-rc-alpine AS build-stage

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/lms-api

# Run tests
FROM build-stage AS test-stage
RUN go test -v ./...

# Intermediate stage with a shell to create user and group
FROM debian:stable-slim AS intermediate-stage

# Install curl
RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*

# Create user and group
RUN groupadd --system --gid 990 lmsapp && \
    useradd --system --uid 990 --gid lmsapp lmsapp

# Final release stage
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /app/

# Copy application binary
COPY --from=build-stage /app/bin/lms-api /app/lms-api

# Copy dependencies
COPY --from=intermediate-stage \
    /usr/lib/x86_64-linux-gnu/libz.so* \
    /lib/x86_64-linux-gnu/ld-linux-x86-64.so* \
    /lib/x86_64-linux-gnu/

# Copy user and group
COPY --from=intermediate-stage /etc/passwd /etc/group /etc/

ENV GIN_MODE=release

USER lmsapp:lmsapp

EXPOSE 7755

# Healthcheck to verify the application is running
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD ["/usr/bin/curl", "-f", "http://localhost:7755/health"] || exit 1

ENTRYPOINT [ "/app/lms-api" ]
