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

// TInteger represents an integer type
type TInteger struct{}

func (t *TInteger) GetTypeParameterCount() int {
	return 0
}

func (t *TInteger) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *ValidateContext) []*ValidationFailure {
	switch v := value.(type) {
	case int, int32, int64, float64:
		// In JSON, numbers come as float64
		if f, ok := value.(float64); ok {
			if f != float64(int64(f)) {
				return []*ValidationFailure{{Path: "", Message: "expected integer"}}
			}
		}
		return nil
	default:
		_ = v
		return []*ValidationFailure{{Path: "", Message: "expected integer"}}
	}
}

func (t *TInteger) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *GenerateContext) interface{} {
	if useBlueprintValue {
		if i, ok := blueprintValue.(int); ok {
			return i
		}
		if f, ok := blueprintValue.(float64); ok {
			return int(f)
		}
	}
	// TODO: Use random generator from context
	return 42
}

func (t *TInteger) GetName(ctx *ValidateContext) string {
	return "Integer"
}
