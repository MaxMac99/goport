/*
 * GoPort API
 *
 * The GoPort API extends the Docker Engine API to connect to remote Hosts by serving a context endpoint. It also adds the ability to manage docker-compose projects.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/errdefs"
	"github.com/gin-gonic/gin"
	"gitlab.com/maxmac99/goport/impl"
	"gitlab.com/maxmac99/goport/models"
)

// ProjectBuild - Build the project
func ProjectBuildHandler(c *gin.Context) {
	var opts models.ProjectBuildOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	responseStream, err := impl.ProjectBuild(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if responseStream != nil {
		c.Stream(responseStream)
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectCreate - Create a project
func ProjectCreateHandler(c *gin.Context) {
	var opts models.ProjectCreateOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	rawData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	opts.Body = rawData
	opts.Format = "json"
	if c.ContentType() == "x-yaml" {
		opts.Format = "yaml"
	}
	err = impl.ProjectCreate(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectDown - Stops containers and removes containers, networks, volumes, and images created by up
func ProjectDownHandler(c *gin.Context) {
	var opts models.ProjectDownOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectDown(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectEvents - Stream container events for every container in the project
func ProjectEventsHandler(c *gin.Context) {
	var opts models.ProjectEventsOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	responseStream, err := impl.ProjectEvents(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if responseStream != nil {
		c.Stream(responseStream)
		return
	}
}

// ProjectImages - List images used by the created containers.
func ProjectImagesHandler(c *gin.Context) {
	var opts models.ProjectImagesOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ProjectImages(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ProjectInspect - Inspect a project
func ProjectInspectHandler(c *gin.Context) {
	var opts models.ProjectInspectOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ProjectInspect(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if response != nil {
		c.Writer.Write(response)
		return
	}
}

// ProjectKill - Forces running containers to stop by sending a SIGKILL signal.
func ProjectKillHandler(c *gin.Context) {
	var opts models.ProjectKillOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectKill(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectList - List projects
func ProjectListHandler(c *gin.Context) {
	response, err := impl.ProjectList(c)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ProjectLogs - Get project logs
func ProjectLogsHandler(c *gin.Context) {
	var opts models.ProjectLogsOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, stream, err := impl.ProjectLogs(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if stream != nil {
		c.Stream(stream)
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ProjectPause - Pauses running containers of a service.
func ProjectPauseHandler(c *gin.Context) {
	var opts models.ProjectPauseOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectPause(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectPs - Lists containers.
func ProjectPsHandler(c *gin.Context) {
	var opts models.ProjectPsOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ProjectPs(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ProjectPull - Pulls images associated with a service.
func ProjectPullHandler(c *gin.Context) {
	var opts models.ProjectPullOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	responseStream, err := impl.ProjectPull(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if responseStream != nil {
		c.Stream(responseStream)
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectPush - Pushes images for services to their respective `registry/repository`.
func ProjectPushHandler(c *gin.Context) {
	var opts models.ProjectPushOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectPush(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectRemove - Removes stopped service containers.
func ProjectRemoveHandler(c *gin.Context) {
	var opts models.ProjectRemoveOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectRemove(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectRestart - Restarts all stopped and running services.
func ProjectRestartHandler(c *gin.Context) {
	var opts models.ProjectRestartOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectRestart(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectRun - Runs a one-time command against a service.
func ProjectRunHandler(c *gin.Context) {
	var opts models.ProjectRunOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	responseStream, response, err := impl.ProjectRun(c, &opts)
	if err != nil {
		errCode := errdefs.GetHTTPErrorStatusCode(err)
		if response == nil {
			c.JSON(errCode, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(errCode, gin.H{
				"output": response,
				"error":  err.Error(),
			})
		}
		return
	}
	if responseStream != nil {
		c.Stream(responseStream)
		return
	}
	c.JSON(200, gin.H{
		"output": response,
	})
}

// ProjectStart - Starts existing containers for a service.
func ProjectStartHandler(c *gin.Context) {
	var opts models.ProjectStartOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectStart(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectStop - Stops running containers without removing them.
func ProjectStopHandler(c *gin.Context) {
	var opts models.ProjectStopOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectStop(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectTop - Displays the running processes.
func ProjectTopHandler(c *gin.Context) {
	var opts models.ProjectTopOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ProjectTop(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ProjectUnpause - Unpauses paused containers of a service.
func ProjectUnpauseHandler(c *gin.Context) {
	var opts models.ProjectUnpauseOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ProjectUnpause(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ProjectUp - Builds, (re)creates, starts, and attaches to containers for a service.
func ProjectUpHandler(c *gin.Context) {
	var opts models.ProjectUpOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	responseStream, err := impl.ProjectUp(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if responseStream != nil {
		c.Stream(responseStream)
		return
	}
	c.JSON(204, gin.H{})
}
