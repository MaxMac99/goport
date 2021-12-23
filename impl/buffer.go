package impl

type channelBuffer struct {
	out chan []byte
}

func newChannelBuffer() channelBuffer {
	return channelBuffer{
		out: make(chan []byte),
	}
}

func (b *channelBuffer) Write(p []byte) (n int, err error) {
	b.out <- p
	return len(p), nil
}

func (b *channelBuffer) Close() error {
	close(b.out)
	return nil
}
