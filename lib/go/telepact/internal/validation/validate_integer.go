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

package validation

import "math"

const integerName = "Integer"

// ValidateInteger validates a value as an integer
func ValidateInteger(value interface{}) []*ValidationFailure {
	switch v := value.(type) {
	case int:
		if v > math.MaxInt64 || v < math.MinInt64 {
			return []*ValidationFailure{
				NewValidationFailure([]interface{}{}, "NumberOutOfRange", map[string]interface{}{}),
			}
		}
		return nil
	case int8, int16, int32, int64:
		// These are always in range
		return nil
	case uint, uint8, uint16, uint32, uint64:
		// Check if value fits in int64
		val := int64(0)
		switch uv := v.(type) {
		case uint:
			if uv > math.MaxInt64 {
				return []*ValidationFailure{
					NewValidationFailure([]interface{}{}, "NumberOutOfRange", map[string]interface{}{}),
				}
			}
			val = int64(uv)
		case uint8:
			val = int64(uv)
		case uint16:
			val = int64(uv)
		case uint32:
			val = int64(uv)
		case uint64:
			if uv > math.MaxInt64 {
				return []*ValidationFailure{
					NewValidationFailure([]interface{}{}, "NumberOutOfRange", map[string]interface{}{}),
				}
			}
			val = int64(uv)
		}
		_ = val
		return nil
	case bool, float32, float64:
		// Reject booleans and floats
		return GetTypeUnexpectedValidationFailure([]interface{}{}, value, integerName)
	default:
		return GetTypeUnexpectedValidationFailure([]interface{}{}, value, integerName)
	}
}
