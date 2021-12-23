package controllers

import (
	"bufio"
	"io"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/context/docker"
	"github.com/docker/cli/cli/context/store"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/maxmac99/goport/context"
	"github.com/maxmac99/goport/goport"
	"github.com/moby/term"
)

func ResolveContext(currentContext string) (client.APIClient, error) {
	_, _, stderr := term.StdStreams()
	configFile := config.LoadDefaultConfigFile(stderr)
	baseContextStore := store.New(config.ContextStoreDir(), command.DefaultContextStoreConfig())
	commonOpts := flags.CommonOptions{
		Debug: true,
	}
	contextStore := &command.ContextStoreWithDefault{
		Store: baseContextStore,
		Resolver: func() (*command.DefaultContext, error) {
			return command.ResolveDefaultContext(&commonOpts, configFile, command.DefaultContextStoreConfig(), stderr)
		},
	}

	contextName := currentContext
	if currentContext == "" {
		contextName = command.DefaultContextName
	}
	ctxMeta, err := contextStore.GetMetadata(contextName)
	if err != nil {
		return nil, err
	}
	epMeta, err := docker.EndpointFromContext(ctxMeta)
	if err != nil {
		return nil, err
	}
	dockerEndpoint, err := docker.WithTLSData(contextStore, contextName, epMeta)
	if err != nil {
		return nil, err
	}

	clientOpts, err := dockerEndpoint.ClientOpts()
	if err != nil {
		return nil, err
	}

	customHeaders := make(map[string]string, len(configFile.HTTPHeaders))
	for k, v := range configFile.HTTPHeaders {
		customHeaders[k] = v
	}
	customHeaders["User-Agent"] = command.UserAgent()
	clientOpts = append(clientOpts, client.WithHTTPHeaders(customHeaders))

	return client.NewClientWithOpts(clientOpts...)
}

func ResolveContexts(contexts []string) (map[string]client.APIClient, error) {
	clients := make(map[string]client.APIClient, len(contexts))
	if len(contexts) == 0 {
		client, err := ResolveContext("default")
		if err != nil {
			return nil, err
		}
		clients["default"] = client
		return clients, nil
	}
	if len(contexts) == 1 && contexts[0] == "all" {
		server := goport.GetGoPort()
		allContexts, err := context.ListContext(server)
		if err != nil {
			return nil, err
		}
		contexts = nil
		for _, context := range *allContexts {
			contexts = append(contexts, context.Name)
		}
	}
	for _, context := range contexts {
		client, err := ResolveContext(context)
		if err != nil {
			return nil, err
		}
		clients[context] = client
	}
	return clients, nil
}

func StreamResponse(c *gin.Context, r io.ReadCloser) func(w io.Writer) bool {
	scanner := bufio.NewScanner(r)
	go func() {
		<-c.Request.Context().Done()
		r.Close()
	}()
	return func(w io.Writer) bool {
		if !scanner.Scan() {
			return false
		}
		content := scanner.Bytes()
		content = append(content, '\n')
		w.Write(content)
		return true
	}
}
