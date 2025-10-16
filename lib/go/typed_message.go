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

// TypedMessage represents a Telepact message with a typed body
// This is useful for type-safe message handling in Go
type TypedMessage[T any] struct {
	Headers map[string]interface{}
	Body    T
}

// NewTypedMessage creates a new TypedMessage
func NewTypedMessage[T any](headers map[string]interface{}, body T) *TypedMessage[T] {
	return &TypedMessage[T]{
		Headers: headers,
		Body:    body,
	}
}

// ToMessage converts a TypedMessage to a regular Message
func (tm *TypedMessage[T]) ToMessage() *Message {
	// Convert the typed body to a generic interface{}
	return NewMessage(tm.Headers, map[string]interface{}{"data": tm.Body})
}

// FromMessage converts a regular Message to a TypedMessage
// Note: This requires type assertion and may panic if types don't match
func FromMessage[T any](msg *Message) *TypedMessage[T] {
	var body T
	// This is a simplified conversion - in practice you'd want better type checking
	if data, ok := msg.Body["data"].(T); ok {
		body = data
	}
	return NewTypedMessage(msg.Headers, body)
}
