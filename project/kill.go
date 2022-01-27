package project

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Kill(project *types.Project, opts api.KillOptions) error {
	buffer := newEmptyStream()
	service := getComposeService(s.apiClient, buffer)
	return service.Kill(s.ctx, project, opts)
}
