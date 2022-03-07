package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Up(project *types.Project, options api.UpOptions) error {
	service := getComposeService(s.apiClient)
	return service.Up(s.ctx, project, options)
}
