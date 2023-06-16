/*
Copyright © 2023 Dex Wood
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ssltool",
	Short: "Various SSL utilities",
	Long:  `Various SSL utilities - Dex Wood`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootExamples = `ssltool details --host www.example.com
ssltool details --host www.example.com --cert`

func init() {
	rootCmd.Example = rootExamples
}
