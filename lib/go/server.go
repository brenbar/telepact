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
	"fmt"

	"github.com/brenbar/telepact/lib/go/internal/binary"
)

// HandlerFunc is a function that handles telepact messages
type HandlerFunc func(ctx context.Context, message *Message) (*Message, error)

// ErrorFunc is a function that handles errors
type ErrorFunc func(error)

// MessageFunc is a function that processes messages (for hooks)
type MessageFunc func(*Message)

// ServerOptions contains configuration options for a Server
type ServerOptions struct {
	OnError       ErrorFunc
	OnRequest     MessageFunc
	OnResponse    MessageFunc
	AuthRequired  bool
	Serialization Serialization
}

// NewServerOptions creates default ServerOptions
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		OnError:       func(error) {},
		OnRequest:     func(*Message) {},
		OnResponse:    func(*Message) {},
		AuthRequired:  true,
		Serialization: NewDefaultSerialization(),
	}
}

// Server represents a Telepact server
type Server struct {
	Handler        HandlerFunc
	OnError        ErrorFunc
	OnRequest      MessageFunc
	OnResponse     MessageFunc
	TelepactSchema *TelepactSchema
	Serializer     *Serializer
}

// NewServer creates a new Server
func NewServer(telepactSchema *TelepactSchema, handler HandlerFunc, options *ServerOptions) (*Server, error) {
	if options == nil {
		options = NewServerOptions()
	}

	// Check for Auth_ struct if auth is required
	if options.AuthRequired {
		if _, hasAuth := telepactSchema.Parsed["struct.Auth_"]; !hasAuth {
			return nil, fmt.Errorf("unauthenticated server: either define a `struct.Auth_` in your schema or set `options.AuthRequired` to false")
		}
	}

	// Create binary encoders
	binaryEncoding := binary.ConstructBinaryEncoding(telepactSchema)
	binaryEncoder := binary.NewServerBinaryEncoder(binaryEncoding)
	base64Encoder := binary.NewServerBase64Encoder()

	// Create serializer
	serializer := NewSerializer(options.Serialization, binaryEncoder, base64Encoder)

	return &Server{
		Handler:        handler,
		OnError:        options.OnError,
		OnRequest:      options.OnRequest,
		OnResponse:     options.OnResponse,
		TelepactSchema: telepactSchema,
		Serializer:     serializer,
	}, nil
}

// Process handles a request message and returns a response
func (s *Server) Process(ctx context.Context, requestMessageBytes []byte, overrideHeaders map[string]interface{}) (*Response, error) {
	// Deserialize request
	requestMessage, err := s.Serializer.Deserialize(requestMessageBytes)
	if err != nil {
		s.OnError(err)
		return nil, err
	}

	// Override headers if provided
	if overrideHeaders != nil {
		for k, v := range overrideHeaders {
			requestMessage.Headers[k] = v
		}
	}

	// Call request hook
	s.OnRequest(requestMessage)

	// Call handler
	responseMessage, err := s.Handler(ctx, requestMessage)
	if err != nil {
		s.OnError(err)
		return nil, err
	}

	// Call response hook
	s.OnResponse(responseMessage)

	// Serialize response
	responseBytes, err := s.Serializer.Serialize(responseMessage)
	if err != nil {
		s.OnError(err)
		return nil, err
	}

	// Determine media type based on headers
	headers := make(map[string]interface{})
	for k, v := range responseMessage.Headers {
		headers[k] = v
	}

	return NewResponse(responseBytes, headers), nil
}
