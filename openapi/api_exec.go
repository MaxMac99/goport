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

// ContainerExec - Create an exec instance
func ContainerExecHandler(c *gin.Context) {
	var opts models.ContainerExecOpts
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
	response, err := impl.ContainerExec(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	if response != nil {
		c.JSON(201, response)
		return
	}
}

// ExecInspect - Inspect an exec instance
func ExecInspectHandler(c *gin.Context) {
	var opts models.ExecInspectOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ExecInspect(c, &opts)
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

// ExecResize - Resize an exec instance
func ExecResizeHandler(c *gin.Context) {
	var opts models.ExecResizeOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ExecResize(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(201, gin.H{})
}

// ExecStart - Start an exec instance
func ExecStartHandler(c *gin.Context) {
	var opts models.ExecStartOpts
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
	err := impl.ExecStart(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.Status(204)
}
