package project

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Run(project *types.Project, options api.RunOptions) (int, error) {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.RunOneOffContainer(s.ctx, project, options)
}
