package project

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Build(project *types.Project, options api.BuildOptions) (Stream, error) {
	if options.Quiet {
		stream := newEmptyStream()
		service := getComposeService(s.apiClient, stream)
		err := service.Build(s.ctx, project, options)
		return nil, err
	}
	stream := newStream()
	service := getComposeService(s.apiClient, stream)
	go func() {
		service.Build(s.ctx, project, options)
		stream.Close()
	}()
	return stream, nil
}
