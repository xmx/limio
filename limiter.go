package limio

import "net"

const (
	Byte int = 1
	KiB      = Byte << 10
	MiB      = KiB << 10
	GiB      = MiB << 10
	TiB      = GiB << 10
)

type ReadLimiter interface {
	ReadLimit() int
	SetReadLimit(int)
}

type WriteLimiter interface {
	WriteLimit() int
	SetWriteLimit(int)
}

type Limiter interface {
	Limit(int)
	ReadLimiter
	WriteLimiter
}

type Counter interface {
	ReadCount() int64
	WriteCount() int64
}

type CountLimiter interface {
	Counter
	Limiter
}

type ConnCountLimiter interface {
	net.Conn
	CountLimiter
}
