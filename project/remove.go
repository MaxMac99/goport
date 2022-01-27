package project

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Remove(project *types.Project, options api.RemoveOptions) error {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.Remove(s.ctx, project, options)
}
