#!/bin/bash
set -x
set -e


# Config
REST_URL="http://localhost:8888/user"
GRPC_TARGET="localhost:8888"
GRPC_PROTO="user.proto"
GRPC_METHOD="user.UserService.GetUser"

# Output files
REST_OUT="rest_benchmark.txt"
GRPC_OUT="grpc_benchmark.json"

# Run wrk for REST benchmark
echo "🔁 Benchmarking REST endpoint..."
wrk -t12 -c100 -d10s $REST_URL > $REST_OUT

# Run ghz for gRPC benchmark
echo "🔁 Benchmarking gRPC endpoint..."
ghz --insecure \
  --proto $GRPC_PROTO \
  --call $GRPC_METHOD \
  -c 100 -n 10000 \
  --format=json \
  --output=$GRPC_OUT > /dev/null

# Extract REST metrics
REST_RPS=$(grep "Requests/sec" $REST_OUT | awk '{print $2}')
REST_LATENCY=$(grep "Latency" $REST_OUT | awk '{print $2}')

# Extract gRPC metrics
GRPC_RPS=$(jq '.rps' $GRPC_OUT)
GRPC_LATENCY=$(jq '.average' $GRPC_OUT)

# Display comparison
echo ""
echo "📊 Benchmark Results:"
echo "----------------------------"
printf "%-15s %-15s %-15s\n" "Type" "RPS (req/sec)" "Avg Latency"
printf "%-15s %-15s %-15s\n" "REST" "$REST_RPS" "$REST_LATENCY"
printf "%-15s %-15s %-15s\n" "gRPC" "$GRPC_RPS" "${GRPC_LATENCY}ms"