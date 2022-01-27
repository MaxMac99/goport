package project

import (
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Top(projectName string, services []string) ([]api.ContainerProcSummary, error) {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	summary, err := service.Top(s.ctx, projectName, services)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
