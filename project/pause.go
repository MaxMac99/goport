package project

import (
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Pause(projectName string, options api.PauseOptions) error {
	service := getComposeService(s.apiClient)
	return service.Pause(s.ctx, projectName, options)
}
