package schema

// MapSchemaParseFailuresToPseudoJson converts a list of schema parse failures
// to a pseudo-JSON format for error reporting
func MapSchemaParseFailuresToPseudoJson(
	schemaParseFailures []*SchemaParseFailure,
	telepactDocumentNameToJson map[string]string,
) []map[string]interface{} {
	pseudoJsonList := make([]map[string]interface{}, 0, len(schemaParseFailures))
	
	for _, f := range schemaParseFailures {
		location := GetPathDocumentCoordinatesPseudoJSON(
			f.Path,
			telepactDocumentNameToJson[f.DocumentName],
		)
		
		pseudoJson := make(map[string]interface{})
		pseudoJson["document"] = f.DocumentName
		pseudoJson["location"] = location
		pseudoJson["path"] = f.Path
		pseudoJson["reason"] = map[string]interface{}{f.Reason: f.Data}
		
		pseudoJsonList = append(pseudoJsonList, pseudoJson)
	}
	
	return pseudoJsonList
}
