import { Client, StatusOK } from 'k6/net/grpc';
import { check, sleep } from 'k6';

// Initialize the gRPC client
const client = new Client();
client.load(['../grpc'], 'user.proto'); // Load the proto file (adjust path as needed)

export const options = {
  vus: 100, // 100 concurrent virtual users (equivalent to -c 100)
  iterations: 100000, // Total number of requests (equivalent to -n 100000)
};

export default function () {
  // Connect to the gRPC server (insecure mode)
  client.connect('localhost:8888', { plaintext: true });

  // Make the gRPC call
  const response = client.invoke('user.UserService/GetUser', {});

  // Check the response status
  check(response, {
    'status is OK': (r) => r.status === StatusOK,
  });

  // Close the connection (optional, k6 manages connections)
  client.close();
  sleep(0.1); // Small delay to prevent overwhelming the server
}