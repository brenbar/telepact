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

// TType is the base interface for all Telepact types
type TType interface {
	GetTypeParameterCount() int
	Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *ValidateContext) []*ValidationFailure
	GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *GenerateContext) interface{}
	GetName(ctx *ValidateContext) string
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

// ValidateContext provides context for validation
type ValidateContext struct {
	// Placeholder
}

// ValidationFailure represents a validation failure
type ValidationFailure struct {
	Path    string
	Message string
}

// GenerateContext provides context for generation
type GenerateContext struct {
	// Placeholder
}
