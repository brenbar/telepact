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

import (
	"sort"
)

// GenerateRandomSelectType generates a random select value (used internally by types package)
func GenerateRandomSelectType(possibleSelects map[string]interface{}, ctx *GenerateContext) interface{} {
	possibleSelect, exists := possibleSelects[ctx.FnScope]
	if !exists {
		return map[string]interface{}{}
	}
	return subSelect(possibleSelect, ctx)
}

// subSelect recursively generates select subsections
func subSelect(possibleSelectSection interface{}, ctx *GenerateContext) interface{} {
	// Check if possible select is a list
	if possibleList, ok := possibleSelectSection.([]interface{}); ok {
		var selectedFieldNames []interface{}
		for _, fieldName := range possibleList {
			if ctx.RandomGenerator.NextBoolean() {
				selectedFieldNames = append(selectedFieldNames, fieldName)
			}
		}
		
		// Sort for deterministic output
		var strFieldNames []string
		for _, fn := range selectedFieldNames {
			if str, ok := fn.(string); ok {
				strFieldNames = append(strFieldNames, str)
			}
		}
		sort.Strings(strFieldNames)
		
		// Convert back to []interface{}
		var result []interface{}
		for _, str := range strFieldNames {
			result = append(result, str)
		}
		return result
	}
	
	// Check if possible select is a map
	if possibleMap, ok := possibleSelectSection.(map[string]interface{}); ok {
		selectedSection := make(map[string]interface{})
		
		// Sort keys for deterministic ordering
		var keys []string
		for key := range possibleMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		
		for _, key := range keys {
			value := possibleMap[key]
			if ctx.RandomGenerator.NextBoolean() {
				result := subSelect(value, ctx)
				// Skip empty maps
				if resultMap, ok := result.(map[string]interface{}); ok && len(resultMap) == 0 {
					continue
				}
				selectedSection[key] = result
			}
		}
		
		return selectedSection
	}
	
	return nil
}
