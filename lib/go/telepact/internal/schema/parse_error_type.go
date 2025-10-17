package schema

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseErrorType parses an error type definition from pseudo-JSON
func ParseErrorType(
	path []interface{},
	errorDefinitionAsParsedJSON map[string]interface{},
	schemaKey string,
	ctx *ParseContext,
) (*types.TError, error) {
	parseFailures := []*SchemaParseFailure{}

	otherKeys := make(map[string]bool)
	for k := range errorDefinitionAsParsedJSON {
		otherKeys[k] = true
	}

	delete(otherKeys, schemaKey)
	delete(otherKeys, "///")

	if len(otherKeys) > 0 {
		for k := range otherKeys {
			loopPath := append(path, k)
			parseFailures = append(parseFailures, &SchemaParseFailure{
				DocumentName: ctx.DocumentName,
				Path:         loopPath,
				Reason:       "ObjectKeyDisallowed",
				Data:         map[string]interface{}{},
			})
		}
	}

	if len(parseFailures) > 0 {
		return nil, &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	errorUnion, err := ParseUnionType(path, errorDefinitionAsParsedJSON, schemaKey, []string{}, []string{}, ctx)
	if err != nil {
		return nil, err
	}

	return types.NewTError(schemaKey, errorUnion), nil
}
