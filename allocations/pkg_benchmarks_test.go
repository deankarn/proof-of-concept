package allocations_test

import (
	"testing"

	"github.com/joeybloggs/proof-of-concept/allocations"
)

// func BenchmarkReflectString(b *testing.B) {

// 	s := "test"

// 	for n := 0; n < b.N; n++ {
// 		allocations.TestReflectAllocations(s)
// 	}
// }

// func BenchmarkReflectFloat64(b *testing.B) {

// 	f := 1.123

// 	for n := 0; n < b.N; n++ {
// 		allocations.TestReflectAllocations(f)
// 	}
// }

// func BenchmarkReflectStringPtr(b *testing.B) {

// 	var s *string
// 	tmp := "test"
// 	s = &tmp

// 	for n := 0; n < b.N; n++ {
// 		allocations.TestReflectAllocations(s)
// 	}
// }

// func BenchmarkReflectFloat64Ptr(b *testing.B) {

// 	var f *float64
// 	tmp := 1.123
// 	f = &tmp

// 	for n := 0; n < b.N; n++ {
// 		allocations.TestReflectAllocations(f)
// 	}
// }

func BenchmarkReflectStructPtr(b *testing.B) {

	type Test struct {
		String  string
		Float64 float64
	}

	t := &Test{
		String:  "test",
		Float64: 1.123,
	}

	// t := new(Test)
	// t.String = "test"
	// t.Float64 = 1.123

	for n := 0; n < b.N; n++ {
		allocations.TestReflectStructAllocations(t)
	}
}

func BenchmarkStructComplexFailureParallel(b *testing.B) {

	type Test struct {
		String  string
		Float64 float64
	}

	t := &Test{
		String:  "test",
		Float64: 1.123,
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			allocations.TestReflectStructAllocations(t)
		}
	})
}

// func BenchmarkBasicString(b *testing.B) {

// 	s := "test"

// 	for n := 0; n < b.N; n++ {
// 		testBasicAllocations(s)
// 	}
// }

// func BenchmarkBasicFloat64(b *testing.B) {

// 	f := 1.123

// 	for n := 0; n < b.N; n++ {
// 		testBasicAllocations(f)
// 	}
// }

// func BenchmarkHybridString(b *testing.B) {

// 	s := "test"

// 	for n := 0; n < b.N; n++ {
// 		testHybridAllocations(s)
// 	}
// }

// func BenchmarkHybridFloat64(b *testing.B) {

// 	f := 1.123

// 	for n := 0; n < b.N; n++ {
// 		testHybridAllocations(f)
// 	}
// }

// func BenchmarkHybridStringPtr(b *testing.B) {

// 	var s *string
// 	tmp := "test"
// 	s = &tmp

// 	for n := 0; n < b.N; n++ {
// 		testHybridAllocations(s)
// 	}
// }

// func BenchmarkHybridFloat64Ptr(b *testing.B) {

// 	var f *float64
// 	tmp := 1.123
// 	f = &tmp

// 	for n := 0; n < b.N; n++ {
// 		testHybridAllocations(f)
// 	}
// }
