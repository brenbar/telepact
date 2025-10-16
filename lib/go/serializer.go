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
	"github.com/brenbar/telepact/lib/go/internal/binary"
)

// Serializer handles serialization and deserialization of messages
type Serializer struct {
	SerializationImpl Serialization
	BinaryEncoder     binary.BinaryEncoder
	Base64Encoder     binary.Base64Encoder
}

// NewSerializer creates a new Serializer
func NewSerializer(serializationImpl Serialization, binaryEncoder binary.BinaryEncoder, base64Encoder binary.Base64Encoder) *Serializer {
	return &Serializer{
		SerializationImpl: serializationImpl,
		BinaryEncoder:     binaryEncoder,
		Base64Encoder:     base64Encoder,
	}
}

// Serialize converts a Message into bytes
func (s *Serializer) Serialize(message *Message) ([]byte, error) {
	// Create a message array [headers, body]
	messageArray := []interface{}{message.Headers, message.Body}

	// Serialize to JSON
	return s.SerializationImpl.ToJSONArray(messageArray)
}

// Deserialize converts bytes into a Message
func (s *Serializer) Deserialize(messageBytes []byte) (*Message, error) {
	// Parse the JSON array
	messageArray, err := s.SerializationImpl.FromJSONArray(messageBytes)
	if err != nil {
		return nil, NewSerializationError("Failed to deserialize message", err)
	}

	if len(messageArray) != 2 {
		return nil, NewSerializationError("Invalid message format: expected array of length 2", nil)
	}

	// Extract headers and body
	headers, ok := messageArray[0].(map[string]interface{})
	if !ok {
		return nil, NewSerializationError("Invalid message format: headers must be an object", nil)
	}

	body, ok := messageArray[1].(map[string]interface{})
	if !ok {
		return nil, NewSerializationError("Invalid message format: body must be an object", nil)
	}

	return NewMessage(headers, body), nil
}
