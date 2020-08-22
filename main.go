package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
)

type res struct {
	Scheme      string `json:"scheme"`
	Hostname    string `json:"host"`
	Port        int    `json:"port"`
	EscapedPath string `json:"path"`
}

func convert(input string) []byte {
	u, err := url.ParseRequestURI(input)
	if err != nil {
		log.Fatal(err)
	}
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		log.Fatal(err)
	}
	res := res{
		Scheme:      u.Scheme,
		Hostname:    u.Hostname(),
		Port:        port,
		EscapedPath: u.EscapedPath(),
	}
	m, err := json.Marshal(&res)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func main() {
	input := os.Args[1]

	fmt.Print(string(convert(input)))
}
