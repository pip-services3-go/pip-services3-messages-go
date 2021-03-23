package build

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
	"github.com/pip-services3-go/pip-services3-messaging-go/queues"
)

// MemoryMessageQueueFactory are creates MemoryMemoryMessageQueue components by their descriptors.
// Name of created message queue is taken from its descriptor.
//
// See Factory
// See MemoryMemoryMessageQueue
type MemoryMessageQueueFactory struct {
	build.Factory
	config     *cconf.ConfigParams
	references cref.IReferences
}

// NewMemoryMessageQueueFactory method are create a new instance of the factory.
func NewMemoryMessageQueueFactory() *MemoryMessageQueueFactory {
	c := MemoryMessageQueueFactory{}

	memoryQueueDescriptor := cref.NewDescriptor("pip-services", "message-queue", "memory", "*", "1.0")

	c.Register(memoryQueueDescriptor, func(locator interface{}) interface{} {
		name := ""
		descriptor, ok := locator.(*cref.Descriptor)
		if ok {
			name = descriptor.Name()
		}

		queue := queues.NewMemoryMessageQueue(name)

		if c.config != nil {
			queue.Configure(c.config)
		}
		if c.references != nil {
			queue.SetReferences(c.references)
		}

		return queue
	})

	return &c
}
