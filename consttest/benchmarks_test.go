package main

import "testing"

func stringTest(path string) {

}

func BenchmarkString(b *testing.B) {

	for n := 0; n < b.N; n++ {
		stringTest("/")
	}
}

func BenchmarkConstString(b *testing.B) {

	for n := 0; n < b.N; n++ {
		stringTest(path)
	}
}

func BenchmarkStringParallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			stringTest("/")
		}
	})
}

func BenchmarkConstStringParallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			stringTest(path)
		}
	})
}

func boolTest(ok bool) {

}

func BenchmarkBool(b *testing.B) {

	for n := 0; n < b.N; n++ {
		boolTest(true)
	}
}

func BenchmarkConstBool(b *testing.B) {

	for n := 0; n < b.N; n++ {
		boolTest(ok)
	}
}

func BenchmarkBoolParallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			boolTest(true)
		}
	})
}

func BenchmarkConstBoolParallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			boolTest(ok)
		}
	})
}
