package project

import (
	"encoding/json"
	"strings"

	"github.com/containerd/console"
)

type Stream interface {
	console.File
	Wait() (string, bool)
}

type streamFile struct {
	title  string
	out    chan map[string]string
	closed bool
}

func NewStream(title string, output chan map[string]string) Stream {
	return &streamFile{
		title:  title,
		out:    output,
		closed: false,
	}
}

func (s *streamFile) Wait() (string, bool) {
	message, ok := <-s.out
	encodedMessage, err := json.Marshal(message)
	if err != nil {
		return "", false
	}
	rawMessage := string(encodedMessage) + "\n"
	return rawMessage, ok
}

func (s *streamFile) Write(p []byte) (n int, err error) {
	if !s.closed {
		buffer := make([]byte, len(p))
		copy(buffer, p)
		messages := string(buffer)
		for _, message := range strings.Split(messages, "\n") {
			output := map[string]string{s.title: message + "\n"}
			s.out <- output
		}
	}
	return len(p), nil
}

func (s *streamFile) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (s *streamFile) Close() error {
	s.closed = true
	return nil
}

func (s *streamFile) Fd() uintptr {
	return 0
}

func (s *streamFile) Name() string {
	return ""
}
