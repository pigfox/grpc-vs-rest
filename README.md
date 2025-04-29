# grpc-vs-rest
A grpc-vs-rest comparison

```In Termianl 1
go run rest/main.go

```In Termianl 2
go run rest/main.go

```Use a benchmarking tool like hey (for HTTP) and ghz (for gRPC)

REST
```
$ hey -n 10000 -c 50 http://localhost:8888/user



gRPC:

