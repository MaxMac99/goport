package impl

import (
	"io"
	"time"
)

type stream struct {
	out   chan map[string]interface{}
	title string
}

func newStream(title string, out chan map[string]interface{}) io.WriteCloser {
	return &stream{
		out:   out,
		title: title,
	}
}

func (b *stream) Write(p []byte) (n int, err error) {
	text := string(p)
	message := map[string]interface{}{
		b.title:     text,
		"timestamp": time.Now().Unix(),
	}
	b.out <- message
	return len(p), nil
}

func (b *stream) Close() error {
	close(b.out)
	return nil
}
