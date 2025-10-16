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
)

// Adapter is a function that sends a message and returns a response
type Adapter func(ctx context.Context, message *Message, serializer *Serializer) (*Message, error)

// ClientOptions contains options for configuring a Client
type ClientOptions struct {
	UseBinary          bool
	AlwaysSendJSON     bool
	TimeoutMsDefault   int
	SerializationImpl  Serialization
}

// NewClientOptions creates default client options
func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		UseBinary:          false,
		AlwaysSendJSON:     true,
		TimeoutMsDefault:   5000,
		SerializationImpl:  &DefaultSerialization{},
	}
}

// Client is a telepact client
type Client struct {
	Adapter            Adapter
	UseBinaryDefault   bool
	AlwaysSendJSON     bool
	TimeoutMsDefault   int
	Serializer         *Serializer
}

// NewClient creates a new telepact client
func NewClient(adapter Adapter, options *ClientOptions) *Client {
	// TODO: Create binary encoding cache, binary encoder, base64 encoder
	// For now, create a basic serializer
	serializer := NewSerializer(options.SerializationImpl, nil, nil)

	return &Client{
		Adapter:            adapter,
		UseBinaryDefault:   options.UseBinary,
		AlwaysSendJSON:     options.AlwaysSendJSON,
		TimeoutMsDefault:   options.TimeoutMsDefault,
		Serializer:         serializer,
	}
}

// Request sends a request message and returns a response message
func (c *Client) Request(ctx context.Context, requestMessage *Message) (*Message, error) {
	// TODO: Implementation
	return nil, &TelepactError{Message: "Not implemented"}
}
