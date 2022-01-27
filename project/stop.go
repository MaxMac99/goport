package project

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Stop(project *types.Project, options api.StopOptions) error {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.Stop(s.ctx, project, options)
}
