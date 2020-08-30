package main

func ExampleMain() {
	cmd := createRootCmd()
	argsCases := [][]string{
		{"https://example.com/aaa"},
		{"--use-ParseRequestURI", "https://example.com/aaa"},
		{"https://example.com?q=v1&q=v2"},
		{"--query-array", "https://example.com?q=v1&q=v2"},
	}
	for _, args := range argsCases {
		cmd.SetArgs(args)
		_ = cmd.Execute()
	}
	// Output: {"host":"example.com","path":"/aaa","scheme":"https"}
	// {"host":"example.com","path":"/aaa","scheme":"https"}
	// {"host":"example.com","query":{"q":"v2"},"rawQuery":"q=v1\u0026q=v2","scheme":"https"}
	// {"host":"example.com","query":{"q":["v1","v2"]},"rawQuery":"q=v1\u0026q=v2","scheme":"https"}
}
