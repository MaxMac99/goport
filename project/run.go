package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Run(project *types.Project, options api.RunOptions) (int, error) {
	service := getComposeService(s.apiClient)
	return service.RunOneOffContainer(s.ctx, project, options)
}
