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
	"github.com/brenbar/telepact/lib/go/telepact/internal/binary"
)

// Serializer converts a Message to and from a serialized form
type Serializer struct {
	serializationImpl Serialization
	binaryEncoder     binary.BinaryEncoder
	base64Encoder     binary.Base64Encoder
}

// NewSerializer creates a new Serializer
func NewSerializer(serializationImpl Serialization, binaryEncoder binary.BinaryEncoder, base64Encoder binary.Base64Encoder) *Serializer {
	return &Serializer{
		serializationImpl: serializationImpl,
		binaryEncoder:     binaryEncoder,
		base64Encoder:     base64Encoder,
	}
}

// Serialize converts a Message into a byte array
func (s *Serializer) Serialize(message *Message) ([]byte, error) {
	// TODO: Implement full serialization logic
	return s.serializationImpl.ToJSON(message)
}

// Deserialize converts a byte array into a Message
func (s *Serializer) Deserialize(messageBytes []byte) (*Message, error) {
	// TODO: Implement full deserialization logic
	result, err := s.serializationImpl.FromJSON(messageBytes)
	if err != nil {
		return nil, err
	}
	
	// Convert result to Message
	msgMap, ok := result.(map[string]interface{})
	if !ok {
		return nil, &SerializationError{Message: "invalid message format"}
	}
	
	headers, _ := msgMap["headers"].(map[string]interface{})
	body, _ := msgMap["body"].(map[string]interface{})
	
	if headers == nil {
		headers = make(map[string]interface{})
	}
	if body == nil {
		body = make(map[string]interface{})
	}
	
	return &Message{
		Headers: headers,
		Body:    body,
	}, nil
}
