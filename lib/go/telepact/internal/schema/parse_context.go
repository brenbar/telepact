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
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseContext holds the state during schema parsing.
type ParseContext struct {
	DocumentName                                 string
	TelepactSchemaDocumentNamesToPseudoJSON      map[string][]interface{}
	TelepactSchemaDocumentNamesToJSON            map[string]string
	SchemaKeysToDocumentName                     map[string]string
	SchemaKeysToIndex                            map[string]int
	ParsedTypes                                  map[string]types.TType
	FnErrorRegexes                               map[string]string
	AllParseFailures                             []*SchemaParseFailure
	FailedTypes                                  map[string]bool
}

// NewParseContext creates a new parse context.
func NewParseContext(
	documentName string,
	telepactSchemaDocumentNamesToPseudoJSON map[string][]interface{},
	telepactSchemaDocumentNamesToJSON map[string]string,
	schemaKeysToDocumentName map[string]string,
	schemaKeysToIndex map[string]int,
	parsedTypes map[string]types.TType,
	fnErrorRegexes map[string]string,
	allParseFailures []*SchemaParseFailure,
	failedTypes map[string]bool,
) *ParseContext {
	return &ParseContext{
		DocumentName:                            documentName,
		TelepactSchemaDocumentNamesToPseudoJSON: telepactSchemaDocumentNamesToPseudoJSON,
		TelepactSchemaDocumentNamesToJSON:       telepactSchemaDocumentNamesToJSON,
		SchemaKeysToDocumentName:                schemaKeysToDocumentName,
		SchemaKeysToIndex:                       schemaKeysToIndex,
		ParsedTypes:                             parsedTypes,
		FnErrorRegexes:                          fnErrorRegexes,
		AllParseFailures:                        allParseFailures,
		FailedTypes:                             failedTypes,
	}
}

// Copy creates a copy of the parse context, optionally with a different document name.
func (ctx *ParseContext) Copy(documentName *string) *ParseContext {
	docName := ctx.DocumentName
	if documentName != nil {
		docName = *documentName
	}
	
	return &ParseContext{
		DocumentName:                            docName,
		TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		TelepactSchemaDocumentNamesToJSON:       ctx.TelepactSchemaDocumentNamesToJSON,
		SchemaKeysToDocumentName:                ctx.SchemaKeysToDocumentName,
		SchemaKeysToIndex:                       ctx.SchemaKeysToIndex,
		ParsedTypes:                             ctx.ParsedTypes,
		FnErrorRegexes:                          ctx.FnErrorRegexes,
		AllParseFailures:                        ctx.AllParseFailures,
		FailedTypes:                             ctx.FailedTypes,
	}
}
