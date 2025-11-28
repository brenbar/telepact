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

// TTypeDeclaration represents a type declaration
type TTypeDeclaration struct {
	Type           TType
	Nullable       bool
	TypeParameters []*TTypeDeclaration
}

// NewTTypeDeclaration creates a new TTypeDeclaration
func NewTTypeDeclaration(ttype TType, nullable bool, typeParameters []*TTypeDeclaration) *TTypeDeclaration {
	return &TTypeDeclaration{
		Type:           ttype,
		Nullable:       nullable,
		TypeParameters: typeParameters,
	}
}

// Validate validates a value against this type declaration
func (td *TTypeDeclaration) Validate(value interface{}, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	return ValidateValueOfType(value, td.Type, td.Nullable, td.TypeParameters, ctx)
}

// GenerateRandomValue generates a random value for this type declaration
func (td *TTypeDeclaration) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, ctx *generation.GenerateContext) interface{} {
	return GenerateRandomValueOfType(blueprintValue, useBlueprintValue, td.Type, td.Nullable, td.TypeParameters, ctx)
}
