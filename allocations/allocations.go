package allocations

import "reflect"

func testReflectAllocations(input interface{}) bool {

	val := reflect.ValueOf(input)

	switch val.Kind() {

	case reflect.String:
		s := val.String()

		return s == "test"

	case reflect.Float64:
		f := val.Float()

		return f == 1.123
	}

	return true
}

// faster at non pointer vars, but cannot handle pointers, interface etc....
func testBasicAllocations(val interface{}) bool {

	// val := reflect.ValueOf(input)

	// fmt.Println(val.(type))

	switch val.(type) {

	case string:
		s := val.(string)

		return s == "test"

	case float64:
		f := val.(float64)

		return f == 1.123
	}

	return true
}

// no gain over testReflectAllocations
func testHybridAllocations(val interface{}) (bool, bool) {

	// val := reflect.ValueOf(input)

	// fmt.Println(val.(type))

	switch val.(type) {

	case string:
		s := val.(string)

		return s == "test", true

	case float64:
		f := val.(float64)

		return f == 1.123, true

		// default:
		// 	testReflectAllocations(val)
		// nVal := reflect.ValueOf(val)
		// // if nVal.Kind() == reflect.Ptr {
		// // 	return testHybridAllocations(nVal.Elem().Interface())
		// // }

		// switch nVal.Kind() {

		// case reflect.String:
		// 	s := nVal.Elem().String()

		// 	return s == "test"

		// case reflect.Float64:
		// 	f := nVal.Elem().Float()

		// 	return f == 1.123
		// }
	}

	return false, false
	// fmt.Println("HERE")
	// return testReflectAllocations(val)
}
