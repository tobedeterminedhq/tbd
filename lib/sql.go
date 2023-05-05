package lib

import (
	"strconv"
	"strings"
)

// AddLimitToSqlSelectStatement adds a limit to a SELECT statement. If the input starts with CREATE it does not add it.
func AddLimitToSqlSelectStatement(s string, limit uint) string {
	trimmed := strings.TrimSpace(s)
	lowerCased := strings.ToLower(trimmed)
	if strings.HasPrefix(lowerCased, "create") {
		return s
	}
	return strings.TrimSuffix(s, ";") + " LIMIT " + strconv.Itoa(int(limit))
}
