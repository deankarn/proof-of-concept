package allocations

import (
	"testing"

	. "gopkg.in/go-playground/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//
//
// go test -cpuprofile cpu.out
// ./validator.test -test.bench=. -test.cpuprofile=cpu.prof
// go tool pprof validator.test cpu.prof
//
//
// go test -memprofile mem.out

func TestBasicFloat64(t *testing.T) {
	f := 1.123
	ok := testBasicAllocations(f)
	Equal(t, ok, true)
}

func TestBasicString(t *testing.T) {
	s := "test"
	ok := testBasicAllocations(s)
	Equal(t, ok, true)
}

func TestReflectFloat64(t *testing.T) {
	f := 1.123
	ok := TestReflectAllocations(f)
	Equal(t, ok, true)
}

func TestReflectString(t *testing.T) {
	s := "test"
	ok := TestReflectAllocations(s)
	Equal(t, ok, true)
}

func TestReflectFloat64Ptr(t *testing.T) {
	var f *float64
	tmp := 1.123
	f = &tmp
	ok := TestReflectAllocations(f)
	Equal(t, ok, true)
}

func TestReflectStringPtr(t *testing.T) {
	var s *string
	tmp := "test"
	s = &tmp
	ok := TestReflectAllocations(s)
	Equal(t, ok, true)
}

func TestHybridFloat64(t *testing.T) {
	f := 1.123
	ok := testHybridAllocations(f)
	Equal(t, ok, true)
}

func TestHybridString(t *testing.T) {
	s := "test"
	ok := testHybridAllocations(s)
	Equal(t, ok, true)
}

func TestHybridFloat64Ptr(t *testing.T) {
	var f *float64
	tmp := 1.123
	f = &tmp
	ok := testHybridAllocations(f)
	Equal(t, ok, true)
}

func TestHybridStringPtr(t *testing.T) {
	var s *string
	tmp := "test"
	s = &tmp
	ok := testHybridAllocations(s)
	Equal(t, ok, true)
}

func TestHybridStructPtr(t *testing.T) {
	type Test struct {
		String  string
		Float64 float64
	}

	s := &Test{
		String:  "test",
		Float64: 1.123,
	}

	ok := TestReflectStructAllocations(s)
	Equal(t, ok, true)
}
