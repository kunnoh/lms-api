#!/bin/bash

# export .env
# export $(cat .env | xargs)

# Configurable variables
APP_CONTAINER_NAME="lms-app"
DB_CONTAINER_NAME="lms-postgres"
IMAGE_NAME="lms-api:latest"
DB_PORT="${DB_PORT:-5432}"
PORT="${PORT:-7755}"

# Colors for output
GREEN="\033[0;32m"
RED="\033[0;31m"
NC="\033[0m" # No color

# Function to run the app locally
dev() {
  echo -e "${GREEN}Running application locally...${NC}"
  go run main.go
}

# Function to build the Go binary
build() {
  echo -e "${GREEN}Building Go binary...${NC}"
  mkdir -p ./bin
  go build -o ./bin/
}

# Function to run tests
test() {
  echo -e "${GREEN}Running tests...${NC}"
  if go test -v ./tests; then
    echo -e "${GREEN}Tests passed${NC}"
  else
    echo -e "${RED}Tests failed${NC}"
  fi
}

# Function to build the Docker image
build_image() {
  echo -e "${GREEN}Building Docker image...${NC}"
  docker build -t "${IMAGE_NAME}" .
}

# Function to run the LMS API container
run_container() {
  echo -e "${GREEN}Running container...${NC}"
  docker run --name "${APP_CONTAINER_NAME}" \
    -e DB_HOST="${DB_HOST}" \
    -e DB_PORT="${DB_PORT}" \
    -e DB_USER="${DB_USER}" \
    -e DB_PASSWORD="${DB_PASSWORD}" \
    -e DB_NAME="${DB_NAME}" \
    -e PORT="${PORT}" \
    -e TOKEN_EXPIRY="${TOKEN_EXPIRY}" \
    -e REFRESH_TOKEN_EXPIRY="${REFRESH_TOKEN_EXPIRY}" \
    -e TOKEN_MAXAGE="${TOKEN_MAXAGE}" \
    -p "${PORT}:${PORT}" \
    "${IMAGE_NAME}"
}

# Function to start the PostgreSQL container
start_db() {
  # DB_CONTAINER_NAME="lms-postgres"
  # if [ "$(docker ps -a -f name=$DB_CONTAINER_NAME)" ]; then
  #   echo -e "${GREEN}$DB_CONTAINER_NAME container is already running.${NC}"
  #   docker rm "${DB_CONTAINER_NAME}"
  # else
    echo -e "${GREEN}Starting a new PostgreSQL container...${NC}"
    docker run --name "${DB_CONTAINER_NAME}" \
      --env-file .env \
      -p "5432:5432" \
      -d postgres:17-alpine
  # fi

  # if [ ! "$(docker ps -a -q -f name=$DB_CONTAINER_NAME)" ]; then
  #   if [ "$(docker ps -aq -f status=exited -f name=$DB_CONTAINER_NAME)" ]; then
  #       # cleanup
  #       docker rm $DB_CONTAINER_NAME
  #   fi
  #   # run your container
  #   docker run --name "${DB_CONTAINER_NAME}" \
  #     -e POSTGRES_USER="${DB_USER}" \
  #     -e POSTGRES_PASSWORD="${DB_PASSWORD}" \
  #     -e POSTGRES_DB="${DB_NAME}" \
  #     -p "${DB_PORT}:5432" -d postgres:17-alpine
  # fi
}

# Function to stop and remove the LMS API container
stop_container() {
  if ["$(docker ps -aq -f name=$APP_CONTAINER_NAME)"]; then
    echo -e "${GREEN}Stopping and removing container ${APP_CONTAINER_NAME}...${NC}"
    docker rm -f "${APP_CONTAINER_NAME}"
  else
    echo -e "${RED}No container named ${APP_CONTAINER_NAME} found.${NC}"
  fi
}

# Function to stop and remove the PostgreSQL container
stop_db() {
  if [ "$(docker ps -aq -f name=$DB_CONTAINER_NAME)" ]; then
    echo -e "${GREEN}Stopping and removing container ${DB_CONTAINER_NAME}...${NC}"
    docker rm -f "${DB_CONTAINER_NAME}"
  else
    echo -e "${RED}No container named ${DB_CONTAINER_NAME} found.${NC}"
  fi
}

# Help menu
help_menu() {
  echo -e "${GREEN}Usage: ./run.sh [command]${NC}"
  echo "Commands:"
  echo "  dev                  Run the application locally"
  echo "  build                Build the Go binary"
  echo "  test                 Run tests"
  echo "  build-image          Build the Docker image"
  echo "  run-app-container    Run the LMS API container"
  echo "  start-db             Start the PostgreSQL container"
  echo "  stop-app-container   Stop and remove the LMS API container"
  echo "  stop-db              Stop and remove the PostgreSQL container"
}

# Main script logic
case "$1" in
  dev) dev ;;
  build) build ;;
  test) test ;;
  build-image) build_image ;;
  run-container) run_container ;;
  start-db) start_db ;;
  stop-container) stop_container ;;
  stop-db) stop_db ;;
  *) help_menu ;;
esac
