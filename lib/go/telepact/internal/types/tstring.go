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

const stringName = "String"

// TString represents a string type
type TString struct{}

// GetTypeParameterCount returns the number of type parameters
func (t *TString) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a string
func (t *TString) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	return validation.ValidateString(value)
}

// GenerateRandomValue generates a random string value
func (t *TString) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	return generation.GenerateRandomString(blueprintValue, useBlueprintValue, ctx)
}

// GetName returns the type name
func (t *TString) GetName(ctx *validation.ValidateContext) string {
	return stringName
}
