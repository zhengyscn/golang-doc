
## Test

```bash
$ go test
$ go test -v .   // 能够输出log
$ go test -v . -run TestSum
$ go test -v . -run TestSum -count 3 -cpu 4 -parallel 2  // parallel指定运行处理器的个数
$ go test -cover  // 覆盖率
$ go test -cover -covermode=count -coverprofile cover.out
$ go tool cover -html=cover1.out

$ go test -coverprofile=c.out
$ go tool cover -html=c.out -o coverage.html
```

## ab

```bash
$ ab -k -c 5 -n 10000 'http://xxxx/api/v1/xxx/'

Requests per second:    0.85 [#/sec] (mean)         // 吞吐率 和并发数有直接关系
Time per request:       5908.457 [ms] (mean)        // 客户端平均响应等待时间
Time per request:       1181.691 [ms] (mean, across all concurrent requests)    // 服务器的处理时间
Transfer rate:          195.92 [Kbytes/sec] received
```