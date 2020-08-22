package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
)

func convert(input string) []byte {
	u, err := url.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	result := map[string]interface{}{
		"scheme":   u.Scheme,
		"host":     u.Hostname(),
		"path":     u.Path,
		"fragment": u.Fragment,
		"rawQuery": u.RawQuery,
	}
	// port
	if u.Port() != "" {
		port, err := strconv.Atoi(u.Port())
		if err != nil {
			log.Fatal(err)
		}
		result["port"] = port
	}
	// query
	queryKv := map[string]string{}
	for k, v := range u.Query() {
		// Last key wins in current implementation
		// But another strategy may be useful
		// SEE: https://stackoverflow.com/a/1746566
		queryKv[k] = v[len(v)-1]
	}
	result["query"] = queryKv

	bin, err := json.Marshal(&result)
	if err != nil {
		log.Fatal(err)
	}
	return bin
}

func main() {
	//TODO handle empty args
	input := os.Args[1]
	fmt.Print(string(convert(input)))
}
