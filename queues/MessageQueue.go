package queues

import (
	"sync"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/auth"
	ccon "github.com/pip-services3-go/pip-services3-components-go/connect"
	ccount "github.com/pip-services3-go/pip-services3-components-go/count"
	clog "github.com/pip-services3-go/pip-services3-components-go/log"
)

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
	IMessageQueue
	Logger             *clog.CompositeLogger
	Counters           *ccount.CompositeCounters
	ConnectionResolver *ccon.ConnectionResolver
	CredentialResolver *auth.CredentialResolver
	Name               string
	Capabilities       *MessagingCapabilities
}

// NewMessageQueue method are creates a new instance of the message queue.
//   - name  (optional) a queue name
func NewMessageQueue(name string) *MessageQueue {
	c := MessageQueue{Name: name}
	c.Logger = clog.NewCompositeLogger()
	c.Counters = ccount.NewCompositeCounters()
	c.ConnectionResolver = ccon.NewEmptyConnectionResolver()
	c.CredentialResolver = auth.NewEmptyCredentialResolver()
	return &c
}

// GetName method are gets the queue name
// Return the queue name.
func (c *MessageQueue) GetName() string { return c.Name }

// GetCapabilities method are gets the queue capabilities
// Return the queue's capabilities object.
func (c *MessageQueue) GetCapabilities() MessagingCapabilities { return *c.Capabilities }

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
func (c *MessageQueue) Open(correlationId string) (err error) {
	var connection *ccon.ConnectionParams
	var credential *auth.CredentialParams

	wg := sync.WaitGroup{}
	var conErr, credErr error

	wg.Add(2)
	go func() {
		result, err := c.ConnectionResolver.Resolve(correlationId)
		connection = result
		conErr = err
		wg.Done()
	}()

	go func() {
		result, err := c.CredentialResolver.Lookup(correlationId)
		credential = result
		credErr = err
		wg.Done()
	}()
	wg.Wait()
	if conErr != nil {
		return conErr
	}
	if credErr != nil {
		return credErr
	}

	return c.OpenWithParams(correlationId, connection, credential)
}

// SendAsObject method are sends an object into the queue.
// Before sending the object is converted into JSON string and wrapped in a MessageEnvelop.
//   - correlationId     (optional) transaction id to trace execution through call chain.
//   - messageType       a message type
//   - value             an object value to be sent
// Returns: error or null for success.
// See Send
func (c *MessageQueue) SendAsObject(correlationId string, messageType string, message interface{}) (err error) {
	envelope := NewMessageEnvelope(correlationId, messageType, "")
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
		c.Listen(correlationId, receiver)
	}()
}

// ToString method are gets a string representation of the object.
// Return a string representation of the object.
func (c *MessageQueue) ToString() string {
	return "[" + c.GetName() + "]"
}
