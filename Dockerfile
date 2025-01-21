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

# Create user and group
RUN groupadd --system --gid 990 lmsapp && \
    useradd --system --uid 990 --gid lmsapp lmsapp

# Final release stage
FROM gcr.io/distroless/base-debian12 AS release-stage

WORKDIR /app/

# Copy application binary
COPY --from=build-stage /app/bin/lms-api /app/lms-api

# Copy user and group
COPY --from=intermediate-stage /etc/passwd /etc/group /etc/

ENV GIN_MODE=release

USER lmsapp:lmsapp

EXPOSE 7755

ENTRYPOINT [ "/app/lms-api" ]
