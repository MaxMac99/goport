package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Build(project *types.Project, options api.BuildOptions) error {
	service := getComposeService(s.apiClient)
	return service.Build(s.ctx, project, options)
}
