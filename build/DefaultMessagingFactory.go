package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	"github.com/pip-services3-go/pip-services3-messaging-go/queues"
)

/*
DefaultMessagingFactory Creates MemoryMessageQueue components by their descriptors.
Name of created message queue is taken from its descriptor.
See Factory
See MemoryMessageQueue
*/
type DefaultMessagingFactory struct {
	cbuild.Factory
	Descriptor            *cref.Descriptor
	MemoryQueueDescriptor *cref.Descriptor
}

// NewDefaultMessagingFactory are create a new instance of the factory.
func NewDefaultMessagingFactory() *DefaultMessagingFactory {
	c := DefaultMessagingFactory{}
	c.Factory = *cbuild.NewFactory()
	c.Descriptor = cref.NewDescriptor("pip-services", "factory", "messaging", "default", "1.0")
	c.MemoryQueueDescriptor = cref.NewDescriptor("pip-services", "message-queue", "memory", "*", "1.0")

	c.Register(c.MemoryQueueDescriptor, func() interface{} {
		return queues.NewMemoryMessageQueue(c.MemoryQueueDescriptor.Name())
	})
	return &c
}
