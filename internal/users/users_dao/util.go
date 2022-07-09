package users_dao

import (
	"fmt"
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

// Close is a wrapper for defer Close() methods
func Close(err *error, c io.Closer) {
	if e := c.Close(); err == nil {
		*err = e
	}
}
