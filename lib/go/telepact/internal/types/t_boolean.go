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

// TBoolean represents a boolean type
type TBoolean struct{}

func (t *TBoolean) GetTypeParameterCount() int {
	return 0
}

func (t *TBoolean) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	if _, ok := value.(bool); !ok {
		return []*validation.ValidationFailure{{Path: "", Message: "expected boolean"}}
	}
	return nil
}

func (t *TBoolean) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	if useBlueprintValue {
		if b, ok := blueprintValue.(bool); ok {
			return b
		}
	}
	// TODO: Use random generator from context
	return true
}

func (t *TBoolean) GetName(ctx *validation.ValidateContext) string {
	return "Boolean"
}
