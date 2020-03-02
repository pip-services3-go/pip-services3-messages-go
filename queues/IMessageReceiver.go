package queues

/*
IMessageReceiver callback interface to receive incoming messages.
Example:

    type MyMessageReceiver struct {
      func (c*MyMessageReceiver) ReceiveMessage(envelop MessageEnvelop, queue IMessageQueue) {
          fmt.Println("Received message: " + envelop.GetMessageAsString());
      }
    }

    messageQueue := NewMemoryMessageQueue();
    messageQueue.Listen("123", NewMyMessageReceiver());

	opnErr := messageQueue.Open("123")
	if opnErr == nil{
       messageQueue.Send("123", NewMessageEnvelop("", "mymessage", "ABC")); // Output in console: "Received message: ABC"
    }
*/
type IMessageReceiver interface {

	// ReceiveMessage method are receives incoming message from the queue.
	// - envelope  an incoming message
	// - queue     a queue where the message comes from
	// - callback  callback function that receives error or null for success.
	// See: MessageEnvelope
	// See: IMessageQueue
	ReceiveMessage(envelope *MessageEnvelope, queue IMessageQueue) (err error)
}
