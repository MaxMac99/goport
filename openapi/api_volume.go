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

// VolumeCreate - Create a volume
func VolumeCreateHandler(c *gin.Context) {
	var opts models.VolumeCreateOpts
	if err := c.ShouldBindQuery(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.VolumeCreate(c, &opts)
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

// VolumeDelete - Remove a volume
func VolumeDeleteHandler(c *gin.Context) {
	var opts models.VolumeDeleteOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	err := impl.VolumeDelete(c, &opts)
	if err != nil {
		code := errdefs.GetHTTPErrorStatusCode(err)
		c.JSON(code, err.Error())
		return
	}
	c.Status(204)
}

// VolumeInspect - Inspect a volume
func VolumeInspectHandler(c *gin.Context) {
	var opts models.VolumeInspectOpts
	if err := c.ShouldBindUri(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.VolumeInspect(c, &opts)
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

// VolumeList - List volumes
func VolumeListHandler(c *gin.Context) {
	var opts models.VolumeListOpts
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.VolumeList(c, &opts)
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

// VolumePrune - Delete unused volumes
func VolumePruneHandler(c *gin.Context) {
	var opts models.VolumePruneOpts
	if err := c.ShouldBind(&opts); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	response, err := impl.VolumePrune(c, &opts)
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
