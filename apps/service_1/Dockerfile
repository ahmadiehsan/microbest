# =========================
# Builder (base)
# =====
FROM golang:1.24-alpine AS builder_base
ENV CGO_ENABLED=0
WORKDIR /build
COPY go.mod go.sum ./

# =========================
# Builder (dev)
# =====
FROM builder_base AS builder_dev
RUN go mod download -x && go install tool

# =========================
# Builder (prod)
# =====
FROM builder_base AS builder_prod
RUN go mod download -x
COPY . ./
RUN go build -v -ldflags="-s -w" -o ./http_server ./cmd/http

# =========================
# App (dev)
# =====
FROM golang:1.24-alpine AS dev
COPY --from=builder_dev /go /go
WORKDIR /app
COPY . ./

# =========================
# App (prod)
# =====
FROM alpine:3.21 AS prod
WORKDIR /app
COPY --from=builder_prod /build ./
