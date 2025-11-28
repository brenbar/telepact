package schema

import (
	"embed"
)

//go:embed json/*.json
var schemaFiles embed.FS

// GetInternalTelepactJson returns the internal Telepact schema JSON.
// This schema defines core types used internally by Telepact such as
// ErrorInvalidRequest_, ErrorInternal_, etc.
func GetInternalTelepactJson() string {
	data, err := schemaFiles.ReadFile("json/internal.telepact.json")
	if err != nil {
		panic("failed to read internal.telepact.json: " + err.Error())
	}
	return string(data)
}
