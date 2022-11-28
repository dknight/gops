#!/bin/bash

# Compiles gops for different architectures.

BINDIR="./bin"

mkdir -pv "$BINDIR"

echo "Compiling..."

GOOS=linux GOARCH=amd64 go build -o "$BINDIR/gops-amd64-linux"
GOOS=darwin GOARCH=amd64 go build -o "$BINDIR/gops-amd64-darwin"
GOOS=darwin GOARCH=arm64 go build -o "$BINDIR/gops-arm64-darwin"
GOOS=windows GOARCH=amd64 go build -o "$BINDIR/gops.exe"
