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

import "fmt"

// InvalidMessage represents an error for invalid messages
type InvalidMessage struct {
	Cause error
}

// NewInvalidMessage creates a new InvalidMessage error
func NewInvalidMessage(cause error) *InvalidMessage {
	return &InvalidMessage{Cause: cause}
}

// Error returns the error message
func (e *InvalidMessage) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return "invalid message"
}

// Unwrap returns the underlying cause
func (e *InvalidMessage) Unwrap() error {
	return e.Cause
}

// Is checks if the target error is an InvalidMessage
func (e *InvalidMessage) Is(target error) bool {
	_, ok := target.(*InvalidMessage)
	return ok
}

// As allows error unwrapping
func (e *InvalidMessage) As(target interface{}) bool {
	if t, ok := target.(**InvalidMessage); ok {
		*t = e
		return true
	}
	return false
}

// Format implements fmt.Formatter
func (e *InvalidMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			if e.Cause != nil {
				fmt.Fprintf(s, "%+v", e.Cause)
			} else {
				fmt.Fprint(s, "invalid message")
			}
		} else {
			fmt.Fprint(s, e.Error())
		}
	case 's':
		fmt.Fprint(s, e.Error())
	}
}
