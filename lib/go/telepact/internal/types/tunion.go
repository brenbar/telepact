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

const unionName = "Object"

// TUnion represents a union type
type TUnion struct {
	Name        string
	Tags        map[string]*TStruct
	TagIndices  map[string]int
}

// NewTUnion creates a new TUnion
func NewTUnion(name string, tags map[string]*TStruct, tagIndices map[string]int) *TUnion {
	return &TUnion{
		Name:       name,
		Tags:       tags,
		TagIndices: tagIndices,
	}
}

// GetTypeParameterCount returns the number of type parameters
func (t *TUnion) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a union
func (t *TUnion) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	// TODO: Implement validateUnion from internal/validation
	return nil
}

// GenerateRandomValue generates a random union value
func (t *TUnion) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	// TODO: Implement generateRandomUnion from internal/generation
	return make(map[string]interface{})
}

// GetName returns the type name
func (t *TUnion) GetName(ctx *validation.ValidateContext) string {
	return unionName
}
