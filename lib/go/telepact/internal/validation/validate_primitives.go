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

// ValidateString validates a string value
func ValidateString(value interface{}) []*ValidationFailure {
	if _, ok := value.(string); ok {
		return nil
	}
	return GetTypeUnexpectedValidationFailure([]string{}, value, "String")
}

// ValidateInteger validates an integer value
func ValidateInteger(value interface{}) []*ValidationFailure {
	switch v := value.(type) {
	case int, int32, int64:
		return nil
	case float64:
		// In JSON, numbers come as float64
		if v != float64(int64(v)) {
			return GetTypeUnexpectedValidationFailure([]string{}, value, "Integer")
		}
		return nil
	default:
		return GetTypeUnexpectedValidationFailure([]string{}, value, "Integer")
	}
}

// ValidateNumber validates a number value
func ValidateNumber(value interface{}) []*ValidationFailure {
	switch value.(type) {
	case int, int32, int64, float32, float64:
		return nil
	default:
		return GetTypeUnexpectedValidationFailure([]string{}, value, "Number")
	}
}

// ValidateBoolean validates a boolean value
func ValidateBoolean(value interface{}) []*ValidationFailure {
	if _, ok := value.(bool); ok {
		return nil
	}
	return GetTypeUnexpectedValidationFailure([]string{}, value, "Boolean")
}

// ValidateBytes validates a bytes value
func ValidateBytes(value interface{}) []*ValidationFailure {
	switch value.(type) {
	case []byte, string:
		return nil
	default:
		return GetTypeUnexpectedValidationFailure([]string{}, value, "Bytes")
	}
}
