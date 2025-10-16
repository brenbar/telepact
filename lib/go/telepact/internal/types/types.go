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
	"github.com/brenbar/telepact/lib/go/telepact/internal/generation"
	"github.com/brenbar/telepact/lib/go/telepact/internal/validation"
)

// TType is the base interface for all Telepact types
type TType interface {
	GetTypeParameterCount() int
	Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure
	GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{}
	GetName(ctx *validation.ValidateContext) string
}

// TTypeDeclaration represents a type declaration with parameters
type TTypeDeclaration struct {
	Type           TType
	Nullable       bool
	TypeParameters []*TTypeDeclaration
}

// Validate validates a value against this type declaration
func (td *TTypeDeclaration) Validate(value interface{}, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	// TODO: Implement full validation with nullable handling
	if value == nil && td.Nullable {
		return nil
	}
	if value == nil && !td.Nullable {
		return []*validation.ValidationFailure{{Path: "", Message: "value cannot be null"}}
	}
	return td.Type.Validate(value, td.TypeParameters, ctx)
}

// GenerateRandomValue generates a random value for this type declaration
func (td *TTypeDeclaration) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, ctx *generation.GenerateContext) interface{} {
	// TODO: Implement full generation with nullable handling
	if td.Nullable {
		// Simplified: always generate non-null for now
	}
	return td.Type.GenerateRandomValue(blueprintValue, useBlueprintValue, td.TypeParameters, ctx)
}

// TFieldDeclaration represents a field declaration in a struct
type TFieldDeclaration struct {
	FieldName       string
	TypeDeclaration *TTypeDeclaration
	Optional        bool
}
