package main

import (
	"encoding/json"
	"testing"
)

var opt = &convertOpt{
	enableQueryValueArray: false,
	useParseRequestURI:    false,
}

func TestBasic(t *testing.T) {
	js, err := convert("https://example.com:80/foo/bar?q1=v1&q2=v2&q3&q2=v3#fr", opt)
	if err != nil {
		t.Error(err)
	}
	var dat map[string]interface{}
	err = json.Unmarshal(js, &dat)
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

func TestLargePort(t *testing.T) {
	js, err := convert("https://example.com:65536", opt)
	if err != nil {
		t.Error(err)
	}
	var dat map[string]interface{}
	json.Unmarshal(js, &dat)
	t.Logf("%v", dat)
	if dat["port"].(float64) != 65536 {
		t.Error(dat["port"])
	}
}

func TestUserPassword(t *testing.T) {
	js, err := convert("https://u:pass@example.com", opt)
	if err != nil {
		t.Error(err)
	}
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

func TestQueryArray(t *testing.T) {
	opt := &convertOpt{
		enableQueryValueArray: true,
		useParseRequestURI:    false,
	}
	js, err := convert("https://example.com?q1=v1&q2=v2&q3&q2=v3", opt)
	if err != nil {
		t.Error(err)
	}
	var dat map[string]interface{}
	json.Unmarshal(js, &dat)
	params := dat["query"].(map[string]interface{})
	q1 := params["q1"].([]interface{})
	if q1[0].(string) != "v1" || len(q1) != 1 {
		t.Error(q1)
	}
	q2 := params["q2"].([]interface{})
	if q2[0].(string) != "v2" || q2[1].(string) != "v3" || len(q2) != 2 {
		t.Error(q2)
	}
	q3 := params["q3"].([]interface{})
	if q3[0].(string) != "" || len(q3) != 1 {
		t.Error(q3)
	}
}

func TestStrictURL(t *testing.T) {
	opt := &convertOpt{
		enableQueryValueArray: false,
		useParseRequestURI:    true,
	}
	_, err := convert("aaa", opt)
	if err == nil {
		t.Error()
	}
}
