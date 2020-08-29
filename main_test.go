package main

func ExampleMain() {
	cmd := createRootCmd()
	argsCases := [][]string{
		[]string{"https://example.com/aaa"},
		[]string{"--use-ParseRequestURI", "https://example.com/aaa"},
		[]string{"https://example.com?q=v1&q=v2"},
		[]string{"--query-array", "https://example.com?q=v1&q=v2"},
	}
	for _, args := range argsCases {
		cmd.SetArgs(args)
		cmd.Execute()
	}
	// Output: {"host":"example.com","path":"/aaa","scheme":"https"}
	// {"host":"example.com","path":"/aaa","scheme":"https"}
	// {"host":"example.com","query":{"q":"v2"},"rawQuery":"q=v1\u0026q=v2","scheme":"https"}
	// {"host":"example.com","query":{"q":["v1","v2"]},"rawQuery":"q=v1\u0026q=v2","scheme":"https"}
}
