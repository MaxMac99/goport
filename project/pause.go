package project

import (
	"io"

	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Pause(projectName string, options api.PauseOptions) (io.WriteCloser, error) {
	buffer := newBufferedFile()
	service := getComposeService(s.apiClient, buffer)
	err := service.Pause(s.ctx, projectName, options)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
