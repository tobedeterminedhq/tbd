package connectionconfig

import (
	"context"
	"fmt"
	"io"

	"github.com/tobedeterminedhq/tbd/lib/databases"
	"github.com/tobedeterminedhq/tbd/lib/databasesImplementation"
	servicev1 "github.com/tobedeterminedhq/tbd/proto_gen/go/tbd/service/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"sigs.k8s.io/yaml"
)

func parseConfig(reader io.Reader) (*servicev1.ConnectionConfig, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	config := &servicev1.ConnectionConfig{}
	jsonBytes, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(jsonBytes, config); err != nil {
		return nil, err
	}
	return config, nil
}

// TODO Think about context throughout the application
func NewConnectionConfig(reader io.Reader) (databases.Database, error) {
	c, err := parseConfig(reader)
	if err != nil {
		return nil, err
	}

	switch c.GetConfig().(type) {
	case *servicev1.ConnectionConfig_Sqlite:
		return NewSqlLite(c.GetSqlite())
	case *servicev1.ConnectionConfig_SqliteInMemory:
		return NewSqlLiteInMemory(c.GetSqliteInMemory())
	case *servicev1.ConnectionConfig_Duckdb:
		return NewDuckDB(c.GetDuckdb())
	case *servicev1.ConnectionConfig_Postgres:
		return NewPostgres(c.GetPostgres())
	case *servicev1.ConnectionConfig_Mysql:
		return NewMySQL(c.GetMysql())
	case *servicev1.ConnectionConfig_BigQuery:
		return NewBigQuery(c.GetBigQuery())
	default:
		return nil, fmt.Errorf("unknown connection config: %v", c)
	}
}

func NewSqlLite(connection *servicev1.ConnectionConfig_ConnectionConfigSqLite) (databases.Database, error) {
	return databasesImplementation.NewSqlLite(connection.GetPath())
}

func NewSqlLiteInMemory(_ *servicev1.ConnectionConfig_ConnectionConfigSqLiteInMemory) (databases.Database, error) {
	return databasesImplementation.NewSqlLiteInMemory()
}

func NewDuckDB(details *servicev1.ConnectionConfig_ConnectionConfigDuckDB) (databases.Database, error) {
	return databasesImplementation.NewDuckDB(details.GetPath(), details.GetParams())
}

func NewPostgres(connection *servicev1.ConnectionConfig_ConnectionConfigPostgres) (databases.Database, error) {
	return databasesImplementation.NewPostgres(postgresToConnectionString(
		connection.GetHost(),
		connection.GetPort(),
		connection.GetDatabase(),
		connection.GetUser(),
		connection.GetPassword(),
		connection.GetParams(),
	))
}

func postgresToConnectionString(host string, port string, database string, user string, password string, params map[string]string) string {
	paramString := ""
	for k, v := range params {
		paramString += fmt.Sprintf("%s=%s ", k, v)
	}
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s "+paramString, host, port, database, user, password)
}

func NewMySQL(connection *servicev1.ConnectionConfig_ConnectionConfigMySql) (databases.Database, error) {
	return databasesImplementation.NewMySql(mysqlToConnectionString(
		connection.GetProtocol(),
		connection.GetHost(),
		connection.GetPort(),
		connection.GetDatabase(),
		connection.GetUsername(),
		connection.GetPassword(),
		connection.GetParams(),
	))
}

func mysqlToConnectionString(protocol string, host string, port string, database string, user string, password string, params map[string]string) string {
	paramString := ""
	for k, v := range params {
		paramString += fmt.Sprintf("%s=%s ", k, v)
	}
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", user, password, protocol, host, port, database, paramString)
}

func NewBigQuery(connection *servicev1.ConnectionConfig_ConnectionConfigBigQuery) (databases.Database, error) {
	return databasesImplementation.NewBigQuery(context.Background(), connection.GetProjectId(), connection.GetDatasetId())
}
