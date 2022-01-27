package project

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.com/maxmac99/compose/pkg/api"
)

func (s *composeService) Pull(project *types.Project, options api.PullOptions) (Stream, error) {
	if options.Quiet {
		stream := newEmptyStream()
		service := getComposeService(s.apiClient, stream)
		err := service.Pull(s.ctx, project, options)
		return nil, err
	}
	stream := newStream()
	service := getComposeService(s.apiClient, stream)
	go func() {
		service.Pull(s.ctx, project, options)
		stream.Close()
	}()
	return stream, nil
}
