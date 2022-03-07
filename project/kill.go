package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Kill(project *types.Project, opts api.KillOptions) error {
	service := getComposeService(s.apiClient)
	return service.Kill(s.ctx, project, opts)
}
