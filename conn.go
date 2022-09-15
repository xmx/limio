package limio

import (
	"context"
	"net"
	"time"

	"golang.org/x/time/rate"
)

func LimitConn(conn net.Conn, rmx, wmx int) ConnLimiter {
	rlm := rate.NewLimiter(rate.Limit(rmx), rmx)
	wlm := rate.NewLimiter(rate.Limit(wmx), wmx)

	return &connect{
		conn: conn,
		rlm:  rlm,
		rmx:  rmx,
		wlm:  wlm,
		wmx:  wmx,
	}
}

type connect struct {
	conn net.Conn

	rlm *rate.Limiter
	rmx int
	rct int64

	wlm *rate.Limiter
	wmx int
	wct int64
}

func (c *connect) Read(b []byte) (int, error) {
	sz := len(b)
	if sz == 0 {
		return 0, nil
	}

	if rmx := c.rmx; sz > rmx {
		sz = rmx
	}

	_ = c.rlm.WaitN(context.Background(), sz)
	n, err := c.conn.Read(b[:sz])
	c.rct += int64(n)

	return n, err
}

func (c *connect) Write(b []byte) (n int, err error) {
	for {
		sz := len(b)
		if sz == 0 {
			break
		}
		if wmx := c.wmx; sz > wmx {
			sz = wmx
		}

		_ = c.wlm.WaitN(context.Background(), sz)
		wn, ex := c.conn.Write(b[:sz])
		if ex != nil {
			err = ex
			break
		}

		c.wct += int64(wn)
		b = b[wn:]
	}

	return
}

func (c *connect) Close() error {
	return c.conn.Close()
}

func (c *connect) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *connect) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *connect) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *connect) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *connect) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func (c *connect) ReadCount() int64 {
	return c.rct
}

func (c *connect) WriteCount() int64 {
	return c.wct
}

func (c *connect) SetLimit(i int) {
	c.SetReadLimit(i)
	c.SetWriteLimit(i)
}

func (c *connect) ReadLimit() int {
	return c.rmx
}

func (c *connect) SetReadLimit(i int) {
	c.rlm.SetBurst(i)
	c.rlm.SetLimit(rate.Limit(i))
	c.rmx = i
}

func (c *connect) WriteLimit() int {
	return c.wmx
}

func (c *connect) SetWriteLimit(i int) {
	c.wlm.SetBurst(i)
	c.wlm.SetLimit(rate.Limit(i))
	c.wmx = i
}
