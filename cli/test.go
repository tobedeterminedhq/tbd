package cli

import (
	"context"
	"fmt"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/tobedeterminedhq/tbd/lib"
	"github.com/tobedeterminedhq/tbd/lib/connectionconfig"
)

var testCmdRenderOnly bool

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.Flags().BoolVarP(&testCmdRenderOnly, "render", "r", false, "render the sql tests and print to the terminal without running them")
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Runs the tests against the sqlite database.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		p, fs, err := parseProject()
		if err != nil {
			return err
		}

		tests, err := lib.ReturnTestsSQL(p, fs)
		if err != nil {
			return fmt.Errorf("error getting tests: %w", err)
		}

		if testCmdRenderOnly {
			for name, test := range tests {
				fmt.Println("-- " + name)
				fmt.Println(test)
				fmt.Println()
			}
			return nil
		}

		reader, err := getConfigFile()
		if err != nil {
			return err
		}
		db, err := connectionconfig.NewConnectionConfig(reader)
		if err != nil {
			return err
		}

		bar := progressbar.Default(int64(len(tests)))
		for name, t := range tests {
			// TODO Need function that can run bank of tests
			err := lib.RunTestSql(ctx, db, t)
			if err != nil {
				return fmt.Errorf("error running test '%s' with sql '%s': %w", name, t, err)
			}
			err = bar.Add(1)
			if err != nil {
				return err
			}
		}
		return nil
	},
}
