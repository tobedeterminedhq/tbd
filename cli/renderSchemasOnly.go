package cli

import (
	"fmt"
	"strings"

	"github.com/tobedeterminedhq/tbd/lib/connectionconfig"

	"github.com/samber/lo"

	"github.com/spf13/cobra"
	"github.com/tobedeterminedhq/tbd/lib"
)

func init() {
	rootCmd.AddCommand(renderSchemasOnlyCmd)
}

var renderSchemasOnlyCmd = &cobra.Command{
	Use:   "render-schema-only",
	Short: "Render the schemas for the databases",
	RunE: func(cmd *cobra.Command, args []string) error {
		p, fs, err := parseProject()
		if err != nil {
			return err
		}

		reader, err := getConfigFile()
		if err != nil {
			return err
		}
		db, err := connectionconfig.NewConnectionConfig(reader)
		if err != nil {
			return err
		}

		sqls, err := lib.ProjectAndFsToSqlForViews(p, fs, db, false, true)
		if err != nil {
			return err
		}

		sqlsOut := lo.Map(sqls, func(sql [2]string, index int) string {
			return sql[1]
		})
		fmt.Println(strings.Join(sqlsOut, ""))

		return nil
	},
}
