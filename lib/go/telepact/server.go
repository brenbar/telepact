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
)

// Handler is a function that processes a message and returns a response message
type Handler func(ctx context.Context, message *Message) (*Message, error)

// ServerOptions contains options for configuring a Server
type ServerOptions struct {
	OnError        func(error)
	OnRequest      func(*Message)
	OnResponse     func(*Message)
	AuthRequired   bool
	Serialization  Serialization
}

// NewServerOptions creates default server options
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		OnError:       func(e error) {},
		OnRequest:     func(m *Message) {},
		OnResponse:    func(m *Message) {},
		AuthRequired:  true,
		Serialization: &DefaultSerialization{},
	}
}

// Server is a telepact server
type Server struct {
	Handler        Handler
	OnError        func(error)
	OnRequest      func(*Message)
	OnResponse     func(*Message)
	TelepactSchema *TelepactSchema
	Serializer     *Serializer
}

// NewServer creates a new telepact server
func NewServer(telepactSchema *TelepactSchema, handler Handler, options *ServerOptions) (*Server, error) {
	// Check for Auth_ requirement
	if options.AuthRequired {
		if _, exists := telepactSchema.Parsed["struct.Auth_"]; !exists {
			return nil, fmt.Errorf("Unauthenticated server. Either define a `struct.Auth_` in your schema or set `options.AuthRequired` to `false`")
		}
	}

	// TODO: Construct binary encoding, binary encoder, base64 encoder
	// For now, create a basic serializer
	serializer := NewSerializer(options.Serialization, nil, nil)

	return &Server{
		Handler:        handler,
		OnError:        options.OnError,
		OnRequest:      options.OnRequest,
		OnResponse:     options.OnResponse,
		TelepactSchema: telepactSchema,
		Serializer:     serializer,
	}, nil
}

// Process processes a telepact request message and returns a telepact response
func (s *Server) Process(ctx context.Context, requestMessageBytes []byte, overrideHeaders map[string]interface{}) (*Response, error) {
	// TODO: Implementation
	return nil, &TelepactError{Message: "Not implemented"}
}
