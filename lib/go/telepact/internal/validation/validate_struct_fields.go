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

// FieldInfo contains information about a field for validation
type FieldInfo struct {
	Name     string
	Optional bool
}

// ValidateStructFields validates struct fields against field declarations
func ValidateStructFields(fieldInfos []FieldInfo, selectedFields []string, actualStruct map[string]interface{}, validateField func(fieldName string, fieldValue interface{}) []*ValidationFailure, ctx *ValidateContext) []*ValidationFailure {
	var validationFailures []*ValidationFailure
	
	// Build a map of field info for easier lookup
	fieldMap := make(map[string]FieldInfo)
	for _, fi := range fieldInfos {
		fieldMap[fi.Name] = fi
	}
	
	// Check for missing required fields
	var missingFields []string
	for _, fieldInfo := range fieldInfos {
		isOptional := fieldInfo.Optional
		isOmittedBySelect := selectedFields != nil && !contains(selectedFields, fieldInfo.Name)
		
		if _, exists := actualStruct[fieldInfo.Name]; !exists && !isOptional && !isOmittedBySelect {
			missingFields = append(missingFields, fieldInfo.Name)
		}
	}
	
	for _, missingField := range missingFields {
		validationFailures = append(validationFailures, NewValidationFailure(
			[]interface{}{},
			"RequiredObjectKeyMissing",
			map[string]interface{}{"key": missingField},
		))
	}
	
	// Validate each field in the actual struct
	for fieldName, fieldValue := range actualStruct {
		_, exists := fieldMap[fieldName]
		if !exists {
			validationFailures = append(validationFailures, NewValidationFailure(
				[]interface{}{fieldName},
				"ObjectKeyDisallowed",
				map[string]interface{}{},
			))
			continue
		}
		
		// Add field name to path
		ctx.Path = append(ctx.Path, fieldName)
		
		// Validate the field value
		nestedValidationFailures := validateField(fieldName, fieldValue)
		
		// Remove field name from path
		ctx.Path = ctx.Path[:len(ctx.Path)-1]
		
		// Add field name to failure paths
		for _, failure := range nestedValidationFailures {
			thisPath := append([]interface{}{fieldName}, failure.Path...)
			validationFailures = append(validationFailures, NewValidationFailure(thisPath, failure.Reason, failure.Data))
		}
	}
	
	return validationFailures
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
