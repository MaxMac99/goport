package context

import (
	"fmt"
	"strings"

	"github.com/maxmac99/goport/goport"
	"github.com/pkg/errors"
)

func RemoveContext(server goport.GoPort, name string, force bool) error {
	var errs []string
	currentCtx := server.CurrentContext()
	if name == "default" {
		errs = append(errs, `default: context "default" cannot be removed`)
	} else if err := doRemove(server, name, name == currentCtx, force); err != nil {
		errs = append(errs, fmt.Sprintf("%s: %s", name, err))
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}

func doRemove(server goport.GoPort, name string, isCurrent, force bool) error {
	if _, err := server.ContextStore().GetMetadata(name); err != nil {
		return err
	}
	if isCurrent {
		if !force {
			return errors.New("context is in use, set -f flag to force remove")
		}
		// fallback to DOCKER_HOST
		cfg := server.ConfigFile()
		cfg.CurrentContext = ""
		if err := cfg.Save(); err != nil {
			return err
		}
	}
	return server.ContextStore().Remove(name)
}
