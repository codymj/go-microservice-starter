package users_dao

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
)

// buildWhereClause using query params, valid list of params and a map of DTO fields to DB fields
func buildWhereClause(params map[string]string, validParams []string, paramToDB map[string]string) string {
	clauses := make([]string, 0)
	for _, validParam := range validParams {
		_, ok := params[validParam]
		if ok {
			databaseField := paramToDB[validParam]
			value := params[validParam]
			clause := fmt.Sprintf("%s='%s'", databaseField, value)
			clauses = append(clauses, clause)
		}
	}

	return strings.Join(clauses, " and ")
}

// hash a plain-text password and returns stringified hash
func hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// Close is a wrapper for defer Close() methods
func Close(err *error, c io.Closer) {
	if e := c.Close(); err == nil {
		*err = e
	}
}
