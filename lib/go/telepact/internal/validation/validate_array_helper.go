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

const arrayName = "Array"

// ValidateArrayType validates a value as an array type (used internally by types package)
func ValidateArrayType(value interface{}, validateElement func(interface{}) []*ValidationFailure, ctx *ValidateContext) []*ValidationFailure {
	// Check if value is a slice
	if arr, ok := value.([]interface{}); ok {
		var validationFailures []*ValidationFailure
		for i, element := range arr {
			// Add "*" to path
			ctx.Path = append(ctx.Path, "*")
			
			// Validate the element
			nestedValidationFailures := validateElement(element)
			
			// Remove "*" from path
			ctx.Path = ctx.Path[:len(ctx.Path)-1]
			
			// Add index to path for failures
			for _, f := range nestedValidationFailures {
				finalPath := append([]interface{}{i}, f.Path...)
				validationFailures = append(validationFailures, NewValidationFailure(finalPath, f.Reason, f.Data))
			}
		}
		
		return validationFailures
	}
	
	return GetTypeUnexpectedValidationFailure([]interface{}{}, value, arrayName)
}
