/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cvvault",
	Short: "An Application to manage json schema based cvs.",
	Long: `CV Vault is an application that helps you manage multiple cvs stored in json files. 

The schema is based on https://jsonresume.org/. For some features there are some transparent extensions for the cvvault project files.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
