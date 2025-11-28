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

// GetOrParseType retrieves or parses a type by name.
func GetOrParseType(path []interface{}, typeName string, ctx *ParseContext) (types.TType, error) {
	// Check if type has already failed to parse
	if ctx.FailedTypes[typeName] {
		return nil, telepact.NewTelepactSchemaParseError([]*SchemaParseFailure{}, ctx.TelepactSchemaDocumentNamesToJSON)
	}
	
	// Check if type is already parsed
	if existingType, ok := ctx.ParsedTypes[typeName]; ok {
		return existingType, nil
	}
	
	// Check for standard types
	regexString := `^(boolean|integer|number|string|any|bytes)|((fn|(union|struct|_ext))\.([a-zA-Z_]\w*))$`
	regex := regexp.MustCompile(regexString)
	
	matches := regex.FindStringSubmatch(typeName)
	if len(matches) == 0 {
		return nil, telepact.NewTelepactSchemaParseError(
			[]*SchemaParseFailure{
				NewSchemaParseFailure(
					ctx.DocumentName,
					path,
					"StringRegexMatchFailed",
					map[string]interface{}{
						"regex": regexString,
					},
				),
			},
			ctx.TelepactSchemaDocumentNamesToJSON,
		)
	}
	
	// Standard type
	standardTypeName := matches[1]
	if standardTypeName != "" {
		switch standardTypeName {
		case "boolean":
			return &types.TBoolean{}, nil
		case "integer":
			return &types.TInteger{}, nil
		case "number":
			return &types.TNumber{}, nil
		case "string":
			return &types.TString{}, nil
		case "bytes":
			return &types.TBytes{}, nil
		case "any":
			return &types.TAny{}, nil
		default:
			return &types.TAny{}, nil
		}
	}
	
	// Custom type (struct, union, fn, _ext)
	customTypeName := matches[2]
	thisIndex, hasIndex := ctx.SchemaKeysToIndex[customTypeName]
	thisDocumentName, hasDoc := ctx.SchemaKeysToDocumentName[customTypeName]
	
	if !hasIndex || !hasDoc {
		return nil, telepact.NewTelepactSchemaParseError(
			[]*SchemaParseFailure{
				NewSchemaParseFailure(
					ctx.DocumentName,
					path,
					"TypeUnknown",
					map[string]interface{}{
						"name": customTypeName,
					},
				),
			},
			ctx.TelepactSchemaDocumentNamesToJSON,
		)
	}
	
	telepactSchemaPseudoJSON := ctx.TelepactSchemaDocumentNamesToPseudoJSON[thisDocumentName]
	definition, ok := telepactSchemaPseudoJSON[thisIndex].(map[string]interface{})
	if !ok {
		return nil, telepact.NewTelepactSchemaParseError(
			[]*SchemaParseFailure{
				NewSchemaParseFailure(
					ctx.DocumentName,
					[]interface{}{thisIndex},
					"TypeUnexpected",
					map[string]interface{}{
						"expected": "Object",
						"actual":   "other",
					},
				),
			},
			ctx.TelepactSchemaDocumentNamesToJSON,
		)
	}
	
	// TODO: Implement actual parsing for struct, union, fn, _ext types
	// For now, return stub implementation
	
	var parsedType types.TType
	
	// Placeholder for actual implementation
	// This would call ParseStructType, ParseUnionType, etc.
	_ = definition
	_ = customTypeName
	_ = thisDocumentName
	
	// Mark as failed for now since full implementation not done
	ctx.FailedTypes[customTypeName] = true
	return nil, telepact.NewTelepactSchemaParseError([]*SchemaParseFailure{}, ctx.TelepactSchemaDocumentNamesToJSON)
	
	// Once implemented, this should:
	// 1. Parse struct types
	// 2. Parse union types
	// 3. Parse function types
	// 4. Parse extension types (_ext.Select_, etc.)
	// And store in ctx.ParsedTypes[customTypeName]
	
	return parsedType, nil
}
