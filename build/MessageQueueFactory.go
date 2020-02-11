package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
	"github.com/pip-services3-go/pip-services3-messages-go/queues"
)

/**
 * Creates [[MemoryMessageQueue]] components by their descriptors.
 * Name of created message queue is taken from its descriptor.
 *
 * @see [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/classes/build.factory.html Factory]]
 * @see [[MemoryMessageQueue]]
 */
type MessageQueueFactory struct {
	build.Factory
	Descriptor            *cref.Descriptor
	MemoryQueueDescriptor *cref.Descriptor
}

/**
 * Create a new instance of the factory.
 */
func NewMessageQueueFactory() *MessageQueueFactory {
	//super();
	mqf := MessageQueueFactory{}
	mqf.Descriptor = cref.NewDescriptor("pip-services", "factory", "message-queue", "default", "1.0")
	mqf.MemoryQueueDescriptor = cref.NewDescriptor("pip-services", "message-queue", "memory", "*", "1.0")
	mqf.Register(mqf.MemoryQueueDescriptor, func() interface{} {
		return queues.NewMemoryMessageQueue(mqf.MemoryQueueDescriptor.Name())
	})
	return &mqf
}
