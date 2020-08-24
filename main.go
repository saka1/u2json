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
			bin, err := convert(args[0], &convertOpt)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Print(string(bin))
		},
	}

	rootCmd.Flags().BoolVarP(
		&convertOpt.enableQueryValueArray,
		"query-array", "", false, "Parse duplicated query params as array",
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
