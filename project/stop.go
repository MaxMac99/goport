package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Stop(project *types.Project, options api.StopOptions) error {
	service := getComposeService(s.apiClient)
	return service.Stop(s.ctx, project, options)
}
