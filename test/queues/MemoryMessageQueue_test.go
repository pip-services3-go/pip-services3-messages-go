package test_queues

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-messages-go/queues"
)

func TestMemoryMessageQueue(t *testing.T) {

	var queue *queues.MemoryMessageQueue
	var fixture *MessageQueueFixture

	queue = queues.NewMemoryMessageQueue("TestQueue")
	fixture = NewMessageQueueFixture(queue)
	queue.Open("")

	defer queue.Close("")

	queue.Clear("")

	t.Run("MemoryMessageQueue:Send Receive Message", fixture.TestSendReceiveMessage)
	t.Run("MemoryMessageQueue:Receive Send Message", fixture.TestReceiveSendMessage)
	t.Run("MemoryMessageQueue:Receive And Complete Message", fixture.TestReceiveCompleteMessage)
	t.Run("MemoryMessageQueue:Receive And Abandon Message", fixture.TestReceiveAbandonMessage)
	t.Run("MemoryMessageQueue:Send Peek Message", fixture.TestSendPeekMessage)
	t.Run("MemoryMessageQueue:Peek No Message", fixture.TestPeekNoMessage)
	t.Run("MemoryMessageQueue:Move To Dead Message", fixture.TestMoveToDeadMessage)
	t.Run("MemoryMessageQueue:On Message", fixture.TestOnMessage)

}
