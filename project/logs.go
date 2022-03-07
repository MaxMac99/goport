package project

import "github.com/docker/compose/v2/pkg/api"

func (s *composeService) Logs(projectName string, consumer api.LogConsumer, options api.LogOptions) error {
	service := getComposeService(s.apiClient)
	return service.Logs(s.ctx, projectName, consumer, options)
}
