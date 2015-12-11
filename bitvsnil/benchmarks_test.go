package main

import "testing"

func BenchmarkBool(b *testing.B) {

	var ok bool

	for n := 0; n < b.N; n++ {
		if ok {
			ok = false
		}
	}
}

func BenchmarkNilErr(b *testing.B) {

	var fn test

	for n := 0; n < b.N; n++ {
		if fn != nil {
			fn = nil
		}
	}
}

func BenchmarkBoolParallel(b *testing.B) {

	var ok bool

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if ok {
				ok = false
			}
		}
	})
}

func BenchmarkNilErrParallel(b *testing.B) {

	var fn test
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if fn != nil {
				fn = nil
			}
		}
	})
}
