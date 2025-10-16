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

package internal

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// SelectStructFields recursively filters struct fields based on selected fields
func SelectStructFields(typeDeclaration *types.TTypeDeclaration, value interface{}, selectedStructFields map[string]interface{}) interface{} {
	typeDeclarationType := typeDeclaration.Type
	typeParams := typeDeclaration.TypeParameters
	
	switch t := typeDeclarationType.(type) {
	case *types.TStruct:
		fields := t.Fields
		structName := t.Name
		
		// Get selected fields for this struct
		var selectedFields []string
		if sel, ok := selectedStructFields[structName]; ok {
			if selList, ok := sel.([]string); ok {
				selectedFields = selList
			} else if selInterfaceList, ok := sel.([]interface{}); ok {
				for _, item := range selInterfaceList {
					if str, ok := item.(string); ok {
						selectedFields = append(selectedFields, str)
					}
				}
			}
		}
		
		valueAsMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}
		
		finalMap := make(map[string]interface{})
		for fieldName, fieldValue := range valueAsMap {
			// Check if field is selected
			if selectedFields == nil || contains(selectedFields, fieldName) {
				if field, exists := fields[fieldName]; exists {
					fieldTypeDeclaration := field.TypeDeclaration
					valueWithSelectedFields := SelectStructFields(fieldTypeDeclaration, fieldValue, selectedStructFields)
					finalMap[fieldName] = valueWithSelectedFields
				}
			}
		}
		
		return finalMap
		
	case *types.TUnion:
		valueAsMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}
		
		// Get the single tag and data
		var unionTag string
		var unionData interface{}
		for k, v := range valueAsMap {
			unionTag = k
			unionData = v
			break
		}
		
		unionDataMap, ok := unionData.(map[string]interface{})
		if !ok {
			return value
		}
		
		unionTags := t.Tags
		unionStructReference, exists := unionTags[unionTag]
		if !exists {
			return value
		}
		
		unionStructRefFields := unionStructReference.Fields
		
		// Build default tags to fields
		defaultTagsToFields := make(map[string][]string)
		for tag, unionStruct := range unionTags {
			var fieldNames []string
			for fieldName := range unionStruct.Fields {
				fieldNames = append(fieldNames, fieldName)
			}
			defaultTagsToFields[tag] = fieldNames
		}
		
		// Get union selected fields
		unionSelectedFields := defaultTagsToFields
		if sel, ok := selectedStructFields[t.Name]; ok {
			if selMap, ok := sel.(map[string]interface{}); ok {
				unionSelectedFields = make(map[string][]string)
				for k, v := range selMap {
					if vList, ok := v.([]string); ok {
						unionSelectedFields[k] = vList
					} else if vInterfaceList, ok := v.([]interface{}); ok {
						var strList []string
						for _, item := range vInterfaceList {
							if str, ok := item.(string); ok {
								strList = append(strList, str)
							}
						}
						unionSelectedFields[k] = strList
					}
				}
			}
		}
		
		thisUnionTagSelectedFieldsDefault := defaultTagsToFields[unionTag]
		selectedFields := unionSelectedFields[unionTag]
		if selectedFields == nil {
			selectedFields = thisUnionTagSelectedFieldsDefault
		}
		
		finalMap := make(map[string]interface{})
		for fieldName, fieldValue := range unionDataMap {
			if selectedFields == nil || contains(selectedFields, fieldName) {
				if field, exists := unionStructRefFields[fieldName]; exists {
					valueWithSelectedFields := SelectStructFields(field.TypeDeclaration, fieldValue, selectedStructFields)
					finalMap[fieldName] = valueWithSelectedFields
				}
			}
		}
		
		return map[string]interface{}{unionTag: finalMap}
		
	case *types.TObject:
		if len(typeParams) == 0 {
			return value
		}
		
		nestedTypeDeclaration := typeParams[0]
		valueAsMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}
		
		finalMap := make(map[string]interface{})
		for key, val := range valueAsMap {
			valueWithSelectedFields := SelectStructFields(nestedTypeDeclaration, val, selectedStructFields)
			finalMap[key] = valueWithSelectedFields
		}
		
		return finalMap
		
	case *types.TArray:
		if len(typeParams) == 0 {
			return value
		}
		
		nestedType := typeParams[0]
		valueAsList, ok := value.([]interface{})
		if !ok {
			return value
		}
		
		var finalList []interface{}
		for _, entry := range valueAsList {
			valueWithSelectedFields := SelectStructFields(nestedType, entry, selectedStructFields)
			finalList = append(finalList, valueWithSelectedFields)
		}
		
		return finalList
		
	default:
		return value
	}
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
