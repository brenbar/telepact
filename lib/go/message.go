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

// Message represents a Telepact message with headers and body
type Message struct {
	Headers map[string]interface{}
	Body    map[string]interface{}
}

// NewMessage creates a new Message
func NewMessage(headers, body map[string]interface{}) *Message {
	return &Message{
		Headers: headers,
		Body:    body,
	}
}

// GetBodyTarget returns the first key in the body (the target function or type)
func (m *Message) GetBodyTarget() string {
	for key := range m.Body {
		return key
	}
	return ""
}

// GetBodyPayload returns the payload associated with the body target
func (m *Message) GetBodyPayload() map[string]interface{} {
	target := m.GetBodyTarget()
	if target == "" {
		return nil
	}
	payload, ok := m.Body[target].(map[string]interface{})
	if !ok {
		return nil
	}
	return payload
}
