package limio

import "net"

const (
	Byte int = 1
	KiB      = Byte << 10
	MiB      = KiB << 10
	GiB      = MiB << 10
	TiB      = GiB << 10
)

// ReadLimiter 读取限速
type ReadLimiter interface {
	// ReadLimit 最大读取速度 byte/s
	ReadLimit() int

	// SetReadLimit 设置最大读取速度 byte/s
	SetReadLimit(int)
}

// WriteLimiter 写入限速
type WriteLimiter interface {
	// WriteLimit 最大写入速度 byte/s
	WriteLimit() int
	// SetWriteLimit 设置最大写入速度 byte/s
	SetWriteLimit(int)
}

// ReadWriteLimiter 读写限速
type ReadWriteLimiter interface {
	// SetLimit 设置读取和写入的最大速度 byte/s
	SetLimit(int)
	ReadLimiter
	WriteLimiter
}

// Counter 读写字节计数器
type Counter interface {
	// ReadCount 读取总字节数
	ReadCount() int64
	// WriteCount 写入总字节数
	WriteCount() int64
}

type ConnLimiter interface {
	net.Conn
	Counter
	ReadWriteLimiter
}
