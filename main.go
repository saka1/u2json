package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("u2json", flag.ContinueOnError)
	fs.SetOutput(stderr)
	queryArray := fs.Bool("query-array", false, "Parse multiple query params as array")
	useParseRequestURI := fs.Bool("use-ParseRequestURI", false, "Check the input with url.ParseRequestURI()")
	if err := fs.Parse(args); err != nil {
		return 1
	}

	urls := fs.Args()
	if len(urls) == 0 {
		_, _ = fmt.Fprintln(stderr, "usage: u2json [flags] URL...")
		return 1
	}

	opt := &convertOpt{
		enableQueryValueArray: *queryArray,
		useParseRequestURI:    *useParseRequestURI,
	}

	exitCode := 0
	for _, url := range urls {
		bin, err := convert(url, opt)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "u2json: %s\n", err)
			exitCode = 1
			continue
		}
		_, _ = fmt.Fprintln(stdout, string(bin))
	}
	return exitCode
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
