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

// TArray represents an array type with element type
type TArray struct{}

func (t *TArray) GetTypeParameterCount() int {
	return 1
}

func (t *TArray) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *ValidateContext) []*ValidationFailure {
	arr, ok := value.([]interface{})
	if !ok {
		return []*ValidationFailure{{Path: "", Message: "expected array"}}
	}
	
	var failures []*ValidationFailure
	
	// TODO: Validate each element against type parameter
	for i := range arr {
		_ = i
		// Validate arr[i] against typeParameters[0]
	}
	
	return failures
}

func (t *TArray) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *GenerateContext) interface{} {
	result := make([]interface{}, 0)
	
	// TODO: Generate random array elements
	if useBlueprintValue {
		if arr, ok := blueprintValue.([]interface{}); ok {
			return arr
		}
	}
	
	return result
}

func (t *TArray) GetName(ctx *ValidateContext) string {
	return "Array"
}
