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
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"gitlab.com/maxmac99/goport/controllers"
	"gitlab.com/maxmac99/goport/models"
)

// ContainerExec - Create an exec instance
func ContainerExec(c *gin.Context, opts *models.ContainerExecOpts) (*models.IdResponse, error) {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	options := types.ExecConfig{
		User:         opts.User,
		Privileged:   opts.Privileged,
		Tty:          opts.Tty,
		AttachStdin:  opts.AttachStdin,
		AttachStderr: opts.AttachStderr,
		AttachStdout: opts.AttachStdout,
		Detach:       opts.Detach,
		DetachKeys:   opts.DetachKeys,
		Env:          opts.Env,
		WorkingDir:   opts.WorkingDir,
		Cmd:          opts.Cmd,
	}
	response, err := client.ContainerExecCreate(c, opts.Id, options)
	if err != nil {
		return nil, err
	}
	return &models.IdResponse{
		Id: response.ID,
	}, nil
}

// ExecInspect - Inspect an exec instance
func ExecInspect(c *gin.Context, opts *models.ExecInspectOpts) (*models.ExecInspectResponse, error) {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return nil, err
	}
	response, err := client.ContainerExecInspect(c, opts.Id)
	if err != nil {
		return nil, err
	}
	return &models.ExecInspectResponse{
		ID:          response.ExecID,
		Running:     response.Running,
		ExitCode:    response.ExitCode,
		ContainerID: response.ContainerID,
		Pid:         response.Pid,
	}, nil
}

// ExecResize - Resize an exec instance
func ExecResize(c *gin.Context, opts *models.ExecResizeOpts) error {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	options := types.ResizeOptions{
		Height: opts.H,
		Width:  opts.W,
	}
	return client.ContainerExecResize(c, opts.Id, options)
}

// ExecStart - Start an exec instance
func ExecStart(c *gin.Context, opts *models.ExecStartOpts) error {
	client, err := controllers.ResolveContext(opts.Context)
	if err != nil {
		return err
	}
	options := types.ExecStartCheck{
		Detach: opts.Detach,
		Tty:    opts.Tty,
	}
	return client.ContainerExecStart(c, opts.Id, options)
}
