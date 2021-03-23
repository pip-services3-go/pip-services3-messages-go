package queues

import (
	"sync"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cauth "github.com/pip-services3-go/pip-services3-components-go/auth"
	ccon "github.com/pip-services3-go/pip-services3-components-go/connect"
	ccount "github.com/pip-services3-go/pip-services3-components-go/count"
	clog "github.com/pip-services3-go/pip-services3-components-go/log"
)

type IMessageQueueOverrides interface {
	IMessageQueue

	// OpenWithParams method are opens the component with given connection and credential parameters.
	//  - correlationId     (optional) transaction id to trace execution through call chain.
	//  - connections        connection parameters
	//  - credential        credential parameters
	// Returns error or nil no errors occured.
	OpenWithParams(correlationId string, connections []*ccon.ConnectionParams, credential *cauth.CredentialParams) error
}

/*
MessageQueue message queue that is used as a basis for specific message queue implementations.

Configuration parameters:

  - name:                        name of the message queue
  - connection(s):
    - discovery_key:             key to retrieve parameters from discovery service
    - protocol:                  connection protocol like http, https, tcp, udp
    - host:                      host name or IP address
    - port:                      port number
    - uri:                       resource URI or connection string with all parameters in it
  - credential(s):
    - store_key:                 key to retrieve parameters from credential store
    - username:                  user name
    - password:                  user password
    - access_id:                 application access id
    - access_key:                application secret key

References:

- *:Logger:*:*:1.0           (optional)  ILogger components to pass log messages
- *:Counters:*:*:1.0         (optional)  ICounters components to pass collected measurements
- *:discovery:*:*:1.0        (optional)  IDiscovery components to discover connection(s)
- *:credential-store:*:*:1.0 (optional)  ICredentialStore componetns to lookup credential(s)
*/
type MessageQueue struct {
	IMessageQueueOverrides
	Logger             *clog.CompositeLogger
	Counters           *ccount.CompositeCounters
	ConnectionResolver *ccon.ConnectionResolver
	CredentialResolver *cauth.CredentialResolver
	Name               string
	Capabilities       *MessagingCapabilities
	Lock               sync.Mutex
}

// NewMessageQueue method are creates a new instance of the message queue.
//   - overrides a message queue overrides
//   - name  (optional) a queue name
func InheritMessageQueue(overrides IMessageQueueOverrides, name string) *MessageQueue {
	c := MessageQueue{
		IMessageQueueOverrides: overrides,
		Name:                   name,
	}
	c.Logger = clog.NewCompositeLogger()
	c.Counters = ccount.NewCompositeCounters()
	c.ConnectionResolver = ccon.NewEmptyConnectionResolver()
	c.CredentialResolver = cauth.NewEmptyCredentialResolver()
	c.Capabilities = NewMessagingCapabilities(false, false, false, false, false, false, false, false, false)
	return &c
}

// GetName method are gets the queue name
// Return the queue name.
func (c *MessageQueue) GetName() string {
	return c.Name
}

// GetCapabilities method are gets the queue capabilities
// Return the queue's capabilities object.
func (c *MessageQueue) GetCapabilities() *MessagingCapabilities {
	return c.Capabilities
}

// Configure method are configures component by passing configuration parameters.
//   - config    configuration parameters to be set.
func (c *MessageQueue) Configure(config *cconf.ConfigParams) {
	c.Name = cconf.NameResolver.ResolveWithDefault(config, c.Name)
	c.Logger.Configure(config)
	c.ConnectionResolver.Configure(config)
	c.CredentialResolver.Configure(config)
}

// SetReferences mmethod are sets references to dependent components.
//   - references 	references to locate the component dependencies.
func (c *MessageQueue) SetReferences(references cref.IReferences) {
	c.Logger.SetReferences(references)
	c.Counters.SetReferences(references)
	c.ConnectionResolver.SetReferences(references)
	c.CredentialResolver.SetReferences(references)
}

// Open method are opens the component.
//   - correlationId 	(optional) transaction id to trace execution through call chain.
// Returns: error or null no errors occured.
func (c *MessageQueue) Open(correlationId string) error {
	connections, err := c.ConnectionResolver.ResolveAll(correlationId)
	if err != nil {
		return err
	}
	if len(connections) == 0 {
		err = cerr.NewConfigError(correlationId, "NO_CONNECTION", "Connection parameters are not set")
		return err
	}

	credential, err := c.CredentialResolver.Lookup(correlationId)
	if err != nil {
		return err
	}

	return c.OpenWithParams(correlationId, connections, credential)
}

// Checks if message queue has been opened
//   - correlationId     (optional) transaction id to trace execution through call chain.
// Returns: error or null for success.
func (c *MessageQueue) CheckOpen(correlationId string) error {
	if !c.IsOpen() {
		err := cerr.NewInvalidStateError(
			correlationId,
			"NOT_OPENED",
			"The queue is not opened",
		)
		return err
	}
	return nil
}

// SendAsObject method are sends an object into the queue.
// Before sending the object is converted into JSON string and wrapped in a MessageEnvelop.
//   - correlationId     (optional) transaction id to trace execution through call chain.
//   - messageType       a message type
//   - value             an object value to be sent
// Returns: error or null for success.
// See Send
func (c *MessageQueue) SendAsObject(correlationId string, messageType string, message interface{}) (err error) {
	envelope := NewMessageEnvelope(correlationId, messageType, nil)
	envelope.SetMessageAsJson(message)
	return c.Send(correlationId, envelope)
}

// BeginListen method are listens for incoming messages without blocking the current thread.
//   - correlationId     (optional) transaction id to trace execution through call chain.
//   - receiver          a receiver to receive incoming messages.
// See Listen
// See IMessageReceiver
func (c *MessageQueue) BeginListen(correlationId string, receiver IMessageReceiver) {
	go func() {
		err := c.Listen(correlationId, receiver)
		if err != nil {
			c.Logger.Error(correlationId, err, "Failed to listed the message queue "+c.Name)
		}
	}()
}

// String method are gets a string representation of the object.
// Return a string representation of the object.
func (c *MessageQueue) String() string {
	return "[" + c.GetName() + "]"
}
