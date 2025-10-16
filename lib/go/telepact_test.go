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
	"testing"
)

func TestMessage(t *testing.T) {
	headers := map[string]interface{}{"key": "value"}
	body := map[string]interface{}{
		"fn.test": map[string]interface{}{
			"arg": "val",
		},
	}

	msg := NewMessage(headers, body)

	if msg.Headers["key"] != "value" {
		t.Errorf("Expected header 'key' to be 'value', got %v", msg.Headers["key"])
	}

	target := msg.GetBodyTarget()
	if target != "fn.test" {
		t.Errorf("Expected target to be 'fn.test', got %s", target)
	}

	payload := msg.GetBodyPayload()
	if payload["arg"] != "val" {
		t.Errorf("Expected payload arg to be 'val', got %v", payload["arg"])
	}
}

func TestSerialization(t *testing.T) {
	serialization := NewDefaultSerialization()

	// Test map serialization
	data := map[string]interface{}{
		"key": "value",
		"num": float64(42),
	}

	bytes, err := serialization.ToJSON(data)
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	result, err := serialization.FromJSON(bytes)
	if err != nil {
		t.Fatalf("Failed to deserialize: %v", err)
	}

	if result["key"] != "value" {
		t.Errorf("Expected key to be 'value', got %v", result["key"])
	}

	if result["num"] != float64(42) {
		t.Errorf("Expected num to be 42, got %v", result["num"])
	}
}

func TestSerializer(t *testing.T) {
	serialization := NewDefaultSerialization()
	serializer := NewSerializer(serialization, nil, nil)

	headers := map[string]interface{}{}
	body := map[string]interface{}{
		"fn.test": map[string]interface{}{
			"arg": "val",
		},
	}
	msg := NewMessage(headers, body)

	// Serialize
	bytes, err := serializer.Serialize(msg)
	if err != nil {
		t.Fatalf("Failed to serialize message: %v", err)
	}

	// Deserialize
	result, err := serializer.Deserialize(bytes)
	if err != nil {
		t.Fatalf("Failed to deserialize message: %v", err)
	}

	if result.GetBodyTarget() != "fn.test" {
		t.Errorf("Expected target to be 'fn.test', got %s", result.GetBodyTarget())
	}
}

func TestServerClientIntegration(t *testing.T) {
	// Create a simple schema
	schema, err := FromJSON("[]")
	if err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	// Create server handler
	handler := func(ctx context.Context, msg *Message) (*Message, error) {
		target := msg.GetBodyTarget()
		if target == "fn.greet" {
			payload := msg.GetBodyPayload()
			name := payload["name"].(string)
			return NewMessage(
				map[string]interface{}{},
				map[string]interface{}{
					"Ok_": map[string]interface{}{
						"message": "Hello " + name,
					},
				},
			), nil
		}
		return nil, NewTelepactError("Unknown function", nil)
	}

	// Create server
	serverOptions := NewServerOptions()
	serverOptions.AuthRequired = false
	server, err := NewServer(schema, handler, serverOptions)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Create client adapter
	adapter := func(ctx context.Context, msg *Message, serializer *Serializer) (*Message, error) {
		// Serialize request
		requestBytes, err := serializer.Serialize(msg)
		if err != nil {
			return nil, err
		}

		// Process through server
		response, err := server.Process(ctx, requestBytes, nil)
		if err != nil {
			return nil, err
		}

		// Deserialize response
		return serializer.Deserialize(response.Bytes)
	}

	// Create client
	clientOptions := NewClientOptions()
	client := NewClient(adapter, clientOptions)

	// Make request
	requestMsg := NewMessage(
		map[string]interface{}{},
		map[string]interface{}{
			"fn.greet": map[string]interface{}{
				"name": "World",
			},
		},
	)

	responseMsg, err := client.Request(context.Background(), requestMsg)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	// Check response
	target := responseMsg.GetBodyTarget()
	if target != "Ok_" {
		t.Errorf("Expected response target to be 'Ok_', got %s", target)
	}

	payload := responseMsg.GetBodyPayload()
	message := payload["message"].(string)
	expected := "Hello World"
	if message != expected {
		t.Errorf("Expected message to be '%s', got '%s'", expected, message)
	}
}

func TestRandomGenerator(t *testing.T) {
	rng := NewRandomGenerator(0, 5)
	rng.SetSeed(42)

	// Test string generation
	str := rng.GenerateString()
	if len(str) == 0 {
		t.Error("Expected non-empty string")
	}

	// Test integer generation
	num := rng.GenerateInteger()
	if num < 0 {
		t.Error("Expected non-negative integer")
	}

	// Test boolean generation
	_ = rng.GenerateBoolean()

	// Test collection length
	length := rng.GenerateCollectionLength()
	if length < 0 || length > 5 {
		t.Errorf("Expected collection length between 0 and 5, got %d", length)
	}
}
