#!/usr/bin/env bash
set -euo pipefail

# Determine project root (directory containing this script's parent)
ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT_DIR"

VERSION="${1:-$(git describe --tags --always)}"
OUTPUT_DIR="$ROOT_DIR/build/bin/$VERSION"

# Clean output directory
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

if [[ -z "${PLATFORMS:-}" ]]; then
    PLATFORMS=("darwin/universal" "windows/amd64" "linux/amd64" "linux/arm64")
else
    read -ra PLATFORMS <<< "$PLATFORMS"
fi

for PLATFORM in "${PLATFORMS[@]}"; do
    echo "==> Building for $PLATFORM"
    wails build -clean -platform "$PLATFORM"

    case "$PLATFORM" in
        darwin/*)  TARGET="macos";;
        windows/*) TARGET="windows";;
        linux/*)   TARGET="linux";;
        *)         TARGET="$PLATFORM";;
    esac

    mkdir -p "$OUTPUT_DIR/$TARGET"
    # Move all files produced by this build into the versioned directory
    mv build/bin/* "$OUTPUT_DIR/$TARGET/"
    # Prepare for next build
    rm -rf build/bin
    mkdir -p build/bin
done

echo "Artifacts stored in $OUTPUT_DIR"
