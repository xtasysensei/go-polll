# Build stage
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/app ./cmd
RUN ls -l /app

# Final stage
FROM golang:1.22-alpine
WORKDIR /app
COPY .env .env
COPY --from=build /app/app .
RUN chmod +x app
EXPOSE 9000
CMD ["./app"]