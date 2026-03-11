# Underscore


```go
// Unused variables are an error in Go.
	// The underscore lets you "use" a variable but discard its value.
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a5, s4, bs
	// Usually you use it to ignore one of the return values of a function
	// For example, in a quick and dirty script you might ignore the
	// error value returned from os.Create, and expect that the file
	// will always be created.
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is how you write to a file, by the way")
	file.Close()

	// Output of course counts as using a variable.
	fmt.Println(s, c, a4, s3, d2, m)
```