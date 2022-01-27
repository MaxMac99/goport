package project

type emptyStream struct {
}

func newEmptyStream() Stream {
	return &emptyStream{}
}

func (s *emptyStream) Wait() (string, bool) {
	return "", true
}

func (s *emptyStream) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (s *emptyStream) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (s *emptyStream) Close() error {
	return nil
}

func (s *emptyStream) Fd() uintptr {
	return 0
}

func (s *emptyStream) Name() string {
	return ""
}
