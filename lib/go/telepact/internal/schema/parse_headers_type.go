package schema

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ParseHeadersType parses a headers type definition from pseudo-JSON
func ParseHeadersType(
	path []interface{},
	headersDefinitionAsParsedJSON map[string]interface{},
	schemaKey string,
	ctx *ParseContext,
) (*types.THeaders, error) {
	parseFailures := []*SchemaParseFailure{}
	requestHeaders := make(map[string]*types.TFieldDeclaration)
	responseHeaders := make(map[string]*types.TFieldDeclaration)

	requestHeadersDef := headersDefinitionAsParsedJSON[schemaKey]
	thisPath := append(path, schemaKey)

	requestHeadersDefMap, ok := requestHeadersDef.(map[string]interface{})
	if !ok {
		branchParseFailures := GetTypeUnexpectedParseFailure(
			ctx.DocumentName,
			thisPath,
			requestHeadersDef,
			"Object",
		)
		parseFailures = append(parseFailures, branchParseFailures...)
	} else {
		requestFields, err := ParseStructFields(
			thisPath, requestHeadersDefMap, true, ctx)

		if err != nil {
			if schemaErr, ok := err.(*TelepactSchemaParseErrorType); ok {
				parseFailures = append(parseFailures, schemaErr.SchemaParseFailures...)
			} else {
				return nil, err
			}
		} else {
			// All headers are optional
			for field := range requestFields {
				requestFields[field].Optional = true
			}
			requestHeaders = requestFields
		}
	}

	responseKey := "->"
	responsePath := append(path, responseKey)

	if _, exists := headersDefinitionAsParsedJSON[responseKey]; !exists {
		parseFailures = append(parseFailures, &SchemaParseFailure{
			DocumentName: ctx.DocumentName,
			Path:         responsePath,
			Reason:       "RequiredObjectKeyMissing",
			Data: map[string]interface{}{
				"key": responseKey,
			},
		})
	}

	responseHeadersDef := headersDefinitionAsParsedJSON[responseKey]

	responseHeadersDefMap, ok := responseHeadersDef.(map[string]interface{})
	if !ok {
		branchParseFailures := GetTypeUnexpectedParseFailure(
			ctx.DocumentName,
			responsePath,
			responseHeadersDef,
			"Object",
		)
		parseFailures = append(parseFailures, branchParseFailures...)
	} else {
		responseFields, err := ParseStructFields(
			responsePath, responseHeadersDefMap, true, ctx)

		if err != nil {
			if schemaErr, ok := err.(*TelepactSchemaParseErrorType); ok {
				parseFailures = append(parseFailures, schemaErr.SchemaParseFailures...)
			} else {
				return nil, err
			}
		} else {
			// All headers are optional
			for field := range responseFields {
				responseFields[field].Optional = true
			}
			responseHeaders = responseFields
		}
	}

	if len(parseFailures) > 0 {
		return nil, &TelepactSchemaParseErrorType{
			SchemaParseFailures:                     parseFailures,
			TelepactSchemaDocumentNamesToPseudoJSON: ctx.TelepactSchemaDocumentNamesToPseudoJSON,
		}
	}

	return types.NewTHeaders(schemaKey, requestHeaders, responseHeaders), nil
}
