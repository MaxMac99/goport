package project

import "gitlab.com/maxmac99/compose/pkg/api"

func (s *composeService) Ps(projectName string, options api.PsOptions) ([]api.ContainerSummary, error) {
	service := getComposeService(s.apiClient, nil)
	return service.Ps(s.ctx, projectName, options)
}
