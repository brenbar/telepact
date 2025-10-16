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

// Message represents a Telepact message with headers and body
type Message struct {
	Headers map[string]interface{}
	Body    map[string]interface{}
}

// GetBodyTarget returns the target (first key) from the message body
func (m *Message) GetBodyTarget() string {
	for key := range m.Body {
		return key
	}
	return ""
}

// GetBodyPayload returns the payload (first value) from the message body
func (m *Message) GetBodyPayload() map[string]interface{} {
	for _, value := range m.Body {
		if payload, ok := value.(map[string]interface{}); ok {
			return payload
		}
		return make(map[string]interface{})
	}
	return make(map[string]interface{})
}
