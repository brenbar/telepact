package schema

import (
	"regexp"
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
)

// ApplyErrorToParsedTypes applies error tags to function result types
func ApplyErrorToParsedTypes(
	errorType *types.TError,
	parsedTypes map[string]types.TType,
	schemaKeysToDocumentNames map[string]string,
	schemaKeysToIndex map[string]int,
	documentNamesToJSON map[string]string,
	fnErrorRegexes map[string]string,
) []*SchemaParseFailure {
	parseFailures := []*SchemaParseFailure{}

	errorKey := errorType.Name
	errorIndex := schemaKeysToIndex[errorKey]
	documentName := schemaKeysToDocumentNames[errorKey]

	// Iterate through parsed types looking for functions
	for parsedTypeName := range parsedTypes {
		// Skip non-function types and result types
		if len(parsedTypeName) < 3 || parsedTypeName[:3] != "fn." {
			continue
		}
		if len(parsedTypeName) >= 3 && parsedTypeName[len(parsedTypeName)-3:] == ".->" {
			continue
		}

		// Get the function result type
		resultTypeName := parsedTypeName + ".->"
		parsedType, ok := parsedTypes[resultTypeName]
		if !ok {
			continue
		}

		fnResult, ok := parsedType.(*types.TUnion)
		if !ok {
			continue
		}

		fnName := parsedTypeName
		fnErrorRegex, ok := fnErrorRegexes[fnName]
		if !ok {
			continue
		}

		// Check if error matches the regex
		regex, err := regexp.Compile(fnErrorRegex)
		if err != nil {
			continue
		}

		if !regex.MatchString(errorKey) {
			continue
		}

		// Apply error tags to function result
		fnResultTags := fnResult.Tags
		errorTags := errorType.Errors.Tags
		errorTagIndices := errorType.Errors.TagIndices

		for errorTagName, errorTag := range errorTags {
			newKey := errorTagName

			// Check for collision
			if _, exists := fnResultTags[newKey]; exists {
				otherPathIndex := schemaKeysToIndex[fnName]
				errorTagIndex := errorTagIndices[newKey]
				otherDocumentName := schemaKeysToDocumentNames[fnName]
				fnErrorTagIndex := fnResult.TagIndices[newKey]

				otherFinalPath := []interface{}{otherPathIndex, "->", fnErrorTagIndex, newKey}
				otherDocumentJSON := documentNamesToJSON[otherDocumentName]
				otherLocationPseudoJSON := GetPathDocumentCoordinatesPseudoJSON(otherFinalPath, otherDocumentJSON)

				parseFailures = append(parseFailures, &SchemaParseFailure{
					DocumentName: documentName,
					Path:         []interface{}{errorIndex, errorKey, errorTagIndex, newKey},
					Reason:       "PathCollision",
					Data: map[string]interface{}{
						"document": otherDocumentName,
						"path":     otherFinalPath,
						"location": otherLocationPseudoJSON,
					},
				})
			}

			// Apply the error tag to the function result
			fnResultTags[newKey] = errorTag
		}
	}

	return parseFailures
}
