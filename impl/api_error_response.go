package impl

import (
	"encoding/json"
	"io"
)

func WriteError(err error, w io.Writer) {
	message := map[string]string{
		"error": err.Error(),
	}
	b, parseErr := json.Marshal(message)
	if parseErr != nil {
		return
	}
	w.Write(b)
}
