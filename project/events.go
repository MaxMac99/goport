package project

import (
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Events(projectName string, options api.EventsOptions) error {
	buffer := newBufferedFile()
	service := getComposeService(s.apiClient, buffer)
	return service.Events(s.ctx, projectName, options)
}
