/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package impl

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/maxmac99/goport/context"
	"gitlab.com/maxmac99/goport/helper/service"
	"gitlab.com/maxmac99/goport/models"
	"gitlab.com/maxmac99/goport/project"
)

// ProjectBuild - Build the project
func ProjectBuild(c *gin.Context, opts *models.ProjectBuildOpts) (func(w io.Writer) bool, error) {
	reader := service.RunBuild(*opts)
	if opts.Quiet {
		return nil, nil
	}
	return StreamResponse(c, reader), nil
}

// ProjectCreate - Create a project
func ProjectCreate(c *gin.Context, opts *models.ProjectCreateOpts) error {
	client, err := context.ResolveContext("default")
	if err != nil {
		return err
	}
	err = project.AddProject(opts.Name, opts.Body, opts.Format)
	if err != nil {
		return err
	}
	p, err := project.GetProject(opts.Name)
	if err != nil {
		project.RemoveProject(opts.Name)
		return err
	}
	service := project.GetProjectService(client, c)
	_, err = service.Convert(p, api.ConvertOptions{
		Format: opts.Format,
		Output: "",
	})
	if err != nil {
		project.RemoveProject(opts.Name)
		return err
	}
	return nil
}

// ProjectDown - Stops containers and removes containers, networks, volumes, and images created by up
func ProjectDown(c *gin.Context, opts *models.ProjectDownOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	}
	name := opts.Name
	if err == os.ErrNotExist {
		err = nil
		name = project.Name
	}
	var timeout *time.Duration
	if opts.Timeout != nil {
		timeoutStr := strconv.FormatInt(*opts.Timeout, 10)
		timeoutValue, err := time.ParseDuration(timeoutStr + "s")
		timeout = &timeoutValue
		if err != nil {
			return err
		}
	}
	downOptions := api.DownOptions{
		RemoveOrphans: opts.Removeorphans,
		Project:       project,
		Timeout:       timeout,
		Images:        opts.Rmi,
		Volumes:       opts.Volumes,
	}
	return service.Down(name, downOptions)
}

// ProjectEvents - Stream container events for every container in the project
func ProjectEvents(c *gin.Context, opts *models.ProjectEventsOpts) (func(w io.Writer) bool, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return nil, err
	}
	name := opts.Name
	if err == nil {
		name = project.Name
	} else {
		err = nil
	}
	outputChan := make(chan api.Event)
	errorChan := make(chan error)
	var errorIn error
	eventsOptions := api.EventsOptions{
		Services: opts.Services,
		Consumer: func(e api.Event) error {
			select {
			case outputChan <- e:
				return errorIn
			case <-c.Request.Context().Done():
				return errors.Errorf("Client disconnected")
			}
		},
	}
	go func() {
		errorIn = service.Events(name, eventsOptions)
		errorChan <- errorIn
	}()
	time.Sleep(200 * time.Millisecond)
	if err != nil {
		return nil, err
	}
	return func(w io.Writer) bool {
		select {
		case e := <-outputChan:
			event := models.ProjectEvent{
				Timestamp:  e.Timestamp.UnixMilli(),
				Service:    e.Service,
				Container:  e.Container,
				Status:     e.Status,
				Attributes: e.Attributes,
			}
			output, errorIn := json.Marshal(event)
			if errorIn != nil {
				return false
			}
			output = append(output, '\n')
			w.Write(output)
			return true
		case err = <-errorChan:
			if err == nil {
				message := err.Error()
				w.Write([]byte(message))
			}
			return false
		case <-c.Request.Context().Done():
			errorIn = errors.Errorf("Client disconnected")
			return false
		}
	}, err
}

// ProjectImages - List images used by the created containers.
func ProjectImages(c *gin.Context, opts *models.ProjectImagesOpts) ([]models.ProjectImagesResponse, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return nil, err
	}
	name := opts.Name
	if err == nil {
		name = project.Name
	} else {
		err = nil
	}
	images, err := service.Images(name, api.ImagesOptions{
		Services: opts.Services,
	})
	if err != nil {
		return nil, err
	}
	var convImages []models.ProjectImagesResponse
	for _, image := range images {
		convImages = append(convImages, models.ProjectImagesResponse{
			ID:            image.ID,
			ContainerName: image.ContainerName,
			Repository:    image.Repository,
			Tag:           image.Tag,
			Size:          image.Size,
		})
	}
	return convImages, nil
}

