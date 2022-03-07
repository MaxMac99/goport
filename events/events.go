package events

import (
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	dockercontext "gitlab.com/maxmac99/goport/context"
)

func Setup() {
	setupDatabase()
	listenToEvents()
}

func listenToEvents() {
	clients, err := dockercontext.ResolveContexts([]string{"all"})
	if err != nil {
		return
	}
	for context, cli := range clients {
		go listenToContextEvents(context, cli)
	}
}

func listenToContextEvents(contextName string, client client.APIClient) {
	eventC, errC := client.Events(ctx, types.EventsOptions{})
	for {
		select {
		case event, ok := <-eventC:
			if !ok {
				log.Println("Events for context \"" + contextName + "\" closed")
				return
			}
			receivedEvent(contextName, event)
		case err, ok := <-errC:
			if !ok {
				log.Println("Error Events for context \"" + contextName + "\" closed")
				return
			}
			log.Println("Received error from context \"" + contextName + "\": " + err.Error())
		}
	}
}
