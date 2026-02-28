#Ste 1: build and app
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

#Step 2: run the app
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY  .env .
EXPOSE 8087
CMD [ "./main" ]