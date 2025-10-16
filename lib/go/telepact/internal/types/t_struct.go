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

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/generation"
	"github.com/brenbar/telepact/lib/go/telepact/internal/validation"
)

// TStruct represents a struct type
type TStruct struct {
	Name   string
	Fields map[string]*TFieldDeclaration
}

func (t *TStruct) GetTypeParameterCount() int {
	return 0
}

func (t *TStruct) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	obj, ok := value.(map[string]interface{})
	if !ok {
		return []*validation.ValidationFailure{{Path: "", Message: "expected object"}}
	}
	
	var failures []*validation.ValidationFailure
	
	// Validate all fields
	for fieldName, fieldDecl := range t.Fields {
		fieldValue, exists := obj[fieldName]
		
		if !exists {
			if !fieldDecl.Optional {
				failures = append(failures, &validation.ValidationFailure{
					Path:    fieldName,
					Message: "required field missing",
				})
			}
			continue
		}
		
		// Validate field value against field type
		fieldFailures := fieldDecl.TypeDeclaration.Validate(fieldValue, ctx)
		for _, failure := range fieldFailures {
			failures = append(failures, &validation.ValidationFailure{
				Path:    fieldName + "." + failure.Path,
				Message: failure.Message,
			})
		}
	}
	
	return failures
}

func (t *TStruct) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	result := make(map[string]interface{})
	
	// Generate random values for each field
	for fieldName, fieldDecl := range t.Fields {
		if !fieldDecl.Optional || ctx.IncludeOptionalFields {
			// Generate value for field
			var blueprintFieldValue interface{}
			if useBlueprintValue {
				if blueprintMap, ok := blueprintValue.(map[string]interface{}); ok {
					blueprintFieldValue = blueprintMap[fieldName]
				}
			}
			result[fieldName] = fieldDecl.TypeDeclaration.GenerateRandomValue(blueprintFieldValue, useBlueprintValue, ctx)
		}
	}
	
	return result
}

func (t *TStruct) GetName(ctx *validation.ValidateContext) string {
	return "Object"
}
