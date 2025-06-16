package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/CoffeeCat1920/Gochi/parser"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Gochi",
	Short: "Apps Launcher written in golang",
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		path := "/usr/share/applications"

		files, err := afero.ReadDir(fs, path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		for _, f := range files {
			if filepath.Ext(f.Name()) == ".desktop" {
				app, err := parser.ParseDesktopFile(filepath.Join(path, f.Name()))
				if err != nil {
					continue
				}
				fmt.Printf("Name: %s, Exec: %s\n", app.Name, app.Exec)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
