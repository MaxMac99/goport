package project

import (
	"io"

	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Kill(project *types.Project, opts api.KillOptions) (io.ReadCloser, error) {
	buffer := newBufferedFile()
	service := getComposeService(s.apiClient, buffer)
	err := service.Kill(s.ctx, project, opts)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
