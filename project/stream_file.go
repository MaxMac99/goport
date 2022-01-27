package project

import "github.com/containerd/console"

type Stream interface {
	console.File
	Wait() (string, bool)
}

type streamFile struct {
	out    chan string
	closed bool
}

func newStream() Stream {
	return &streamFile{
		out:    make(chan string),
		closed: false,
	}
}

func (s *streamFile) Wait() (string, bool) {
	message, ok := <-s.out
	return message, ok
}

func (s *streamFile) Write(p []byte) (n int, err error) {
	if !s.closed {
		buffer := make([]byte, len(p))
		copy(buffer, p)
		s.out <- string(buffer)
	}
	return len(p), nil
}

func (s *streamFile) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (s *streamFile) Close() error {
	if !s.closed {
		close(s.out)
	}
	s.closed = true
	return nil
}

func (s *streamFile) Fd() uintptr {
	return 0
}

func (s *streamFile) Name() string {
	return ""
}
