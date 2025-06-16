package cmd

import (
	"fmt"
	"os"

	"github.com/CoffeeCat1920/Gochi/loader"
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

		for _, e := range entries {
			fmt.Printf("Name: %s, Command:%s \n", e.Name, e.Exec)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
