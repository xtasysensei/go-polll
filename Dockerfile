# Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/app ./cmd

# Final stage
FROM golang:1.22-alpine

WORKDIR /app

COPY --from=build /app/app /app/app

EXPOSE 9000

CMD ["./app"]