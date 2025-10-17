package schema

// GetAuthTelepactJson returns the auth Telepact schema JSON.
// This schema defines struct.Auth_ and related authentication types.
func GetAuthTelepactJson() string {
	data, err := schemaFiles.ReadFile("json/auth.telepact.json")
	if err != nil {
		panic("failed to read auth.telepact.json: " + err.Error())
	}
	return string(data)
}
