package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tobedeterminedhq/tbd/lib"
	"github.com/tobedeterminedhq/tbd/lib/connectionconfig"
)

var runCmdModelsOnly bool

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&runCmdModelsOnly, "models-only", "m", false, "only do the models and not the seeds")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the seeds and models.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

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

		sqls, err := lib.ProjectAndFsToSqlForViews(p, fs, db, runCmdModelsOnly, false)
		if err != nil {
			return err
		}

		bar := progressbar.Default(int64(len(sqls)))
		for _, s := range sqls {
			_, err := db.ExecContext(ctx, s[1])
			if err != nil {
				return fmt.Errorf("error running sql applying '%s': %w", s, err)
			}
			err = bar.Add(1)
			if err != nil {
				return err
			}
		}

		return nil
	},
}
