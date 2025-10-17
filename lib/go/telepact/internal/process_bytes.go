package internal

import (
	"context"
	
	"github.com/brenbar/telepact/lib/go/telepact"
)

// ProcessBytes is the main server entry point for processing request bytes
// It parses the request, validates it, calls the handler, and returns a response
func ProcessBytes(
	ctx context.Context,
	requestMessageBytes []byte,
	overrideHeaders map[string]interface{},
	serializer *telepact.Serializer,
	telepactSchema *telepact.TelepactSchema,
	onError func(error),
	onRequest func(*telepact.Message),
	onResponse func(*telepact.Message),
	handler func(context.Context, *telepact.Message) (*telepact.Message, error),
) (*telepact.Response, error) {
	// TODO: Implement full ProcessBytes logic
	// This is a stub implementation
	
	// Parse request message
	requestMessage, err := ParseRequestMessage(requestMessageBytes, serializer, telepactSchema, onError)
	if err != nil {
		// On error, return ErrorUnknown_
		if onError != nil {
			onError(err)
		}
		errorMsg := telepact.NewMessage(map[string]interface{}{}, map[string]interface{}{
			"ErrorUnknown_": map[string]interface{}{},
		})
		responseBytes, _ := serializer.Serialize(errorMsg)
		return &telepact.Response{
			Bytes:   responseBytes,
			Headers: map[string]interface{}{},
		}, nil
	}
	
	// Call onRequest callback
	if onRequest != nil {
		func() {
			defer func() { recover() }()
			onRequest(requestMessage)
		}()
	}
	
	// Handle message
	responseMessage, err := HandleMessage(ctx, requestMessage, overrideHeaders, telepactSchema, handler, onError)
	if err != nil {
		// On error, return ErrorUnknown_
		if onError != nil {
			onError(err)
		}
		errorMsg := telepact.NewMessage(map[string]interface{}{}, map[string]interface{}{
			"ErrorUnknown_": map[string]interface{}{},
		})
		responseBytes, _ := serializer.Serialize(errorMsg)
		return &telepact.Response{
			Bytes:   responseBytes,
			Headers: map[string]interface{}{},
		}, nil
	}
	
	// Call onResponse callback
	if onResponse != nil {
		func() {
			defer func() { recover() }()
			onResponse(responseMessage)
		}()
	}
	
	// Serialize response
	responseBytes, err := serializer.Serialize(responseMessage)
	if err != nil {
		if onError != nil {
			onError(err)
		}
		errorMsg := telepact.NewMessage(map[string]interface{}{}, map[string]interface{}{
			"ErrorUnknown_": map[string]interface{}{},
		})
		responseBytes, _ := serializer.Serialize(errorMsg)
		return &telepact.Response{
			Bytes:   responseBytes,
			Headers: map[string]interface{}{},
		}, nil
	}
	
	return &telepact.Response{
		Bytes:   responseBytes,
		Headers: responseMessage.Headers,
	}, nil
}
