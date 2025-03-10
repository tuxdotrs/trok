#!/bin/sh
# Copyright 2025 t@tux.rs. All rights reserved.

set -e

repo="tuxdotrs/trok"
release_url="https://api.github.com/repos/$repo/releases/latest"

if [ "$OS" = "Windows_NT" ]; then
    echo "Error: this installer only works on linux & macOS." 1>&2
    exit 1
else
    case $(uname -sm) in
    "Darwin x86_64") target="darwin-x86_64" ;;
    "Darwin arm64") target="darwin-aarch64" ;;
    "Linux x86_64") target="linux-x86_64" ;;
    "Linux arm64"|"Linux aarch64") target="linux-aarch64" ;;
    *) target="linux-x86_64" ;;
    esac
fi

echo "Downloading $repo for $target"
release_target_url=$(
    curl -s "$release_url" |
    grep "browser_download_url" |
    grep "$target" |
    sed -re 's/.*: "([^"]+)".*/\1/' \
)

curl -sL "$release_target_url" | tar xz
