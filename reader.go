package limio

import (
	"context"
	"io"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	err := rd.lim.WaitN(ctx, sz)
	if cancel(); err != nil {
		return 0, err
	}

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
