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
 message queue that is used as a basis for specific message queue implementations.

 Configuration parameters

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

 References

- *:Logger:*:*:1.0           (optional)  ILogger components to pass log messages
- *:Counters:*:*:1.0         (optional)  ICounters components to pass collected measurements
- *:discovery:*:*:1.0        (optional)  IDiscovery components to discover connection(s)
- *:credential-store:*:*:1.0 (optional)  ICredentialStore componetns to lookup credential(s)
*/
// implements IMessageQueue, IReferenceable, IConfigurable
type MessageQueue struct {
	IMessageQueue
	Logger             *clog.CompositeLogger
	Counters           *ccount.CompositeCounters
	ConnectionResolver *ccon.ConnectionResolver
	CredentialResolver *auth.CredentialResolver

	Name         string
	Capabilities *MessagingCapabilities
}

/*
Creates a new instance of the message queue.
 *
- name  (optional) a queue name
*/
func NewMessageQueue(name string) *MessageQueue {
	mq := MessageQueue{Name: name}
	mq.Logger = clog.NewCompositeLogger()
	mq.Counters = ccount.NewCompositeCounters()
	mq.ConnectionResolver = ccon.NewEmptyConnectionResolver()
	mq.CredentialResolver = auth.NewEmptyCredentialResolver()
	return &mq
}

/*
Gets the queue name
 *
Return the queue name.
*/
func (c *MessageQueue) GetName() string { return c.Name }

/*
Gets the queue capabilities
 *
Return the queue's capabilities object.
*/
func (c *MessageQueue) GetCapabilities() MessagingCapabilities { return *c.Capabilities }

/*
Configures component by passing configuration parameters.
 *
- config    configuration parameters to be set.
*/
func (c *MessageQueue) Configure(config *cconf.ConfigParams) {
	c.Name = cconf.NameResolver.ResolveWithDefault(config, c.Name)
	c.Logger.Configure(config)
	c.ConnectionResolver.Configure(config)
	c.CredentialResolver.Configure(config)
}

/*
Sets references to dependent components.
 *
- references 	references to locate the component dependencies.
*/
func (c *MessageQueue) SetReferences(references cref.IReferences) {
	c.Logger.SetReferences(references)
	c.Counters.SetReferences(references)
	c.ConnectionResolver.SetReferences(references)
	c.CredentialResolver.SetReferences(references)
}

/*
Checks if the component is opened.
 *
Return true if the component has been opened and false otherwise.
*/
// func (c *MessageQueue) IsOpen() bool {
// 	return true
// }

/*
	Opens the component.
	 *
	- correlationId 	(optional) transaction id to trace execution through call chain.
    - callback 			callback function that receives error or null no errors occured.
*/
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

/*
Opens the component with given connection and credential parameters.

- correlationId     (optional) transaction id to trace execution through call chain.
- connection        connection parameters
- credential        credential parameters
- callback 			callback function that receives error or null no errors occured.
*/

// func (c *MessageQueue) openWithParams(correlationId string,
// 	connection *ccon.ConnectionParams, credential *auth.CredentialParams) error

/*
	Closes component and frees used resources.

	- correlationId 	(optional) transaction id to trace execution through call chain.
    - callback 			callback function that receives error or null no errors occured.
*/
// func (c* MessageQueue) Close(correlationId string) (err error)

/*
	Clears component state.

	- correlationId 	(optional) transaction id to trace execution through call chain.
    - callback 			callback function that receives error or null no errors occured.
*/
//func (c* MessageQueue)  Clear(correlationId string) (err error)

/*
Reads the current number of messages in the queue to be delivered.

- callback      callback function that receives number of messages or error.
*/
//func (c* MessageQueue)  ReadMessageCount()(count int64, err error)

/*
Sends a message into the queue.

- correlationId     (optional) transaction id to trace execution through call chain.
- envelope          a message envelop to be sent.
- callback          (optional) callback function that receives error or null for success.
*/
//func (c* MessageQueue)  Send(correlationId string, envelope *MessageEnvelope) (err error)

