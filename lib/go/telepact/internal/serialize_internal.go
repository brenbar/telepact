package internal

import (
	"github.com/brenbar/telepact/lib/go/telepact"
)

// SerializeInternal serializes a message for internal telepact use
// This handles the full message serialization with headers and body
func SerializeInternal(
	serializer telepact.Serializer,
	message *telepact.Message,
) ([]byte, error) {
	// TODO: Implement full SerializeInternal logic
	// This is a stub implementation
	
	// For now, just serialize the message using the serializer
	return serializer.Serialize(message)
}

// DeserializeInternal deserializes bytes into a message for internal telepact use
// This handles the full message deserialization with headers and body
func DeserializeInternal(
	serializer *telepact.Serializer,
	messageBytes []byte,
) (*telepact.Message, error) {
	// TODO: Implement full DeserializeInternal logic
	// This is a stub implementation
	
	return serializer.Deserialize(messageBytes)
}
