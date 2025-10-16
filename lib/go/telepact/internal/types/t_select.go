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

// TSelect represents a select type (runtime type selection)
type TSelect struct {
	PossibleSelects map[string]interface{}
}

func (t *TSelect) GetTypeParameterCount() int {
	return 0
}

func (t *TSelect) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *ValidateContext) []*ValidationFailure {
	// TODO: Implement select validation
	return nil
}

func (t *TSelect) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *GenerateContext) interface{} {
	// TODO: Implement select random generation
	return make(map[string]interface{})
}

func (t *TSelect) GetName(ctx *ValidateContext) string {
	return "Object"
}
