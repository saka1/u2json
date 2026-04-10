package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type convertOpt struct {
	enableQueryValueArray bool
	useParseRequestURI    bool
}

type urlResult struct {
	Fragment string `json:"fragment,omitempty"`
	Host     string `json:"host,omitempty"`
	Password string `json:"password,omitempty"`
	Path     string `json:"path,omitempty"`
	Port     int    `json:"port,omitempty"`
	Query    any    `json:"query,omitempty"`
	RawQuery string `json:"rawQuery,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	User     string `json:"user,omitempty"`
}

func convert(input string, opt *convertOpt) ([]byte, error) {
	var u *url.URL
	var err error
	if opt.useParseRequestURI {
		u, err = url.ParseRequestURI(input)
	} else {
		u, err = url.Parse(input)
	}
	if err != nil {
		return nil, err
	}

	result := urlResult{}

	result.Scheme = u.Scheme

	if u.User != nil {
		result.User = u.User.Username()
		if p, ok := u.User.Password(); ok {
			result.Password = p
		}
	}

	result.Host = u.Hostname()

	if u.Port() != "" {
		port, err := strconv.Atoi(u.Port())
		if err != nil {
			return nil, fmt.Errorf("fail to parse port: %w", err)
		}
		result.Port = port
	}

	result.Path = u.Path

	if u.RawQuery != "" {
		result.RawQuery = u.RawQuery
		if opt.enableQueryValueArray {
			result.Query = u.Query()
		} else {
			queryKv := map[string]string{}
			for k, v := range u.Query() {
				queryKv[k] = v[len(v)-1]
			}
			result.Query = queryKv
		}
	}

	result.Fragment = u.Fragment

	bin, err := json.Marshal(&result)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal to JSON: %w", err)
	}
	return bin, nil
}
