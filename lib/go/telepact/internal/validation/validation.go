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

// ValidateContext provides context for validation
type ValidateContext struct {
	Path            []string
	Select          map[string]interface{}
	Fn              string
	CoerceBase64    bool
	Base64Coercions map[string]interface{}
	BytesCoercions  map[string]interface{}
}

// NewValidateContext creates a new ValidateContext
func NewValidateContext(selectFields map[string]interface{}, fn string, coerceBase64 bool) *ValidateContext {
	return &ValidateContext{
		Path:            make([]string, 0),
		Select:          selectFields,
		Fn:              fn,
		CoerceBase64:    coerceBase64,
		Base64Coercions: make(map[string]interface{}),
		BytesCoercions:  make(map[string]interface{}),
	}
}

// ValidationFailure represents a validation failure
type ValidationFailure struct {
	Path    string
	Message string
	Value   interface{}
}

// InvalidMessage represents an invalid message error
type InvalidMessage struct {
	Message string
}

func (e *InvalidMessage) Error() string {
	return "Invalid message: " + e.Message
}

// InvalidMessageBody represents an invalid message body error
type InvalidMessageBody struct {
	Message string
}

func (e *InvalidMessageBody) Error() string {
	return "Invalid message body: " + e.Message
}

// GetTypeUnexpectedValidationFailure creates a validation failure for unexpected type
func GetTypeUnexpectedValidationFailure(path []string, value interface{}, expectedType string) []*ValidationFailure {
	pathStr := ""
	if len(path) > 0 {
		for i, p := range path {
			if i > 0 {
				pathStr += "."
			}
			pathStr += p
		}
	}
	
	return []*ValidationFailure{{
		Path:    pathStr,
		Message: "Expected " + expectedType,
		Value:   value,
	}}
}
