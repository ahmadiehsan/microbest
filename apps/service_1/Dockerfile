# =========================
# Builder (base)
# =====
FROM golang:1.24-alpine AS builder_base
RUN apk add --no-cache curl

WORKDIR /build/go_supervisor
RUN curl -L -o supervisord.tar.gz https://github.com/ochinchina/supervisord/releases/download/v0.7.3/supervisord_0.7.3_Linux_64-bit.tar.gz && \
    tar -xzf supervisord.tar.gz --strip-components=1 && \
    chmod +x supervisord

WORKDIR /build/app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./

# =========================
# Builder (dev)
# =====
FROM builder_base AS builder_dev
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download -x && \
    go install tool && \
    cp -a /go/pkg/mod /go/pkg_mod_cache

# =========================
# Builder (prod)
# =====
FROM builder_base AS builder_prod
COPY . ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download -x && \
    go build -v -ldflags="-s -w" -o ./apis ./cmd/apis && \
    go build -v -ldflags="-s -w" -o ./events ./cmd/events

# =========================
# App (dev)
# =====
FROM golang:1.24-alpine AS dev
COPY --from=builder_dev /build/go_supervisor/supervisord /usr/local/bin/supervisord
COPY --from=builder_dev /go/bin /go/bin
COPY --from=builder_dev /go/pkg_mod_cache /go/pkg/mod
WORKDIR /app
COPY . ./

# =========================
# App (prod)
# =====
FROM alpine:3.21 AS prod
COPY --from=builder_prod /build/go_supervisor/supervisord /usr/local/bin/supervisord
WORKDIR /app
COPY --from=builder_prod /build/app ./
