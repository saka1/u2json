package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	convertOpt := convertOpt{
		enableQueryValueArray: false,
		enableStrictURL:       false,
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
		&convertOpt.enableStrictURL,
		"strict-url", "", false, "Check the input is URL",
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
