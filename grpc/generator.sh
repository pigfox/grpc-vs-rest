#!/bin/bash
set -x
set -e

#!/bin/bash
set -e

# Create output directory
mkdir -p user

# Generate Go code directly into ./user (not nested)
# Ensure you have the necessary tools installed
protoc --go_out=. --go-grpc_out=. user.proto



