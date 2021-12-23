package project

import (
	"io"

	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Down(projectName string, options api.DownOptions) (io.ReadCloser, error) {
	buffer := newBufferedFile()
	service := getComposeService(s.apiClient, buffer)
	err := service.Down(s.ctx, projectName, options)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
