package project

import (
	"io"

	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Push(project *types.Project, options api.PushOptions) (io.ReadCloser, error) {
	buffer := newBufferedFile()
	service := getComposeService(s.apiClient, buffer)
	err := service.Push(s.ctx, project, options)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
