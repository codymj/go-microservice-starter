package validate

import (
	"bytes"
	"encoding/json"
)

// compactJson returns a compacted JSON byte array for logging
func compactJson(jsn []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, jsn); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

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
