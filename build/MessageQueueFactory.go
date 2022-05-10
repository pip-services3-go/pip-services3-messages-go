package build

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/build"
)

// MessageQueueFactory are creates MessageQueue components by their descriptors.
// Name of created message queue is taken from its descriptor.
//
// See Factory
// See MessageQueue
type MessageQueueFactory struct {
	build.Factory
	Config     *cconf.ConfigParams
	References cref.IReferences
}

// NewMessageQueueFactory method creates a new instance of the factory.
func InheritMessageQueueFactory() *MessageQueueFactory {
	c := MessageQueueFactory{
		Factory: *build.NewFactory(),
	}
	return &c
}

func (c *MessageQueueFactory) Configure(config *cconf.ConfigParams) {
	c.Config = config
}

func (c *MessageQueueFactory) SetReferences(references cref.IReferences) {
	c.References = references
}
