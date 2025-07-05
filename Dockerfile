FROM golang:1.22-bullseye AS build

# Install Node.js and npm
RUN apt-get update && apt-get install -y nodejs npm && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy go work files and modules
COPY go.work go.work.sum ./
COPY cmd/go.mod cmd/go.sum ./cmd/
COPY internal/go.mod internal/go.sum ./internal/
COPY internal/pdf/go.mod internal/pdf/go.sum ./internal/pdf/

# Synchronize workspace modules
RUN go work sync

# Copy the rest of the repository
COPY . .

# Install frontend dependencies and build the UI
RUN npm ci --prefix internal/ui
RUN npm run build --prefix internal/ui

# Install Wails CLI and build the application
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest
RUN wails build -clean

FROM debian:bullseye-slim AS final
WORKDIR /app

# Copy built binary
COPY --from=build /app/build/bin/ /app/

CMD ["./Bari$teuer"]
