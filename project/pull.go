package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Pull(project *types.Project, options api.PullOptions) error {
	service := getComposeService(s.apiClient)
	return service.Pull(s.ctx, project, options)
}
