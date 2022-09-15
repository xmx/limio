# limio

I/O 限速组件，如：

限制 socket 上传速度为 200 k/s，下载速度为 500 k/s

```go
func ExampleLimitConn(conn net.Conn) {
	newConn := limio.LimitConn(conn, 200*limio.KiB, 500*limio.KiB)
}
```

支持实时调整速率
