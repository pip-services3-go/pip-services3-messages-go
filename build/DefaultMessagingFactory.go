package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	"github.com/pip-services3-go/pip-services3-messages-go/queues"
)

/**
 * Creates [[MemoryMessageQueue]] components by their descriptors.
 * Name of created message queue is taken from its descriptor.
 *
 * @see [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/classes/build.factory.html Factory]]
 * @see [[MemoryMessageQueue]]
 */
type DefaultMessagingFactory struct {
	cbuild.Factory
	Descriptor            *cref.Descriptor
	MemoryQueueDescriptor *cref.Descriptor
}

/**
 * Create a new instance of the factory.
 */
func NewDefaultMessagingFactory() *DefaultMessagingFactory {
	//super();
	dmf := DefaultMessagingFactory{}
	dmf.Factory = *cbuild.NewFactory()
	dmf.Descriptor = cref.NewDescriptor("pip-services", "factory", "messaging", "default", "1.0")
	dmf.MemoryQueueDescriptor = cref.NewDescriptor("pip-services", "message-queue", "memory", "*", "1.0")

	dmf.Register(dmf.MemoryQueueDescriptor, func() interface{} {
		return queues.NewMemoryMessageQueue(dmf.MemoryQueueDescriptor.Name())
	})
	return &dmf
}
