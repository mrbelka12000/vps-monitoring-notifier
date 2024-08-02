package vps_monitoring_notifier

type serviceName string

const (
	transcripterBot serviceName = "TranscripterBot"
	mognoDB         serviceName = "MongoDB"
	postgresql      serviceName = "PostgreSQL"
	goalsScheduler  serviceName = "GoalsScheduler"
	mockServer      serviceName = "MockServer"
)
