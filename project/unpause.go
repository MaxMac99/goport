package project

import (
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Unpause(projectName string, options api.PauseOptions) error {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.UnPause(s.ctx, projectName, options)
}
