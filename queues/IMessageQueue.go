package queues

import (
	"time" /** @module queues */
)

/**
 * Interface for asynchronous message queues.
 *
 * Not all queues may implement all the methods.
 * Attempt to call non-supported method will result in NotImplemented exception.
 * To verify if specific method is supported consult with [[MessagingCapabilities]].
 *
 * @see [[MessageEnvelop]]
 * @see [[MessagingCapabilities]]
 */
// extends IOpenable, IClosable
type IMessageQueue interface {
	/**
	 * Gets the queue name
	 *
	 * @returns the queue name.
	 */
	GetName() string

	/**
	 * Gets the queue capabilities
	 *
	 * @returns the queue's capabilities object.
	 */
	GetCapabilities() MessagingCapabilities

	/**
	 * Reads the current number of messages in the queue to be delivered.
	 *
	 * @param callback      callback function that receives number of messages or error.
	 */
	ReadMessageCount() (count int64, err error)

	/**
	 * Sends a message into the queue.
	 *
	 * @param correlationId     (optional) transaction id to trace execution through call chain.
	 * @param envelope          a message envelop to be sent.
	 * @param callback          (optional) callback function that receives error or null for success.
	 */
	Send(correlationId string, envelope MessageEnvelope) (err error)

	/**
	 * Sends an object into the queue.
	 * Before sending the object is converted into JSON string and wrapped in a [[MessageEnvelop]].
	 *
	 * @param correlationId     (optional) transaction id to trace execution through call chain.
	 * @param messageType       a message type
	 * @param value             an object value to be sent
	 * @param callback          (optional) callback function that receives error or null for success.
	 *
	 * @see [[send]]
	 */
	SendAsObject(correlationId string, messageType string, value interface{}) (err error)

	/**
	 * Peeks a single incoming message from the queue without removing it.
	 * If there are no messages available in the queue it returns null.
	 *
	 * @param correlationId     (optional) transaction id to trace execution through call chain.
	 * @param callback          callback function that receives a message or error.
	 */
	Peek(correlationId string) (result MessageEnvelope, err error)

	/**
	 * Peeks multiple incoming messages from the queue without removing them.
	 * If there are no messages available in the queue it returns an empty list.
	 *
	 * @param correlationId     (optional) transaction id to trace execution through call chain.
	 * @param messageCount      a maximum number of messages to peek.
	 * @param callback          callback function that receives a list with messages or error.
	 */
	PeekBatch(correlationId string, messageCount int64) (result []MessageEnvelope, err error)

	/**
	 * Receives an incoming message and removes it from the queue.
	 *
	 * @param correlationId     (optional) transaction id to trace execution through call chain.
	 * @param waitTimeout       a timeout in milliseconds to wait for a message to come.
	 * @param callback          callback function that receives a message or error.
	 */
	Receive(correlationId string, waitTimeout time.Duration) (result MessageEnvelope, err error)

	/**
	 * Renews a lock on a message that makes it invisible from other receivers in the queue.
	 * This method is usually used to extend the message processing time.
	 *
	 * @param message       a message to extend its lock.
	 * @param lockTimeout   a locking timeout in milliseconds.
	 * @param callback      (optional) callback function that receives an error or null for success.
	 */
	RenewLock(message MessageEnvelope, lockTimeout time.Duration) (err error)

	/**
	 * Permanently removes a message from the queue.
	 * This method is usually used to remove the message after successful processing.
	 *
	 * @param message   a message to remove.
	 * @param callback  (optional) callback function that receives an error or null for success.
	 */
	Complete(message MessageEnvelope) (err error)

	/**
	 * Returnes message into the queue and makes it available for all subscribers to receive it again.
	 * This method is usually used to return a message which could not be processed at the moment
	 * to repeat the attempt. Messages that cause unrecoverable errors shall be removed permanently
	 * or/and send to dead letter queue.
	 *
	 * @param message   a message to return.
	 * @param callback  (optional) callback function that receives an error or null for success.
	 */
	Abandon(message MessageEnvelope) (err error)

	/**
	 * Permanently removes a message from the queue and sends it to dead letter queue.
	 *
	 * @param message   a message to be removed.
	 * @param callback  (optional) callback function that receives an error or null for success.
	 */
	MoveToDeadLetter(message MessageEnvelope) (err error)

	/**
	 * Listens for incoming messages and blocks the current thread until queue is closed.
	 *
	 * @param correlationId     (optional) transaction id to trace execution through call chain.
	 * @param receiver          a receiver to receive incoming messages.
	 *
	 * @see [[IMessageReceiver]]
	 * @see [[receive]]
	 */
	Listen(correlationId string, receiver IMessageReceiver)

	/**
	 * Listens for incoming messages without blocking the current thread.
	 *
	 * @param correlationId     (optional) transaction id to trace execution through call chain.
	 * @param receiver          a receiver to receive incoming messages.
	 *
	 * @see [[listen]]
	 * @see [[IMessageReceiver]]
	 */
	BeginListen(correlationId string, receiver IMessageReceiver)

	/**
	 * Ends listening for incoming messages.
	 * When this method is call [[listen]] unblocks the thread and execution continues.
	 *
	 * @param correlationId     (optional) transaction id to trace execution through call chain.
	 */
	EndListen(correlationId string)
}
