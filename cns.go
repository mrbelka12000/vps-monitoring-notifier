package vps_monitoring_notifier

type serviceName string

const (
	transcripterBot serviceName = "TranscripterBot"
	mognoDB         serviceName = "MongoDB"
	postgreSQL      serviceName = "PostgreSQL"
	goalsScheduler  serviceName = "GoalsScheduler"
	mockServer      serviceName = "MockServer"
	redisDB         serviceName = "Redis"
)
