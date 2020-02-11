package queues

/**
 * Message queue that sends and receives messages within the same process by using shared memory.
 *
 * This queue is typically used for testing to mock real queues.
 *
 * ### Configuration parameters ###
 *
 * - name:                        name of the message queue
 *
 * ### References ###
 *
 * - <code>\*:logger:\*:\*:1.0</code>           (optional) [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/interfaces/log.ilogger.html ILogger]] components to pass log messages
 * - <code>\*:counters:\*:\*:1.0</code>         (optional) [[https://rawgit.com/pip-services-node/pip-services3-components-node/master/doc/api/interfaces/count.icounters.html ICounters]] components to pass collected measurements
 *
 * @see [[MessageQueue]]
 * @see [[MessagingCapabilities]]
 *
 * ### Example ###
 *
 *     let queue = new MessageQueue("myqueue");
 *
 *     queue.send("123", new MessageEnvelop(null, "mymessage", "ABC"));
 *
 *     queue.receive("123", (err, message) => {
 *         if (message != null) {
 *            ...
 *            queue.complete("123", message);
 *         }
 *     });
 */
type MemoryMessageQueue struct {
	MessageQueue
	messages          []MessageEnvelope
	lockTokenSequence int
	//lockedMessages { [id: number]: LockedMessage; } = {};
	opened bool
	/** Used to stop the listening process. */
	cancel bool
}

/**
 * Creates a new instance of the message queue.
 *
 * @param name  (optional) a queue name.
 *
 * @see [[MessagingCapabilities]]
 */
func NewMemoryMessageQueue(name string) *MemoryMessageQueue {
	mmq := MemoryMessageQueue{}
	mmq.MessageQueue = *NewMessageQueue(name)
	mmq.messages = make([]MessageEnvelope, 0)
	mmq.lockTokenSequence = 0
	//mmq.lockedMessages { [id: number]: LockedMessage; } = {};
	mmq.opened = false
	mmq.cancel = false
	mmq.Capabilities = NewMessagingCapabilities(true, true, true, true, true, true, true, false, true)
	return &mmq
}

// /**
//  * Checks if the component is opened.
//  *
//  * @returns true if the component has been opened and false otherwise.
//  */
// public isOpen(): boolean {
//     return this.opened;
// }

// /**
//  * Opens the component with given connection and credential parameters.
//  *
//  * @param correlationId     (optional) transaction id to trace execution through call chain.
//  * @param connection        connection parameters
//  * @param credential        credential parameters
//  * @param callback 			callback function that receives error or null no errors occured.
//  */
// protected openWithParams(correlationId: string, connection: ConnectionParams, credential: CredentialParams, callback: (err: any) => void): void {
//     this.opened = true;
//     callback(null);
// }

// /**
//  * Closes component and frees used resources.
//  *
//  * @param correlationId 	(optional) transaction id to trace execution through call chain.
//  * @param callback 			callback function that receives error or null no errors occured.
//  */
// public close(correlationId: string, callback: (err: any) => void): void {
//     this.opened = false;
//     this.cancel = true;
//     this._logger.trace(correlationId, "Closed queue %s", this);
//     callback(null);
// }

// /**
//  * Clears component state.
//  *
//  * @param correlationId 	(optional) transaction id to trace execution through call chain.
//  * @param callback 			callback function that receives error or null no errors occured.
//  */
// public clear(correlationId: string, callback: (err?: any) => void): void {
//     this.messages = [];
//     this.lockedMessages = {};
//     this.cancel = false;

//     callback();
// }

// /**
//  * Reads the current number of messages in the queue to be delivered.
//  *
//  * @param callback      callback function that receives number of messages or error.
//  */
// public readMessageCount(callback: (err: any, count: number) => void): void {
//     let count = this.messages.length;
//     callback(null, count);
// }

// /**
//  * Sends a message into the queue.
//  *
//  * @param correlationId     (optional) transaction id to trace execution through call chain.
//  * @param envelope          a message envelop to be sent.
//  * @param callback          (optional) callback function that receives error or null for success.
//  */
// public send(correlationId: string, envelope: MessageEnvelope, callback?: (err: any) => void): void {
//     try {
//         envelope.sent_time = new Date();
//         // Add message to the queue
//         this.messages.push(envelope);

//         this._counters.incrementOne("queue." + this.getName() + ".sentmessages");
//         this._logger.debug(envelope.correlation_id, "Sent message %s via %s", envelope.toString(), this.toString());

//         if (callback) callback(null);
//     } catch (ex) {
//         if (callback) callback(ex);
//         else throw ex;
//     }
// }

