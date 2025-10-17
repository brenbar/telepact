package schema

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/brenbar/telepact/lib/go/telepact"
)

// GetSchemaFileMap loads all .telepact.json files from a directory
// Returns a map of relative file paths to their JSON content
// Raises TelepactSchemaParseError if there are any issues (directories, invalid filenames)
func GetSchemaFileMap(directory string) (map[string]string, error) {
	finalJsonDocuments := make(map[string]string)
	schemaParseFailures := []*SchemaParseFailure{}
	
	// Walk the directory tree
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Get relative path
		relativePath, err := filepath.Rel(directory, path)
		if err != nil {
			return err
		}
		
		// Skip the root directory itself
		if relativePath == "." {
			return nil
		}
		
		// Check if it's a directory
		if info.IsDir() {
			schemaParseFailures = append(schemaParseFailures, &SchemaParseFailure{
				DocumentName: relativePath,
				Path:         []interface{}{},
				Reason:       "DirectoryDisallowed",
				Data:         map[string]interface{}{},
			})
			finalJsonDocuments[relativePath] = "[]"
			return nil
		}
		
		// Read the file
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		
		// Check if filename ends with .telepact.json
		if !strings.HasSuffix(relativePath, ".telepact.json") {
			schemaParseFailures = append(schemaParseFailures, &SchemaParseFailure{
				DocumentName: relativePath,
				Path:         []interface{}{},
				Reason:       "FileNamePatternInvalid",
				Data:         map[string]interface{}{"expected": "*.telepact.json"},
			})
		}
		
		finalJsonDocuments[relativePath] = string(content)
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// If there were parse failures, return error
	if len(schemaParseFailures) > 0 {
		return nil, &telepact.TelepactSchemaParseError{
			SchemaParseFailures: schemaParseFailures,
			DocumentNamesToJSON: finalJsonDocuments,
		}
	}
	
	return finalJsonDocuments, nil
}
