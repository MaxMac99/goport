package events

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventType string

const (
	AllEventType       EventType = "all"
	ContainerEventType EventType = "container"
	ImageEventType     EventType = "image"
	VolumeEventType    EventType = "volume"
	NetworkEventType   EventType = "network"
	BuilderEventType   EventType = "builder"
	DaemonEventType    EventType = "daemon"
)

type registerNotificationRequest struct {
	DeviceToken string  `uri:"token" binding:"required"`
	EventType   string  `uri:"event" binding:"required"`
	Context     *string `form:"context"`
	Action      *string `form:"action"`
	Id          *string `form:"id"`
}

func RegisterNotificationHandler(c *gin.Context) {
	var request registerNotificationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := RegisterNotificationListener(request.DeviceToken, request.EventType, request.Context, request.Action, request.Id)
	if err != nil {
		switch e := err.(type) {
		case *ListenerExistsError:
			fmt.Println("Done with error: ", e)
			c.AbortWithError(http.StatusBadRequest, e)
			return
		default:
			fmt.Println("Done with error: ", e)
			c.AbortWithError(http.StatusInternalServerError, e)
			return
		}
	}
	c.Status(204)
}

func UnregisterNotificationHandler(c *gin.Context) {
	var request registerNotificationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := UnregisterNotificationListener(request.DeviceToken, request.EventType, request.Context, request.Action, request.Id)
	if err != nil {
		switch e := err.(type) {
		case *ListenerExistsError:
			fmt.Println("Done with error: ", e)
			c.AbortWithError(http.StatusBadRequest, e)
			return
		case *NoListenersToExcludeError:
			fmt.Println("Done with error: ", e)
			c.AbortWithError(http.StatusBadRequest, e)
			return
		default:
			fmt.Println("Done with error: ", e)
			c.AbortWithError(http.StatusInternalServerError, e)
			return
		}
	}
	c.Status(204)
}

func GetNotificationsHandler(c *gin.Context) {
	var request registerNotificationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	listeners, err := GetNotificationListenersWithToken(request.DeviceToken, request.EventType, request.Context, request.Action, request.Id)
	if err != nil {
		fmt.Println("Done with error: ", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, listeners)
}
