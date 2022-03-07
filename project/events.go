package project

import (
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Events(projectName string, options api.EventsOptions) error {
	service := getComposeService(s.apiClient)
	return service.Events(s.ctx, projectName, options)
}
