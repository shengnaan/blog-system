FROM golang:1.22-alpine AS builder
LABEL authors="Pelyovin Pyotr"

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags "-s -w -extldflags '-static'" -o /app /src/cmd/api

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app /app
EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/app"]
