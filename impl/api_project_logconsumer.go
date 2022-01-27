package impl

import "gitlab.com/maxmac99/goport/models"

type LogConsumer struct {
	logChan          chan models.LogObjectLog
	statusChan       chan models.LogObjectStatus
	registrationChan chan models.LogObjectRegister
}

func (l *LogConsumer) Log(service, container, message string) {
	l.logChan <- models.LogObjectLog{
		Service:   service,
		Container: container,
		Message:   message,
	}
}

func (l *LogConsumer) Status(container, message string) {
	l.statusChan <- models.LogObjectStatus{
		Container: container,
		Message:   message,
	}
}

func (l *LogConsumer) Register(container string) {
	l.registrationChan <- models.LogObjectRegister{
		Container: container,
	}
}
