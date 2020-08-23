package main

import (
	"encoding/json"
	"testing"
)

func TestBasic(t *testing.T) {
	js := convert("https://example.com:80/foo/bar?q1=v1&q2=v2&q3&q2=v3#fr")
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
	if dat["fragment"] != "fr" {
		t.Error(dat["fragment"])
	}
	if dat["rawQuery"] != "q1=v1&q2=v2&q3&q2=v3" {
		t.Error(dat["rawQuery"])
	}

	queryParams := dat["query"].(map[string]interface{})
	t.Logf("%v", queryParams)
	if queryParams["q1"].(string) != "v1" {
		t.Error(queryParams["q1"])
	}
	if queryParams["q2"].(string) != "v3" {
		t.Error(queryParams["q2"])
	}
	if queryParams["q3"].(string) != "" {
		t.Error(queryParams["q3"])
	}
}

func TestSimpleUrl(t *testing.T) {
	js := convert("https://example.com/")
	var dat map[string]interface{}
	err := json.Unmarshal(js, &dat)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", dat)
	if dat["host"] != "example.com" {
		t.Error(dat["host"])
	}
	if dat["path"] != "/" {
		t.Error(dat["path"])
	}
}

func TestLargePort(t *testing.T) {
	js := convert("https://example.com:65536")
	var dat map[string]interface{}
	json.Unmarshal(js, &dat)
	t.Logf("%v", dat)
	if dat["port"].(float64) != 65536 {
		t.Error(dat["port"])
	}
}

func TestUserPassword(t *testing.T) {
	js := convert("https://u:pass@example.com")
	var dat map[string]interface{}
	json.Unmarshal(js, &dat)
	t.Logf("%v", dat)
	if dat["user"] != "u" {
		t.Error(dat["user"])
	}
	if dat["password"] != "pass" {
		t.Error(dat["password"])
	}
}
