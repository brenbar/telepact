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

const unionName = "Object"

// TUnion represents a union type
type TUnion struct {
	Name        string
	Tags        map[string]*TStruct
	TagIndices  map[string]int
}

// NewTUnion creates a new TUnion
func NewTUnion(name string, tags map[string]*TStruct, tagIndices map[string]int) *TUnion {
	return &TUnion{
		Name:       name,
		Tags:       tags,
		TagIndices: tagIndices,
	}
}

// GetTypeParameterCount returns the number of type parameters
func (t *TUnion) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a union
func (t *TUnion) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	return validation.ValidateUnionType(value, t.Name, func(obj map[string]interface{}) []*validation.ValidationFailure {
		var selectedTags map[string]interface{}
		if ctx.Select != nil {
			// Check if it's a function union (starts with "fn.")
			if len(t.Name) > 3 && t.Name[:3] == "fn." {
				if selValue := ctx.Select[t.Name]; selValue != nil {
					selectedTags = map[string]interface{}{t.Name: selValue}
				}
			} else {
				if selValue, ok := ctx.Select[t.Name].(map[string]interface{}); ok {
					selectedTags = selValue
				}
			}
		}
		
		// Build sorted tag list
		var tags []string
		for tag := range t.Tags {
			tags = append(tags, tag)
		}
		sort.Strings(tags)
		
		return validation.ValidateUnionTags(tags, selectedTags, obj, func(tag string, payload map[string]interface{}, selTags map[string]interface{}) []*validation.ValidationFailure {
			unionStruct := t.Tags[tag]
			
			var selectedFields []string
			if selTags != nil {
				if fields, ok := selTags[tag].([]string); ok {
					selectedFields = fields
				} else if fieldsList, ok := selTags[tag].([]interface{}); ok {
					// Convert []interface{} to []string
					for _, f := range fieldsList {
						if str, ok := f.(string); ok {
							selectedFields = append(selectedFields, str)
						}
					}
				}
			}
			
			// Add union tag to path
			ctx.Path = append(ctx.Path, tag)
			
			// Build field infos
			var fieldInfos []validation.FieldInfo
			for fieldName, fieldDecl := range unionStruct.Fields {
				fieldInfos = append(fieldInfos, validation.FieldInfo{
					Name:     fieldName,
					Optional: fieldDecl.Optional,
				})
			}
			
			result := validation.ValidateStructFields(fieldInfos, selectedFields, payload, func(fieldName string, fieldValue interface{}) []*validation.ValidationFailure {
				fieldDecl := unionStruct.Fields[fieldName]
				return fieldDecl.TypeDeclaration.Validate(fieldValue, ctx)
			}, ctx)
			
			// Remove union tag from path
			ctx.Path = ctx.Path[:len(ctx.Path)-1]
			
			return result
		}, ctx)
	}, ctx)
}

// GenerateRandomValue generates a random union value
func (t *TUnion) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	// Build sorted tag list
	var tags []string
	for tag := range t.Tags {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	
	return generation.GenerateRandomUnionType(blueprintValue, useBlueprintValue, tags, func(tag string, bpValue interface{}, useBp bool) interface{} {
		unionStruct := t.Tags[tag]
		
		// Build field infos in sorted order
		var fieldNames []string
		for fieldName := range unionStruct.Fields {
			fieldNames = append(fieldNames, fieldName)
		}
		sort.Strings(fieldNames)
		
		var fieldInfos []generation.FieldInfo
		for _, fieldName := range fieldNames {
			fieldDecl := unionStruct.Fields[fieldName]
			fieldInfos = append(fieldInfos, generation.FieldInfo{
				Name:     fieldName,
				Optional: fieldDecl.Optional,
			})
		}
		
		return generation.GenerateRandomStructType(bpValue, useBp, fieldInfos, func(fieldName string, bpVal interface{}, useBpVal bool) interface{} {
			fieldDecl := unionStruct.Fields[fieldName]
			return fieldDecl.TypeDeclaration.GenerateRandomValue(bpVal, useBpVal, ctx)
		}, ctx)
	}, ctx)
}

// GetName returns the type name
func (t *TUnion) GetName(ctx *validation.ValidateContext) string {
	return unionName
}
