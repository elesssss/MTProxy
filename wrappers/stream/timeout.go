package stream

import (
	"net"
	"time"

	"github.com/9seconds/mtg/conntypes"
	"go.uber.org/zap"
)

const (
	timeoutRead  = 2 * time.Minute
	timeoutWrite = 2 * time.Minute
)

type wrapperTimeout struct {
	parent conntypes.StreamReadWriteCloser
}

func (w *wrapperTimeout) WriteTimeout(p []byte, timeout time.Duration) (int, error) {
	return w.parent.WriteTimeout(p, timeout) // nolint: wrapcheck
}

func (w *wrapperTimeout) Write(p []byte) (int, error) {
	return w.parent.WriteTimeout(p, timeoutWrite) // nolint: wrapcheck
}

func (w *wrapperTimeout) ReadTimeout(p []byte, timeout time.Duration) (int, error) {
	return w.parent.ReadTimeout(p, timeout) // nolint: wrapcheck
}

func (w *wrapperTimeout) Read(p []byte) (int, error) {
	return w.parent.ReadTimeout(p, timeoutRead) // nolint: wrapcheck
}

func (w *wrapperTimeout) Close() error {
	return w.parent.Close() // nolint: wrapcheck
}

func (w *wrapperTimeout) Conn() net.Conn {
	return w.parent.Conn()
}

func (w *wrapperTimeout) Logger() *zap.SugaredLogger {
	return w.parent.Logger().Named("timeout")
}

func (w *wrapperTimeout) LocalAddr() *net.TCPAddr {
	return w.parent.LocalAddr()
}

func (w *wrapperTimeout) RemoteAddr() *net.TCPAddr {
	return w.parent.RemoteAddr()
}

func NewTimeout(parent conntypes.StreamReadWriteCloser) conntypes.StreamReadWriteCloser {
	return &wrapperTimeout{
		parent: parent,
	}
}
