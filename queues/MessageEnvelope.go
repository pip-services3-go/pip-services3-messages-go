package queues

import (
	"encoding/json"
	"time"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

/*
MessageEnvelope allows adding additional information to messages. A correlation id, message id, and a message type
are added to the data being sent/received. Additionally, a MessageEnvelope can reference a lock token.
Side note: a MessageEnvelope"s message is stored as a buffer, so strings are converted
using utf8 conversions.
*/
type MessageEnvelope struct {
	reference interface{}

	//The unique business transaction id that is used to trace calls across components.
	Correlation_id string
	// The message"s auto-generated ID.
	Message_id string
	// String value that defines the stored message"s type.
	Message_type string
	// The time at which the message was sent.
	Sent_time time.Time
	//The stored message.
	Message string
}

// NewMessageEnvelope method are creates a new MessageEnvelope, which adds a correlation id, message id, and a type to the
// data being sent/received.
// - correlationId     (optional) transaction id to trace execution through call chain.
// - messageType       a string value that defines the message"s type.
// - message           the data being sent/received.
// Returns: *MessageEnvelope new instance
func NewMessageEnvelope(correlationId string, messageType string, message string) *MessageEnvelope {
	me := MessageEnvelope{}
	me.Correlation_id = correlationId
	me.Message_type = messageType
	me.Message = message
	me.Message_id = cdata.IdGenerator.NextLong()
	return &me
}

//GetReference method are returns the lock token that this MessageEnvelope references.
func (c *MessageEnvelope) GetReference() interface{} {
	return c.reference
}

// SetReference method are sets a lock token reference for this MessageEnvelope.
// - value     the lock token to reference.
func (c *MessageEnvelope) SetReference(value interface{}) {
	c.reference = value
}

// GetMessageAsString method are returns the information stored in this message as a string.
func (c *MessageEnvelope) GetMessageAsString() string {
	return c.Message
}

// SetMessageAsString method are stores the given string.
// - value    the string to set. Will be converted to a bufferg.
func (c *MessageEnvelope) SetMessageAsString(value string) {
	c.Message = value
}

// GetMessageAsJson method are returns the value that was stored in this message as a JSON string.
// See  SetMessageAsJson
func (c *MessageEnvelope) GetMessageAsJson() interface{} {
	if c.Message == "" {
		return nil
	}
	temp := []byte(c.Message)
	var result interface{}
	umErr := json.Unmarshal(temp, &result)
	if umErr != nil {
		return nil
	}
	return result
}

// SetMessageAsJson method are stores the given value as a JSON string.
// - value     the value to convert to JSON and store in this message.
// See  GetMessageAsJson
func (c *MessageEnvelope) SetMessageAsJson(value interface{}) {
	if value == nil {
		c.Message = ""
	} else {
		temp, mErr := json.Marshal(value)
		if mErr == nil {
			c.Message = string(temp)
		}
	}
}

// ToString method are convert"s this MessageEnvelope to a string, using the following format:
// <correlation_id>,<Message_type>,<message.toString>
// If any of the values are nil, they will be replaced with ---.
// Returns the generated string.
func (c *MessageEnvelope) ToString() string {
	builder := "["
	if c.Correlation_id == "" {
		builder += "---"
	} else {
		builder += c.Correlation_id
	}
	builder += ","
	if c.Message_type == "" {
		builder += "---"
	} else {
		builder += c.Message_type
	}
	builder += ","
	if c.Message == "" {
		builder += "---"
	} else {
		builder += c.Message
	}
	builder += "]"
	return builder
}
