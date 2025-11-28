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

package types

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/validation"
)

// ValidateValueOfType validates a value against a type, handling nullable types
func ValidateValueOfType(value interface{}, thisType TType, nullable bool, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	if value == nil {
		if !nullable {
			return validation.GetTypeUnexpectedValidationFailure([]interface{}{}, value, thisType.GetName(ctx))
		}
		return nil
	}
	
	return thisType.Validate(value, typeParameters, ctx)
}
