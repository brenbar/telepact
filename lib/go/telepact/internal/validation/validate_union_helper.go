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

const unionName = "Object"

// ValidateUnionType validates a value as a union type (used internally by types package)
func ValidateUnionType(value interface{}, name string, validateTags func(map[string]interface{}) []*ValidationFailure, ctx *ValidateContext) []*ValidationFailure {
	if obj, ok := value.(map[string]interface{}); ok {
		return validateTags(obj)
	}
	return GetTypeUnexpectedValidationFailure([]interface{}{}, value, unionName)
}
