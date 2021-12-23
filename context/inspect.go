package context

import (
	"github.com/docker/cli/cli/context/store"
	"github.com/maxmac99/goport/goport"
)

func InspectContext(server goport.GoPort, name string) (*Context, error) {
	c, err := server.ContextStore().GetMetadata(name)
	if err != nil {
		return nil, err
	}
	tlsListing, err := server.ContextStore().ListTLSFiles(name)
	if err != nil {
		return nil, err
	}
	return &Context{
		Metadata:    c,
		TLSMaterial: tlsListing,
		Storage:     server.ContextStore().GetStorageInfo(name),
	}, nil
}

type Context struct {
	store.Metadata
	TLSMaterial map[string]store.EndpointFiles
	Storage     store.StorageInfo
}
