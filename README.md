# limio

I/O 限速组件，如：

```go
// 限制 net.Conn 最大下载速度 200k/s，最大上传速度 100k/s
nc := limio.LimitConn(conn, 200*limio.KiB, 100*limio.KiB)
// 限制 io.Reader 最大读取速度为 1m/s
nr := limio.LimitReader(r, limio.MiB)
// 调整 net.Conn 上传和下载最大速度为 1m/s
nc.SetLimit(limio.MiB)
// 调整 io.Reader 最大读取速度为 2m/s
nr.SetReadLimit(2 * limio.MiB)

_ = nc.ReadLimit()  // 获取读取（下载）速度上限
_ = nc.WriteLimit() // 获取写入（上传）速度上限
_ = nc.ReadCount()  // 已经读取（下载）的总字节数
_ = nc.WriteCount() // 已经写入（上传）的总字节数
```
