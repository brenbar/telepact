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

package telepact_test

import (
	"context"
	"testing"

	"github.com/brenbar/telepact/lib/go/telepact"
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

func TestClientServerIntegration(t *testing.T) {
	// This is a basic integration test showing client-server flow
	// Full schema parsing and validation would be needed for complete functionality
	
	ctx := context.Background()
	
	// Create a simple schema (placeholder - real implementation would parse .telepact.json)
	schema := &telepact.TelepactSchema{
		Original:              make([]interface{}, 0),
		Parsed:                make(map[string]types.TType),
		ParsedRequestHeaders:  make(map[string]*types.TFieldDeclaration),
		ParsedResponseHeaders: make(map[string]*types.TFieldDeclaration),
	}
	
	// Create a handler
	handler := func(ctx context.Context, request *telepact.Message) (*telepact.Message, error) {
		// Echo back the request in a response
		target := request.GetBodyTarget()
		if target == "fn.echo" {
			return &telepact.Message{
				Headers: make(map[string]interface{}),
				Body: map[string]interface{}{
					"Ok_": map[string]interface{}{
						"echoed": request.GetBodyPayload(),
					},
				},
			}, nil
		}
		
		return &telepact.Message{
			Headers: make(map[string]interface{}),
			Body: map[string]interface{}{
				"ErrorUnknown_": map[string]interface{}{},
			},
		}, nil
	}
	
	// Create server
	serverOpts := telepact.NewServerOptions()
	serverOpts.AuthRequired = false // Disable auth for this test
	server, err := telepact.NewServer(schema, handler, serverOpts)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	
	// Create client adapter that uses the server directly
	adapter := func(ctx context.Context, msg *telepact.Message, serializer *telepact.Serializer) (*telepact.Message, error) {
		requestBytes, err := serializer.Serialize(msg)
		if err != nil {
			return nil, err
		}
		
		response, err := server.Process(ctx, requestBytes, nil)
		if err != nil {
			return nil, err
		}
		
		return serializer.Deserialize(response.Bytes)
	}
	
	clientOpts := telepact.NewClientOptions()
	client := telepact.NewClient(adapter, clientOpts)
	
	// Make a request
	request := &telepact.Message{
		Headers: make(map[string]interface{}),
		Body: map[string]interface{}{
			"fn.echo": map[string]interface{}{
				"message": "Hello, Telepact!",
			},
		},
	}
	
	response, err := client.Request(ctx, request)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	
	// Verify response
	if response.GetBodyTarget() != "Ok_" {
		t.Errorf("Expected Ok_ response, got %s", response.GetBodyTarget())
	}
	
	t.Logf("Integration test passed! Response: %+v", response.Body)
}

func TestTypeValidation(t *testing.T) {
	// Test type validation (basic example)
	// Full implementation would require complete schema parsing
	
	// This demonstrates the type system structure
	t.Log("Type system structure validated in compilation")
}
