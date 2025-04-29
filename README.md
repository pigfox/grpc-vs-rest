# grpc-vs-rest
A grpc-vs-rest comparison

```In Terminal
go run ./rest/main.go
run test
go run ./grpc/main.go
run test


```Use a benchmarking tool like hey (for HTTP)

REST
```
$ hey -n 100000 -c 50 http://localhost:8888/user
Summary:
  Total:	1.6221 secs
  Slowest:	0.0117 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	61648.7015
  
  Total data:	6000000 bytes
  Size/request:	60 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [83134]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [13678]	|■■■■■■■
  0.004 [2326]	|■
  0.005 [601]	|
  0.006 [178]	|
  0.007 [58]	|
  0.008 [19]	|
  0.009 [1]	|
  0.011 [2]	|
  0.012 [2]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0006 secs
  75% in 0.0010 secs
  90% in 0.0016 secs
  95% in 0.0021 secs
  99% in 0.0034 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0117 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0023 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0103 secs
  resp wait:	0.0007 secs, 0.0000 secs, 0.0095 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0100 secs

Status code distribution:
  [200]	100000 responses

gRPC:
$ hey -n 100000 -c 50 http://localhost:8888/user

Summary:
  Total:	30.0228 secs
  Slowest:	0.1103 secs
  Fastest:	0.0011 secs
  Average:	0.0149 secs
  Requests/sec:	3330.7998
  
  Total data:	6000000 bytes
  Size/request:	60 bytes

Response time histogram:
  0.001 [1]	|
  0.012 [35855]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.023 [53315]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.034 [9773]	|■■■■■■■
  0.045 [882]	|■
  0.056 [117]	|
  0.067 [26]	|
  0.077 [21]	|
  0.088 [7]	|
  0.099 [2]	|
  0.110 [1]	|


Latency distribution:
  10% in 0.0073 secs
  25% in 0.0103 secs
  50% in 0.0141 secs
  75% in 0.0185 secs
  90% in 0.0233 secs
  95% in 0.0265 secs
  99% in 0.0341 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.1103 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0047 secs
  resp wait:	0.0148 secs, 0.0010 secs, 0.1102 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0049 secs

Status code distribution:
  [200]	100000 responses
