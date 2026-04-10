package main

import (
	"encoding/json"
	"testing"
)

type testURLResult struct {
	Fragment string          `json:"fragment,omitempty"`
	Host     string          `json:"host,omitempty"`
	Password string          `json:"password,omitempty"`
	Path     string          `json:"path,omitempty"`
	Port     int             `json:"port,omitempty"`
	Query    json.RawMessage `json:"query,omitempty"`
	RawQuery string          `json:"rawQuery,omitempty"`
	Scheme   string          `json:"scheme,omitempty"`
	User     string          `json:"user,omitempty"`
}

var opt = &convertOpt{
	enableQueryValueArray: false,
	useParseRequestURI:    false,
}

func TestBasic(t *testing.T) {
	js, err := convert("https://example.com:80/foo/bar?q1=v1&q2=v2&q3&q2=v3#fr", opt)
	if err != nil {
		t.Fatal(err)
	}
	var got testURLResult
	if err := json.Unmarshal(js, &got); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", got)

	if got.Scheme != "https" {
		t.Errorf("scheme: got %q, want %q", got.Scheme, "https")
	}
	if got.Host != "example.com" {
		t.Errorf("host: got %q, want %q", got.Host, "example.com")
	}
	if got.Port != 80 {
		t.Errorf("port: got %d, want %d", got.Port, 80)
	}
	if got.Path != "/foo/bar" {
		t.Errorf("path: got %q, want %q", got.Path, "/foo/bar")
	}
	if got.Fragment != "fr" {
		t.Errorf("fragment: got %q, want %q", got.Fragment, "fr")
	}
	if got.RawQuery != "q1=v1&q2=v2&q3&q2=v3" {
		t.Errorf("rawQuery: got %q, want %q", got.RawQuery, "q1=v1&q2=v2&q3&q2=v3")
	}

	var query map[string]string
	if err := json.Unmarshal(got.Query, &query); err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", query)
	if query["q1"] != "v1" {
		t.Errorf("query q1: got %q, want %q", query["q1"], "v1")
	}
	if query["q2"] != "v3" {
		t.Errorf("query q2: got %q, want %q", query["q2"], "v3")
	}
	if query["q3"] != "" {
		t.Errorf("query q3: got %q, want %q", query["q3"], "")
	}
}

func TestLargePort(t *testing.T) {
	js, err := convert("https://example.com:65536", opt)
	if err != nil {
		t.Fatal(err)
	}
	var got testURLResult
	if err := json.Unmarshal(js, &got); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", got)
	if got.Port != 65536 {
		t.Errorf("port: got %d, want %d", got.Port, 65536)
	}
}

func TestUserPassword(t *testing.T) {
	js, err := convert("https://u:pass@example.com", opt)
	if err != nil {
		t.Fatal(err)
	}
	var got testURLResult
	if err := json.Unmarshal(js, &got); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", got)
	if got.User != "u" {
		t.Errorf("user: got %q, want %q", got.User, "u")
	}
	if got.Password != "pass" {
		t.Errorf("password: got %q, want %q", got.Password, "pass")
	}
}

func TestQueryArray(t *testing.T) {
	opt := &convertOpt{
		enableQueryValueArray: true,
		useParseRequestURI:    false,
	}
	js, err := convert("https://example.com?q1=v1&q2=v2&q3&q2=v3", opt)
	if err != nil {
		t.Fatal(err)
	}
	var got testURLResult
	if err := json.Unmarshal(js, &got); err != nil {
		t.Fatal(err)
	}

	var params map[string][]string
	if err := json.Unmarshal(got.Query, &params); err != nil {
		t.Fatal(err)
	}
	q1 := params["q1"]
	if len(q1) != 1 || q1[0] != "v1" {
		t.Errorf("query q1: got %v, want [v1]", q1)
	}
	q2 := params["q2"]
	if len(q2) != 2 || q2[0] != "v2" || q2[1] != "v3" {
		t.Errorf("query q2: got %v, want [v2 v3]", q2)
	}
	q3 := params["q3"]
	if len(q3) != 1 || q3[0] != "" {
		t.Errorf("query q3: got %v, want []", q3)
	}
}

func TestStrictURL(t *testing.T) {
	opt := &convertOpt{
		enableQueryValueArray: false,
		useParseRequestURI:    true,
	}
	_, err := convert("aaa", opt)
	if err == nil {
		t.Fatal("expected error for invalid URL with ParseRequestURI, got nil")
	}
}
