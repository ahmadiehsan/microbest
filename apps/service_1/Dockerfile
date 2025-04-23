# =========================
# Base
# =====
FROM python:3.12-slim-bullseye AS base

# App configurations
ENV APP_PATH="/app"
ENV APP_VENV_SETUP_PATH="/venv_setup"

# =========================
# Builder (base)
# =====
FROM base AS builder_base

# Python configurations
ENV PYTHONDONTWRITEBYTECODE=1

# PIP configurations
ENV PIP_NO_CACHE_DIR=1
ENV PIP_DISABLE_PIP_VERSION_CHECK=1

# Machine dependencies
RUN apt-get update && \
    apt-get install curl git -y --no-install-recommends && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# UV installation
RUN curl -LsSf https://astral.sh/uv/0.6.14/install.sh | sh
ENV PATH="/root/.local/bin/:$PATH"

# UV configurations
ENV UV_PYTHON_DOWNLOADS=0
ENV UV_LINK_MODE=copy

# =========================
# Builder (dev)
# =====
FROM builder_base AS builder_dev

# UV configurations
ENV UV_COMPILE_BYTECODE=0

# Python dependencies
WORKDIR $APP_VENV_SETUP_PATH
COPY uv.lock pyproject.toml ./
RUN --mount=type=cache,target=/root/.cache/uv uv sync --frozen

# =========================
# Builder (prod)
# =====
FROM builder_base AS builder_prod

# UV configurations
ENV UV_COMPILE_BYTECODE=1

# Python dependencies
WORKDIR $APP_VENV_SETUP_PATH
COPY uv.lock pyproject.toml ./
RUN --mount=type=cache,target=/root/.cache/uv uv sync --frozen --no-dev

# =========================
# App (base)
# =====
FROM base AS app_base

# Python configurations
ENV PYTHONUNBUFFERED=1
ENV PYTHONPATH=.

# Project directory
WORKDIR $APP_PATH

# Project content
COPY . .

# Virtualenv activation
ENV PATH="$APP_VENV_SETUP_PATH/.venv/bin:$PATH"

# =========================
# App (dev)
# =====
FROM app_base AS dev

# Copy builder setups
COPY --from=builder_dev $APP_VENV_SETUP_PATH $APP_VENV_SETUP_PATH

# =========================
# App (prod)
# =====
FROM app_base AS prod

# Copy builder setups
COPY --from=builder_prod $APP_VENV_SETUP_PATH $APP_VENV_SETUP_PATH
