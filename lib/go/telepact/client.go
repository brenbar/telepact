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

// Adapter is a function that sends a message and returns a response
type Adapter func(ctx context.Context, message *Message, serializer *Serializer) (*Message, error)

// ClientOptions contains options for the Client
type ClientOptions struct {
	UseBinary         bool
	AlwaysSendJSON    bool
	TimeoutMsDefault  int
	SerializationImpl Serialization
}

// NewClientOptions creates default ClientOptions
func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		UseBinary:         false,
		AlwaysSendJSON:    true,
		TimeoutMsDefault:  5000,
		SerializationImpl: &DefaultSerialization{},
	}
}

// Client is a Telepact client
type Client struct {
	adapter           Adapter
	useBinaryDefault  bool
	alwaysSendJSON    bool
	timeoutMsDefault  int
	serializer        *Serializer
}

// NewClient creates a new Client
func NewClient(adapter Adapter, options *ClientOptions) *Client {
	binaryEncoder := &binary.ClientBinaryEncoder{}
	base64Encoder := &binary.ClientBase64Encoder{}
	serializer := NewSerializer(options.SerializationImpl, binaryEncoder, base64Encoder)
	
	return &Client{
		adapter:           adapter,
		useBinaryDefault:  options.UseBinary,
		alwaysSendJSON:    options.AlwaysSendJSON,
		timeoutMsDefault:  options.TimeoutMsDefault,
		serializer:        serializer,
	}
}

// Request sends a request message and returns a response
func (c *Client) Request(ctx context.Context, requestMessage *Message) (*Message, error) {
	// TODO: Implement full request logic with binary encoding support
	return c.adapter(ctx, requestMessage, c.serializer)
}
