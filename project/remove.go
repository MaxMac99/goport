package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Remove(project *types.Project, options api.RemoveOptions) error {
	service := getComposeService(s.apiClient)
	return service.Remove(s.ctx, project, options)
}
