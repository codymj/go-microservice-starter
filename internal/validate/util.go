package validate

// getPostHelloSchema returns json validator schema for post /greeting endpoint
func getPostHelloSchema() []byte {
	return []byte(`
	{
	  "$schema": "http://json-schema.org/draft-07/schema#",
	  "$id": "PostHelloSchema",
	  "type": "object",
	  "required": [
		"name"
	  ],
	  "properties": {
		"name": {
		  "type": "string",
		  "minLength": 1
		}
	  }
	}
	`)
}
