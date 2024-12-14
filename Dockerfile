# Build stage
FROM golang:1.23-alpine AS build-stage

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

RUN groupadd --system --gid 1001 lmsapp && \
    useradd --system --uid 1001 --gid lmsapp lmsapp

# Final release stage
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /app/

COPY --from=build-stage /app/bin/lms-api /app/lms-api
COPY --from=intermediate-stage /etc/passwd /etc/passwd
COPY --from=intermediate-stage /etc/group /etc/group

ENV GIN_MODE=release

USER lmsapp:lmsapp

EXPOSE 7755

ENTRYPOINT [ "/app/lms-api" ]
