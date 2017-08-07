//Package expect lets you write better assertions
package expect

import (
	"reflect"
	"testing"
)

var (
	t *testing.T
)

//Info store the value and the expect methods
type Info struct {
	value interface{}
}

//New will initialize
func New(teste *testing.T) func(i interface{}) Info {
	t = teste
	return valueToAssert
}

//valueToAssert recive the data to use for assurance
func valueToAssert(i interface{}) Info {
	return Info{value: i}
}

//ToBe asserts that object is strictly equal to value.
func (i Info) ToBe(theExpectedValue interface{}) {
	valueType := reflect.TypeOf(i.value).Kind()
	theExpectedValueType := reflect.TypeOf(theExpectedValue).Kind()

	if valueType != theExpectedValueType {
		t.Errorf("The expected value is a %v and the value is a %v", theExpectedValueType, valueType)
	} else {
		switch valueType {
		case reflect.Slice:
			s := reflect.ValueOf(theExpectedValue)
			value := reflect.ValueOf(i.value)

			for j := 0; j < s.Len(); j++ {
				if s.Index(j).Interface() != value.Index(j).Interface() {
					t.Errorf("Expected the value %v to be %v", i.value, theExpectedValue)
					break
				}
			}
		default:
			if i.value != theExpectedValue {
				t.Errorf("Expected the value %v to be %v", i.value, theExpectedValue)
			}
		}
	}
}

//ToNotBe asserts that object is not strictly equal to value.
func (i Info) ToNotBe(theNotExpectedValue interface{}) {
	valueType := reflect.TypeOf(i.value).Kind()
	theNotExpectedValueType := reflect.TypeOf(theNotExpectedValue).Kind()

	if valueType != theNotExpectedValueType {
		t.Errorf("The not expected value is a %v and the value is a %v", theNotExpectedValueType, valueType)
	} else {
		switch valueType {
		case reflect.Slice:
			s := reflect.ValueOf(theNotExpectedValue)
			value := reflect.ValueOf(i.value)

			hasSomeValueDifferent := false
			for j := 0; j < s.Len(); j++ {
				if s.Index(j).Interface() != value.Index(j).Interface() {
					hasSomeValueDifferent = true
					break
				}
			}

			if !hasSomeValueDifferent {
				t.Errorf("Expected the value %v to not be %v", i.value, theNotExpectedValue)
			}
		default:
			if i.value == theNotExpectedValue {
				t.Errorf("Expected the value %v to not be %v", i.value, theNotExpectedValue)
			}
		}
	}
}

//ToBeLessThanOrEqualTo asserts the given number is less than or equal to value.
func (i Info) ToBeLessThanOrEqualTo(theNotExpectedValue int) {
	value, ok := i.value.(int)
	if ok {
		if value > theNotExpectedValue {
			t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
		}
	} else {
		t.Errorf("The parameter %v is not of type int", i.value)
	}
}

//ToBeLessThan asserts the given number is less than the value.
func (i Info) ToBeLessThan(theNotExpectedValue int) {
	value, ok := i.value.(int)
	if ok {
		if value >= theNotExpectedValue {
			t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
		}
	} else {
		t.Errorf("The parameter %v is not of type int", i.value)
	}
}

//ToBeGreaterThanOrEqualTo asserts the given number is greater than or equal to value.
func (i Info) ToBeGreaterThanOrEqualTo(theNotExpectedValue int) {
	value, ok := i.value.(int)
	if ok {
		if value < theNotExpectedValue {
			t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
		}
	} else {
		t.Errorf("The parameter %v is not of type int", i.value)
	}
}

//ToBeGreaterThan asserts the given number is greater than the value.
func (i Info) ToBeGreaterThan(theNotExpectedValue int) {
	value, ok := i.value.(int)
	if ok {
		if value <= theNotExpectedValue {
			t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
		}
	} else {
		t.Errorf("The parameter %v is not of type int", i.value)
	}
}

//ToExist asserts the given object is not nil.
func (i Info) ToExist() {
	if i.value == nil {
		t.Errorf("Expected to exist but got %v", i.value)
	}
}

//ToNotExist asserts the given object is nil.
func (i Info) ToNotExist() {
	if i.value != nil {
		t.Errorf("Expected to not exist but got %v", i.value)
	}
}

//ToInclude asserts the given array contains the informed value
func (i Info) ToInclude(theExpectedValue interface{}) {
	valueType := reflect.TypeOf(i.value).Kind()
	theExpectedValueType := reflect.TypeOf(theExpectedValue).Kind()

	value := reflect.ValueOf(i.value)

	if valueType != reflect.Slice {
		t.Errorf("The value infromed is of type %v", valueType)
	} else if value.Len() == 0 {
		t.Errorf("The value informed is an empty slice")
	} else {

		sliceType := reflect.TypeOf(value.Index(0).Interface()).Kind()

		if sliceType != theExpectedValueType {
			t.Errorf("The expected value informed is of type %v and the slice is of type %v", theExpectedValueType, sliceType)
		} else {

			found := false

			for j := 0; j < value.Len(); j++ {
				if value.Index(j).Interface() == theExpectedValue {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Informed value %v not found in slice %v", theExpectedValue, i.value)
			}
		}
	}
}