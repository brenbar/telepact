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

// TObject represents a map/object type with value type
type TObject struct{}

func (t *TObject) GetTypeParameterCount() int {
	return 1
}

func (t *TObject) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *ValidateContext) []*ValidationFailure {
	obj, ok := value.(map[string]interface{})
	if !ok {
		return []*ValidationFailure{{Path: "", Message: "expected object"}}
	}
	
	var failures []*ValidationFailure
	
	// TODO: Validate each value against type parameter
	for key := range obj {
		_ = key
		// Validate obj[key] against typeParameters[0]
	}
	
	return failures
}

func (t *TObject) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *GenerateContext) interface{} {
	result := make(map[string]interface{})
	
	// TODO: Generate random object values
	if useBlueprintValue {
		if obj, ok := blueprintValue.(map[string]interface{}); ok {
			return obj
		}
	}
	
	return result
}

func (t *TObject) GetName(ctx *ValidateContext) string {
	return "Object"
}
