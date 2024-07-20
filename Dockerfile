FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o /bin

FROM alpine:latest
WORKDIR $HOME
COPY --from=build /bin .
EXPOSE 7755
CMD ["./lms-api"]
