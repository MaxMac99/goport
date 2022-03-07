package service

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
	goportcontext "gitlab.com/maxmac99/goport/context"
	"gitlab.com/maxmac99/goport/goport"
	"gitlab.com/maxmac99/goport/models"
	"gitlab.com/maxmac99/goport/project"
)

func RunBuild(opts models.ProjectBuildOpts) project.Stream {
	args := buildBuildArgs(opts)
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

func Build() error {
	defaultcontext := context.Background()
	opts := parseBuildArgs()

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
	progress := "plain"
	if opts.Quiet {
		progress = "quiet"
	}
	var args map[string]*string
	if opts.BuildArgs != "" {
		err = json.Unmarshal([]byte(opts.BuildArgs), &args)
		if err != nil {
			return err
		}
	}
	err = service.Build(defaultcontext, project, api.BuildOptions{
		Pull:     opts.Pull,
		Progress: progress,
		Args:     args,
		NoCache:  opts.NoCache,
		Quiet:    opts.Quiet,
		Services: opts.Services,
	})
	if err != nil {
		return err
	}
	return nil
}

func buildBuildArgs(opts models.ProjectBuildOpts) []string {
	args := []string{"build", opts.Name, opts.Context}
	if opts.NoCache {
		args = append(args, "--nocache")
	}
	if opts.Pull {
		args = append(args, "--pull")
	}
	if opts.Quiet {
		args = append(args, "--quiet")
	}
	if len(opts.Services) > 0 {
		args = append(args, "--services")
		args = append(args, opts.Services...)
	}
	if opts.BuildArgs != "" {
		args = append(args, "--buildargs", opts.BuildArgs)
	}
	return args
}

func parseBuildArgs() models.ProjectBuildOpts {
	opts := models.ProjectBuildOpts{
		Name:      os.Args[2],
		Context:   "",
		Services:  []string{},
		BuildArgs: "",
		NoCache:   false,
		Pull:      false,
		Quiet:     false,
	}
	var currentArg string
	for _, arg := range os.Args[3:] {
		if arg == "--context" {
			currentArg = "context"
		} else if arg == "--nocache" {
			opts.NoCache = true
		} else if arg == "--pull" {
			opts.Pull = true
		} else if arg == "--quiet" {
			opts.Quiet = true
		} else if arg == "--services" {
			currentArg = "services"
		} else if arg == "--buildargs" {
			currentArg = "buildargs"
		} else if currentArg == "service" {
			opts.Services = append(opts.Services, arg)
		} else if currentArg == "buildargs" {
			if opts.BuildArgs == "" {
				opts.BuildArgs = arg
			} else {
				opts.BuildArgs = opts.BuildArgs + " " + arg
			}
		} else if currentArg == "context" {
			opts.Context = arg
		}
	}
	return opts
}
