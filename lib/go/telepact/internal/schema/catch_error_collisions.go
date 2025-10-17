package schema

import (
	"sort"
)

// CatchErrorCollisions checks for collisions between error definitions
func CatchErrorCollisions(
	telepactSchemaNameToPseudoJSON map[string][]interface{},
	errorKeys []string,
	keysToIndex map[string]int,
	schemaKeysToDocumentName map[string]string,
	documentNamesToJSON map[string]string,
) []*SchemaParseFailure {
	parseFailures := []*SchemaParseFailure{}

	// Sort error keys by document name and index
	errorKeysList := make([]string, len(errorKeys))
	copy(errorKeysList, errorKeys)
	sort.Slice(errorKeysList, func(i, j int) bool {
		docI := schemaKeysToDocumentName[errorKeysList[i]]
		docJ := schemaKeysToDocumentName[errorKeysList[j]]
		if docI != docJ {
			return docI < docJ
		}
		return keysToIndex[errorKeysList[i]] < keysToIndex[errorKeysList[j]]
	})

	// Check each pair of error definitions
	for i := 0; i < len(errorKeysList); i++ {
		for j := i + 1; j < len(errorKeysList); j++ {
			defKey := errorKeysList[i]
			otherDefKey := errorKeysList[j]

			index := keysToIndex[defKey]
			otherIndex := keysToIndex[otherDefKey]

			documentName := schemaKeysToDocumentName[defKey]
			otherDocumentName := schemaKeysToDocumentName[otherDefKey]

			telepactSchemaPseudoJSON := telepactSchemaNameToPseudoJSON[documentName]
			otherTelepactSchemaPseudoJSON := telepactSchemaNameToPseudoJSON[otherDocumentName]

			def := telepactSchemaPseudoJSON[index].(map[string]interface{})
			otherDef := otherTelepactSchemaPseudoJSON[otherIndex].(map[string]interface{})

			errDef := def[defKey].([]interface{})
			otherErrDef := otherDef[otherDefKey].([]interface{})

			// Check each error tag in both definitions
			for k := 0; k < len(errDef); k++ {
				thisErrDef := errDef[k].(map[string]interface{})
				thisErrDefKeys := make(map[string]bool)
				for key := range thisErrDef {
					if key != "///" {
						thisErrDefKeys[key] = true
					}
				}

				for l := 0; l < len(otherErrDef); l++ {
					thisOtherErrDef := otherErrDef[l].(map[string]interface{})
					thisOtherErrDefKeys := make(map[string]bool)
					for key := range thisOtherErrDef {
						if key != "///" {
							thisOtherErrDefKeys[key] = true
						}
					}

					// Check if keys are identical (collision)
					if len(thisErrDefKeys) == len(thisOtherErrDefKeys) {
						match := true
						for key := range thisErrDefKeys {
							if !thisOtherErrDefKeys[key] {
								match = false
								break
							}
						}

						if match {
							// Get the first key from each set
							var thisErrorDefKey string
							for key := range thisErrDefKeys {
								thisErrorDefKey = key
								break
							}
							var thisOtherErrorDefKey string
							for key := range thisOtherErrDefKeys {
								thisOtherErrorDefKey = key
								break
							}

							finalThisPath := []interface{}{index, defKey, k, thisErrorDefKey}
							finalThisDocumentJSON := documentNamesToJSON[documentName]
							finalThisLocationPseudoJSON := GetPathDocumentCoordinatesPseudoJSON(finalThisPath, finalThisDocumentJSON)

							parseFailures = append(parseFailures, &SchemaParseFailure{
								DocumentName: otherDocumentName,
								Path:         []interface{}{otherIndex, otherDefKey, l, thisOtherErrorDefKey},
								Reason:       "PathCollision",
								Data: map[string]interface{}{
									"document": documentName,
									"path":     finalThisPath,
									"location": finalThisLocationPseudoJSON,
								},
							})
						}
					}
				}
			}
		}
	}

	return parseFailures
}
