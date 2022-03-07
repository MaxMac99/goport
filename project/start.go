package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Start(project *types.Project, options api.StartOptions) error {
	service := getComposeService(s.apiClient)
	return service.Start(s.ctx, project, options)
}
