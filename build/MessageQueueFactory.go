package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
	"github.com/pip-services3-go/pip-services3-messaging-go/queues"
)

// MessageQueueFactory are creates MemoryMessageQueue components by their descriptors.
// Name of created message queue is taken from its descriptor.
//
// See Factory
// See MemoryMessageQueue
type MessageQueueFactory struct {
	build.Factory
}

// NewMessageQueueFactory method are create a new instance of the factory.
func NewMessageQueueFactory() *MessageQueueFactory {
	c := MessageQueueFactory{}

	memoryQueueDescriptor := cref.NewDescriptor("pip-services", "message-queue", "memory", "*", "1.0")

	c.Register(memoryQueueDescriptor, func(locator interface{}) interface{} {
		name := ""
		descriptor, ok := locator.(*cref.Descriptor)
		if ok {
			name = descriptor.Name()
		}

		return queues.NewMemoryMessageQueue(name)
	})

	return &c
}
