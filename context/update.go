package context

import (
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/context/docker"
	"github.com/docker/cli/cli/context/store"
	"github.com/pkg/errors"
	"gitlab.com/maxmac99/goport/goport"
)

type UpdateOptions struct {
	Name        string
	Description string
	Docker      map[string]string
}

func UpdateContext(server goport.GoPort, o *UpdateOptions) error {
	if err := store.ValidateContextName(o.Name); err != nil {
		return err
	}
	s := server.ContextStore()
	c, err := s.GetMetadata(o.Name)
	if err != nil {
		return err
	}
	dockerContext, err := command.GetDockerContext(c)
	if err != nil {
		return err
	}
	if o.Description != "" {
		dockerContext.Description = o.Description
	}

	c.Metadata = dockerContext

	tlsDataToReset := make(map[string]*store.EndpointTLSData)

	if o.Docker != nil {
		dockerEP, dockerTLS, err := getDockerEndpointMetadataAndTLS(server, o.Docker)
		if err != nil {
			return errors.Wrap(err, "unable to create docker endpoint config")
		}
		c.Endpoints[docker.DockerEndpoint] = dockerEP
		tlsDataToReset[docker.DockerEndpoint] = dockerTLS
	}
	if err := s.CreateOrUpdate(c); err != nil {
		return err
	}
	for ep, tlsData := range tlsDataToReset {
		if err := s.ResetEndpointTLSMaterial(o.Name, ep, tlsData); err != nil {
			return err
		}
	}
	return nil
}
