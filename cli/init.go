package cli

import (
	"errors"
	"os"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tobedeterminedhq/tbd/lib"
	"github.com/tobedeterminedhq/tbd/lib/filesystemhelpers"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

// TODO Add test works and add tests works for folders with a dotfile.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project. Creates a new directory and initializes a new project in it.",
	Long: `Initialize a new project. Creates a new directory and initializes a new project in it.

If the directory is not empty, the command will fail.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if the directory is empty
		isEmpty, err := isEmptyBarDotfiles(".")
		if err != nil {
			return err
		}
		if !isEmpty {
			return errors.New("directory is not empty")
		}

		// Write the project to the current directory
		filesystem := lib.Init()
		return filesystemhelpers.WriteFSToDisk(filesystem, "init", ".")
	},
}

func isEmptyBarDotfiles(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()
	names, err := f.Readdirnames(-1) // Or f.Readdir(1)
	if err != nil {
		return false, err
	}
	filtered := lo.Filter(names, func(name string, _ int) bool {
		return name[:1] != "."
	})
	if len(filtered) > 1 {
		return false, nil
	}
	return true, err // Either not empty or error, suits both cases
}
