package impl

type log struct {
	service   string
	container string
	message   string
}

type status struct {
	container string
	message   string
}

type registration struct {
	container string
}

type logConsumer struct {
	logChan          chan log
	statusChan       chan status
	registrationChan chan registration
}

func (l *logConsumer) Log(service, container, message string) {
	l.logChan <- log{
		service:   service,
		container: container,
		message:   message,
	}
}

func (l *logConsumer) Status(container, message string) {
	l.statusChan <- status{
		container: container,
		message:   message,
	}
}

func (l *logConsumer) Register(container string) {
	l.registrationChan <- registration{
		container: container,
	}
}
