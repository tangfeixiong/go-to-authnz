/*
   Borrowed from https://github.com/coreos/dex/blob/master/cmd/dex/main.go
*/
package cmd

// package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "dex",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(2)
		},
	}
	rootCmd.AddCommand(commandServe())
	rootCmd.AddCommand(commandVersion())
	return rootCmd
}

func dex_main() {
	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
