package project

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Restart(project *types.Project, options api.RestartOptions) error {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.Restart(s.ctx, project, options)
}
