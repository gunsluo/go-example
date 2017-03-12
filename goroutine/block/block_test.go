package main

import "testing"

func BenchmarkTest(b *testing.B) {
	test()
}

func BenchmarkTestBlock(b *testing.B) {
	testBlock()
}
