package project

import (
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Down(projectName string, options api.DownOptions) error {
	service := getComposeService(s.apiClient)
	return service.Down(s.ctx, projectName, options)
}
