package main

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
)

func convert(input string) []byte {
	u, err := url.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	result := map[string]interface{}{}
	// scheme
	if u.Scheme != "" {
		result["scheme"] = u.Scheme
	}
	// user
	if u.User != nil {
		usr := u.User
		if usr.Username() != "" {
			result["user"] = usr.Username()
		}
		password, passwordSet := usr.Password()
		if passwordSet {
			result["password"] = password
		}
	}
	// host
	if u.Hostname() != "" {
		result["host"] = u.Hostname()
	}
	// port
	if u.Port() != "" {
		port, err := strconv.Atoi(u.Port())
		if err != nil {
			log.Fatal(err)
		}
		result["port"] = port
	}
	// path
	if u.Path != "" {
		result["path"] = u.Path
	}
	// query
	if u.RawQuery != "" {
		result["rawQuery"] = u.RawQuery
	}
	if u.RawQuery != "" {
		queryKv := map[string]string{}
		for k, v := range u.Query() {
			// Last key wins in current implementation
			// But another strategy may be useful
			// SEE: https://stackoverflow.com/a/1746566
			queryKv[k] = v[len(v)-1]
		}
		result["query"] = queryKv
	}
	// fragment
	if u.Fragment != "" {
		result["fragment"] = u.Fragment
	}

	bin, err := json.Marshal(&result)
	if err != nil {
		log.Fatal(err)
	}
	return bin
}
