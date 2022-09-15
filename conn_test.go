package limio_test

import (
	"net"
	"testing"

	"github.com/xmx/limio"
)

func TestConn(t *testing.T) {
	var conn net.Conn
	limio.LimitConn(conn, 100*limio.KiB, 300*limio.KiB)

	limio.LimitReader(nil, limio.MiB)
}
