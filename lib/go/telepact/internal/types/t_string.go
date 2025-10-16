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

// TString represents a string type
type TString struct{}

func (t *TString) GetTypeParameterCount() int {
	return 0
}

func (t *TString) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *ValidateContext) []*ValidationFailure {
	if _, ok := value.(string); !ok {
		return []*ValidationFailure{{Path: "", Message: "expected string"}}
	}
	return nil
}

func (t *TString) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *GenerateContext) interface{} {
	if useBlueprintValue {
		if str, ok := blueprintValue.(string); ok {
			return str
		}
	}
	// TODO: Use random generator from context
	return "alpha"
}

func (t *TString) GetName(ctx *ValidateContext) string {
	return "String"
}
