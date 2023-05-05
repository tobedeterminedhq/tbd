package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLimitToSqlSelectStatement(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		limit uint
		want  string
	}{
		{
			name:  "simple SELECT example with semi colon",
			s:     "SELECT * FROM users;",
			limit: 100,
			want:  "SELECT * FROM users LIMIT 100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AddLimitToSqlSelectStatement(tt.s, tt.limit), "AddLimitToSqlSelectStatement(%v, %v)", tt.s, tt.limit)
		})
	}
}