// ProjectInspect - Inspect a project
func ProjectInspect(c *gin.Context, opts *models.ProjectInspectOpts) ([]byte, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return nil, err
	}
	return service.Convert(project, api.ConvertOptions{
		Format: opts.Format,
		Output: "",
	})
}

// ProjectKill - Forces running containers to stop by sending a SIGKILL signal.
func ProjectKill(c *gin.Context, opts *models.ProjectKillOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	}
	return service.Kill(project, api.KillOptions{
		Services: opts.Services,
		Signal:   opts.Signal,
	})
}

// ProjectList - List projects
func ProjectList(c *gin.Context, opts *models.ProjectListOpts) (map[string]interface{}, error) {
	clients, err := context.ResolveContexts(opts.Context)
	if err != nil {
		return nil, err
	}
	numObjects := 1
	if opts.Stored {
		numObjects = 2
	}
	output := make(map[string]interface{}, numObjects)

	if opts.Stored {
		stacks, err := project.GetStacks()
		if err != nil {
			return nil, err
		}
		output["Stored"] = stacks
	}

	remoteOutput := make(map[string][]project.Stack)
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(len(clients))
	for context, cli := range clients {
		go func(context string, cli client.APIClient) {
			service := project.GetProjectService(cli, c)
			stacks, err := service.GetActiveStacks(api.ListOptions{
				All: opts.All,
			})
			if err != nil {
				wg.Done()
				return
			}
			mutex.Lock()
			remoteOutput[context] = stacks
			mutex.Unlock()
			wg.Done()
		}(context, cli)
	}
	wg.Wait()
	output["Remote"] = remoteOutput
	return output, nil
}

// ProjectLogs - Get project logs
func ProjectLogs(c *gin.Context, opts *models.ProjectLogsOpts) ([]models.LogObject, func(w io.Writer) bool, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, nil, err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return nil, nil, err
	}
	name := opts.Name
	if err == nil {
		name = project.Name
	} else {
		err = nil
	}
	if opts.Follow {
		response := ProjectLogsStream(c, name, service, opts)
		return nil, response, nil
	}
	response, err := ProjectLogsResponse(name, service, opts)
	return response, nil, err
}

func ProjectLogsStream(c *gin.Context, name string, service project.ProjectService, opts *models.ProjectLogsOpts) func(w io.Writer) bool {
	doneChan := make(chan bool)
	errChan := make(chan error)
	consumer := LogConsumer{
		logChan:          make(chan models.LogObjectLog),
		statusChan:       make(chan models.LogObjectStatus),
		registrationChan: make(chan models.LogObjectRegister),
	}
	go func() {
		err := service.Logs(name, &consumer, api.LogOptions{
			Services:   opts.Services,
			Tail:       opts.Tail,
			Since:      "",
			Until:      "",
			Follow:     opts.Follow,
			Timestamps: opts.Timestamps,
		})
		if err != nil {
			errChan <- err
		}
		doneChan <- true
	}()
	return func(w io.Writer) bool {
		var message models.LogObject
		select {
		case err := <-errChan:
			WriteError(err, w)
			return false
		case <-doneChan:
			return false
		case logItem := <-consumer.logChan:
			message = models.LogObject{
				Log: &logItem,
			}
		case statusItem := <-consumer.statusChan:
			message = models.LogObject{
				Status: &statusItem,
			}
		case registerItem := <-consumer.registrationChan:
			message = models.LogObject{
				Register: &registerItem,
			}
		case <-c.Request.Context().Done():
			return false
		}
		b, err := json.Marshal(message)
		b = append(b, '\n')
		if err != nil {
			WriteError(err, w)
		} else {
			w.Write(b)
		}
		return true
	}
}

