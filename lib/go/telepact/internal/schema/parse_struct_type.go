package schema

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseStructType parses a struct type definition from pseudo-JSON
func ParseStructType(
	path []interface{},
	structDefinitionAsPseudoJSON map[string]interface{},
	schemaKey string,
	ignoreKeys []string,
	ctx *ParseContext,
) (*types.TStruct, error) {
	parseFailures := []*SchemaParseFailure{}
	otherKeys := make(map[string]bool)

	// Collect all keys
	for k := range structDefinitionAsPseudoJSON {
		otherKeys[k] = true
	}

	// Remove allowed keys
	delete(otherKeys, schemaKey)
	delete(otherKeys, "///")
	delete(otherKeys, "_ignoreIfDuplicate")
	for _, ignoreKey := range ignoreKeys {
		delete(otherKeys, ignoreKey)
	}

	// Report disallowed keys
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

	thisPath := append(path, schemaKey)
	defInit, exists := structDefinitionAsPseudoJSON[schemaKey]
	if !exists {
		parseFailures = append(parseFailures, &SchemaParseFailure{
			DocumentName: ctx.DocumentName,
			Path:         thisPath,
			Reason:       "ObjectKeyMissing",
			Data:         map[string]interface{}{"key": schemaKey},
		})
	}

	var definition map[string]interface{}
	if defInitMap, ok := defInit.(map[string]interface{}); !ok {
		branchParseFailures := GetTypeUnexpectedParseFailure(
			ctx.DocumentName, thisPath, defInit, "Object")
		parseFailures = append(parseFailures, branchParseFailures...)
	} else {
		definition = defInitMap
	}

	if len(parseFailures) > 0 {
		return nil, &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	fields, err := ParseStructFields(thisPath, definition, false, ctx)
	if err != nil {
		return nil, err
	}

	return types.NewTStruct(schemaKey, fields), nil
}
