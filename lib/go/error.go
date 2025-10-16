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

import "fmt"

// TelepactError represents an error in the Telepact library
type TelepactError struct {
	Message string
	Cause   error
}

// Error implements the error interface
func (e *TelepactError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *TelepactError) Unwrap() error {
	return e.Cause
}

// NewTelepactError creates a new TelepactError
func NewTelepactError(message string, cause error) *TelepactError {
	return &TelepactError{
		Message: message,
		Cause:   cause,
	}
}

// TelepactSchemaParseError represents a schema parsing error
type TelepactSchemaParseError struct {
	Message string
	Cause   error
}

// Error implements the error interface
func (e *TelepactSchemaParseError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Schema parse error: %s: %v", e.Message, e.Cause)
	}
	return fmt.Sprintf("Schema parse error: %s", e.Message)
}

// Unwrap returns the underlying error
func (e *TelepactSchemaParseError) Unwrap() error {
	return e.Cause
}

// NewTelepactSchemaParseError creates a new TelepactSchemaParseError
func NewTelepactSchemaParseError(message string, cause error) *TelepactSchemaParseError {
	return &TelepactSchemaParseError{
		Message: message,
		Cause:   cause,
	}
}

// SerializationError represents a serialization/deserialization error
type SerializationError struct {
	Message string
	Cause   error
}

// Error implements the error interface
func (e *SerializationError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Serialization error: %s: %v", e.Message, e.Cause)
	}
	return fmt.Sprintf("Serialization error: %s", e.Message)
}

// Unwrap returns the underlying error
func (e *SerializationError) Unwrap() error {
	return e.Cause
}

// NewSerializationError creates a new SerializationError
func NewSerializationError(message string, cause error) *SerializationError {
	return &SerializationError{
		Message: message,
		Cause:   cause,
	}
}
