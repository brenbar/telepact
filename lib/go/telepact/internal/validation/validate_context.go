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
	// TODO: Add fields as needed for validation context
}

// NewValidateContext creates a new ValidateContext
func NewValidateContext() *ValidateContext {
	return &ValidateContext{}
}

// ValidationFailure represents a validation failure
type ValidationFailure struct {
	Path   []interface{}
	Reason string
	Data   map[string]interface{}
}

// NewValidationFailure creates a new ValidationFailure
func NewValidationFailure(path []interface{}, reason string, data map[string]interface{}) *ValidationFailure {
	return &ValidationFailure{
		Path:   path,
		Reason: reason,
		Data:   data,
	}
}
