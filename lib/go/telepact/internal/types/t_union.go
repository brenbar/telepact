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

// TUnion represents a union type with multiple variants
type TUnion struct {
	Name     string
	Variants map[string]*TFieldDeclaration
}

func (t *TUnion) GetTypeParameterCount() int {
	return 0
}

func (t *TUnion) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *ValidateContext) []*ValidationFailure {
	obj, ok := value.(map[string]interface{})
	if !ok {
		return []*ValidationFailure{{Path: "", Message: "expected union object"}}
	}
	
	// A union should have exactly one key
	if len(obj) != 1 {
		return []*ValidationFailure{{Path: "", Message: "union must have exactly one variant"}}
	}
	
	var failures []*ValidationFailure
	
	// TODO: Validate the variant value
	for variantName := range obj {
		if _, exists := t.Variants[variantName]; !exists {
			failures = append(failures, &ValidationFailure{
				Path:    variantName,
				Message: "unknown union variant",
			})
		}
	}
	
	return failures
}

func (t *TUnion) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *GenerateContext) interface{} {
	if useBlueprintValue {
		if obj, ok := blueprintValue.(map[string]interface{}); ok {
			return obj
		}
	}
	
	// TODO: Generate random union variant
	result := make(map[string]interface{})
	
	// Pick first variant for now
	for variantName := range t.Variants {
		result[variantName] = make(map[string]interface{})
		break
	}
	
	return result
}

func (t *TUnion) GetName(ctx *ValidateContext) string {
	return "Union"
}
