package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
	"github.com/pip-services3-go/pip-services3-messaging-go/queues"
)

/*
MessageQueueFactory are creates MemoryMessageQueue components by their descriptors.
Name of created message queue is taken from its descriptor.

See Factory
See MemoryMessageQueue
*/
type MessageQueueFactory struct {
	build.Factory
	Descriptor            *cref.Descriptor
	MemoryQueueDescriptor *cref.Descriptor
}

// NewMessageQueueFactory method are create a new instance of the factory.
func NewMessageQueueFactory() *MessageQueueFactory {
	c := MessageQueueFactory{}
	c.Descriptor = cref.NewDescriptor("pip-services", "factory", "message-queue", "default", "1.0")
	c.MemoryQueueDescriptor = cref.NewDescriptor("pip-services", "message-queue", "memory", "*", "1.0")
	c.Register(c.MemoryQueueDescriptor, func() interface{} {
		return queues.NewMemoryMessageQueue(c.MemoryQueueDescriptor.Name())
	})
	return &c
}
