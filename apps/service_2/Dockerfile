# =========================
# Base
# =====
FROM python:3.12-alpine AS base
ENV APP_PATH="/app" \
    APP_VENV_SETUP_PATH="/venv_setup"

# =========================
# Builder (base)
# =====
FROM base AS builder_base
RUN apk add --no-cache curl

ENV PYTHONDONTWRITEBYTECODE=1 \
    PIP_NO_CACHE_DIR=1 \
    PIP_DISABLE_PIP_VERSION_CHECK=1

RUN curl -LsSf https://astral.sh/uv/0.8.10/install.sh | sh
ENV PATH="/root/.local/bin/:$PATH" \
    UV_PYTHON_DOWNLOADS=0 \
    UV_LINK_MODE=copy

WORKDIR $APP_VENV_SETUP_PATH
COPY uv.lock pyproject.toml ./

# =========================
# Builder (dev)
# =====
FROM builder_base AS builder_dev
ENV UV_COMPILE_BYTECODE=0
RUN --mount=type=cache,target=/root/.cache/uv \
    uv sync --frozen

# =========================
# Builder (prod)
# =====
FROM builder_base AS builder_prod
ENV UV_COMPILE_BYTECODE=1
RUN --mount=type=cache,target=/root/.cache/uv \
    uv sync --frozen --no-dev

# =========================
# App (base)
# =====
FROM base AS app_base
ENV PYTHONUNBUFFERED=1 \
    PYTHONPATH=. \
    PATH="$APP_VENV_SETUP_PATH/.venv/bin:$PATH"
WORKDIR /app
COPY . ./

# =========================
# App (dev)
# =====
FROM app_base AS dev
COPY --from=builder_dev $APP_VENV_SETUP_PATH $APP_VENV_SETUP_PATH

# =========================
# App (prod)
# =====
FROM app_base AS prod
COPY --from=builder_prod $APP_VENV_SETUP_PATH $APP_VENV_SETUP_PATH
