package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tobedeterminedhq/tbd/lib"
)

func init() {
	rootCmd.AddCommand(renderCmd)
}

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Returns the shape of the project as a dot viz.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, _, err := parseProject()
		if err != nil {
			return err
		}
		g, err := lib.ProjectToGraph(p)
		if err != nil {
			return err
		}
		vis, err := g.ToDotViz()
		if err != nil {
			return err
		}

		fmt.Println(string(vis))
		return nil
	},
}
