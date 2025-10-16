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

// TTypeDeclaration represents a type declaration
type TTypeDeclaration struct {
	// Placeholder
}

// TFieldDeclaration represents a field declaration
type TFieldDeclaration struct {
	Name         string
	Type         TType
	Optional     bool
	DefaultValue interface{}
}
