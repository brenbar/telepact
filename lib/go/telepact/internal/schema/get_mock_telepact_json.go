package schema

// GetMockTelepactJson returns the mock Telepact schema JSON.
// This schema defines _ext.Call_ and _ext.Stub_ types for testing.
func GetMockTelepactJson() string {
	data, err := schemaFiles.ReadFile("json/mock_internal.telepact.json")
	if err != nil {
		panic("failed to read mock_internal.telepact.json: " + err.Error())
	}
	return string(data)
}
