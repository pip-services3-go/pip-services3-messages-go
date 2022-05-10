package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-messaging-go/queues"
)

// MemoryMessageQueueFactory are creates MemoryMemoryMessageQueue components by their descriptors.
// Name of created message queue is taken from its descriptor.
//
// See Factory
// See MemoryMemoryMessageQueue
type MemoryMessageQueueFactory struct {
	MessageQueueFactory
}

// NewMemoryMessageQueueFactory method are create a new instance of the factory.
func NewMemoryMessageQueueFactory() *MemoryMessageQueueFactory {
	c := MemoryMessageQueueFactory{
		MessageQueueFactory: *InheritMessageQueueFactory(),
	}

	memoryQueueDescriptor := cref.NewDescriptor("pip-services", "message-queue", "memory", "*", "1.0")

	c.Register(memoryQueueDescriptor, func(locator interface{}) interface{} {
		name := ""
		descriptor, ok := locator.(*cref.Descriptor)
		if ok {
			name = descriptor.Name()
		}
		return c.CreateQueue(name)
	})

	return &c
}

// Creates a message queue component and assigns its name.
//
// Parameters:
//   - name: a name of the created message queue.
func (c *MemoryMessageQueueFactory) CreateQueue(name string) queues.IMessageQueue {
	queue := queues.NewMemoryMessageQueue(name)

	if c.Config != nil {
		queue.Configure(c.Config)
	}
	if c.References != nil {
		queue.SetReferences(c.References)
	}

	return queue
}
