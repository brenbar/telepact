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

package validation

// ValidateSelectType validates a select value
func ValidateSelectType(givenObj interface{}, possibleFnSelects map[string]interface{}, ctx *ValidateContext) []*ValidationFailure {
	if _, ok := givenObj.(map[string]interface{}); !ok {
		return GetTypeUnexpectedValidationFailure([]interface{}{}, givenObj, "Object")
	}
	
	fnScope := ""
	if ctx.Fn != nil {
		fnScope = *ctx.Fn
	}
	
	possibleSelect, exists := possibleFnSelects[fnScope]
	if !exists {
		return []*ValidationFailure{
			NewValidationFailure([]interface{}{}, "FunctionNotFound", map[string]interface{}{"fn": fnScope}),
		}
	}
	
	return isSubSelect([]interface{}{}, givenObj, possibleSelect)
}

// isSubSelect recursively validates select subsections
func isSubSelect(path []interface{}, givenObj interface{}, possibleSelectSection interface{}) []*ValidationFailure {
	// Check if possible select is a list
	if possibleList, ok := possibleSelectSection.([]interface{}); ok {
		// Given object should be a list
		givenList, ok := givenObj.([]interface{})
		if !ok {
			return GetTypeUnexpectedValidationFailure(path, givenObj, "Array")
		}
		
		var validationFailures []*ValidationFailure
		for index, element := range givenList {
			// Check if element is in possible list
			found := false
			for _, possibleElement := range possibleList {
				if element == possibleElement {
					found = true
					break
				}
			}
			if !found {
				validationFailures = append(validationFailures, NewValidationFailure(
					append(path, index),
					"ArrayElementDisallowed",
					map[string]interface{}{},
				))
			}
		}
		return validationFailures
	}
	
	// Check if possible select is a dict
	if possibleMap, ok := possibleSelectSection.(map[string]interface{}); ok {
		// Given object should be a map
		givenMap, ok := givenObj.(map[string]interface{})
		if !ok {
			return GetTypeUnexpectedValidationFailure(path, givenObj, "Object")
		}
		
		var validationFailures []*ValidationFailure
		for key, value := range givenMap {
			if possibleValue, exists := possibleMap[key]; exists {
				innerFailures := isSubSelect(append(path, key), value, possibleValue)
				validationFailures = append(validationFailures, innerFailures...)
			} else {
				validationFailures = append(validationFailures, NewValidationFailure(
					append(path, key),
					"ObjectKeyDisallowed",
					map[string]interface{}{},
				))
			}
		}
		return validationFailures
	}
	
	return nil
}
