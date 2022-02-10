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
	DeviceToken string   `uri:"token" binding:"required"`
	EventType   string   `uri:"event" binding:"required"`
	Context     []string `form:"context"`
	Action      []string `form:"action"`
	Ids         []string `form:"ids"`
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
	fmt.Println("Register notification: ", request)
	if err := RegisterNotificationListener(NotificationListener{
		DeviceToken: request.DeviceToken,
		EventType:   request.EventType,
		Contexts:    request.Context,
		Actions:     request.Action,
		Ids:         request.Ids,
	}); err != nil {
		fmt.Println("Done with error: ", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		fmt.Println("Success")
		c.Status(204)
	}
}
