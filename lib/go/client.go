//|
//|  Copyright The Telepact Authors
//|
//|  Licensed under the Apache License, Version 2.0 (the "License");
//|  you may not use this file except in compliance with the License.
//|  You may obtain a copy of the License at
//|
//|  https://www.apache.org/licenses/LICENSE-2.0
//|
//|  Unless required by applicable law or agreed to in writing, software
//|  distributed under the License is distributed on an "AS IS" BASIS,
//|  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//|  See the License for the specific language governing permissions and
//|  limitations under the License.
//|

package telepact

import (
	"context"
	"time"

	"github.com/brenbar/telepact/lib/go/internal/binary"
)

// AdapterFunc is a function that sends a message and returns a response
type AdapterFunc func(ctx context.Context, message *Message, serializer *Serializer) (*Message, error)

// ClientOptions contains configuration options for a Client
type ClientOptions struct {
	UseBinary        bool
	AlwaysSendJSON   bool
	TimeoutMsDefault int
	Serialization    Serialization
}

// NewClientOptions creates default ClientOptions
func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		UseBinary:        false,
		AlwaysSendJSON:   true,
		TimeoutMsDefault: 5000,
		Serialization:    NewDefaultSerialization(),
	}
}

// Client represents a Telepact client
type Client struct {
	Adapter          AdapterFunc
	UseBinaryDefault bool
	AlwaysSendJSON   bool
	TimeoutMsDefault int
	Serializer       *Serializer
}

// NewClient creates a new Client
func NewClient(adapter AdapterFunc, options *ClientOptions) *Client {
	if options == nil {
		options = NewClientOptions()
	}

	// Create binary encoders with cache
	binaryEncodingCache := binary.NewDefaultBinaryEncodingCache()
	binaryEncoder := binary.NewClientBinaryEncoder(&binaryEncodingCache.BinaryEncodingCache)
	base64Encoder := binary.NewClientBase64Encoder()

	// Create serializer
	serializer := NewSerializer(options.Serialization, binaryEncoder, base64Encoder)

	return &Client{
		Adapter:          adapter,
		UseBinaryDefault: options.UseBinary,
		AlwaysSendJSON:   options.AlwaysSendJSON,
		TimeoutMsDefault: options.TimeoutMsDefault,
		Serializer:       serializer,
	}
}

// Request sends a request message and returns the response message
func (c *Client) Request(ctx context.Context, requestMessage *Message) (*Message, error) {
	// Apply timeout if no deadline is set
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(c.TimeoutMsDefault)*time.Millisecond)
		defer cancel()
	}

	// Call adapter
	return c.Adapter(ctx, requestMessage, c.Serializer)
}
