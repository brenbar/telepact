package schema

import (
	"sort"
)

// CatchHeaderCollisions checks for collisions between header definitions
func CatchHeaderCollisions(
	telepactSchemaNameToPseudoJSON map[string][]interface{},
	headerKeys []string,
	keysToIndex map[string]int,
	schemaKeysToDocumentNames map[string]string,
	documentNamesToJSON map[string]string,
) []*SchemaParseFailure {
	parseFailures := []*SchemaParseFailure{}

	// Sort header keys by document name and index
	headerKeysList := make([]string, len(headerKeys))
	copy(headerKeysList, headerKeys)
	sort.Slice(headerKeysList, func(i, j int) bool {
		docI := schemaKeysToDocumentNames[headerKeysList[i]]
		docJ := schemaKeysToDocumentNames[headerKeysList[j]]
		if docI != docJ {
			return docI < docJ
		}
		return keysToIndex[headerKeysList[i]] < keysToIndex[headerKeysList[j]]
	})

	// Check each pair of header definitions
	for i := 0; i < len(headerKeysList); i++ {
		for j := i + 1; j < len(headerKeysList); j++ {
			defKey := headerKeysList[i]
			otherDefKey := headerKeysList[j]

			index := keysToIndex[defKey]
			otherIndex := keysToIndex[otherDefKey]

			documentName := schemaKeysToDocumentNames[defKey]
			otherDocumentName := schemaKeysToDocumentNames[otherDefKey]

			telepactSchemaPseudoJSON := telepactSchemaNameToPseudoJSON[documentName]
			otherTelepactSchemaPseudoJSON := telepactSchemaNameToPseudoJSON[otherDocumentName]

			def := telepactSchemaPseudoJSON[index].(map[string]interface{})
			otherDef := otherTelepactSchemaPseudoJSON[otherIndex].(map[string]interface{})

			headerDef := def[defKey].(map[string]interface{})
			otherHeaderDef := otherDef[otherDefKey].(map[string]interface{})

			// Check for request header collisions
			for headerKey := range headerDef {
				if headerKey == "->" || headerKey == "///" {
					continue
				}
				if _, exists := otherHeaderDef[headerKey]; exists {
					thisPath := []interface{}{index, defKey, headerKey}
					thisDocumentJSON := documentNamesToJSON[documentName]
					thisLocation := GetPathDocumentCoordinatesPseudoJSON(thisPath, thisDocumentJSON)

					parseFailures = append(parseFailures, &SchemaParseFailure{
						DocumentName: otherDocumentName,
						Path:         []interface{}{otherIndex, otherDefKey, headerKey},
						Reason:       "PathCollision",
						Data: map[string]interface{}{
							"document": documentName,
							"path":     thisPath,
							"location": thisLocation,
						},
					})
				}
			}

			// Check for response header collisions
			resHeaderDef, ok1 := def["->"].(map[string]interface{})
			otherResHeaderDef, ok2 := otherDef["->"].(map[string]interface{})
			if ok1 && ok2 {
				for resHeaderKey := range resHeaderDef {
					if resHeaderKey == "///" {
						continue
					}
					if _, exists := otherResHeaderDef[resHeaderKey]; exists {
						thisPath := []interface{}{index, defKey, "->", resHeaderKey}
						thisDocumentJSON := documentNamesToJSON[documentName]
						thisLocation := GetPathDocumentCoordinatesPseudoJSON(thisPath, thisDocumentJSON)

						parseFailures = append(parseFailures, &SchemaParseFailure{
							DocumentName: otherDocumentName,
							Path:         []interface{}{otherIndex, otherDefKey, "->", resHeaderKey},
							Reason:       "PathCollision",
							Data: map[string]interface{}{
								"document": documentName,
								"path":     thisPath,
								"location": thisLocation,
							},
						})
					}
				}
			}
		}
	}

	return parseFailures
}
