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
	"sort"
	"github.com/brenbar/telepact/lib/go/telepact/internal/validation"
	"github.com/brenbar/telepact/lib/go/telepact/internal/generation"
)

const structName = "Object"

// TStruct represents a struct type
type TStruct struct {
	Name   string
	Fields map[string]*TFieldDeclaration
}

// NewTStruct creates a new TStruct
func NewTStruct(name string, fields map[string]*TFieldDeclaration) *TStruct {
	return &TStruct{
		Name:   name,
		Fields: fields,
	}
}

// GetTypeParameterCount returns the number of type parameters
func (t *TStruct) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a struct
func (t *TStruct) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	return validation.ValidateStructType(value, t.Name, func(obj map[string]interface{}) []*validation.ValidationFailure {
		var selectedFields []string
		if ctx.Select != nil {
			if fields, ok := ctx.Select[t.Name].([]string); ok {
				selectedFields = fields
			} else if fieldsList, ok := ctx.Select[t.Name].([]interface{}); ok {
				// Convert []interface{} to []string
				for _, f := range fieldsList {
					if str, ok := f.(string); ok {
						selectedFields = append(selectedFields, str)
					}
				}
			}
		}
		
		// Build field infos
		var fieldInfos []validation.FieldInfo
		for fieldName, fieldDecl := range t.Fields {
			fieldInfos = append(fieldInfos, validation.FieldInfo{
				Name:     fieldName,
				Optional: fieldDecl.Optional,
			})
		}
		
		return validation.ValidateStructFields(fieldInfos, selectedFields, obj, func(fieldName string, fieldValue interface{}) []*validation.ValidationFailure {
			fieldDecl := t.Fields[fieldName]
			return fieldDecl.TypeDeclaration.Validate(fieldValue, ctx)
		}, ctx)
	}, ctx)
}

// GenerateRandomValue generates a random struct value
func (t *TStruct) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	// Build field infos in sorted order
	var fieldNames []string
	for fieldName := range t.Fields {
		fieldNames = append(fieldNames, fieldName)
	}
	sort.Strings(fieldNames)
	
	var fieldInfos []generation.FieldInfo
	for _, fieldName := range fieldNames {
		fieldDecl := t.Fields[fieldName]
		fieldInfos = append(fieldInfos, generation.FieldInfo{
			Name:     fieldName,
			Optional: fieldDecl.Optional,
		})
	}
	
	return generation.GenerateRandomStructType(blueprintValue, useBlueprintValue, fieldInfos, func(fieldName string, bpValue interface{}, useBp bool) interface{} {
		fieldDecl := t.Fields[fieldName]
		return fieldDecl.TypeDeclaration.GenerateRandomValue(bpValue, useBp, ctx)
	}, ctx)
}

// GetName returns the type name
func (t *TStruct) GetName(ctx *validation.ValidateContext) string {
	return structName
}
