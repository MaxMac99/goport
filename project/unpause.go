package project

import (
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Unpause(projectName string, options api.PauseOptions) error {
	service := getComposeService(s.apiClient)
	return service.UnPause(s.ctx, projectName, options)
}
