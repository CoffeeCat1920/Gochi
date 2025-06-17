package cmd

import (
	"fmt"
	"os"

	"github.com/CoffeeCat1920/Gochi/loader"
	"github.com/CoffeeCat1920/Gochi/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Gochi",
	Short: "Apps Launcher written in golang",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := loader.GetEntries()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		tui.Run(entries)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
