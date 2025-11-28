package schema

import (
	"strings"

	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseFunctionResultType parses a function result type (union with Ok_ tag required)
func ParseFunctionResultType(
	path []interface{},
	functionDefinitionAsParsedJSON map[string]interface{},
	schemaKey string,
	ctx *ParseContext,
) (*types.TUnion, error) {
	parseFailures := []*SchemaParseFailure{}

	resultSchemaKey := "->"

	var resultType *types.TUnion
	if _, exists := functionDefinitionAsParsedJSON[resultSchemaKey]; !exists {
		parseFailures = append(parseFailures, &SchemaParseFailure{
			DocumentName: ctx.DocumentName,
			Path:         path,
			Reason:       "RequiredObjectKeyMissing",
			Data:         map[string]interface{}{"key": resultSchemaKey},
		})
	} else {
		// Get all keys for ignore list
		ignoreKeys := make([]string, 0, len(functionDefinitionAsParsedJSON))
		for k := range functionDefinitionAsParsedJSON {
			ignoreKeys = append(ignoreKeys, k)
		}

		result, err := ParseUnionType(path, functionDefinitionAsParsedJSON,
			resultSchemaKey, ignoreKeys, []string{"Ok_"}, ctx)
		if err != nil {
			if schemaErr, ok := err.(*TelepactSchemaParseErrorType); ok {
				parseFailures = append(parseFailures, schemaErr.SchemaParseFailures...)
			} else {
				return nil, err
			}
		} else {
			resultType = result
		}
	}

	if len(parseFailures) > 0 {
		return nil, &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	// Derive possible select and add to TSelect
	fnSelectType := DerivePossibleSelect(schemaKey, resultType)
	selectType, err := GetOrParseType([]interface{}{}, "_ext.Select_", ctx)
	if err != nil {
		return nil, err
	}
	if sel, ok := selectType.(*types.TSelect); ok {
		sel.PossibleSelects[schemaKey] = fnSelectType
	}

	return resultType, nil
}

// ParseFunctionErrorsRegex parses the _errors regex from a function definition
func ParseFunctionErrorsRegex(
	path []interface{},
	functionDefinitionAsParsedJSON map[string]interface{},
	schemaKey string,
	ctx *ParseContext,
) (string, error) {
	parseFailures := []*SchemaParseFailure{}

	errorsRegexKey := "_errors"
	regexPath := append(path, errorsRegexKey)

	var errorsRegex string
	if _, exists := functionDefinitionAsParsedJSON[errorsRegexKey]; exists && !strings.HasSuffix(schemaKey, "_") {
		parseFailures = append(parseFailures, &SchemaParseFailure{
			DocumentName: ctx.DocumentName,
			Path:         regexPath,
			Reason:       "ObjectKeyDisallowed",
			Data:         map[string]interface{}{},
		})
	} else {
		errorsRegexInit := functionDefinitionAsParsedJSON[errorsRegexKey]
		if errorsRegexInit == nil {
			errorsRegex = "^errors\\..*$"
		} else if errorsRegexStr, ok := errorsRegexInit.(string); !ok {
			thisParseFailures := GetTypeUnexpectedParseFailure(
				ctx.DocumentName, regexPath, errorsRegexInit, "String")
			parseFailures = append(parseFailures, thisParseFailures...)
		} else {
			errorsRegex = errorsRegexStr
		}
	}

	if len(parseFailures) > 0 {
		return "", &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	return errorsRegex, nil
}
