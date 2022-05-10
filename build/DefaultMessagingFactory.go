package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	"github.com/pip-services3-go/pip-services3-messaging-go/queues"
)

// DefaultMessagingFactory Creates MemoryMessageQueue components by their descriptors.
// Name of created message queue is taken from its descriptor.
//
// See Factory
// See MemoryMessageQueue
type DefaultMessagingFactory struct {
	cbuild.Factory
}

// NewDefaultMessagingFactory are create a new instance of the factory.
func NewDefaultMessagingFactory() *DefaultMessagingFactory {
	c := DefaultMessagingFactory{}
	c.Factory = *cbuild.NewFactory()

	memoryQueueDescriptor := cref.NewDescriptor("pip-services", "message-queue", "memory", "*", "1.0")
	memoryQueueFactoryDescriptor := cref.NewDescriptor("pip-services", "queue-factory", "memory", "*", "1.0")

	c.Register(memoryQueueDescriptor, func(locator interface{}) interface{} {
		name := ""
		descriptor, ok := locator.(*cref.Descriptor)
		if ok {
			name = descriptor.Name()
		}

		return queues.NewMemoryMessageQueue(name)
	})
	c.RegisterType(memoryQueueFactoryDescriptor, NewMemoryMessageQueueFactory)

	return &c
}
