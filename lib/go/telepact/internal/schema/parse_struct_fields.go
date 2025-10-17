package schema

import (
	"fmt"
	"strings"

	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseStructFields parses struct field definitions from pseudo-JSON
func ParseStructFields(
	path []interface{},
	referenceStruct map[string]interface{},
	isHeader bool,
	ctx *ParseContext,
) (map[string]*types.TFieldDeclaration, error) {
	parseFailures := []*SchemaParseFailure{}
	fields := make(map[string]*types.TFieldDeclaration)

	for fieldDeclaration, typeDeclarationValue := range referenceStruct {
		// Check for field collision (same name with/without optional marker)
		for existingField := range fields {
			existingFieldNoOpt := strings.Split(existingField, "!")[0]
			fieldNoOpt := strings.Split(fieldDeclaration, "!")[0]
			if fieldNoOpt == existingFieldNoOpt {
				finalPath := append(path, fieldDeclaration)
				finalOtherPath := append(path, existingField)
				// Note: document JSON is stored as []interface{} in ParseContext
				finalOtherLocationPseudoJSON := map[string]interface{}{
					"path": finalOtherPath,
				}
				parseFailures = append(parseFailures, &SchemaParseFailure{
					DocumentName: ctx.DocumentName,
					Path:         finalPath,
					Reason:       "PathCollision",
					Data: map[string]interface{}{
						"document": ctx.DocumentName,
						"path":     finalOtherPath,
						"location": finalOtherLocationPseudoJSON,
					},
				})
			}
		}

		// Parse the field
		parsedField, err := ParseField(path, fieldDeclaration, typeDeclarationValue, isHeader, ctx)
		if err != nil {
			if schemaErr, ok := err.(*TelepactSchemaParseErrorType); ok {
				parseFailures = append(parseFailures, schemaErr.SchemaParseFailures...)
			} else {
				return nil, err
			}
		} else {
			fieldName := parsedField.FieldName
			fields[fieldName] = parsedField
		}
	}

	if len(parseFailures) > 0 {
		return nil, &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	return fields, nil
}

// TelepactSchemaParseErrorType represents a schema parse error
type TelepactSchemaParseErrorType struct {
	SchemaParseFailures                     []*SchemaParseFailure
	TelepactSchemaDocumentNamesToPseudoJSON map[string][]interface{}
}

func (e *TelepactSchemaParseErrorType) Error() string {
	return fmt.Sprintf("Schema parse error: %d failures", len(e.SchemaParseFailures))
}
