package project

import "github.com/docker/compose/v2/pkg/api"

func (s *composeService) Images(projectName string, options api.ImagesOptions) ([]api.ImageSummary, error) {
	service := getComposeService(s.apiClient)
	return service.Images(s.ctx, projectName, options)
}
