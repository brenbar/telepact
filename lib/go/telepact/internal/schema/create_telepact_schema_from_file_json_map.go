package schema

import (
	"regexp"

	"github.com/brenbar/telepact/lib/go/telepact"
)

// CreateTelepactSchemaFromFileJsonMap creates a TelepactSchema from a map of
// document names to JSON content strings.
//
// This function automatically adds the internal Telepact schema ("internal_")
// and conditionally adds the auth schema ("auth_") if any document references
// "struct.Auth_".
//
// Parameters:
//   - jsonDocuments: Map of document name to JSON content string
//
// Returns:
//   - *telepact.TelepactSchema: The parsed schema
//   - error: Any parsing errors encountered
func CreateTelepactSchemaFromFileJsonMap(jsonDocuments map[string]string) (*telepact.TelepactSchema, error) {
	// Copy the input map to avoid modifying the original
	finalJsonDocuments := make(map[string]string)
	for k, v := range jsonDocuments {
		finalJsonDocuments[k] = v
	}

	// Always add the internal schema
	finalJsonDocuments["internal_"] = GetInternalTelepactJson()

	// Determine if we need to add the auth schema by checking if any document
	// references "struct.Auth_"
	authRegex := regexp.MustCompile(`"struct\.Auth_"\s*:`)
	for _, jsonContent := range jsonDocuments {
		if authRegex.MatchString(jsonContent) {
			finalJsonDocuments["auth_"] = GetAuthTelepactJson()
			break
		}
	}

	// Parse the schema
	return ParseTelepactSchema(finalJsonDocuments)
}
