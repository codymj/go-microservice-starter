package validate

// getPostGreetingSchema returns json validator schema for POST /greeting
func getPostGreetingSchema() []byte {
	return []byte(`
	{
	  "$schema": "http://json-schema.org/draft-07/schema#",
	  "$id": "PostGreetingSchema",
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
