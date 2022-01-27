package project

import "gitlab.com/maxmac99/compose/pkg/api"

func (s *composeService) Logs(projectName string, consumer api.LogConsumer, options api.LogOptions) error {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.Logs(s.ctx, projectName, consumer, options)
}
