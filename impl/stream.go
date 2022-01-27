package impl

import (
	"bufio"
	"io"

	"github.com/gin-gonic/gin"
	"gitlab.com/maxmac99/goport/project"
)

func StreamResponse(c *gin.Context, r project.Stream) func(w io.Writer) bool {
	go func() {
		<-c.Request.Context().Done()
		r.Close()
	}()
	return func(w io.Writer) bool {
		content, ok := r.Wait()
		if !ok {
			return false
		}
		w.Write([]byte(content))
		return true
	}
}

func StreamReadingResponse(c *gin.Context, r io.ReadCloser) func(w io.Writer) bool {
	scanner := bufio.NewScanner(r)
	go func() {
		<-c.Request.Context().Done()
		r.Close()
	}()
	return func(w io.Writer) bool {
		if !scanner.Scan() {
			return false
		}
		content := scanner.Bytes()
		content = append(content, '\n')
		w.Write(content)
		return true
	}
}
