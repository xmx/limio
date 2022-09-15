# limio

I/O 限速组件，如：

```go
func ExampleLimitConn(conn net.Conn) {
	// 限制下载速度 200k/s 上传速度 500k/s
	newConn := limio.LimitConn(conn, 200*limio.KiB, 500*limio.KiB)
}
```

支持实时调整速率
