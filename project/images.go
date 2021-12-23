package project

import "gitlab.com/maxmac99/compose/pkg/api"

func (s *composeService) Images(projectName string, options api.ImagesOptions) ([]api.ImageSummary, error) {
	buffer := newBufferedFile()
	service := getComposeService(s.apiClient, buffer)
	return service.Images(s.ctx, projectName, options)
}
