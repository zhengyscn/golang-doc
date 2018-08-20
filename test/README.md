
## Test
```
$ go test
$ go test -v .   // 能够输出log
$ go test -v . -run TestSum
$ go test -v . -run TestSum -count 3 -cpu 4 -parallel 2  // parallel指定运行处理器的个数

$ mockgen -source=hellomock.go > mock_hellomock/mock_talker.go
```