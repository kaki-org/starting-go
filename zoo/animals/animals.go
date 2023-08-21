package animals

func FooFunc(n int) int {
	return internalFunc(n)
}

func internalFunc(n int) int {
	return n + 1
}