// /**
//  * Peeks a single incoming message from the queue without removing it.
//  * If there are no messages available in the queue it returns null.
//  *
//  * @param correlationId     (optional) transaction id to trace execution through call chain.
//  * @param callback          callback function that receives a message or error.
//  */
// public peek(correlationId: string, callback: (err: any, result: MessageEnvelope) => void): void {
//     try {
//         let message: MessageEnvelope = null;

//         // Pick a message
//         if (this.messages.length > 0)
//             message = this.messages[0];

//         if (message != null)
//             this._logger.trace(message.correlation_id, "Peeked message %s on %s", message, this.toString());

//         callback(null, message);
//     } catch (ex) {
//         callback(ex, null);
//     }
// }

// /**
//  * Peeks multiple incoming messages from the queue without removing them.
//  * If there are no messages available in the queue it returns an empty list.
//  *
//  * @param correlationId     (optional) transaction id to trace execution through call chain.
//  * @param messageCount      a maximum number of messages to peek.
//  * @param callback          callback function that receives a list with messages or error.
//  */
// public peekBatch(correlationId: string, messageCount: number, callback: (err: any, result: MessageEnvelope[]) => void): void {
//     try {
//         let messages = this.messages.slice(0, messageCount);

//         this._logger.trace(correlationId, "Peeked %d messages on %s", messages.length, this.toString());

//         callback(null, messages);
//     } catch (ex) {
//         callback(ex, null);
//     }
// }

// /**
//  * Receives an incoming message and removes it from the queue.
//  *
//  * @param correlationId     (optional) transaction id to trace execution through call chain.
//  * @param waitTimeout       a timeout in milliseconds to wait for a message to come.
//  * @param callback          callback function that receives a message or error.
//  */
// public receive(correlationId: string, waitTimeout: number, callback: (err: any, result: MessageEnvelope) => void): void {
//     let err: any = null;
//     let message: MessageEnvelope = null;
//     let messageReceived: boolean = false;

//     let checkIntervalMs = 100;
//     let i = 0;
//     async.whilst(
//         () => {
//             return i < waitTimeout && !messageReceived;
//         },
//         (whilstCallback) => {
//             i = i + checkIntervalMs;

//             setTimeout(() => {
//                 if (this.messages.length == 0) {
//                     whilstCallback();
//                     return;
//                 }

//                 try {
//                     // Get message the the queue
//                     message = this.messages.shift();

//                     if (message != null) {
//                         // Generate and set locked token
//                         var lockedToken = this.lockTokenSequence++;
//                         message.setReference(lockedToken);

//                         // Add messages to locked messages list
//                         let lockedMessage: LockedMessage = new LockedMessage();
//                         let now: Date = new Date();
//                         now.setMilliseconds(now.getMilliseconds() + waitTimeout);
//                         lockedMessage.expirationTime = now;
//                         lockedMessage.message = message;
//                         lockedMessage.timeout = waitTimeout;
//                         this.lockedMessages[lockedToken] = lockedMessage;
//                     }

//                     if (message != null) {
//                         this._counters.incrementOne("queue." + this.getName() + ".receivedmessages");
//                         this._logger.debug(message.correlation_id, "Received message %s via %s", message, this.toString());
//                     }
//                 } catch (ex) {
//                     err = ex;
//                 }

//                 messageReceived = true;
//                 whilstCallback();
//             }, checkIntervalMs);
//         },
//         (err) => {
//             callback(err, message);
//         }
//     );
// }

// /**
//  * Renews a lock on a message that makes it invisible from other receivers in the queue.
//  * This method is usually used to extend the message processing time.
//  *
//  * @param message       a message to extend its lock.
//  * @param lockTimeout   a locking timeout in milliseconds.
//  * @param callback      (optional) callback function that receives an error or null for success.
//  */
// public renewLock(message: MessageEnvelope, lockTimeout: number, callback?: (err: any) => void): void {
//     if (message.getReference() == null) {
//         if (callback) callback(null);
//         return;
//     }

//     // Get message from locked queue
//     try {
//         let lockedToken: number = message.getReference();
//         let lockedMessage: LockedMessage = this.lockedMessages[lockedToken];

//         // If lock is found, extend the lock
//         if (lockedMessage) {
//             let now: Date = new Date();
//             // Todo: Shall we skip if the message already expired?
//             if (lockedMessage.expirationTime > now) {
//                 now.setMilliseconds(now.getMilliseconds() + lockedMessage.timeout);
//                 lockedMessage.expirationTime = now;
//             }
//         }

//         this._logger.trace(message.correlation_id, "Renewed lock for message %s at %s", message, this.toString());

//         if (callback) callback(null);
//     } catch (ex) {
//         if (callback) callback(ex);
//         else throw ex;
//     }
// }

