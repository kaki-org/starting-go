package main

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

func TestingBenchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// do something
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkAppendAllocateEveryTime(b *testing.B) {
	base := []string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base = append(base, fmt.Sprintf("no%d", i))
	}
}

func BenchmarkAppendAllocateOnce(b *testing.B) {
	base := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base[i] = fmt.Sprintf("no%d", i)
	}
}

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
