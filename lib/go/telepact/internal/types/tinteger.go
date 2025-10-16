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

const integerName = "Integer"

// TInteger represents an integer type
type TInteger struct{}

// GetTypeParameterCount returns the number of type parameters
func (t *TInteger) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as an integer
func (t *TInteger) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	return validation.ValidateInteger(value)
}

// GenerateRandomValue generates a random integer value
func (t *TInteger) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	return generation.GenerateRandomInteger(blueprintValue, useBlueprintValue, ctx)
}

// GetName returns the type name
func (t *TInteger) GetName(ctx *validation.ValidateContext) string {
	return integerName
}
