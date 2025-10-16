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

const mockStubName = "_ext.Stub_"

// TMockStub represents a mock stub type for testing
type TMockStub struct {
	Types map[string]TType
}

// NewTMockStub creates a new TMockStub
func NewTMockStub(types map[string]TType) *TMockStub {
	return &TMockStub{
		Types: types,
	}
}

// GetTypeParameterCount returns the number of type parameters
func (t *TMockStub) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a mock stub
func (t *TMockStub) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	// TODO: Implement validate_mock_stub from internal/validation
	return nil
}

// GenerateRandomValue generates a random mock stub value
func (t *TMockStub) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	// TODO: Implement generate_random_mock_stub from internal/generation
	return make(map[string]interface{})
}

// GetName returns the type name
func (t *TMockStub) GetName(ctx *validation.ValidateContext) string {
	return mockStubName
}
