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
	"net/http"

	"github.com/docker/docker/errdefs"
	"github.com/gin-gonic/gin"
	"gitlab.com/maxmac99/goport/impl"
	"gitlab.com/maxmac99/goport/models"
)

// ContainerChanges - Get changes on a container’s filesystem
func ContainerChangesHandler(c *gin.Context) {
	var opts models.ContainerChangesOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerChanges(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ContainerCreate - Create a container
func ContainerCreateHandler(c *gin.Context) {
	var opts models.ContainerCreateOpts
	if err := c.ShouldBindQuery(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerCreate(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.JSON(201, response)
		return
	}
}

// ContainerDelete - Remove a container
func ContainerDeleteHandler(c *gin.Context) {
	var opts models.ContainerDeleteOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerDelete(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(204)
}

// ContainerExport - Export a container
func ContainerExportHandler(c *gin.Context) {
	var opts models.ContainerExportOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerExport(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.Header("Content-Type", "application/json")
		c.Writer.Flush()
		c.Stream(response)
		return
	}
}

// ContainerInspect - Inspect a container
func ContainerInspectHandler(c *gin.Context) {
	var opts models.ContainerInspectOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerInspect(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ContainerKill - Kill a container
func ContainerKillHandler(c *gin.Context) {
	var opts models.ContainerKillOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerKill(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(204)
}

// ContainerList - List containers
func ContainerListHandler(c *gin.Context) {
	var opts models.ContainerListOpts
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerList(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ContainerLogs - Get container logs
func ContainerLogsHandler(c *gin.Context) {
	var opts models.ContainerLogsOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerLogs(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.Header("Content-Type", "application/json")
		c.Writer.Flush()
		c.Stream(response)
		return
	}
}

// ContainerPause - Pause a container
func ContainerPauseHandler(c *gin.Context) {
	var opts models.ContainerPauseOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerPause(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(204)
}

// ContainerPrune - Delete stopped containers
func ContainerPruneHandler(c *gin.Context) {
	var opts models.ContainerPruneOpts
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerPrune(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ContainerRename - Rename a container
func ContainerRenameHandler(c *gin.Context) {
	var opts models.ContainerRenameOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerRename(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(204)
}

// ContainerResize - Resize a container TTY
func ContainerResizeHandler(c *gin.Context) {
	var opts models.ContainerResizeOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerResize(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{})
}

// ContainerRestart - Restart a container
func ContainerRestartHandler(c *gin.Context) {
	var opts models.ContainerRestartOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerRestart(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(204)
}

// ContainerStart - Start a container
func ContainerStartHandler(c *gin.Context) {
	var opts models.ContainerStartOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerStart(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(204)
}

// ContainerStats - Get container stats based on resource usage
func ContainerStatsHandler(c *gin.Context) {
	var opts models.ContainerStatsOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, stream, err := impl.ContainerStats(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.Data(200, "application/json", response)
		return
	}
	if stream != nil {
		c.Header("Content-Type", "application/json")
		c.Writer.Flush()
		c.Stream(stream)
		return
	}
}

// ContainerStop - Stop a container
func ContainerStopHandler(c *gin.Context) {
	var opts models.ContainerStopOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerStop(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(204)
}

// ContainerTop - List processes running inside a container
func ContainerTopHandler(c *gin.Context) {
	var opts models.ContainerTopOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerTop(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ContainerUnpause - Unpause a container
func ContainerUnpauseHandler(c *gin.Context) {
	var opts models.ContainerUnpauseOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContainerUnpause(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(204)
}

// ContainerUpdate - Update a container
func ContainerUpdateHandler(c *gin.Context) {
	var opts models.ContainerUpdateOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBindQuery(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerUpdate(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}

// ContainerWait - Wait for a container
func ContainerWaitHandler(c *gin.Context) {
	var opts models.ContainerWaitOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContainerWait(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if response != nil {
		c.JSON(200, response)
		return
	}
}
