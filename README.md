# GO HTTP SERVER COMPARISION

To compare http server performance in local environment

## Frameworks

### Kami

```
$ wrk -t12 -c400 -d30s http://127.0.0.1:8088/
Running 30s test @ http://127.0.0.1:8088/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.65ms    3.32ms 124.30ms   89.10%
    Req/Sec     3.60k     2.32k   12.81k    53.52%
  1291501 requests in 30.10s, 158.89MB read
  Socket errors: connect 155, read 5, write 0, timeout 0
Requests/sec:  42908.55
Transfer/sec:      5.28MB
```

### FastHttpServer

```
$ wrk -t12 -c400 -d30s http://127.0.0.1:8088/
Running 30s test @ http://127.0.0.1:8088/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.38ms    2.69ms 126.05ms   97.68%
    Req/Sec     5.98k     3.36k   16.57k    51.40%
  2146905 requests in 30.10s, 307.12MB read
  Socket errors: connect 155, read 0, write 0, timeout 0
Requests/sec:  71322.45
Transfer/sec:     10.20MB
```

### FastHttpServerRouting

```
$ wrk -t12 -c400 -d30s http://127.0.0.1:8088/
Running 30s test @ http://127.0.0.1:8088/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.37ms    0.86ms  98.93ms   87.81%
    Req/Sec     7.00k     3.92k   13.96k    70.17%
  2092796 requests in 30.08s, 295.39MB read
  Socket errors: connect 155, read 0, write 0, timeout 0
Requests/sec:  69579.18
Transfer/sec:      9.82MB
```

## Run benchmark

Start http server

```
$ go get
$ go run raw.go
```

While runing server, you start benchmark tool.

```
$ wrk -t12 -c400 -d30s http://127.0.0.1:8088/
```
