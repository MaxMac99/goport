package events

import (
	"log"

	"github.com/docker/docker/api/types/events"
	"gitlab.com/maxmac99/goport/notifications"
)

func receivedEvent(contextName string, event events.Message) {
	if event.Action == "destroy" {
		DeleteNotificationRegistrationForId(event.Actor.ID)
	}
	deviceTokens, err := GetDeviceTokensForEvent(contextName, event.Type, event.Action, event.Actor.ID)
	if err != nil {
		log.Println("Could not get deviceTokens: " + err.Error())
		return
	}
	if len(deviceTokens) == 0 {
		return
	}
	notification := buildNotification(deviceTokens, contextName, event)
	if notification == nil {
		return
	}
	notifications.SendNotification(*notification)
}
