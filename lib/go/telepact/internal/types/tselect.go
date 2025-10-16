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

const selectName = "Object"

// TSelect represents a select type
type TSelect struct {
	PossibleSelects map[string]interface{}
}

// NewTSelect creates a new TSelect
func NewTSelect(possibleSelects map[string]interface{}) *TSelect {
	return &TSelect{
		PossibleSelects: possibleSelects,
	}
}

// GetTypeParameterCount returns the number of type parameters
func (t *TSelect) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a select
func (t *TSelect) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	// TODO: Implement validateSelect from internal/validation
	return nil
}

// GenerateRandomValue generates a random select value
func (t *TSelect) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	// TODO: Implement generateRandomSelect from internal/generation
	return nil
}

// GetName returns the type name
func (t *TSelect) GetName(ctx *validation.ValidateContext) string {
	return selectName
}
