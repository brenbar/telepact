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
	"github.com/brenbar/telepact/lib/go/telepact/internal/generation"
)

const numberName = "Number"

// TNumber represents a number (float) type
type TNumber struct{}

// GetTypeParameterCount returns the number of type parameters
func (t *TNumber) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a number
func (t *TNumber) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	// TODO: Implement validateNumber from internal/validation
	switch value.(type) {
	case float32, float64:
		return nil
	default:
		return []*validation.ValidationFailure{
			validation.NewValidationFailure("Expected number", []string{}),
		}
	}
}

// GenerateRandomValue generates a random number value
func (t *TNumber) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	// TODO: Implement generateRandomNumber from internal/generation
	return 3.14
}

// GetName returns the type name
func (t *TNumber) GetName(ctx *validation.ValidateContext) string {
	return numberName
}
