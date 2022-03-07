package project

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func (s *composeService) Convert(project *types.Project, options api.ConvertOptions) ([]byte, error) {
	service := getComposeService(s.apiClient)
	return service.Convert(s.ctx, project, options)
}
