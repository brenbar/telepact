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

// Serialization defines the interface for custom serialization implementations
type Serialization interface {
	// ToJSON converts a map to JSON bytes
	ToJSON(data map[string]interface{}) ([]byte, error)

	// FromJSON parses JSON bytes into a map
	FromJSON(data []byte) (map[string]interface{}, error)

	// ToJSONArray converts a slice to JSON bytes (for array-based messages)
	ToJSONArray(data []interface{}) ([]byte, error)

	// FromJSONArray parses JSON bytes into a slice
	FromJSONArray(data []byte) ([]interface{}, error)
}

// DefaultSerialization provides a default JSON serialization implementation
type DefaultSerialization struct{}

// NewDefaultSerialization creates a new DefaultSerialization
func NewDefaultSerialization() *DefaultSerialization {
	return &DefaultSerialization{}
}

// ToJSON converts a map to JSON bytes
func (s *DefaultSerialization) ToJSON(data map[string]interface{}) ([]byte, error) {
	return Marshal(data)
}

// FromJSON parses JSON bytes into a map
func (s *DefaultSerialization) FromJSON(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := Unmarshal(data, &result)
	return result, err
}

// ToJSONArray converts a slice to JSON bytes
func (s *DefaultSerialization) ToJSONArray(data []interface{}) ([]byte, error) {
	return Marshal(data)
}

// FromJSONArray parses JSON bytes into a slice
func (s *DefaultSerialization) FromJSONArray(data []byte) ([]interface{}, error) {
	var result []interface{}
	err := Unmarshal(data, &result)
	return result, err
}
