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
	"fmt"

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
	headers := make(map[string]interface{})
	for k, v := range message.Headers {
		headers[k] = v
	}
	
	serializeAsBinary := false
	if val, ok := headers["@binary_"]; ok {
		if b, isBool := val.(bool); isBool && b {
			serializeAsBinary = true
		}
		delete(headers, "@binary_")
	}
	
	messageAsPseudoJSON := []interface{}{headers, message.Body}
	
	if serializeAsBinary {
		// Try binary encoding first
		encodedMessage, err := s.binaryEncoder.Encode(messageAsPseudoJSON)
		if err == nil {
			return s.serializationImpl.ToMsgpack(encodedMessage)
		}
	}
	
	// Default: direct JSON serialization (simplified for placeholder implementation)
	// In full implementation, would use base64 encoding
	return s.serializationImpl.ToJSON(messageAsPseudoJSON)
}

// Deserialize converts a byte array into a Message
func (s *Serializer) Deserialize(messageBytes []byte) (*Message, error) {
	var messageAsPseudoJSON interface{}
	var isMsgPack bool
	var err error
	
	// Check if it's MessagePack (starts with 0x92 for array of 2)
	if len(messageBytes) > 0 && messageBytes[0] == 0x92 {
		isMsgPack = true
		messageAsPseudoJSON, err = s.serializationImpl.FromMsgpack(messageBytes)
	} else {
		isMsgPack = false
		messageAsPseudoJSON, err = s.serializationImpl.FromJSON(messageBytes)
	}
	
	if err != nil {
		return nil, &SerializationError{Message: "invalid message format", Cause: err}
	}
	
	messageList, ok := messageAsPseudoJSON.([]interface{})
	if !ok {
		return nil, &SerializationError{Message: "message must be array"}
	}
	
	if len(messageList) != 2 {
		return nil, &SerializationError{Message: "message must have 2 elements"}
	}
	
	// For placeholder implementation, skip base64/binary decoding
	// In full implementation, would decode based on isMsgPack flag
	finalMessageList := messageList
	_ = isMsgPack // Will be used in full implementation
	
	headers, ok := finalMessageList[0].(map[string]interface{})
	if !ok {
		return nil, &SerializationError{Message: "headers must be object"}
	}
	
	body, ok := finalMessageList[1].(map[string]interface{})
	if !ok {
		return nil, &SerializationError{Message: "body must be object"}
	}
	
	if len(body) != 1 {
		return nil, &SerializationError{Message: "body must have exactly one entry"}
	}
	
	// Validate body value is an object
	for _, v := range body {
		if _, ok := v.(map[string]interface{}); !ok {
			return nil, &SerializationError{Message: fmt.Sprintf("body value must be object, got %T", v)}
		}
	}
	
	return &Message{
		Headers: headers,
		Body:    body,
	}, nil
}
