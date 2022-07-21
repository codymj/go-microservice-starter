package validate

import (
	"bytes"
	"encoding/json"
)

const (
	InfoBeginValidatePostUsers  = "begin request validation of POST /users"
	InfoEndValidatePostUsers    = "finished request validation of POST /users"
	InfoBeginValidatePutUsersId = "begin request validation of PUT /users/{id}"
	InfoEndValidatePutUsersId   = "finished request validation of PUT /users/{id}"
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
		},
		"additionalProperties": false
	}
	`)
}

// getPutUsersIdSchema returns json validator schema for PUT /users/{id}
func getPutUsersIdSchema() []byte {
	return []byte(`
	{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"$id": "PutUsersIdSchema",
		"type": "object",
		"required": [
			"id",
			"username",
			"email",
			"isVerified",
			"createdOn",
			"updatedOn"
		],
		"properties": {
			"id": {
				"$id": "#/properties/id",
				"type": "string",
				"minLength": 36,
				"maxLength": 36
			},
			"username": {
				"$id": "#/properties/username",
				"type": "string",
				"minLength": 3
			},
			"email": {
				"$id": "#/properties/email",
				"type": "string",
				"format": "email"
			},
			"isVerified": {
				"$id": "#/properties/isVerified",
				"type": "boolean"
			},
			"createdOn": {
				"$id": "#/properties/createdOn",
				"type": "number"
			},
			"updatedOn": {
				"$id": "#/properties/updatedOn",
				"type": "number"
			}
		},
		"additionalProperties": false
	}
	`)
}
