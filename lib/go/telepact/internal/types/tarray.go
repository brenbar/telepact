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

const arrayName = "Array"

// TArray represents an array type
type TArray struct{}

// GetTypeParameterCount returns the number of type parameters
func (t *TArray) GetTypeParameterCount() int {
	return 1
}

// Validate validates a value as an array
func (t *TArray) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	// TODO: Implement validateArray from internal/validation
	return nil
}

// GenerateRandomValue generates a random array value
func (t *TArray) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	// TODO: Implement generateRandomArray from internal/generation
	return []interface{}{}
}

// GetName returns the type name
func (t *TArray) GetName(ctx *validation.ValidateContext) string {
	return arrayName
}
