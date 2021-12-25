package context

import (
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/context/docker"
	"github.com/docker/docker/client"
	"gitlab.com/maxmac99/goport/goport"
)

func GetClientForContext(server goport.GoPort, context string) (client.APIClient, error) {
	contextName := context
	if context == "" {
		contextName = command.DefaultContextName
	}

	ctxMeta, err := server.ContextStore().GetMetadata(contextName)
	if err != nil {
		return nil, err
	}
	epMeta, err := docker.EndpointFromContext(ctxMeta)
	if err != nil {
		return nil, err
	}
	dockerEndpoint, err := docker.WithTLSData(server.ContextStore(), contextName, epMeta)
	if err != nil {
		return nil, err
	}

	clientOpts, err := dockerEndpoint.ClientOpts()
	if err != nil {
		return nil, err
	}

	customHeaders := make(map[string]string, len(server.ConfigFile().HTTPHeaders))
	for k, v := range server.ConfigFile().HTTPHeaders {
		customHeaders[k] = v
	}
	customHeaders["User-Agent"] = command.UserAgent()
	clientOpts = append(clientOpts, client.WithHTTPHeaders(customHeaders))

	return client.NewClientWithOpts(clientOpts...)
}
