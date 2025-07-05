#!/bin/sh
# Run service benchmarks with CPU and memory profiling
set -e
cd "$(git rev-parse --show-toplevel)"
PROFILE_DIR=${PROFILE_DIR:-bench_profiles}
mkdir -p "$PROFILE_DIR"
go test ./internal/service -bench=Mass -benchmem -cpuprofile "$PROFILE_DIR/cpu.out" -memprofile "$PROFILE_DIR/mem.out"
cat <<MSG
Profiles written to $PROFILE_DIR/cpu.out and $PROFILE_DIR/mem.out
Use 'go tool pprof' to analyze the results, e.g.:
  go tool pprof $PROFILE_DIR/cpu.out
MSG
