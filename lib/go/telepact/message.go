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

// Message represents a telepact message with headers and body
type Message struct {
	Headers map[string]interface{}
	Body    map[string]interface{}
}

// NewMessage creates a new Message
func NewMessage(headers map[string]interface{}, body map[string]interface{}) *Message {
	return &Message{
		Headers: headers,
		Body:    body,
	}
}

// GetBodyTarget returns the first key in the body (the target function name)
func (m *Message) GetBodyTarget() string {
	for key := range m.Body {
		return key
	}
	return ""
}

// GetBodyPayload returns the payload associated with the first body key
func (m *Message) GetBodyPayload() map[string]interface{} {
	for _, value := range m.Body {
		if payload, ok := value.(map[string]interface{}); ok {
			return payload
		}
	}
	return nil
}

// TypedMessage represents a typed message with generic body type
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
