package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Restart(project *types.Project, options api.RestartOptions) error {
	service := getComposeService(s.apiClient)
	return service.Restart(s.ctx, project, options)
}
