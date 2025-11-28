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

package generation

// FieldInfo contains information about a field for generation
type FieldInfo struct {
	Name     string
	Optional bool
}

// GenerateRandomStructType generates a random struct value (used internally by types package)
func GenerateRandomStructType(blueprintValue interface{}, useBlueprintValue bool, fieldInfos []FieldInfo, generateField func(fieldName string, blueprintValue interface{}, useBlueprintValue bool) interface{}, ctx *GenerateContext) map[string]interface{} {
	startingStruct := make(map[string]interface{})
	if useBlueprintValue {
		if sm, ok := blueprintValue.(map[string]interface{}); ok {
			startingStruct = sm
		}
	}
	
	obj := make(map[string]interface{})
	for _, fieldInfo := range fieldInfos {
		fieldName := fieldInfo.Name
		thisUseBlueprintValue := false
		var thisBlueprintValue interface{}
		
		if val, exists := startingStruct[fieldName]; exists {
			thisUseBlueprintValue = true
			thisBlueprintValue = val
		}
		
		if thisUseBlueprintValue {
			value := generateField(fieldName, thisBlueprintValue, thisUseBlueprintValue)
			obj[fieldName] = value
		} else {
			if !fieldInfo.Optional {
				if !ctx.AlwaysIncludeRequiredFields && ctx.RandomGenerator.NextBoolean() {
					continue
				}
				value := generateField(fieldName, nil, false)
				obj[fieldName] = value
			} else {
				if !ctx.IncludeOptionalFields || (ctx.RandomizeOptionalFields && ctx.RandomGenerator.NextBoolean()) {
					continue
				}
				value := generateField(fieldName, nil, false)
				obj[fieldName] = value
			}
		}
	}
	
	return obj
}