// /**
//  * Permanently removes a message from the queue.
//  * This method is usually used to remove the message after successful processing.
//  *
//  * @param message   a message to remove.
//  * @param callback  (optional) callback function that receives an error or null for success.
//  */
// public complete(message: MessageEnvelope, callback: (err: any) => void): void {
//     if (message.getReference() == null) {
//         if (callback) callback(null);
//         return;
//     }

//     try {
//         let lockKey: number = message.getReference();
//         delete this.lockedMessages[lockKey];
//         message.setReference(null);

//         this._logger.trace(message.correlation_id, "Completed message %s at %s", message, this.toString());

//         if (callback) callback(null);
//     } catch (ex) {
//         if (callback) callback(ex);
//         else throw ex;
//     }
// }

// /**
//  * Returnes message into the queue and makes it available for all subscribers to receive it again.
//  * This method is usually used to return a message which could not be processed at the moment
//  * to repeat the attempt. Messages that cause unrecoverable errors shall be removed permanently
//  * or/and send to dead letter queue.
//  *
//  * @param message   a message to return.
//  * @param callback  (optional) callback function that receives an error or null for success.
//  */
// public abandon(message: MessageEnvelope, callback: (err: any) => void): void {
//     if (message.getReference() == null) {
//         if (callback) callback(null);
//         return;
//     }

//     try {
//         // Get message from locked queue
//         let lockedToken: number = message.getReference();
//         let lockedMessage: LockedMessage = this.lockedMessages[lockedToken];
//         if (lockedMessage) {
//             // Remove from locked messages
//             delete this.lockedMessages[lockedToken];
//             message.setReference(null);

//             // Skip if it is already expired
//             if (lockedMessage.expirationTime <= new Date()) {
//                 callback(null);
//                 return;
//             }
//         }
//         // Skip if it absent
//         else {
//             if (callback) callback(null);
//             return;
//         }

//         this._logger.trace(message.correlation_id, "Abandoned message %s at %s", message, this.toString());

//         if (callback) callback(null);
//     } catch (ex) {
//         if (callback) callback(ex);
//         else throw ex;
//     }

//     this.send(message.correlation_id, message, null);
// }

// /**
//  * Permanently removes a message from the queue and sends it to dead letter queue.
//  *
//  * @param message   a message to be removed.
//  * @param callback  (optional) callback function that receives an error or null for success.
//  */
// public moveToDeadLetter(message: MessageEnvelope, callback: (err: any) => void): void {
//     if (message.getReference() == null) {
//         if (callback) callback(null);
//         return;
//     }

//     try {
//         let lockedToken: number = message.getReference();
//         delete this.lockedMessages[lockedToken];
//         message.setReference(null);

//         this._counters.incrementOne("queue." + this.getName() + ".deadmessages");
//         this._logger.trace(message.correlation_id, "Moved to dead message %s at %s", message, this.toString());

//         if (callback) callback(null);
//     } catch (ex) {
//         if (callback) callback(ex);
//         else throw ex;
//     }
// }

// /**
//  * Listens for incoming messages and blocks the current thread until queue is closed.
//  *
//  * @param correlationId     (optional) transaction id to trace execution through call chain.
//  * @param receiver          a receiver to receive incoming messages.
//  *
//  * @see [[IMessageReceiver]]
//  * @see [[receive]]
//  */
// public listen(correlationId: string, receiver: IMessageReceiver): void {
//     let timeoutInterval = 1000;

//     this._logger.trace(null, "Started listening messages at %s", this.toString());

//     this.cancel = false;

//     async.whilst(
//         () => {
//             return !this.cancel;
//         },
//         (whilstCallback) => {
//             let message: MessageEnvelope;

//             async.series([
//                 (callback) => {
//                     this.receive(correlationId, timeoutInterval, (err, result) => {
//                         message = result;
//                         if (err) this._logger.error(correlationId, err, "Failed to receive the message");
//                         callback();
//                     })
//                 },
//                 (callback) => {
//                     if (message != null && !this.cancel) {
//                         receiver.receiveMessage(message, this, (err) => {
//                             if (err) this._logger.error(correlationId, err, "Failed to process the message");
//                             callback();
//                         });
//                     }
//                 },
//             ]);

//             async.series([
//                 (callback) => {
//                     setTimeout(callback, timeoutInterval);
//                 }
//             ], whilstCallback);
//         },
//         (err) => {
//             if (err) this._logger.error(correlationId, err, "Failed to process the message");
//         }
//     );
// }

// /**
//  * Ends listening for incoming messages.
//  * When this method is call [[listen]] unblocks the thread and execution continues.
//  *
//  * @param correlationId     (optional) transaction id to trace execution through call chain.
//  */
// public endListen(correlationId: string): void {
//     this.cancel = true;
// }
