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

// MockServerOptions contains options for configuring a MockServer
type MockServerOptions struct {
	OnError                              func(error)
	EnableMessageResponseGeneration      bool
	EnableOptionalFieldGeneration        bool
	RandomizeOptionalFieldGeneration     bool
	GeneratedCollectionLengthMin         int
	GeneratedCollectionLengthMax         int
}

// NewMockServerOptions creates default mock server options
func NewMockServerOptions() *MockServerOptions {
	return &MockServerOptions{
		OnError:                              func(e error) {},
		EnableMessageResponseGeneration:      true,
		EnableOptionalFieldGeneration:        true,
		RandomizeOptionalFieldGeneration:     true,
		GeneratedCollectionLengthMin:         0,
		GeneratedCollectionLengthMax:         3,
	}
}

// MockServer is a mock instance of a telepact server
type MockServer struct {
	Random                               *RandomGenerator
	EnableGeneratedDefaultStub           bool
	EnableOptionalFieldGeneration        bool
	RandomizeOptionalFieldGeneration     bool
	Server                               *Server
}

// NewMockServer creates a new mock server
func NewMockServer(mockTelepactSchema *MockTelepactSchema, options *MockServerOptions) (*MockServer, error) {
	random := NewRandomGenerator(options.GeneratedCollectionLengthMin, options.GeneratedCollectionLengthMax)

	serverOptions := NewServerOptions()
	serverOptions.OnError = options.OnError
	serverOptions.AuthRequired = false

	// Create a handler that will mock responses
	handler := func(ctx context.Context, message *Message) (*Message, error) {
		// TODO: Implementation of mock handling
		return nil, &TelepactError{Message: "Not implemented"}
	}

	server, err := NewServer(&mockTelepactSchema.TelepactSchema, handler, serverOptions)
	if err != nil {
		return nil, err
	}

	return &MockServer{
		Random:                               random,
		EnableGeneratedDefaultStub:           options.EnableMessageResponseGeneration,
		EnableOptionalFieldGeneration:        options.EnableOptionalFieldGeneration,
		RandomizeOptionalFieldGeneration:     options.RandomizeOptionalFieldGeneration,
		Server:                               server,
	}, nil
}

// Process processes a telepact request message into a response
func (m *MockServer) Process(ctx context.Context, message []byte) ([]byte, error) {
	response, err := m.Server.Process(ctx, message, nil)
	if err != nil {
		return nil, err
	}
	return response.Bytes, nil
}
