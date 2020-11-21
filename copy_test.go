package gocopy

import (
	"testing"
)

func BenchmarkCopy(b *testing.B) {
	var o = make(mytype, 100)
	for i := 0; i < b.N; i++ {
		o.Copy()
	}
}

func BenchmarkCopyWithMake(b *testing.B) {
	var o = make(mytype, 100)
	for i := 0; i < b.N; i++ {
		o.CopyWithMake()
	}
}
