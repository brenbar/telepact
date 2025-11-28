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

// ParseField parses a field declaration from a schema definition.
func ParseField(
	path []interface{},
	fieldDeclaration string,
	typeDeclarationValue interface{},
	isHeader bool,
	ctx *ParseContext,
) (*types.TFieldDeclaration, error) {
	headerRegexString := `^@[a-z][a-zA-Z0-9_]*$`
	regexString := `^([a-z][a-zA-Z0-9_]*)(!)?$`
	regexToUse := regexString
	if isHeader {
		regexToUse = headerRegexString
	}
	
	regex := regexp.MustCompile(regexToUse)
	
	matches := regex.FindStringSubmatch(fieldDeclaration)
	if len(matches) == 0 {
		finalPath := append(path, fieldDeclaration)
		return nil, telepact.NewTelepactSchemaParseError(
			[]*SchemaParseFailure{
				NewSchemaParseFailure(
					ctx.DocumentName,
					finalPath,
					"KeyRegexMatchFailed",
					map[string]interface{}{"regex": regexToUse},
				),
			},
			ctx.TelepactSchemaDocumentNamesToJSON,
		)
	}
	
	fieldName := matches[0]
	optional := true
	if !isHeader && len(matches) > 2 {
		optional = matches[2] != ""
	}
	
	thisPath := append(path, fieldName)
	
	typeDeclaration, err := ParseTypeDeclaration(thisPath, typeDeclarationValue, ctx)
	if err != nil {
		return nil, err
	}
	
	return types.NewTFieldDeclaration(fieldName, typeDeclaration, optional), nil
}
