package project

import (
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Pause(projectName string, options api.PauseOptions) error {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.Pause(s.ctx, projectName, options)
}
