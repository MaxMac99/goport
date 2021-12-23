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
	"github.com/maxmac99/goport/impl"
	"github.com/maxmac99/goport/models"
)

// ContextCreate - Create a context
func ContextCreateHandler(c *gin.Context) {
	var opts models.ContextCreateOpts
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContextCreate(c, &opts)
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

// ContextDelete - Remove a context
func ContextDeleteHandler(c *gin.Context) {
	var opts models.ContextDeleteOpts
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContextDelete(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}

// ContextInspect - Inspect a context
func ContextInspectHandler(c *gin.Context) {
	var opts models.ContextInspectOpts
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.ContextInspect(c, &opts)
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

// ContextList - List contexts
func ContextListHandler(c *gin.Context) {
	response, err := impl.ContextList(c)
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

// ContextUpdate - Update a context
func ContextUpdateHandler(c *gin.Context) {
	var opts models.ContextUpdateOpts
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.ContextUpdate(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(204, gin.H{})
}
