#!/bin/sh
# Run service benchmarks with CPU and memory profiling and measure frontend metrics
set -e

ROOT_DIR="$(git rev-parse --show-toplevel)"
cd "$ROOT_DIR"

PROFILE_DIR=${PROFILE_DIR:-$ROOT_DIR/bench_profiles}
mkdir -p "$PROFILE_DIR"

# Backend benchmarks
go test ./internal/service -bench=Mass -benchmem \
    -cpuprofile "$PROFILE_DIR/cpu.out" \
    -memprofile "$PROFILE_DIR/mem.out"

# Frontend build size and load time
cd internal/ui
npm ci --silent
START=$(date +%s)
npm run build >/dev/null
BUILD_TIME=$(( $(date +%s) - START ))
BUILD_SIZE=$(du -sh dist | awk '{print $1}')
npx vite preview --port 5173 >/dev/null 2>&1 &
SERVER_PID=$!
sleep 2
curl -o /dev/null -s -w "%{time_total}\n" http://localhost:5173 > "$PROFILE_DIR/frontend_load_seconds.txt"
kill $SERVER_PID
echo "$BUILD_SIZE" > "$PROFILE_DIR/frontend_build_size.txt"
echo "$BUILD_TIME" > "$PROFILE_DIR/frontend_build_seconds.txt"
cd - >/dev/null

cat <<MSG
Profiles written to $PROFILE_DIR/cpu.out and $PROFILE_DIR/mem.out
Frontend build size written to $PROFILE_DIR/frontend_build_size.txt
Build time written to $PROFILE_DIR/frontend_build_seconds.txt
Page load time written to $PROFILE_DIR/frontend_load_seconds.txt
Use 'go tool pprof' to analyze the Go profiles, e.g.:
  go tool pprof $PROFILE_DIR/cpu.out
MSG
