# limio

> 本代码仅供学习和参考，__不建议__ 将本库引入到项目中。
> 
> 本代码为抛砖引玉之用，如你有类似的 I/O 限流需求。可以参考本代码，然后自己
> 动手实现一个限流器，本代码尚未经过严格的测试验证，只是本人在个人小项目中简
> 单使用和测试，不够客观严格。

## I/O 限流

使用示例：

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
