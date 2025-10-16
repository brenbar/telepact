// Copyright The Telepact Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"strings"

	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
	"github.com/brenbar/telepact/lib/go/telepact/internal/validation"
)

// ValidateHeaders validates request headers against parsed header definitions.
func ValidateHeaders(headers map[string]interface{}, parsedRequestHeaders map[string]*types.TFieldDeclaration, functionName string) []*validation.ValidationFailure {
	validationFailures := []*validation.ValidationFailure{}
	
	for header, headerValue := range headers {
		// Check if header starts with @
		if !strings.HasPrefix(header, "@") {
			validationFailures = append(validationFailures, &validation.ValidationFailure{
				Path:   []interface{}{header},
				Reason: "RequiredObjectKeyPrefixMissing",
				Data:   map[string]interface{}{"prefix": "@"},
			})
		}
		
		// Validate against header definition if it exists
		if field, ok := parsedRequestHeaders[header]; ok {
			fnName := &functionName
			ctx := &validation.ValidateContext{Fn: fnName}
			thisValidationFailures := field.TypeDeclaration.Validate(headerValue, ctx)
			
			// Prepend header name to path
			for _, failure := range thisValidationFailures {
				pathWithHeader := append([]interface{}{header}, failure.Path...)
				validationFailures = append(validationFailures, &validation.ValidationFailure{
					Path:   pathWithHeader,
					Reason: failure.Reason,
					Data:   failure.Data,
				})
			}
		}
	}
	
	return validationFailures
}
