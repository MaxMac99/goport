package project

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Convert(project *types.Project, options api.ConvertOptions) ([]byte, error) {
	service := getComposeService(s.apiClient, nil)
	return service.Convert(s.ctx, project, options)
}
