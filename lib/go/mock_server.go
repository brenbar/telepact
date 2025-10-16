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

// MockTelepactSchema represents a mock schema for testing
type MockTelepactSchema struct {
	*TelepactSchema
}

// NewMockTelepactSchema creates a new MockTelepactSchema
func NewMockTelepactSchema() *MockTelepactSchema {
	schema, _ := FromJSON("[]")
	return &MockTelepactSchema{
		TelepactSchema: schema,
	}
}

// MockServerOptions contains configuration options for a MockServer
type MockServerOptions struct {
	OnError                          ErrorFunc
	EnableMessageResponseGeneration  bool
	EnableOptionalFieldGeneration    bool
	RandomizeOptionalFieldGeneration bool
	GeneratedCollectionLengthMin     int
	GeneratedCollectionLengthMax     int
}

// NewMockServerOptions creates default MockServerOptions
func NewMockServerOptions() *MockServerOptions {
	return &MockServerOptions{
		OnError:                          func(error) {},
		EnableMessageResponseGeneration:  true,
		EnableOptionalFieldGeneration:    true,
		RandomizeOptionalFieldGeneration: true,
		GeneratedCollectionLengthMin:     0,
		GeneratedCollectionLengthMax:     3,
	}
}

// MockServer provides a mock implementation of a Telepact server for testing
type MockServer struct {
	Server                           *Server
	Random                           *RandomGenerator
	EnableGeneratedDefaultStub       bool
	EnableOptionalFieldGeneration    bool
	RandomizeOptionalFieldGeneration bool
}

// NewMockServer creates a new MockServer
func NewMockServer(mockSchema *MockTelepactSchema, options *MockServerOptions) (*MockServer, error) {
	if options == nil {
		options = NewMockServerOptions()
	}

	random := NewRandomGenerator(
		options.GeneratedCollectionLengthMin,
		options.GeneratedCollectionLengthMax,
	)

	// Create a simple handler that generates mock responses
	handler := func(ctx context.Context, message *Message) (*Message, error) {
		// Simple mock response - just echo back with Ok_
		return NewMessage(
			map[string]interface{}{},
			map[string]interface{}{
				"Ok_": map[string]interface{}{
					"result": "mock response",
				},
			},
		), nil
	}

	serverOptions := NewServerOptions()
	serverOptions.OnError = options.OnError
	serverOptions.AuthRequired = false

	server, err := NewServer(mockSchema.TelepactSchema, handler, serverOptions)
	if err != nil {
		return nil, err
	}

	return &MockServer{
		Server:                           server,
		Random:                           random,
		EnableGeneratedDefaultStub:       options.EnableMessageResponseGeneration,
		EnableOptionalFieldGeneration:    options.EnableOptionalFieldGeneration,
		RandomizeOptionalFieldGeneration: options.RandomizeOptionalFieldGeneration,
	}, nil
}

// Process handles a request message using the mock server
func (ms *MockServer) Process(ctx context.Context, messageBytes []byte) ([]byte, error) {
	response, err := ms.Server.Process(ctx, messageBytes, nil)
	if err != nil {
		return nil, err
	}
	return response.Bytes, nil
}
