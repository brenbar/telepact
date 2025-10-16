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

import "errors"

// TelepactError indicates critical failure in telepact processing logic.
type TelepactError struct {
	Message string
	Cause   error
}

func (e *TelepactError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

func (e *TelepactError) Unwrap() error {
	return e.Cause
}

// NewTelepactError creates a new TelepactError
func NewTelepactError(message string) *TelepactError {
	return &TelepactError{Message: message}
}

// NewTelepactErrorWithCause creates a new TelepactError with a cause
func NewTelepactErrorWithCause(message string, cause error) *TelepactError {
	return &TelepactError{Message: message, Cause: cause}
}

// SerializationError represents a serialization error
var SerializationError = errors.New("serialization error")

// TelepactSchemaParseError represents a schema parsing error
type TelepactSchemaParseError struct {
	Message string
	Details map[string]interface{}
}

func (e *TelepactSchemaParseError) Error() string {
	return e.Message
}

// NewTelepactSchemaParseError creates a new TelepactSchemaParseError
func NewTelepactSchemaParseError(message string, details map[string]interface{}) *TelepactSchemaParseError {
	return &TelepactSchemaParseError{Message: message, Details: details}
}
