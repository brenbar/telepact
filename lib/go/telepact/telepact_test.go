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
	"testing"
)

func TestMessage(t *testing.T) {
	headers := map[string]interface{}{"auth": "token123"}
	body := map[string]interface{}{"fn.greet": map[string]interface{}{"subject": "World"}}
	
	msg := NewMessage(headers, body)
	
	if msg.Headers["auth"] != "token123" {
		t.Errorf("Expected auth header to be 'token123', got %v", msg.Headers["auth"])
	}
	
	if msg.GetBodyTarget() != "fn.greet" {
		t.Errorf("Expected body target to be 'fn.greet', got %v", msg.GetBodyTarget())
	}
}

func TestResponse(t *testing.T) {
	bytes := []byte{1, 2, 3, 4}
	headers := map[string]interface{}{"status": "ok"}
	
	resp := NewResponse(bytes, headers)
	
	if len(resp.Bytes) != 4 {
		t.Errorf("Expected bytes length to be 4, got %d", len(resp.Bytes))
	}
	
	if resp.Headers["status"] != "ok" {
		t.Errorf("Expected status header to be 'ok', got %v", resp.Headers["status"])
	}
}

func TestRandomGenerator(t *testing.T) {
	rg := NewRandomGenerator(1, 5)
	
	// Test that SetSeed works
	rg.SetSeed(42)
	
	// Test NextInt
	val1 := rg.NextInt()
	if val1 < 0 {
		t.Errorf("NextInt should return non-negative value, got %d", val1)
	}
	
	// Test NextBoolean
	boolVal := rg.NextBoolean()
	_ = boolVal // Just verify it doesn't panic
	
	// Test NextString
	strVal := rg.NextString()
	if strVal == "" {
		t.Error("NextString should return non-empty string")
	}
	
	// Test NextDouble
	dblVal := rg.NextDouble()
	if dblVal < 0.0 || dblVal > 1.0 {
		t.Errorf("NextDouble should return value between 0 and 1, got %f", dblVal)
	}
	
	// Test NextCollectionLength
	lenVal := rg.NextCollectionLength()
	if lenVal < 1 || lenVal >= 5 {
		t.Errorf("NextCollectionLength should return value between 1 and 4, got %d", lenVal)
	}
}

func TestDefaultSerialization(t *testing.T) {
	ser := NewDefaultSerialization()
	
	testData := map[string]interface{}{
		"name": "test",
		"value": float64(42),
	}
	
	// Test ToJSON
	jsonBytes, err := ser.ToJSON(testData)
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	
	if len(jsonBytes) == 0 {
		t.Error("ToJSON returned empty bytes")
	}
	
	// Test FromJSON
	result, err := ser.FromJSON(jsonBytes)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}
	
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("FromJSON did not return a map")
	}
	
	if resultMap["name"] != "test" {
		t.Errorf("Expected name to be 'test', got %v", resultMap["name"])
	}
	
	// Test ToMsgpack
	msgpackBytes, err := ser.ToMsgpack(testData)
	if err != nil {
		t.Fatalf("ToMsgpack failed: %v", err)
	}
	
	if len(msgpackBytes) == 0 {
		t.Error("ToMsgpack returned empty bytes")
	}
	
	// Test FromMsgpack
	result2, err := ser.FromMsgpack(msgpackBytes)
	if err != nil {
		t.Fatalf("FromMsgpack failed: %v", err)
	}
	
	resultMap2, ok := result2.(map[string]interface{})
	if !ok {
		t.Fatal("FromMsgpack did not return a map")
	}
	
	if resultMap2["name"] != "test" {
		t.Errorf("Expected name to be 'test', got %v", resultMap2["name"])
	}
}

func TestTelepactError(t *testing.T) {
	err := NewTelepactError("test error")
	if err.Error() != "test error" {
		t.Errorf("Expected error message 'test error', got %v", err.Error())
	}
	
	cause := NewTelepactError("cause")
	errWithCause := NewTelepactErrorWithCause("wrapper", cause)
	if errWithCause.Unwrap() != cause {
		t.Error("Expected Unwrap to return the cause")
	}
}
