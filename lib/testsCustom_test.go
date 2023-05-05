package lib

import (
	"testing"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTestNameCustomColumn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		test *servicev1.TestCustomColumn
		want string
	}{
		{
			name: "sample",
			test: &servicev1.TestCustomColumn{
				Model:    "user",
				Column:   "name",
				Path:     "user_123",
				TestName: "valid",
			},
			want: "test_user_name_valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GenerateTestNameCustomColumn(tt.test), "GenerateTestNameCustomColumn(%v)", tt.test)
		})
	}
}

// TODO Need a test for custom rendering of the Model so it takes the Path and not the model name