func ProjectLogsResponse(name string, service project.ProjectService, opts *models.ProjectLogsOpts) ([]models.LogObject, error) {
	doneChan := make(chan bool)
	errChan := make(chan error)
	consumer := LogConsumer{
		logChan:          make(chan models.LogObjectLog),
		statusChan:       make(chan models.LogObjectStatus),
		registrationChan: make(chan models.LogObjectRegister),
	}
	go func() {
		err := service.Logs(name, &consumer, api.LogOptions{
			Services:   opts.Services,
			Tail:       opts.Tail,
			Since:      "",
			Until:      "",
			Follow:     opts.Follow,
			Timestamps: opts.Timestamps,
		})
		if err != nil {
			errChan <- err
		}
		doneChan <- true
	}()
	var allLogs []models.LogObject
	for {
		select {
		case err := <-errChan:
			return nil, err
		case <-doneChan:
			return allLogs, nil
		case logItem := <-consumer.logChan:
			message := models.LogObject{
				Log: &logItem,
			}
			allLogs = append(allLogs, message)
		case statusItem := <-consumer.statusChan:
			message := models.LogObject{
				Status: &statusItem,
			}
			allLogs = append(allLogs, message)
		case registerItem := <-consumer.registrationChan:
			message := models.LogObject{
				Register: &registerItem,
			}
			allLogs = append(allLogs, message)
		}
	}
}

// ProjectPause - Pauses running containers of a service.
func ProjectPause(c *gin.Context, opts *models.ProjectPauseOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	}
	name := opts.Name
	if err == nil {
		name = project.Name
	}
	return service.Pause(name, api.PauseOptions{
		Services: opts.Services,
	})
}

// ProjectPs - Lists containers.
func ProjectPs(c *gin.Context, opts *models.ProjectPsOpts) (*[]models.ProjectContainerSummaryItem, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return nil, err
	}
	name := opts.Name
	if err == nil {
		name = project.Name
	} else {
		err = nil
	}
	psResponse, err := service.Ps(name, api.PsOptions{
		All:      opts.All,
		Services: opts.Services,
	})
	if err != nil {
		return nil, err
	}

	var response []models.ProjectContainerSummaryItem
	for _, item := range psResponse {
		var publishers []models.PortPublisher
		for _, publisher := range item.Publishers {
			publishers = append(publishers, models.PortPublisher{
				URL:           publisher.URL,
				TargetPort:    publisher.TargetPort,
				PublishedPort: publisher.PublishedPort,
				Protocol:      publisher.Protocol,
			})
		}
		response = append(response, models.ProjectContainerSummaryItem{
			Id:         item.ID,
			Name:       item.Name,
			Command:    item.Command,
			Project:    item.Project,
			Service:    item.Service,
			State:      item.State,
			Health:     item.Health,
			ExitCode:   item.ExitCode,
			Publishers: publishers,
		})
	}
	return &response, err
}

// ProjectPull - Pulls images associated with a service.
func ProjectPull(c *gin.Context, opts *models.ProjectPullOpts) (func(w io.Writer) bool, error) {
	reader := service.RunPull(*opts)
	if opts.Quiet {
		return nil, nil
	}
	return StreamResponse(c, reader), nil
}

// ProjectPush - Pushes images for services to their respective `registry/repository`.
func ProjectPush(c *gin.Context, opts *models.ProjectPushOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	}
	return service.Push(project, api.PushOptions{
		IgnoreFailures: opts.IgnorePushFailures,
	})
}

// ProjectRemove - Removes stopped service containers.
func ProjectRemove(c *gin.Context, opts *models.ProjectRemoveOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	}
	return service.Remove(project, api.RemoveOptions{
		Volumes:  opts.Volumes,
		Force:    true,
		Services: opts.Services,
	})
}

// ProjectRestart - Restarts all stopped and running services.
func ProjectRestart(c *gin.Context, opts *models.ProjectRestartOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	} else if err == os.ErrNotExist {
		err = nil
	}
	var timeout *time.Duration
	if opts.Timeout != nil {
		timeoutStr := strconv.FormatInt(*opts.Timeout, 10)
		timeoutValue, err := time.ParseDuration(timeoutStr + "s")
		timeout = &timeoutValue
		if err != nil {
			return err
		}
	}
	return service.Restart(project, api.RestartOptions{
		Timeout:  timeout,
		Services: opts.Services,
	})
}

// ProjectRun - Runs a one-time command against a service.
func ProjectRun(c *gin.Context, opts *models.ProjectRunOpts) (func(w io.Writer) bool, map[string]interface{}, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, nil, err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return nil, nil, err
	} else if err == os.ErrNotExist {
		err = nil
	}
	if opts.Detach {
		response, err := ProjectRunResponse(c, service, project, opts)
		return nil, response, err
	}
	return ProjectRunStream(c, service, project, opts), nil, err
}

