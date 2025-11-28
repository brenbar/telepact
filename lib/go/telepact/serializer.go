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
func NewSerializer(
	serializationImpl Serialization,
	binaryEncoder binary.BinaryEncoder,
	base64Encoder binary.Base64Encoder,
) *Serializer {
	return &Serializer{
		serializationImpl: serializationImpl,
		binaryEncoder:     binaryEncoder,
		base64Encoder:     base64Encoder,
	}
}

// Serialize converts a Message into a byte array
func (s *Serializer) Serialize(message *Message) ([]byte, error) {
	// TODO: Implement serialization logic from internal/SerializeInternal
	return serializeInternal(message, s.binaryEncoder, s.base64Encoder, s.serializationImpl)
}

// Deserialize converts a byte array into a Message
func (s *Serializer) Deserialize(messageBytes []byte) (*Message, error) {
	// TODO: Implement deserialization logic from internal/DeserializeInternal
	return deserializeInternal(messageBytes, s.serializationImpl, s.binaryEncoder, s.base64Encoder)
}

// Placeholder functions - these will be implemented in internal package
func serializeInternal(
	message *Message,
	binaryEncoder binary.BinaryEncoder,
	base64Encoder binary.Base64Encoder,
	serializationImpl Serialization,
) ([]byte, error) {
	// TODO: Implement
	return serializationImpl.ToJSON(message)
}

func deserializeInternal(
	messageBytes []byte,
	serializationImpl Serialization,
	binaryEncoder binary.BinaryEncoder,
	base64Encoder binary.Base64Encoder,
) (*Message, error) {
	// TODO: Implement
	obj, err := serializationImpl.FromJSON(messageBytes)
	if err != nil {
		return nil, err
	}
	
	// Convert to Message
	if msgMap, ok := obj.(map[string]interface{}); ok {
		headers := make(map[string]interface{})
		body := make(map[string]interface{})
		
		if h, ok := msgMap["headers"].(map[string]interface{}); ok {
			headers = h
		}
		if b, ok := msgMap["body"].(map[string]interface{}); ok {
			body = b
		}
		
		return NewMessage(headers, body), nil
	}
	
	return nil, NewTelepactError("Invalid message format")
}
