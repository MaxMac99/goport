package context

import (
	"fmt"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/context/docker"
	"github.com/docker/cli/cli/context/kubernetes"
	"gitlab.com/maxmac99/goport/goport"
)

type ContextSummary struct {
	Name               string
	Description        string
	DockerEndpoint     string
	KubernetesEndpoint string
	StackOrchestrator  string
}

func ListContext(server goport.GoPort) (*[]ContextSummary, error) {
	contextMap, err := server.ContextStore().List()
	if err != nil {
		return nil, err
	}
	var contexts []ContextSummary
	for _, rawMeta := range contextMap {
		meta, err := command.GetDockerContext(rawMeta)
		if err != nil {
			return nil, err
		}
		dockerEndpoint, err := docker.EndpointFromContext(rawMeta)
		if err != nil {
			return nil, err
		}
		kubernetesEndpoint := kubernetes.EndpointFromContext(rawMeta)
		kubEndpointText := ""
		if kubernetesEndpoint != nil {
			kubEndpointText = fmt.Sprintf("%s (%s)", kubernetesEndpoint.Host, kubernetesEndpoint.DefaultNamespace)
		}
		if rawMeta.Name == command.DefaultContextName {
			meta.Description = "Current DOCKER_HOST based configuration"
		}
		desc := ContextSummary{
			Name:               rawMeta.Name,
			Description:        meta.Description,
			StackOrchestrator:  string(meta.StackOrchestrator),
			DockerEndpoint:     dockerEndpoint.Host,
			KubernetesEndpoint: kubEndpointText,
		}
		contexts = append(contexts, desc)
	}
	return &contexts, nil
}
