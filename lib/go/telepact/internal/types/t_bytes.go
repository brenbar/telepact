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

// TBytes represents a bytes type
type TBytes struct{}

func (t *TBytes) GetTypeParameterCount() int {
	return 0
}

func (t *TBytes) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	return validation.ValidateBytes(value)
}

func (t *TBytes) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	if useBlueprintValue {
		if b, ok := blueprintValue.([]byte); ok {
			return b
		}
		if s, ok := blueprintValue.(string); ok {
			return []byte(s)
		}
	}
	// TODO: Use random generator from context
	return []byte{0x01, 0x02, 0x03, 0x04}
}

func (t *TBytes) GetName(ctx *validation.ValidateContext) string {
	return "Bytes"
}
