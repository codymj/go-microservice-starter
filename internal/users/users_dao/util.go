package users_dao

import (
	"fmt"
	"io"
	"strings"
)

// buildWhereClause using query params, valid list of params and a map of DTO fields to DB fields
func buildWhereClause(params map[string]string) (string, []any) {
	clauses := make([]string, 0)
	vals := make([]any, 0)
	i := 1
	for _, validParam := range validUserParams {
		_, ok := params[validParam]
		if ok {
			databaseField := paramToColumn[validParam]
			vals = append(vals, params[validParam])
			clause := fmt.Sprintf("%s = $%d", databaseField, i)
			clauses = append(clauses, clause)
			i++
		}
	}

	return strings.Join(clauses, " and "), vals
}

// Close is a wrapper for defer Close() methods
func Close(err *error, c io.Closer) {
	if e := c.Close(); err == nil {
		*err = e
	}
}
