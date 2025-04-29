# grpc-vs-rest
A grpc-vs-rest comparison

```
go run ./rest/main.go
run test
go run ./grpc/main.go
run test
```
Use a benchmarking tool like hey (for HTTP)

## REST
```
$ hey -n 100000 -c 100 http://localhost:8888/user

Summary:
  Total:	1.7160 secs
  Slowest:	0.0140 secs
  Fastest:	0.0001 secs
  Average:	0.0017 secs
  Requests/sec:	58275.8352
  
  Total data:	6000000 bytes
  Size/request:	60 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [56471]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [26096]	|■■■■■■■■■■■■■■■■■■
  0.004 [10494]	|■■■■■■■
  0.006 [4441]	|■■■
  0.007 [1643]	|■
  0.008 [582]	|
  0.010 [188]	|
  0.011 [59]	|
  0.013 [11]	|
  0.014 [14]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0006 secs
  50% in 0.0012 secs
  75% in 0.0023 secs
  90% in 0.0037 secs
  95% in 0.0047 secs
  99% in 0.0068 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0140 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0020 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0065 secs
  resp wait:	0.0014 secs, 0.0001 secs, 0.0130 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0091 secs

Status code distribution:
  [200]	100000 responses
```
## gRPC
```
./grpc$ ghz --insecure \
  --proto user.proto \
  --call user.UserService.GetUser \
  -c 100 -n 100000 \
  localhost:8888

Summary:
  Count:	100000
  Total:	5.10 s
  Slowest:	21.72 ms
  Fastest:	0.19 ms
  Average:	3.35 ms
  Requests/sec:	19590.53

Response time histogram:
  0.189  [1]     |
  2.343  [28158] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  4.496  [51769] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  6.650  [16535] |∎∎∎∎∎∎∎∎∎∎∎∎∎
  8.803  [2964]  |∎∎
  10.957 [421]   |
  13.110 [108]   |
  15.264 [34]    |
  17.417 [7]     |
  19.571 [2]     |
  21.724 [1]     |

Latency distribution:
  10 % in 1.51 ms 
  25 % in 2.21 ms 
  50 % in 3.16 ms 
  75 % in 4.21 ms 
  90 % in 5.40 ms 
  95 % in 6.24 ms 
  99 % in 8.04 ms 

Status code distribution:
  [OK]   100000 responses   
```