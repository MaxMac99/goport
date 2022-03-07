package events

type ListenerExistsError struct {
	Listener NotificationListener
}

func (l *ListenerExistsError) Error() string {
	return "The listener already exists"
}

type NoListenersToExcludeError struct {
}

func (l *NoListenersToExcludeError) Error() string {
	return "There are no listeners to exclude"
}
