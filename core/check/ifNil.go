package check

import (
	"fmt"
	"reflect"
)

// IfNil tests if the provided interface pointer or underlying object is nil
func IfNil(checker NilInterfaceChecker) bool {
	if checker == nil {
		return true
	}
	return checker.IsInterfaceNil()
}

// IfNilReflect tests if the provided interface pointer or underlying pointer receiver is nil
func IfNilReflect(i interface{}) bool {
	if v := reflect.ValueOf(i); v.IsValid() {
		if v.Kind() == reflect.Ptr && v.IsNil() {
			return true
		}
		return false
	}
	return true
}

// AssertNotNil throws a programmer error (go panic) if the object is nil
func AssertNotNil(object NilInterfaceChecker, objectName string) {
	if object == nil || object.IsInterfaceNil() {
		panic(fmt.Sprintf("ProgrammerError: %s is NIL", objectName))
	}
}
