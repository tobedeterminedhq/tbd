package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tobedeterminedhq/tbd/lib"
)

func init() {
	rootCmd.AddCommand(compileCmd)
}

var compileCmd = &cobra.Command{
	Use:   "compile [model]",
	Short: "Compile a sql command for a model",
	Long:  `Compile a sql command for a model that can be used, generally always uses SELECT statements`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		modelName := args[0]

		{
			err := lib.ValidateModelName(modelName)
			if err != nil {
				return fmt.Errorf("invalid model name: %w", err)
			}
		}

		p, fs, err := parseProject()
		if err != nil {
			return err
		}

		sql, err := lib.ProjectAndFsToQuerySql(p, fs, modelName)
		if err != nil {
			return err
		}

		cmd.Print(sql)
		return nil
	},
}
