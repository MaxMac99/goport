package project

import (
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Down(projectName string, options api.DownOptions) error {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.Down(s.ctx, projectName, options)
}
