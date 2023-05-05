package cli

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tobedeterminedhq/tbd/lib"
	servicev1 "github.com/tobedeterminedhq/tbd/proto_gen/go/tbd/service/v1"
)

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file location, default is '.tbd.config.yaml', then '$HOME/.tbd.config.yaml'")
}

var rootCmd = &cobra.Command{
	Use:   "tbd",
	Short: "tbd is a very fast and useful database tool.",
}

// Execute is the entry point for the cli tool. To be imported with cli.Execute()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getConfigFile() (io.Reader, error) {
	if cfgFile != "" {
		return os.Open(cfgFile)
	}
	if _, err := os.Stat(".tbd.config.yaml"); err == nil {
		return os.Open(".tbd.config.yaml")
	}
	if _, err := os.Stat("$HOME/.tbd.config.yaml"); err == nil {
		return os.Open("$HOME/.tbd.config.yaml")
	}
	return nil, fmt.Errorf("no config file found")
}

func parseProject() (*servicev1.Project, fs.FS, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}
	c, err := lib.ParseConfigFromPath(filepath.Join(dir, "project.yml"))
	if err != nil {
		return nil, nil, err
	}
	fs := os.DirFS(".")
	p, err := lib.ParseProject(c, fs, "")
	if err != nil {
		return nil, nil, err
	}
	return p, fs, nil
}
