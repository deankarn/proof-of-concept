package allocations

import (
	"reflect"
	"sync"
)

type cachedField struct {
	Idx  int
	Name string
}

type cachedStruct struct {
	Name   string
	fields map[int]cachedField
}

type structCacheMap struct {
	lock sync.RWMutex
	m    map[string]*structCache
}

func (s *structCacheMap) Get(key string) (*cachedStruct, bool) {
	s.lock.RLock()
	value, ok := s.m[key]
	s.lock.RUnlock()
	return value, ok
}

func (s *structCacheMap) Set(key string, value *cachedStruct) {
	s.lock.Lock()
	s.m[key] = value
	s.lock.Unlock()
}

// TestReflectStructAllocations ...
func TestReflectStructAllocations(input interface{}) bool {

	typ := reflect.TypeOf(input)
	// val := reflect.ValueOf(input)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	sName := typ.Name()
	s, ok := sCache.Get(sName)

	if ok {

		for range s.fields {

		}

	} else {
		s = &structCache{Name: sName, fields: map[int]fieldCache{}}
		// tp := val.Type()
		numFields := typ.NumField()

		var fld reflect.StructField

		for i := 0; i < numFields; i++ {
			fld = typ.Field(i)

			s.fields[i] = fieldCache{Idx: i, Name: fld.Name}
			// fmt.Println(fld.)

			// if fld.PkgPath != "" {
			// 	continue
			// }
			// fmt.Println(fld.Name)
			if len(fld.PkgPath) != 0 {
				continue
			}

			// s.fields[i] = fieldCache{Idx: i, Name: fld.Name}

			// if !unicode.IsUpper(rune(fld.Name[0])) {
			// 	continue
			// }

			// v.traverseField(topStruct, currentStruct, current.Field(i), errPrefix, errs, true, fld.Tag.Get(v.tagName), fld.Name, customName, partial, exclude, includeExclude)
		}
	}

	if !ok {
		sCache.Set(sName, s)
	}

	return true
}

// TestReflectAllocations ...
func TestReflectAllocations(input interface{}) bool {

	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {

	case reflect.String:
		s := val.String()

		return s == "test"

	case reflect.Float64:
		f := val.Float()

		return f == 1.123
	}

	return false
}

// faster at non pointer vars, but cannot handle pointers, interface etc....
func testBasicAllocations(val interface{}) bool {

	switch val.(type) {

	case string:

		s := val.(string)

		return s == "test"

	case float64:
		f := val.(float64)

		return f == 1.123
	}

	return false
}

// no gain over testReflectAllocations
func testHybridAllocations(input interface{}) bool {

	// switch input.(type) {
	// case string:
	// 	// s := input.(string)
	// 	return testHybridReflectAllocations(input.(string))

	// case float64:
	// 	// f := input.(float64)
	// 	return testHybridReflectAllocations(input.(float64))
	// }

	// reflect.ValueOf(input)

	// // if val.Kind() == reflect.Ptr {
	return testHybridReflectAllocations(input)
	// }

	// return true

	// return testBasicValue(input)
}

func testHybridReflectAllocations(input interface{}) bool {

	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {

	case reflect.String:
		s := val.String()

		return s == "test"

	case reflect.Float64:
		f := val.Float()

		return f == 1.123
	}

	return false
}

func testBasicValue(val interface{}) bool {
	switch val.(type) {

	case string:
		s := val.(string)

		return s == "test"

	case float64:
		f := val.(float64)

		return f == 1.123
	}

	return false
}
