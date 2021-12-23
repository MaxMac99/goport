package context

import (
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/context/docker"
	"github.com/docker/cli/cli/context/store"
	"github.com/maxmac99/goport/goport"
	"github.com/pkg/errors"
)

type CreateOptions struct {
	Name        string
	Description string
	Docker      map[string]string
	From        string
}

func CreateContext(server goport.GoPort, opts *CreateOptions) error {
	s := server.ContextStore()
	err := checkContextNameForCreation(s, opts.Name)
	if err != nil {
		return err
	}
	stackOrchestrator := command.Orchestrator("")
	switch {
	case opts.From == "" && opts.Docker == nil:
		err = createFromExistingContext(s, "default", stackOrchestrator, opts)
	case opts.From != "":
		err = createFromExistingContext(s, opts.From, stackOrchestrator, opts)
	default:
		err = createNewContext(opts, stackOrchestrator, server, s)
	}
	return err
}

func createNewContext(o *CreateOptions, stackOrchestrator command.Orchestrator, server goport.GoPort, s store.Writer) error {
	if o.Docker == nil {
		return errors.New("docker endpoint configuration is required")
	}
	contextMetadata := newContextMetadata(stackOrchestrator, o)
	contextTLSData := store.ContextTLSData{
		Endpoints: make(map[string]store.EndpointTLSData),
	}
	dockerEP, dockerTLS, err := getDockerEndpointMetadataAndTLS(server, o.Docker)
	if err != nil {
		return errors.Wrap(err, "unable to create docker endpoint config")
	}
	contextMetadata.Endpoints[docker.DockerEndpoint] = dockerEP
	if dockerTLS != nil {
		contextTLSData.Endpoints[docker.DockerEndpoint] = *dockerTLS
	}
	if err := s.CreateOrUpdate(contextMetadata); err != nil {
		return err
	}
	if err := s.ResetTLSMaterial(o.Name, &contextTLSData); err != nil {
		return err
	}
	return nil
}

func checkContextNameForCreation(s store.Reader, name string) error {
	if err := store.ValidateContextName(name); err != nil {
		return err
	}
	if _, err := s.GetMetadata(name); !store.IsErrContextDoesNotExist(err) {
		if err != nil {
			return errors.Wrap(err, "error while getting existing contexts")
		}
		return errors.Errorf("context %q already exists", name)
	}
	return nil
}

type descriptionAndOrchestratorStoreDecorator struct {
	store.Reader
	description  string
	orchestrator command.Orchestrator
}

func createFromExistingContext(s store.ReaderWriter, fromContextName string, stackOrchestrator command.Orchestrator, o *CreateOptions) error {
	if len(o.Docker) != 0 {
		return errors.New("cannot use --docker flags when --from is set")
	}
	reader := store.Export(fromContextName, &descriptionAndOrchestratorStoreDecorator{
		Reader:       s,
		description:  o.Description,
		orchestrator: stackOrchestrator,
	})
	defer reader.Close()
	return store.Import(o.Name, s, reader)
}

func newContextMetadata(stackOrchestrator command.Orchestrator, o *CreateOptions) store.Metadata {
	return store.Metadata{
		Endpoints: make(map[string]interface{}),
		Metadata: command.DockerContext{
			Description:       o.Description,
			StackOrchestrator: stackOrchestrator,
		},
		Name: o.Name,
	}
}
