package project

import (
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Top(projectName string, services []string) ([]api.ContainerProcSummary, error) {
	service := getComposeService(s.apiClient)
	summary, err := service.Top(s.ctx, projectName, services)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