func ProjectRunResponse(c *gin.Context, service project.ProjectService, project *types.Project, opts *models.ProjectRunOpts) (map[string]interface{}, error) {
	stream := newSingleStringReader()
	code, err := service.Run(project, api.RunOptions{
		Name:              opts.ContainerName,
		Service:           opts.Service,
		Command:           opts.Command,
		Entrypoint:        opts.Entrypoint,
		Detach:            opts.Detach,
		AutoRemove:        opts.AutoRemove,
		Stdin:             nil,
		Stdout:            &stream,
		Stderr:            &stream,
		Tty:               opts.Tty,
		WorkingDir:        opts.Workdir,
		User:              opts.User,
		Environment:       opts.Environment,
		Labels:            opts.Labels,
		UseNetworkAliases: opts.UseAliases,
		NoDeps:            !opts.Deps,
		Index:             0,
	})
	containerId := stream.out
	if last := len(containerId) - 1; last >= 0 && containerId[last] == '\n' {
		containerId = containerId[:last]
	}
	output := map[string]interface{}{
		"containerId": containerId,
		"return":      code,
	}
	return output, err
}

type runResponse struct {
	code int
	err  error
}

func ProjectRunStream(c *gin.Context, service project.ProjectService, project *types.Project, opts *models.ProjectRunOpts) func(w io.Writer) bool {
	inReader := newEmptyReader()
	outBuffer := newChannelBuffer()
	errBuffer := newChannelBuffer()
	returnChan := make(chan runResponse)

	go func() {
		code, err := service.Run(project, api.RunOptions{
			Name:              opts.ContainerName,
			Service:           opts.Service,
			Command:           opts.Command,
			Entrypoint:        opts.Entrypoint,
			Detach:            opts.Detach,
			AutoRemove:        opts.AutoRemove,
			Stdin:             &inReader,
			Stdout:            &outBuffer,
			Stderr:            &errBuffer,
			Tty:               opts.Tty,
			WorkingDir:        opts.Workdir,
			User:              opts.User,
			Environment:       opts.Environment,
			Labels:            opts.Labels,
			UseNetworkAliases: opts.UseAliases,
			NoDeps:            !opts.Deps,
			Index:             0,
		})
		returnChan <- runResponse{
			code: code,
			err:  err,
		}
	}()

	return func(w io.Writer) bool {
		var message map[string]interface{}
		done := false
		select {
		case output := <-outBuffer.out:
			message = map[string]interface{}{
				"stdout": output,
			}
		case errorMessage := <-errBuffer.out:
			message = map[string]interface{}{
				"stderr": errorMessage,
			}
		case returnMessage := <-returnChan:
			outBuffer.Close()
			errBuffer.Close()
			message = map[string]interface{}{
				"returnCode": returnMessage.code,
			}
			if returnMessage.err != nil {
				message["error"] = returnMessage.err.Error()
			}
			done = true
		case <-c.Request.Context().Done():
			outBuffer.Close()
			errBuffer.Close()
			return false
		}
		b, err := json.Marshal(message)
		b = append(b, '\n')
		if err != nil {
			WriteError(err, w)
		} else {
			w.Write(b)
		}
		return !done
	}
}

// ProjectStart - Starts existing containers for a service.
func ProjectStart(c *gin.Context, opts *models.ProjectStartOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	}

	if len(opts.Services) > 0 {
		err = project.ForServices(opts.Services)
		if err != nil {
			return err
		}
	}
	return service.Start(project, api.StartOptions{Wait: true})
}

// ProjectStop - Stops running containers without removing them.
func ProjectStop(c *gin.Context, opts *models.ProjectStopOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	}

	if len(opts.Services) > 0 {
		err = project.ForServices(opts.Services)
		if err != nil {
			return err
		}
	}
	var timeout *time.Duration
	if opts.Timeout != nil {
		timeoutStr := strconv.FormatInt(*opts.Timeout, 10)
		timeoutValue, err := time.ParseDuration(timeoutStr + "s")
		timeout = &timeoutValue
		if err != nil {
			return err
		}
	}
	return service.Stop(project, api.StopOptions{
		Timeout:  timeout,
		Services: opts.Services,
	})
}

