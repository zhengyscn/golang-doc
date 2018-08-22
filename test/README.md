
## Test
```Bash
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