package queues

import "time"

/*
LockedMessage data object used to store and lock incoming messages in MemoryMessageQueue.
See: MemoryMessageQueue
*/
type LockedMessage struct {

	//The incoming message.
	Message *MessageEnvelope

	// The expiration time for the message lock.
	// If it is null then the message is not locked.
	ExpirationTime time.Time

	//The lock timeout in milliseconds.
	Timeout time.Duration
}
