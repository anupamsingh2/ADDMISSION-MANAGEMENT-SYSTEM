# ---- Build stage ----
FROM golang:1.24 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o admission-system

# ---- Runtime stage ----
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/admission-system .
COPY .env .

CMD ["./admission-system"]