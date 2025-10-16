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
	"testing"
)

func TestMessageCreation(t *testing.T) {
	msg := &Message{
		Headers: map[string]interface{}{"test": "header"},
		Body:    map[string]interface{}{"fn.test": map[string]interface{}{"arg": "value"}},
	}

	if msg.GetBodyTarget() != "fn.test" {
		t.Errorf("Expected body target 'fn.test', got '%s'", msg.GetBodyTarget())
	}
}

func TestRandomGenerator(t *testing.T) {
	rg := NewRandomGenerator(0, 3)
	rg.SetSeed(42)

	// Test that it generates consistent values with the same seed
	val1 := rg.NextInt()
	rg.SetSeed(42)
	val2 := rg.NextInt()

	if val1 != val2 {
		t.Errorf("Expected consistent random values, got %d and %d", val1, val2)
	}
}

func TestClientOptionsCreation(t *testing.T) {
	options := NewClientOptions()
	if options.TimeoutMsDefault != 5000 {
		t.Errorf("Expected default timeout 5000, got %d", options.TimeoutMsDefault)
	}
}

func TestServerOptionsCreation(t *testing.T) {
	options := NewServerOptions()
	if !options.AuthRequired {
		t.Error("Expected auth required to be true by default")
	}
}
