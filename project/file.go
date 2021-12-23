package project

import (
	"bytes"

	"github.com/containerd/console"
)

type bufferedFile struct {
	bytes.Buffer
}

func newBufferedFile() console.File {
	return &bufferedFile{}
}

func (b *bufferedFile) Close() error {
	b.Buffer.Reset()
	return nil
}

func (b *bufferedFile) Fd() uintptr {
	return 0
}

func (b *bufferedFile) Name() string {
	return ""
}
