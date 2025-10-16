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
	"testing"

	"github.com/brenbar/telepact/lib/go/telepact"
)

func TestMessage(t *testing.T) {
	msg := &telepact.Message{
		Headers: map[string]interface{}{"auth": "token"},
		Body:    map[string]interface{}{"fn.greet": map[string]interface{}{"subject": "world"}},
	}

	target := msg.GetBodyTarget()
	if target != "fn.greet" {
		t.Errorf("Expected target 'fn.greet', got '%s'", target)
	}

	payload := msg.GetBodyPayload()
	if payload["subject"] != "world" {
		t.Errorf("Expected subject 'world', got '%v'", payload["subject"])
	}
}

func TestRandomGenerator(t *testing.T) {
	rng := telepact.NewRandomGenerator(0, 3)
	rng.SetSeed(42)

	// Test deterministic generation
	val1 := rng.NextInt()
	rng.SetSeed(42)
	val2 := rng.NextInt()

	if val1 != val2 {
		t.Errorf("Random generator not deterministic: %d != %d", val1, val2)
	}

	// Test string generation
	str := rng.NextString()
	if str == "" {
		t.Error("Expected non-empty string")
	}

	// Test boolean generation
	_ = rng.NextBoolean()

	// Test bytes generation
	bytes := rng.NextBytes()
	if len(bytes) != 4 {
		t.Errorf("Expected 4 bytes, got %d", len(bytes))
	}
}

func TestSerialization(t *testing.T) {
	ser := &telepact.DefaultSerialization{}

	msg := map[string]interface{}{
		"hello": "world",
		"count": 42,
	}

	// Test JSON
	jsonBytes, err := ser.ToJSON(msg)
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	decoded, err := ser.FromJSON(jsonBytes)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	decodedMap := decoded.(map[string]interface{})
	if decodedMap["hello"] != "world" {
		t.Errorf("Expected 'world', got %v", decodedMap["hello"])
	}

	// Test MessagePack
	msgpackBytes, err := ser.ToMsgpack(msg)
	if err != nil {
		t.Fatalf("ToMsgpack failed: %v", err)
	}

	decoded2, err := ser.FromMsgpack(msgpackBytes)
	if err != nil {
		t.Fatalf("FromMsgpack failed: %v", err)
	}

	decodedMap2 := decoded2.(map[string]interface{})
	if decodedMap2["hello"] != "world" {
		t.Errorf("Expected 'world', got %v", decodedMap2["hello"])
	}
}
