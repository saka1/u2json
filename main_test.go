package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	t.Parallel()

	cases := []struct {
		args     []string
		want     string
		wantCode int
	}{
		{
			[]string{"https://example.com/aaa"},
			`{"host":"example.com","path":"/aaa","scheme":"https"}`,
			0,
		},
		{
			[]string{"--use-ParseRequestURI", "https://example.com/aaa"},
			`{"host":"example.com","path":"/aaa","scheme":"https"}`,
			0,
		},
		{
			[]string{"https://example.com?q=v1&q=v2"},
			`{"host":"example.com","query":{"q":"v2"},"rawQuery":"q=v1\u0026q=v2","scheme":"https"}`,
			0,
		},
		{
			[]string{"--query-array", "https://example.com?q=v1&q=v2"},
			`{"host":"example.com","query":{"q":["v1","v2"]},"rawQuery":"q=v1\u0026q=v2","scheme":"https"}`,
			0,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.args), func(t *testing.T) {
			t.Parallel()
			var stdout, stderr bytes.Buffer
			code := run(tc.args, &stdout, &stderr)
			if code != tc.wantCode {
				t.Fatalf("exit code: got %d, want %d (stderr: %s)", code, tc.wantCode, stderr.String())
			}
			got := strings.TrimSpace(stdout.String())
			if got != tc.want {
				t.Fatalf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestNoArgs(t *testing.T) {
	t.Parallel()
	var stdout, stderr bytes.Buffer
	code := run(nil, &stdout, &stderr)
	if code != 1 {
		t.Fatalf("exit code: got %d, want 1", code)
	}
	if !strings.Contains(stderr.String(), "usage:") {
		t.Fatalf("expected usage message on stderr, got: %s", stderr.String())
	}
}

func TestInvalidURL(t *testing.T) {
	t.Parallel()
	var stdout, stderr bytes.Buffer
	code := run([]string{"--use-ParseRequestURI", "not a url"}, &stdout, &stderr)
	if code != 1 {
		t.Fatalf("exit code: got %d, want 1", code)
	}
	if stderr.Len() == 0 {
		t.Fatal("expected error on stderr")
	}
}
