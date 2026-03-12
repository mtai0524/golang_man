# Multi-stage Dockerfile for Wails (Svelte)
# 1. Build stage
FROM golang:1.23-bookworm AS build

# Install Node.js and dependencies for Wails
RUN apt-get update && apt-get install -y \
    curl \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    && curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# Install Wails CLI
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

COPY frontend/package*.json ./frontend/
RUN cd frontend && npm install

# Copy source code
COPY . .

# Build the application
RUN wails build -platform linux/amd64

# 2. Runtime stage
FROM debian:bookworm-slim

# Install runtime dependencies for GTK/WebKit, Xvfb, and D-Bus
RUN apt-get update && apt-get install -y \
    libgtk-3-0 \
    libwebkit2gtk-4.0-37 \
    ca-certificates \
    xvfb \
    libnss3 \
    libasound2 \
    libgbm1 \
    dbus \
    dbus-x11 \
    xauth \
    xfonts-base \
    libx11-xcb1 \
    libxcomposite1 \
    libxcursor1 \
    libxdamage1 \
    libxi6 \
    libxtst6 \
    libgl1-mesa-dri \
    libgl1-mesa-glx \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/build/bin/svelte-desktop .

# Start Xvfb, wait for it to be ready, then launch app with a D-Bus session
CMD ["sh", "-c", "\
    Xvfb :99 -screen 0 1024x768x24 -ac +extension GLX +render -noreset & \
    export DISPLAY=:99 && \
    sleep 1 && \
    dbus-launch --exit-with-session ./svelte-desktop \
    "]