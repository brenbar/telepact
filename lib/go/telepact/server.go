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
	"fmt"

	"github.com/brenbar/telepact/lib/go/telepact/internal/binary"
)

// Handler is a function that handles a Telepact message
type Handler func(ctx context.Context, requestMessage *Message) (*Message, error)

// ServerOptions contains options for the Server
type ServerOptions struct {
	OnError        func(error)
	OnRequest      func(*Message)
	OnResponse     func(*Message)
	AuthRequired   bool
	Serialization  Serialization
}

// NewServerOptions creates default ServerOptions
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		OnError:       func(e error) {},
		OnRequest:     func(m *Message) {},
		OnResponse:    func(m *Message) {},
		AuthRequired:  true,
		Serialization: &DefaultSerialization{},
	}
}

// Server is a Telepact server
type Server struct {
	handler        Handler
	onError        func(error)
	onRequest      func(*Message)
	onResponse     func(*Message)
	telepactSchema *TelepactSchema
	serializer     *Serializer
}

// NewServer creates a new Server
func NewServer(telepactSchema *TelepactSchema, handler Handler, options *ServerOptions) (*Server, error) {
	binaryEncoder := &binary.ServerBinaryEncoder{}
	base64Encoder := &binary.ServerBase64Encoder{}
	serializer := NewSerializer(options.Serialization, binaryEncoder, base64Encoder)
	
	// Check for auth requirement
	if options.AuthRequired {
		if _, exists := telepactSchema.Parsed["struct.Auth_"]; !exists {
			return nil, fmt.Errorf("unauthenticated server. Either define a `struct.Auth_` in your schema or set `options.AuthRequired` to `false`")
		}
	}
	
	return &Server{
		handler:        handler,
		onError:        options.OnError,
		onRequest:      options.OnRequest,
		onResponse:     options.OnResponse,
		telepactSchema: telepactSchema,
		serializer:     serializer,
	}, nil
}

// Process processes a Telepact Request Message into a Telepact Response Message
func (s *Server) Process(ctx context.Context, requestMessageBytes []byte, overrideHeaders map[string]interface{}) (*Response, error) {
	// TODO: Implement full process logic
	// This is a placeholder implementation
	
	requestMessage, err := s.serializer.Deserialize(requestMessageBytes)
	if err != nil {
		s.onError(err)
		return nil, err
	}
	
	// Merge override headers
	if overrideHeaders != nil {
		for k, v := range overrideHeaders {
			requestMessage.Headers[k] = v
		}
	}
	
	s.onRequest(requestMessage)
	
	responseMessage, err := s.handler(ctx, requestMessage)
	if err != nil {
		s.onError(err)
		return nil, err
	}
	
	s.onResponse(responseMessage)
	
	responseBytes, err := s.serializer.Serialize(responseMessage)
	if err != nil {
		s.onError(err)
		return nil, err
	}
	
	return &Response{
		Bytes:   responseBytes,
		Headers: responseMessage.Headers,
	}, nil
}
