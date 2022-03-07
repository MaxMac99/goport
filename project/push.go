package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Push(project *types.Project, options api.PushOptions) error {
	service := getComposeService(s.apiClient)
	return service.Push(s.ctx, project, options)
}
