package project

import (
	"context"
	"io"

	"github.com/compose-spec/compose-go/types"
	"github.com/containerd/console"
	"github.com/docker/docker/client"
	"github.com/maxmac99/goport/goport"
	"gitlab.com/maxmac99/compose/pkg/api"
	"gitlab.com/maxmac99/compose/pkg/compose"
)

type ProjectService interface {
	GetContext() context.Context
	Build(project *types.Project, options api.BuildOptions) (io.ReadCloser, error)
	GetActiveStacks(opts api.ListOptions) ([]Stack, error)
	Convert(project *types.Project, options api.ConvertOptions) ([]byte, error)
	Down(projectName string, options api.DownOptions) (io.ReadCloser, error)
	Events(projectName string, options api.EventsOptions) error
	Images(projectName string, options api.ImagesOptions) ([]api.ImageSummary, error)
	Kill(project *types.Project, opts api.KillOptions) (io.ReadCloser, error)
	Logs(projectName string, consumer api.LogConsumer, options api.LogOptions) error
	Pause(projectName string, options api.PauseOptions) (io.WriteCloser, error)
	Ps(projectName string, options api.PsOptions) ([]api.ContainerSummary, error)
	Pull(project *types.Project, options api.PullOptions) (io.ReadCloser, error)
	Push(project *types.Project, options api.PushOptions) (io.ReadCloser, error)
	Remove(project *types.Project, options api.RemoveOptions) (io.ReadCloser, error)
	Restart(project *types.Project, options api.RestartOptions) (io.ReadCloser, error)
	Run(project *types.Project, options api.RunOptions) (int, error)
	Start(project *types.Project, options api.StartOptions) (io.ReadCloser, error)
	Stop(project *types.Project, options api.StopOptions) (io.ReadCloser, error)
	Top(projectName string, services []string) ([]api.ContainerProcSummary, error)
	Unpause(projectName string, options api.PauseOptions) (io.WriteCloser, error)
	Up(project *types.Project, options api.UpOptions) (io.ReadCloser, error)
}

type composeService struct {
	apiClient client.APIClient
	ctx       context.Context
}

func (s *composeService) GetContext() context.Context {
	return s.ctx
}

func GetProjectService(apiClient client.APIClient, ctx context.Context) ProjectService {
	return &composeService{
		apiClient: apiClient,
		ctx:       ctx,
	}
}

func getComposeService(apiClient client.APIClient, writer console.File) api.Service {
	server := goport.GetGoPort()
	return compose.NewComposeService(apiClient, server.ConfigFile(), writer, writer)
}
