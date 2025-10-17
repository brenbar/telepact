package internal

import (
	"github.com/brenbar/telepact/lib/go/telepact"
)

// ParseRequestMessage parses request bytes into a Message
// It handles deserialization and basic validation
func ParseRequestMessage(
	requestMessageBytes []byte,
	serializer *telepact.Serializer,
	telepactSchema *telepact.TelepactSchema,
	onError func(error),
) (*telepact.Message, error) {
	// TODO: Implement full ParseRequestMessage logic
	// This is a stub implementation
	
	// Deserialize the message
	message, err := serializer.Deserialize(requestMessageBytes)
	if err != nil {
		if onError != nil {
			onError(err)
		}
		return nil, err
	}
	
	// TODO: Perform validation
	// TODO: Handle base64 decoding
	// TODO: Store parse failures in headers
	
	_ = telepactSchema
	
	return message, nil
}
