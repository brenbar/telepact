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
	"encoding/json"

	"github.com/vmihailenco/msgpack/v5"
)

// DefaultSerialization provides default JSON and MessagePack serialization
type DefaultSerialization struct{}

// ToJSON converts a message to JSON bytes
func (d *DefaultSerialization) ToJSON(message interface{}) ([]byte, error) {
	return json.Marshal(message)
}

// ToMsgpack converts a message to MessagePack bytes
func (d *DefaultSerialization) ToMsgpack(message interface{}) ([]byte, error) {
	return msgpack.Marshal(message)
}

// FromJSON converts JSON bytes to a message
func (d *DefaultSerialization) FromJSON(bytes []byte) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal(bytes, &result)
	return result, err
}

// FromMsgpack converts MessagePack bytes to a message
func (d *DefaultSerialization) FromMsgpack(bytes []byte) (interface{}, error) {
	var result interface{}
	err := msgpack.Unmarshal(bytes, &result)
	return result, err
}
