/*
Copyright © 2023 Dex Wood
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssltool",
	Short: "Various SSL utilities",
	Long:  `Various SSL utilities - Dex Wood`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Args", args)
		if len(args) > 0 {
			// Check if first argument is a known subcommand
			var isSubcommand bool
			for _, c := range cmd.Commands() {
				if c.Name() == args[0] || c.HasAlias(args[0]) {
					isSubcommand = true
					break
				}
			}
			if !isSubcommand {
				// Delegate to details command with host argument
				newArgs := []string{"details", "--host", args[0]}
				newArgs = append(newArgs, args[1:]...)
				cmd.SetArgs(newArgs)
				if err := cmd.Execute(); err != nil {
					os.Exit(1)
				}
				return
			}
		}
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootExamples = `
ssltool -- www.example.com
ssltool details --host www.example.com
ssltool details --host www.example.com --cert`

func init() {
	rootCmd.Example = rootExamples
}
