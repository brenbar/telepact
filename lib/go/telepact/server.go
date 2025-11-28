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

// MessageHandler is a function that handles incoming telepact messages
type MessageHandler func(context.Context, *Message) (*Message, error)

// ServerOptions contains options for configuring a Server
type ServerOptions struct {
	OnError       func(error)
	OnRequest     func(*Message)
	OnResponse    func(*Message)
	AuthRequired  bool
	Serialization Serialization
}

// Server is a telepact server
type Server struct {
	telepactSchema *TelepactSchema
	handler        MessageHandler
	onError        func(error)
	onRequest      func(*Message)
	onResponse     func(*Message)
	serializer     *Serializer
}

// NewServer creates a new Server with the given schema and handler
func NewServer(
	telepactSchema *TelepactSchema,
	handler MessageHandler,
	options *ServerOptions,
) (*Server, error) {
	if options == nil {
		options = &ServerOptions{
			OnError:       func(e error) {},
			OnRequest:     func(m *Message) {},
			OnResponse:    func(m *Message) {},
			AuthRequired:  true,
			Serialization: NewDefaultSerialization(),
		}
	}

	// TODO: Construct binary encoding
	// binaryEncoding := constructBinaryEncoding(telepactSchema)
	// binaryEncoder := NewServerBinaryEncoder(binaryEncoding)
	// base64Encoder := NewServerBase64Encoder()
	
	// For now, create placeholder encoders
	var binaryEncoder binary.BinaryEncoder = nil  // TODO
	var base64Encoder binary.Base64Encoder = nil  // TODO
	
	serializer := NewSerializer(options.Serialization, binaryEncoder, base64Encoder)

	if options.AuthRequired {
		if _, exists := telepactSchema.Parsed["struct.Auth_"]; !exists {
			return nil, NewTelepactError(
				"Unauthenticated server. Either define a `struct.Auth_` in your schema or set `options.AuthRequired` to `false`.",
			)
		}
	}

	return &Server{
		telepactSchema: telepactSchema,
		handler:        handler,
		onError:        options.OnError,
		onRequest:      options.OnRequest,
		onResponse:     options.OnResponse,
		serializer:     serializer,
	}, nil
}

// Process processes a telepact Request Message into a telepact Response Message
func (s *Server) Process(
	ctx context.Context,
	requestMessageBytes []byte,
	overrideHeaders map[string]interface{},
) (*Response, error) {
	// TODO: Implement from internal/ProcessBytes
	return processBytes(
		ctx,
		requestMessageBytes,
		overrideHeaders,
		s.serializer,
		s.telepactSchema,
		s.onError,
		s.onRequest,
		s.onResponse,
		s.handler,
	)
}

// Placeholder function - will be implemented in internal package
func processBytes(
	ctx context.Context,
	requestMessageBytes []byte,
	overrideHeaders map[string]interface{},
	serializer *Serializer,
	telepactSchema *TelepactSchema,
	onError func(error),
	onRequest func(*Message),
	onResponse func(*Message),
	handler MessageHandler,
) (*Response, error) {
	// TODO: Implement
	return NewResponse([]byte{}, map[string]interface{}{}), nil
}
