package impl

type LogObject struct {
	Log          *LogMessage      `json:"Log,omitempty"`
	Status       *LogStatus       `json:"Status,omitempty"`
	Registration *LogRegistration `json:"Register,omitempty"`
}

type LogMessage struct {
	Service   string `json:"Service,omitempty"`
	Container string `json:"Container,omitempty"`
	Message   string `json:"Message,omitempty"`
}

type LogStatus struct {
	Container string `json:"Container,omitempty"`
	Message   string `json:"Message,omitempty"`
}

type LogRegistration struct {
	Container string `json:"Container,omitempty"`
}

type LogConsumer struct {
	logChan          chan LogMessage
	statusChan       chan LogStatus
	registrationChan chan LogRegistration
}

func (l *LogConsumer) Log(service, container, message string) {
	l.logChan <- LogMessage{
		Service:   service,
		Container: container,
		Message:   message,
	}
}

func (l *LogConsumer) Status(container, message string) {
	l.statusChan <- LogStatus{
		Container: container,
		Message:   message,
	}
}

func (l *LogConsumer) Register(container string) {
	l.registrationChan <- LogRegistration{
		Container: container,
	}
}
