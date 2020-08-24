package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type convertOpt struct {
	enableQueryValueArray bool
	enableStrictURL       bool
}

func convert(input string, opt *convertOpt) ([]byte, error) {
	var u *url.URL
	var err error
	if opt.enableStrictURL {
		u, err = url.ParseRequestURI(input)
	} else {
		u, err = url.Parse(input)
	}
	if err != nil {
		return nil, fmt.Errorf("Fail to parse as URL: %s", err)
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
			return nil, fmt.Errorf("Fail to parse port: %s", err)
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
		if opt.enableQueryValueArray {
			result["query"] = u.Query()
		} else {
			queryKv := map[string]string{}
			for k, v := range u.Query() {
				// Last key wins if enableQueryValueArray is false
				// But another strategy may be useful
				// SEE: https://stackoverflow.com/a/1746566
				queryKv[k] = v[len(v)-1]
			}
			result["query"] = queryKv
		}
	}
	// fragment
	if u.Fragment != "" {
		result["fragment"] = u.Fragment
	}

	bin, err := json.Marshal(&result)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal to JSON: %s", err)
	}
	return bin, nil
}