// ProjectTop - Displays the running processes.
func ProjectTop(c *gin.Context, opts *models.ProjectTopOpts) (*[]models.ProjectTopResponse, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return nil, err
	}
	name := opts.Name
	if err == nil {
		name = project.Name
	} else {
		err = nil
	}
	response, err := service.Top(name, opts.Services)
	if err != nil {
		return nil, err
	}
	var topResponse []models.ProjectTopResponse
	for _, item := range response {
		topResponse = append(topResponse, models.ProjectTopResponse{
			Id:        item.ID,
			Name:      item.Name,
			Titles:    item.Titles,
			Processes: item.Processes,
		})
	}
	return &topResponse, nil
}

// ProjectUnpause - Unpauses paused containers of a service.
func ProjectUnpause(c *gin.Context, opts *models.ProjectUnpauseOpts) error {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return err
	}
	name := opts.Name
	if err == nil {
		name = project.Name
	} else {
		err = nil
	}
	return service.Unpause(name, api.PauseOptions{
		Services: opts.Services,
	})
}

// ProjectUp - Builds, (re)creates, starts, and attaches to containers for a service.
func ProjectUp(c *gin.Context, opts *models.ProjectUpOpts) (func(w io.Writer) bool, error) {
	client, err := context.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	service := project.GetProjectService(client, c)
	project, err := project.GetProject(opts.Name)
	if err != nil && err != os.ErrNotExist {
		return nil, err
	}

	recreate := api.RecreateDiverged
	recreateDependencies := api.RecreateDiverged
	if opts.ForceRecreate {
		recreate = api.RecreateForce
	}
	if opts.AlwaysRecreateDeps {
		recreateDependencies = api.RecreateForce
	}
	if !opts.Recreate {
		recreate = api.RecreateNever
		recreateDependencies = api.RecreateNever
	}
	var timeout *time.Duration
	if opts.Timeout != nil {
		timeoutStr := strconv.FormatInt(*opts.Timeout, 10)
		timeoutValue, err := time.ParseDuration(timeoutStr + "s")
		timeout = &timeoutValue
		if err != nil {
			return nil, err
		}
	}
	upOpts := api.UpOptions{
		Create: api.CreateOptions{
			Services:             opts.Services,
			RemoveOrphans:        opts.RemoveOrphans,
			IgnoreOrphans:        strings.ToLower(project.Environment["COMPOSE_IGNORE_ORPHANS"]) == "true",
			Recreate:             recreate,
			RecreateDependencies: recreateDependencies,
			Inherit:              !opts.RenewAnonVolumes,
			Timeout:              timeout,
			QuietPull:            opts.QuietPull,
		},
		Start: api.StartOptions{
			Attach:       nil,
			AttachTo:     opts.Attach,
			CascadeStop:  opts.AbortOnContainerExit,
			ExitCodeFrom: opts.ExitCodeFrom,
		},
	}

	if opts.Detach {
		err = ProjectUpResponse(service, project, upOpts)
		return nil, err
	}
	return ProjectUpStream(c, service, project, upOpts), nil
}

func ProjectUpResponse(service project.ProjectService, project *types.Project, opts api.UpOptions) error {
	return service.Up(project, opts)
}

func ProjectUpStream(c *gin.Context, service project.ProjectService, project *types.Project, opts api.UpOptions) func(w io.Writer) bool {
	doneChan := make(chan bool)
	errChan := make(chan error)
	consumer := LogConsumer{
		logChan:          make(chan models.LogObjectLog),
		statusChan:       make(chan models.LogObjectStatus),
		registrationChan: make(chan models.LogObjectRegister),
	}
	opts.Start.Attach = &consumer
	go func() {
		err := service.Up(project, opts)
		if err != nil {
			errChan <- err
		}
		doneChan <- true
	}()
	return func(w io.Writer) bool {
		var message models.LogObject
		select {
		case err := <-errChan:
			WriteError(err, w)
			return false
		case <-doneChan:
			return false
		case logItem := <-consumer.logChan:
			message = models.LogObject{
				Log: &logItem,
			}
		case statusItem := <-consumer.statusChan:
			message = models.LogObject{
				Status: &statusItem,
			}
		case registerItem := <-consumer.registrationChan:
			message = models.LogObject{
				Register: &registerItem,
			}
		case <-c.Request.Context().Done():
			return false
		}
		b, err := json.Marshal(message)
		b = append(b, '\n')
		if err != nil {
			WriteError(err, w)
		} else {
			w.Write(b)
		}
		return true
	}
}