/*
Sends an object into the queue.
Before sending the object is converted into JSON string and wrapped in a MessageEnvelop.

- correlationId     (optional) transaction id to trace execution through call chain.
- messageType       a message type
- value             an object value to be sent
- callback          (optional) callback function that receives error or null for success.
 *
See send
*/
func (c *MessageQueue) SendAsObject(correlationId string, messageType string, message interface{}) (err error) {
	envelope := NewMessageEnvelope(correlationId, messageType, message)
	return c.Send(correlationId, envelope)
}

/*
Peeks a single incoming message from the queue without removing it.
If there are no messages available in the queue it returns null.

- correlationId     (optional) transaction id to trace execution through call chain.
- callback          callback function that receives a message or error.
*/
//  func (c* MessageQueue)  Peek(correlationId string) callback: (err: any, result: MessageEnvelope) => void);

/*
Peeks multiple incoming messages from the queue without removing them.
If there are no messages available in the queue it returns an empty list.

- correlationId     (optional) transaction id to trace execution through call chain.
- messageCount      a maximum number of messages to peek.
- callback          callback function that receives a list with messages or error.
*/
//func (c* MessageQueue)  PeekBatch(correlationId string, messageCount int64) (result []MessageEnvelope, err error)

/*
Receives an incoming message and removes it from the queue.

- correlationId     (optional) transaction id to trace execution through call chain.
- waitTimeout       a timeout in milliseconds to wait for a message to come.
- callback          callback function that receives a message or error.
*/
//func (c* MessageQueue)  Receive(correlationId string, waitTimeout time.Duration) (result *MessageEnvelope, err error)

/*
Renews a lock on a message that makes it invisible from other receivers in the queue.
This method is usually used to extend the message processing time.

- message       a message to extend its lock.
- lockTimeout   a locking timeout in milliseconds.
- callback      (optional) callback function that receives an error or null for success.
*/
//func (c* MessageQueue)  RenewLock(message MessageEnvelope, lockTimeout time.Duration)(err error)

/*
Permanently removes a message from the queue.
This method is usually used to remove the message after successful processing.

- message   a message to remove.
- callback  (optional) callback function that receives an error or null for success.
*/
//func (c* MessageQueue)  Complete(message MessageEnvelope) (err error)

/*
Returnes message into the queue and makes it available for all subscribers to receive it again.
This method is usually used to return a message which could not be processed at the moment
to repeat the attempt. Messages that cause unrecoverable errors shall be removed permanently
or/and send to dead letter queue.

- message   a message to return.
- callback  (optional) callback function that receives an error or null for success.
*/
//func (c* MessageQueue)  Abandon(message MessageEnvelope) (err error)

/*
Permanently removes a message from the queue and sends it to dead letter queue.

- message   a message to be removed.
- callback  (optional) callback function that receives an error or null for success.
*/
//func (c* MessageQueue)  MoveToDeadLetter(message MessageEnvelope) (err error)

/*
Listens for incoming messages and blocks the current thread until queue is closed.

- correlationId     (optional) transaction id to trace execution through call chain.
- receiver          a receiver to receive incoming messages.

See IMessageReceiver
See receive
*/
//func (c* MessageQueue)  Listen(correlationId string, receiver IMessageReceiver);

/*
Ends listening for incoming messages.
When c method is call listen unblocks the thread and execution continues.

- correlationId     (optional) transaction id to trace execution through call chain.
*/
//func (c* MessageQueue)  EndListen(correlationId string);

/*
Listens for incoming messages without blocking the current thread.

- correlationId     (optional) transaction id to trace execution through call chain.
- receiver          a receiver to receive incoming messages.

See listen
See IMessageReceiver
*/
func (c *MessageQueue) BeginListen(correlationId string, receiver IMessageReceiver) {
	go func() {
		c.Listen(correlationId, receiver)
	}()
}

/*
Gets a string representation of the object.

Return a string representation of the object.
*/
func (c *MessageQueue) ToString() string {
	return "[" + c.GetName() + "]"
}
