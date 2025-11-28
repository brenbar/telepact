// Copyright The Telepact Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"regexp"

	"github.com/brenbar/telepact/lib/go/telepact"
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseTypeDeclaration parses a type declaration from pseudo-JSON.
func ParseTypeDeclaration(path []interface{}, typeDeclarationObject interface{}, ctx *ParseContext) (*types.TTypeDeclaration, error) {
	// String type declaration (e.g., "string", "integer?", "struct.MyType")
	if typeDeclarationString, ok := typeDeclarationObject.(string); ok {
		regexString := `^(.*?)(\?)?$`
		regex := regexp.MustCompile(regexString)
		
		matches := regex.FindStringSubmatch(typeDeclarationString)
		if len(matches) == 0 {
			return nil, telepact.NewTelepactSchemaParseError(
				[]*SchemaParseFailure{
					NewSchemaParseFailure(
						ctx.DocumentName,
						path,
						"StringRegexMatchFailed",
						map[string]interface{}{"regex": regexString},
					),
				},
				ctx.TelepactSchemaDocumentNamesToJSON,
			)
		}
		
		typeName := matches[1]
		nullable := matches[2] != ""
		
		tType, err := GetOrParseType(path, typeName, ctx)
		if err != nil {
			return nil, err
		}
		
		if tType.GetTypeParameterCount() != 0 {
			return nil, telepact.NewTelepactSchemaParseError(
				[]*SchemaParseFailure{
					NewSchemaParseFailure(
						ctx.DocumentName,
						path,
						"ArrayLengthUnexpected",
						map[string]interface{}{
							"actual":   1,
							"expected": tType.GetTypeParameterCount() + 1,
						},
					),
				},
				ctx.TelepactSchemaDocumentNamesToJSON,
			)
		}
		
		return types.NewTTypeDeclaration(tType, nullable, []*types.TTypeDeclaration{}), nil
	}
	
	// Array type declaration (e.g., ["string"], [{"string": "integer"}])
	if listObject, ok := typeDeclarationObject.([]interface{}); ok {
		if len(listObject) != 1 {
			return nil, telepact.NewTelepactSchemaParseError(
				[]*SchemaParseFailure{
					NewSchemaParseFailure(
						ctx.DocumentName,
						path,
						"ArrayLengthUnexpected",
						map[string]interface{}{
							"actual":   len(listObject),
							"expected": 1,
						},
					),
				},
				ctx.TelepactSchemaDocumentNamesToJSON,
			)
		}
		
		elementTypeDeclaration := listObject[0]
		newPath := append(path, 0)
		
		arrayType := &types.TArray{}
		parsedElementType, err := ParseTypeDeclaration(newPath, elementTypeDeclaration, ctx)
		if err != nil {
			return nil, err
		}
		
		return types.NewTTypeDeclaration(arrayType, false, []*types.TTypeDeclaration{parsedElementType}), nil
	}
	
	// Object/map type declaration (e.g., {"string": "integer"})
	if mapObject, ok := typeDeclarationObject.(map[string]interface{}); ok {
		if len(mapObject) != 1 {
			return nil, telepact.NewTelepactSchemaParseError(
				[]*SchemaParseFailure{
					NewSchemaParseFailure(
						ctx.DocumentName,
						path,
						"ObjectSizeUnexpected",
						map[string]interface{}{
							"actual":   len(mapObject),
							"expected": 1,
						},
					),
				},
				ctx.TelepactSchemaDocumentNamesToJSON,
			)
		}
		
		// Get the single key-value pair
		var key string
		var value interface{}
		for k, v := range mapObject {
			key = k
			value = v
			break
		}
		
		if key != "string" {
			keyPath := append(path, key)
			return nil, telepact.NewTelepactSchemaParseError(
				[]*SchemaParseFailure{
					NewSchemaParseFailure(
						ctx.DocumentName,
						path,
						"RequiredObjectKeyMissing",
						map[string]interface{}{"key": "string"},
					),
					NewSchemaParseFailure(
						ctx.DocumentName,
						keyPath,
						"ObjectKeyDisallowed",
						map[string]interface{}{},
					),
				},
				ctx.TelepactSchemaDocumentNamesToJSON,
			)
		}
		
		newPath := append(path, key)
		
		objectType := &types.TObject{}
		parsedValueType, err := ParseTypeDeclaration(newPath, value, ctx)
		if err != nil {
			return nil, err
		}
		
		return types.NewTTypeDeclaration(objectType, false, []*types.TTypeDeclaration{parsedValueType}), nil
	}
	
	// Invalid type
	failures := GetTypeUnexpectedParseFailure(
		ctx.DocumentName,
		path,
		typeDeclarationObject,
		"StringOrArrayOrObject",
	)
	return nil, telepact.NewTelepactSchemaParseError(failures, ctx.TelepactSchemaDocumentNamesToJSON)
}
