package queues

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

//import { IdGenerator } from "pip-services3-commons-node";
//TODO: UTF-8 important?
/**
 * Allows adding additional information to messages. A correlation id, message id, and a message type
 * are added to the data being sent/received. Additionally, a MessageEnvelope can reference a lock token.
 *
 * Side note: a MessageEnvelope"s message is stored as a buffer, so strings are converted
 * using utf8 conversions.
 */
type MessageEnvelope struct {
	reference interface{}

	/** The unique business transaction id that is used to trace calls across components. */
	Correlation_id string
	/** The message"s auto-generated ID. */
	Message_id string
	/** String value that defines the stored message"s type. */
	Message_type string
	/** The time at which the message was sent. */
	Sent_time time.Time
	/** The stored message. */
	Message *bytes.Buffer
}

/**
 * Creates a new MessageEnvelope, which adds a correlation id, message id, and a type to the
 * data being sent/received.
 *
 * @param correlationId     (optional) transaction id to trace execution through call chain.
 * @param messageType       a string value that defines the message"s type.
 * @param message           the data being sent/received.
 */
func NewMessageEnvelope(correlationId string, messageType string, message interface{}) *MessageEnvelope {
	me := MessageEnvelope{}
	me.Correlation_id = correlationId
	me.Message_type = messageType
	if message != nil {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(message)
		if err == nil {
			me.Message = &buf
		}
	}
	me.Message_id = cdata.IdGenerator.NextLong()
	return &me
}

/**
 * @returns the lock token that c MessageEnvelope references.
 */
func (c *MessageEnvelope) GetReference() interface{} {
	return c.reference
}

/**
 * Sets a lock token reference for c MessageEnvelope.
 *
 * @param value     the lock token to reference.
 */
func (c *MessageEnvelope) SetReference(value interface{}) {
	c.reference = value
}

/**
 * @returns the information stored in c message as a UTF-8 encoded string.
 */
func (c *MessageEnvelope) GetMessageAsString() string {
	if c.Message != nil {
		return c.Message.String()
	}
	return ""
}

/**
 * Stores the given string.
 *
 * @param value     the string to set. Will be converted to
 *                  a buffer, using UTF-8 encoding.
 */
func (c *MessageEnvelope) SetMessageAsString(value string) {
	c.Message = bytes.NewBufferString(value)
}

/**
 * @returns the value that was stored in c message
 *          as a JSON string.
 *
 * @see [[setMessageAsJson]]
 */
func (c *MessageEnvelope) GetMessageAsJson() interface{} {
	if c.Message == nil {
		return nil
	}
	temp := c.Message.Bytes()
	var result interface{}
	umErr := json.Unmarshal(temp, &result)
	if umErr != nil {
		return nil
	}
	return result
}

/**
 * Stores the given value as a JSON string.
 *
 * @param value     the value to convert to JSON and store in
 *                  c message.
 *
 * @see [[getMessageAsJson]]
 */
func (c *MessageEnvelope) SetMessageAsJson(value interface{}) {
	if value == nil {
		c.Message = nil
	} else {
		temp, mErr := json.Marshal(value)
		if mErr == nil {
			c.Message = bytes.NewBuffer(temp)
		}
	}
}

/**
 * Convert"s c MessageEnvelope to a string, using the following format:
 *
 * <code>"[<correlation_id>,<Message_type>,<message.toString>]"</code>.
 *
 * If any of the values are <code>nil</code>, they will be replaced with <code>---</code>.
 *
 * @returns the generated string.
 */
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
	if c.Message == nil {
		builder += "---"
	} else {
		builder += c.Message.String()
	}
	builder += "]"
	return builder
}
