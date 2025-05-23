#!/bin/bash

echo "📦 Loading environment..."
export $(grep -v '^#' .env | xargs)

echo "🛠️ Starting Operary backend..."
go run main.go
