package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	convertOpt := convertOpt{
		enableQueryValueArray: false,
		useParseRequestURI:    false,
	}
	var rootCmd = &cobra.Command{
		Use:   "u2json",
		Short: "Convert URL to JSON",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			errorFound := false
			for _, url := range args {
				bin, err := convert(url, &convertOpt)
				if err != nil {
					fmt.Fprintf(os.Stderr, "u2json: %s\n", err)
					errorFound = true
					continue
				}
				fmt.Println(string(bin))
			}
			if errorFound {
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().BoolVarP(
		&convertOpt.enableQueryValueArray,
		"query-array", "", false, "Parse multiple query params as array",
	)
	rootCmd.Flags().BoolVarP(
		&convertOpt.useParseRequestURI,
		"use-ParseRequestURI", "", false, "Check the input with url.ParseRequestURI()",
	)
	return rootCmd
}

func main() {
	rootCmd := createRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
