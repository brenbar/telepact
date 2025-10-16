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

const objectName = "Object"

// TObject represents an object (map) type
type TObject struct{}

// GetTypeParameterCount returns the number of type parameters
func (t *TObject) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as an object
func (t *TObject) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	// TODO: Implement validateObject from internal/validation
	if _, ok := value.(map[string]interface{}); !ok {
		return validation.GetTypeUnexpectedValidationFailure([]interface{}{}, value, objectName)
	}
	return nil
}

// GenerateRandomValue generates a random object value
func (t *TObject) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	// TODO: Implement generateRandomObject from internal/generation
	return make(map[string]interface{})
}

// GetName returns the type name
func (t *TObject) GetName(ctx *validation.ValidateContext) string {
	return objectName
}
