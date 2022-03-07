package events

import (
	"github.com/docker/docker/api/types/events"
	"github.com/sideshow/apns2/payload"
	"gitlab.com/maxmac99/goport/notifications"
)

func buildNotification(devices []string, contextName string, event events.Message) *notifications.NotificationRequest {
	payload := buildNotificationPayload(contextName, event)
	if payload == nil {
		return nil
	}
	return &notifications.NotificationRequest{
		AppName:     "goport",
		Priority:    notifications.PriorityHigh,
		PushType:    notifications.PushTypeAlert,
		Devices:     devices,
		Payload:     payload,
		Development: true,
	}
}

func buildNotificationPayload(contextName string, event events.Message) *payload.Payload {
	switch event.Type {
	case string(ContainerEventType):
		return buildContainerNotificationPayload(contextName, event)
	default:
		return nil
	}
}

func buildContainerNotificationPayload(contextName string, event events.Message) *payload.Payload {
	payload := payload.NewPayload()
	name := event.Actor.Attributes["name"]
	image := event.Actor.Attributes["image"]
	switch event.Action {
	case "create":
		payload.AlertTitle("Container created").AlertBody("Created container \"" + name + "\" with image \"" + image + "\" on " + contextName)
	case "destroy":
		payload.AlertTitle("Container removed").AlertBody("Removed container \"" + name + "\" with image \"" + image + "\" on " + contextName)
	case "die":
		payload.AlertTitle("Container stopped").AlertBody("Stopped container \"" + name + "\" with image \"" + image + "\" on " + contextName)
	case "pause":
		payload.AlertTitle("Container paused").AlertBody("Paused container \"" + name + "\" with image \"" + image + "\" on " + contextName)
	case "unpause":
		payload.AlertTitle("Container unpaused").AlertBody("Continue container \"" + name + "\" with image \"" + image + "\" on " + contextName)
	case "start":
		payload.AlertTitle("Container started").AlertBody("Started container \"" + name + "\" with image \"" + image + "\" on " + contextName)
	case "stop":
		payload.AlertTitle("Container stopped").AlertBody("Stopped container \"" + name + "\" with image \"" + image + "\" on " + contextName)
	case "restart":
		payload.AlertTitle("Container restarted").AlertBody("Restarted container \"" + name + "\" with image \"" + image + "\" on " + contextName)
	}
	return payload
}
