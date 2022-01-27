package impl

type channelBuffer struct {
	out    chan string
	closed bool
}

func newChannelBuffer() channelBuffer {
	return channelBuffer{
		out:    make(chan string),
		closed: false,
	}
}

func (s *channelBuffer) Wait() (string, bool) {
	message, ok := <-s.out
	return message, ok
}

func (b *channelBuffer) Write(p []byte) (n int, err error) {
	if !b.closed {
		buffer := make([]byte, len(p))
		copy(buffer, p)
		b.out <- string(buffer)
	}
	return len(p), nil
}

func (b *channelBuffer) Close() error {
	if !b.closed {
		close(b.out)
	}
	b.closed = true
	return nil
}

type singleStringReader struct {
	out string
}

func newSingleStringReader() singleStringReader {
	return singleStringReader{
		out: "",
	}
}

func (b *singleStringReader) Write(p []byte) (n int, err error) {
	buffer := make([]byte, len(p))
	copy(buffer, p)
	b.out = string(buffer)
	return len(p), nil
}

func (b *singleStringReader) Close() error {
	return nil
}

type emptyReader struct {
}

func newEmptyReader() emptyReader {
	return emptyReader{}
}

func (b *emptyReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (b *emptyReader) Close() error {
	return nil
}
