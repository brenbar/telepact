//
//  Copyright The Telepact Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package telepact

import (
	"context"
	
	"github.com/brenbar/telepact/lib/go/telepact/internal/binary"
)

// MessageAdapter is a function that sends a message and returns a response
type MessageAdapter func(context.Context, *Message, *Serializer) (*Message, error)

// ClientOptions contains options for configuring a Client
type ClientOptions struct {
	UseBinary        bool
	AlwaysSendJSON   bool
	TimeoutMsDefault int
	Serialization    Serialization
}

// Client is a telepact client
type Client struct {
	adapter          MessageAdapter
	useBinaryDefault bool
	alwaysSendJSON   bool
	timeoutMsDefault int
	serializer       *Serializer
}

// NewClient creates a new Client with the given adapter and options
func NewClient(adapter MessageAdapter, options *ClientOptions) *Client {
	if options == nil {
		options = &ClientOptions{
			UseBinary:        false,
			AlwaysSendJSON:   true,
			TimeoutMsDefault: 5000,
			Serialization:    NewDefaultSerialization(),
		}
	}

	// TODO: Create binary encoding cache and encoders
	// binaryEncodingCache := NewDefaultBinaryEncodingCache()
	// binaryEncoder := NewClientBinaryEncoder(binaryEncodingCache)
	// base64Encoder := NewClientBase64Encoder()
	
	// For now, create placeholder encoders
	var binaryEncoder binary.BinaryEncoder = nil  // TODO
	var base64Encoder binary.Base64Encoder = nil  // TODO

	serializer := NewSerializer(options.Serialization, binaryEncoder, base64Encoder)

	return &Client{
		adapter:          adapter,
		useBinaryDefault: options.UseBinary,
		alwaysSendJSON:   options.AlwaysSendJSON,
		timeoutMsDefault: options.TimeoutMsDefault,
		serializer:       serializer,
	}
}

// Request sends a request message and returns the response
func (c *Client) Request(ctx context.Context, requestMessage *Message) (*Message, error) {
	// TODO: Implement from internal/ClientHandleMessage
	return clientHandleMessage(
		ctx,
		requestMessage,
		c.adapter,
		c.serializer,
		c.timeoutMsDefault,
		c.useBinaryDefault,
		c.alwaysSendJSON,
	)
}

// Placeholder function - will be implemented in internal package
func clientHandleMessage(
	ctx context.Context,
	requestMessage *Message,
	adapter MessageAdapter,
	serializer *Serializer,
	timeoutMsDefault int,
	useBinaryDefault bool,
	alwaysSendJSON bool,
) (*Message, error) {
	// TODO: Implement
	return adapter(ctx, requestMessage, serializer)
}
