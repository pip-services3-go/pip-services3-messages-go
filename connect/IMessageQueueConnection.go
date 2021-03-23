package connect

// IMessageQueueConnection Interface for queue connections
type IMessageQueueConnection interface {
	GetQueueNames() ([]string, error)
}
