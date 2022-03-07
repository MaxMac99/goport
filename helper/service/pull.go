package service

import (
	"context"
	"os"
	"os/exec"

	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
	goportcontext "gitlab.com/maxmac99/goport/context"
	"gitlab.com/maxmac99/goport/goport"
	"gitlab.com/maxmac99/goport/models"
	"gitlab.com/maxmac99/goport/project"
)

func RunPull(opts models.ProjectPullOpts) project.Stream {
	args := buildPullArgs(opts)
	cmd := exec.Command("helper/helper", args...)
	output := make(chan map[string]string)
	outStream := project.NewStream("stdout", output)
	errStream := project.NewStream("stderr", output)
	cmd.Stdout = outStream
	cmd.Stderr = errStream

	go func() {
		err := cmd.Run()
		if err != nil {
			output <- map[string]string{"error": err.Error()}
		}
		outStream.Close()
		errStream.Close()
		close(output)
	}()
	return outStream
}

func Pull() error {
	defaultcontext := context.Background()
	opts := parsePullArgs()

	client, err := goportcontext.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	server := goport.GetGoPort()
	service := compose.NewComposeService(client, server.ConfigFile())
	project, err := project.GetProject(opts.Name)
	if err != nil {
		return err
	}
	err = service.Pull(defaultcontext, project, api.PullOptions{
		Quiet:          opts.Quiet,
		IgnoreFailures: opts.Ignorepullfailures,
	})
	if err != nil {
		return err
	}
	return nil
}

func buildPullArgs(opts models.ProjectPullOpts) []string {
	args := []string{"pull", opts.Name, opts.Context}
	if opts.Ignorepullfailures {
		args = append(args, "--ignorepullfailures")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	return args
}

func parsePullArgs() models.ProjectPullOpts {
	opts := models.ProjectPullOpts{
		Name:               os.Args[2],
		Context:            "",
		Ignorepullfailures: false,
		Quiet:              false,
	}
	var currentArg string
	for _, arg := range os.Args[3:] {
		if arg == "--context" {
			currentArg = "context"
		} else if arg == "--ignorepullfailures" {
			opts.Ignorepullfailures = true
		} else if arg == "--quiet" {
			opts.Quiet = true
		} else if currentArg == "context" {
			opts.Context = arg
		}
	}
	return opts
}
