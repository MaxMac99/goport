package project

import (
	"context"

	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
	"github.com/docker/docker/client"
	"gitlab.com/maxmac99/goport/goport"
)

type ProjectService interface {
	GetContext() context.Context
	Build(project *types.Project, options api.BuildOptions) error
	GetActiveStacks(opts api.ListOptions) ([]Stack, error)
	Convert(project *types.Project, options api.ConvertOptions) ([]byte, error)
	Down(projectName string, options api.DownOptions) error
	Events(projectName string, options api.EventsOptions) error
	Images(projectName string, options api.ImagesOptions) ([]api.ImageSummary, error)
	Kill(project *types.Project, opts api.KillOptions) error
	Logs(projectName string, consumer api.LogConsumer, options api.LogOptions) error
	Pause(projectName string, options api.PauseOptions) error
	Ps(projectName string, options api.PsOptions) ([]api.ContainerSummary, error)
	Pull(project *types.Project, options api.PullOptions) error
	Push(project *types.Project, options api.PushOptions) error
	Remove(project *types.Project, options api.RemoveOptions) error
	Restart(project *types.Project, options api.RestartOptions) error
	Run(project *types.Project, options api.RunOptions) (int, error)
	Start(project *types.Project, options api.StartOptions) error
	Stop(project *types.Project, options api.StopOptions) error
	Top(projectName string, services []string) ([]api.ContainerProcSummary, error)
	Unpause(projectName string, options api.PauseOptions) error
	Up(project *types.Project, options api.UpOptions) error
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

func getComposeService(apiClient client.APIClient) api.Service {
	server := goport.GetGoPort()
	return compose.NewComposeService(apiClient, server.ConfigFile())
}
