#!/bin/bash

# Configurable variables
APP_CONTAINER_NAME="lms-app"
DB_CONTAINER_NAME="lms-postgres"
IMAGE_NAME="lms-api:latest"
DB_PORT="${DB_PORT:-5432}"
PORT="${PORT:-7755}"

# Colors for output
GREEN="\033[0;32m"
HED="\033[1;${GREEN}"
RED="\033[0;31m"
BROWN="\033[38;5;94m"
NC="\033[0m" # No color

# Function to run the app locally
dev() {
  echo -e "${GREEN}Running application locally...${NC}"
  go run main.go
}

# Function to build the Go binary
build_app() {
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
build_app_image() {
  echo ""
  echo -e "${GREEN}Building Docker image...${NC}"
  docker build -t "${IMAGE_NAME}" .
}

# Function to run the LMS API container
run_app() {
  echo -e "${GREEN}Starting ${APP_CONTAINER_NAME} container...${NC}"
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
    --network lms-network \
    "${IMAGE_NAME}"
}

# Function to stop and remove the LMS API container
stop_app() {
  if [ "$(docker ps -aq -f name=$APP_CONTAINER_NAME)" ]; then
    echo -e "${GREEN}Stopping ${APP_CONTAINER_NAME}...${NC}"
    docker stop "${APP_CONTAINER_NAME}"
    echo -e "${GREEN}removing container ${APP_CONTAINER_NAME}...${NC}"
    docker rm -f "${APP_CONTAINER_NAME}"
  else
    echo -e "${RED}No container named ${APP_CONTAINER_NAME} found.${NC}"
  fi
}

# Function to start the PostgreSQL container
start_db() {
  if docker inspect "${DB_CONTAINER_NAME}" > /dev/null 2>&1; then
    echo -e "${GREEN}The container $DB_CONTAINER_NAME exists.${NC}"

    if $(docker inspect -f '{{.State.Status}}' "${DB_CONTAIRE_NAME}" | grep -q "exited"); then
      echo -e "${GREEN}$DB_CONTAINER_NAME container is not running.${NC}"
      echo -e "${GREEN}Restarting $DB_CONTAINER_NAME container.${NC}"
      docker start lms-postgres
    else
      echo -e "${GREEN}$DB_CONTAINER_NAME container is already running.${NC}"
    fi
  else
    echo -e "${GREEN}Starting ${DB_CONTAINER_NAME} container...${NC}"
    docker run --name "${DB_CONTAINER_NAME}" \
      --env-file .env \
      -e POSTGRES_USER=$(grep DB_USER .env | cut -d '=' -f2) \
      -e POSTGRES_PASSWORD="${DB_PASSWORD}" \
      -e POSTGRES_DB=$(grep DB_NAME .env | cut -d '=' -f2) \
      -p "$(grep DB_PORT .env | cut -d '=' -f2):5432" \
      --network lms-network \
      -d postgres:17-alpine
  fi
}

# Function to stop and remove the PostgreSQL container
stop_db() {
  if [ "$(docker ps -aq -f name=$DB_CONTAINER_NAME)" ]; then
    echo -e "${GREEN}Stopping ${DB_CONTAINER_NAME}...${NC}"
    docker stop "${DB_CONTAINER_NAME}"
    echo -e "${GREEN}Removing container ${DB_CONTAINER_NAME}...${NC}"
    docker rm -f "${DB_CONTAINER_NAME}"
  else
    echo -e "${RED}No container named ${DB_CONTAINER_NAME} found.${NC}"
  fi
}

# Help menu
help_menu() {
  echo -e "${GREEN}Usage: ./run.sh [command]${NC}"
  echo -e "Commands:"
  echo -e "${HED}dev          ${BROWN}Run the application locally${NC}"
  echo -e "${HED}build-app    ${BROWN}Build the Go binary for the app${NC}"
  echo -e "${HED}test-app     ${BROWN}Run tests${NC}"
  echo -e "${HED}build-image  ${BROWN}Build the Docker image${NC}"
  echo -e "${HED}start-app    ${BROWN}Start the LMS API container${NC}"
  echo -e "${HED}stop-app     ${BROWN}Stop and remove the LMS API container${NC}"
  echo -e "${HED}start-db     ${BROWN}Start the PostgreSQL container${NC}"
  echo -e "${HED}stop-db      ${BROWN}Stop and remove the PostgreSQL container${NC}"
}

# Main script logic
case "$1" in
  dev) dev ;;
  build-app) build_app ;;
  test) test ;;
  build-image) build_app_image ;;
  start-app) run_app ;;
  stop-app) stop_app ;;
  start-db) start_db ;;
  stop-db) stop_db ;;
  *) help_menu ;;
esac