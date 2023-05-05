package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tobedeterminedhq/tbd/lib"
)

// var renderSqlCmdFull bool
func init() {
	rootCmd.AddCommand(renderSqlCmd)

	// TODO Implement a select with not just the parents
	// renderSqlCmd.Flags().BoolVarP(&renderSqlCmdFull, "full", "f", false, "render full sql instead of just the model that refers to other models")
}

var renderSqlCmd = &cobra.Command{
	Use:   "sql [model]",
	Short: "Returns the sql select for a model.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		p, fs, err := parseProject()
		if err != nil {
			return err
		}
		{
			err := lib.ValidateModelName(args[0])
			if err != nil {
				return fmt.Errorf("invalid model name: %w", err)
			}
		}
		s, err := lib.ProjectAndFsToQuerySql(p, fs, args[0])
		if err != nil {
			return err
		}
		fmt.Println(s)
		return nil
	},
}
