package connectionconfig

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	servicev1 "github.com/tobedeterminedhq/tbd/proto_gen/go/tbd/service/v1"
)

func Test_parseConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		reader  io.Reader
		want    *servicev1.ConnectionConfig
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "empty",
			reader: strings.NewReader("---\nname: test\n"),
			want: &servicev1.ConnectionConfig{
				Name: "test",
			},
			wantErr: assert.NoError,
		},
		{
			name:   "sqlite",
			reader: strings.NewReader("---\nname: test\nsqlite:\n    path: /tmp/test.db\n"),
			want: &servicev1.ConnectionConfig{
				Name: "test",
				Config: &servicev1.ConnectionConfig_Sqlite{
					Sqlite: &servicev1.ConnectionConfig_ConnectionConfigSqLite{
						Path: "/tmp/test.db",
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name:   "sqlite_in_memory",
			reader: strings.NewReader("---\nname: test\nsqlite_in_memory: {}\n"),
			want: &servicev1.ConnectionConfig{
				Name: "test",
				Config: &servicev1.ConnectionConfig_SqliteInMemory{
					SqliteInMemory: &servicev1.ConnectionConfig_ConnectionConfigSqLiteInMemory{},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "mysql",
			reader: strings.NewReader(`
---
name: test
mysql:
    host: localhost
    port: "2000"
`),
			want: &servicev1.ConnectionConfig{
				Name: "test",
				Config: &servicev1.ConnectionConfig_Mysql{
					Mysql: &servicev1.ConnectionConfig_ConnectionConfigMySql{
						Host: "localhost",
						Port: "2000",
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfig(tt.reader)

			tt.wantErr(t, err)
			// TODO: Improve this test and how the values are compared
			assert.Equalf(t, tt.want.String(), got.String(), "parseConfig() = %v, want %v", got, tt.want)
		})
	}
}
