package main

import (
	"encoding/json"
	"testing"
)

func TestConvert(t *testing.T) {
	js := convert("https://example.com:80/foo/bar")
	var dat map[string]interface{}
	err := json.Unmarshal(js, &dat)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", dat)
	if dat["scheme"] != "https" {
		t.Error(dat["scheme"])
	}
	if dat["host"] != "example.com" {
		t.Error(dat["host"])
	}
	if dat["port"].(float64) != 80 {
		t.Error(dat["port"])
	}
	if dat["path"] != "/foo/bar" {
		t.Error(dat["path"])
	}
}
