import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 100, // 100 concurrent virtual users (equivalent to -c 100)
  iterations: 100000, // Total number of requests (equivalent to -n 100000)
  // Alternatively, use duration for time-based tests
};

export default function () {
  http.get('http://localhost:8888/user'); // HTTP GET request to the target
  sleep(0.1); // Optional: Small delay to prevent overwhelming the system
}