package queues

/**
 * Data object that contains supported capabilities of a message queue.
 * If certain capability is not supported a queue will throw NotImplemented exception.
 */
type MessagingCapabilities struct {
	canMessageCount bool
	canSend         bool
	canReceive      bool
	canPeek         bool
	canPeekBatch    bool
	canRenewLock    bool
	canAbandon      bool
	canDeadLetter   bool
	canClear        bool
}

/**
 * Creates a new instance of the capabilities object.
 *
 * @param canMessageCount   true if queue supports reading message count.
 * @param canSend           true if queue is able to send messages.
 * @param canReceive        true if queue is able to receive messages.
 * @param canPeek           true if queue is able to peek messages.
 * @param canPeekBatch      true if queue is able to peek multiple messages in one batch.
 * @param canRenewLock      true if queue is able to renew message lock.
 * @param canAbandon        true if queue is able to abandon messages.
 * @param canDeadLetter     true if queue is able to send messages to dead letter queue.
 * @param canClear          true if queue can be cleared.
 */

func NewMessagingCapabilities(canMessageCount bool, canSend bool, canReceive bool,
	canPeek bool, canPeekBatch bool, canRenewLock bool, canAbandon bool,
	canDeadLetter bool, canClear bool) *MessagingCapabilities {

	mc := MessagingCapabilities{}
	mc.canMessageCount = canMessageCount
	mc.canSend = canSend
	mc.canReceive = canReceive
	mc.canPeek = canPeek
	mc.canPeekBatch = canPeekBatch
	mc.canRenewLock = canRenewLock
	mc.canAbandon = canAbandon
	mc.canDeadLetter = canDeadLetter
	mc.canClear = canClear
	return &mc
}

/**
 * Informs if the queue is able to read number of messages.
 *
 * @returns true if queue supports reading message count.
 */
func (c *MessagingCapabilities) CanMessageCount() bool { return c.canMessageCount }

/**
 * Informs if the queue is able to send messages.
 *
 * @returns true if queue is able to send messages.
 */
func (c *MessagingCapabilities) CanSend() bool { return c.canSend }

/**
 * Informs if the queue is able to receive messages.
 *
 * @returns true if queue is able to receive messages.
 */
func (c *MessagingCapabilities) CanReceive() bool { return c.canReceive }

/**
 * Informs if the queue is able to peek messages.
 *
 * @returns true if queue is able to peek messages.
 */
func (c *MessagingCapabilities) CanPeek() bool { return c.canPeek }

/**
 * Informs if the queue is able to peek multiple messages in one batch.
 *
 * @returns true if queue is able to peek multiple messages in one batch.
 */
func (c *MessagingCapabilities) CanPeekBatch() bool { return c.canPeekBatch }

/**
 * Informs if the queue is able to renew message lock.
 *
 * @returns true if queue is able to renew message lock.
 */
func (c *MessagingCapabilities) CanRenewLock() bool { return c.canRenewLock }

/**
 * Informs if the queue is able to abandon messages.
 *
 * @returns true if queue is able to abandon.
 */
func (c *MessagingCapabilities) CanAbandon() bool { return c.canAbandon }

/**
 * Informs if the queue is able to send messages to dead letter queue.
 *
 * @returns true if queue is able to send messages to dead letter queue.
 */
func (c *MessagingCapabilities) CanDeadLetter() bool { return c.canDeadLetter }

/**
 * Informs if the queue can be cleared.
 *
 * @returns true if queue can be cleared.
 */
func (c *MessagingCapabilities) CanClear() bool { return c.canClear }
