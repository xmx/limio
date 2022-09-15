package limio

import (
	"context"
	"io"

	"golang.org/x/time/rate"
)

func LimitReader(r io.Reader, n int) ReadLimiter {
	lim := rate.NewLimiter(rate.Limit(n), n)

	return &limitReader{
		r:   r,
		lim: lim,
		max: n,
	}
}

type limitReader struct {
	r   io.Reader
	lim *rate.Limiter
	max int
}

func (rd *limitReader) Read(p []byte) (int, error) {
	sz := len(p)
	if sz == 0 {
		return 0, nil
	}

	if max := rd.max; sz > max {
		sz = max
	}

	_ = rd.lim.WaitN(context.Background(), sz)

	return rd.r.Read(p[:sz])
}

func (rd *limitReader) ReadLimit() int {
	return rd.max
}

func (rd *limitReader) SetReadLimit(n int) {
	rd.lim.SetBurst(n)
	rd.lim.SetLimit(rate.Limit(n))
	rd.max = n
}
