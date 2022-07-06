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

// getPostUsersSchema returns json validator schema for POST /users
func getPostUsersSchema() []byte {
	return []byte(`
	{
	  "$schema": "http://json-schema.org/draft-07/schema#",
	  "$id": "PostUsersSchema",
	  "type": "object",
	  "required": [
		"username",
		"password",
		"email"
	  ],
	  "properties": {
		"username": {
		  "$id": "#/properties/username",
		  "type": "string",
		  "minLength": 3
		},
		"password": {
		  "$id": "#/properties/password",
		  "type": "string",
		  "minLength": 6
		},
		"email": {
		  "$id": "#/properties/email",
		  "type": "string",
		  "format": "email"
		}
	  }
	}
	`)
}
