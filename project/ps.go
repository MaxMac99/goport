package project

import "github.com/docker/compose/v2/pkg/api"

func (s *composeService) Ps(projectName string, options api.PsOptions) ([]api.ContainerSummary, error) {
	service := getComposeService(s.apiClient)
	return service.Ps(s.ctx, projectName, options)
}
