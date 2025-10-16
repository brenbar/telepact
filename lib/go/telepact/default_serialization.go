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
	"encoding/json"

	"github.com/vmihailenco/msgpack/v5"
)

// DefaultSerialization provides the default implementation of Serialization
type DefaultSerialization struct{}

// ToJSON serializes an object to JSON
func (d *DefaultSerialization) ToJSON(telepactMessage interface{}) ([]byte, error) {
	return json.Marshal(telepactMessage)
}

// ToMsgpack serializes an object to MessagePack
func (d *DefaultSerialization) ToMsgpack(telepactMessage interface{}) ([]byte, error) {
	return msgpack.Marshal(telepactMessage)
}

// FromJSON deserializes JSON bytes to an object
func (d *DefaultSerialization) FromJSON(bytes []byte) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal(bytes, &result)
	return result, err
}

// FromMsgpack deserializes MessagePack bytes to an object
func (d *DefaultSerialization) FromMsgpack(bytes []byte) (interface{}, error) {
	var result interface{}
	err := msgpack.Unmarshal(bytes, &result)
	return result, err
}
