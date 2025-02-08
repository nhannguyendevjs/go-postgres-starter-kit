#!/bin/bash

# Exit on error
set -e

# Load environment variables (if using .env)
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

# Run the application
echo "Starting the application..."
exec go run cmd/app/main.go
