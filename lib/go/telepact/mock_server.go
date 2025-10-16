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
)

// MockServerOptions contains options for the MockServer
type MockServerOptions struct {
	OnError                             func(error)
	EnableMessageResponseGeneration     bool
	EnableOptionalFieldGeneration       bool
	RandomizeOptionalFieldGeneration    bool
	GeneratedCollectionLengthMin        int
	GeneratedCollectionLengthMax        int
}

// NewMockServerOptions creates default MockServerOptions
func NewMockServerOptions() *MockServerOptions {
	return &MockServerOptions{
		OnError:                             func(e error) {},
		EnableMessageResponseGeneration:     true,
		EnableOptionalFieldGeneration:       true,
		RandomizeOptionalFieldGeneration:    true,
		GeneratedCollectionLengthMin:        0,
		GeneratedCollectionLengthMax:        3,
	}
}

// MockServer is a mock instance of a Telepact server
type MockServer struct {
	random                              *RandomGenerator
	enableGeneratedDefaultStub          bool
	enableOptionalFieldGeneration       bool
	randomizeOptionalFieldGeneration    bool
	stubs                               []interface{} // TODO: Use proper MockStub type
	invocations                         []interface{} // TODO: Use proper MockInvocation type
	server                              *Server
}

// NewMockServer creates a new MockServer
func NewMockServer(mockTelepactSchema *MockTelepactSchema, options *MockServerOptions) (*MockServer, error) {
	random := NewRandomGenerator(options.GeneratedCollectionLengthMin, options.GeneratedCollectionLengthMax)
	
	serverOptions := NewServerOptions()
	serverOptions.OnError = options.OnError
	serverOptions.AuthRequired = false
	
	telepactSchema := &TelepactSchema{
		Original:              mockTelepactSchema.Original,
		Parsed:                mockTelepactSchema.Parsed,
		ParsedRequestHeaders:  mockTelepactSchema.ParsedRequestHeaders,
		ParsedResponseHeaders: mockTelepactSchema.ParsedResponseHeaders,
	}
	
	ms := &MockServer{
		random:                           random,
		enableGeneratedDefaultStub:       options.EnableMessageResponseGeneration,
		enableOptionalFieldGeneration:    options.EnableOptionalFieldGeneration,
		randomizeOptionalFieldGeneration: options.RandomizeOptionalFieldGeneration,
		stubs:                            make([]interface{}, 0),
		invocations:                      make([]interface{}, 0),
	}
	
	server, err := NewServer(telepactSchema, ms.handle, serverOptions)
	if err != nil {
		return nil, err
	}
	
	ms.server = server
	return ms, nil
}

// Process processes a Telepact Request Message into a Telepact Response Message
func (ms *MockServer) Process(ctx context.Context, messageBytes []byte) ([]byte, error) {
	response, err := ms.server.Process(ctx, messageBytes, nil)
	if err != nil {
		return nil, err
	}
	return response.Bytes, nil
}

func (ms *MockServer) handle(ctx context.Context, requestMessage *Message) (*Message, error) {
	// TODO: Implement mock handling logic
	// This is a placeholder
	return &Message{
		Headers: make(map[string]interface{}),
		Body:    make(map[string]interface{}),
	}, nil
}
