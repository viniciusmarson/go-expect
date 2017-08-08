//Package expect lets you write better assertions
package expect

import (
	"reflect"
	"strings"
	"testing"
)

var (
	t *testing.T
)

type info struct {
	value interface{}
}

func valueToAssert(i interface{}) info {
	return info{value: i}
}

//New will initialize the expect function
func New(teste *testing.T) func(i interface{}) info {
	t = teste
	return valueToAssert
}

//ToBeTrue asserts that the value is true
func (i info) ToBeTrue() {
	valueType := reflect.TypeOf(i.value).Kind()
	if valueType != reflect.Bool {
		t.Errorf("The value is a %v and not a boolean", valueType)
		return
	}

	if i.value.(bool) != true {
		t.Errorf("Expected to be true but get false")
	}
}

//ToBeFalse asserts that the value is false
func (i info) ToBeFalse() {
	valueType := reflect.TypeOf(i.value).Kind()
	if valueType != reflect.Bool {
		t.Errorf("The value is a %v and not a boolean", valueType)
		return
	}

	if i.value.(bool) != false {
		t.Errorf("Expected to be false but get true")
	}
}

//ToBe asserts that object is strictly equal to the informed value.
func (i info) ToBe(theExpectedValue interface{}) {
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

//ToNotBe asserts that object is not strictly equal to the informed value.
func (i info) ToNotBe(theNotExpectedValue interface{}) {
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

//Contains will check if the string contains some informed value
func (i info) Contains(contains interface{}) {
	valueType := reflect.TypeOf(i.value).Kind()
	containsType := reflect.TypeOf(contains).Kind()

	if valueType != reflect.String || containsType != reflect.String {
		t.Errorf("The value is a %v and the contains value is a %v. Must be a string", valueType, containsType)
		return
	}

	value := i.value.(string)
	containsValue := contains.(string)

	if !strings.Contains(value, containsValue) {
		t.Errorf("%s not contains %s", value, containsValue)
	}
}

//NotContains will check if the string not contains some informed value
func (i info) NotContains(contains interface{}) {
	valueType := reflect.TypeOf(i.value).Kind()
	containsType := reflect.TypeOf(contains).Kind()

	if valueType != reflect.String || containsType != reflect.String {
		t.Errorf("The value is a %v and the contains value is a %v. Must be a string", valueType, containsType)
		return
	}

	value := i.value.(string)
	containsValue := contains.(string)

	if strings.Contains(value, containsValue) {
		t.Errorf("%s contains %s", value, containsValue)
	}
}

//ToBeLessThanOrEqualTo asserts the given number is less than or equal to the informed value.
func (i info) ToBeLessThanOrEqualTo(theNotExpectedValue int) {
	value, ok := i.value.(int)
	if ok {
		if value > theNotExpectedValue {
			t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
		}
	} else {
		t.Errorf("The parameter %v is not of type int", i.value)
	}
}

//ToBeLessThan asserts the given number is less than the informed value.
func (i info) ToBeLessThan(theNotExpectedValue int) {
	value, ok := i.value.(int)
	if ok {
		if value >= theNotExpectedValue {
			t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
		}
	} else {
		t.Errorf("The parameter %v is not of type int", i.value)
	}
}

//ToBeGreaterThanOrEqualTo asserts the given number is greater than or equal to the informed value.
func (i info) ToBeGreaterThanOrEqualTo(theNotExpectedValue int) {
	value, ok := i.value.(int)
	if ok {
		if value < theNotExpectedValue {
			t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
		}
	} else {
		t.Errorf("The parameter %v is not of type int", i.value)
	}
}

//ToBeGreaterThan asserts the given number is greater than the informed value.
func (i info) ToBeGreaterThan(theNotExpectedValue int) {
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
func (i info) ToExist() {
	if i.value == nil {
		t.Errorf("Expected to exist but got %v", i.value)
	}
}

//ToNotExist asserts the given object is nil.
func (i info) ToNotExist() {
	if i.value != nil {
		t.Errorf("Expected to not exist but got %v", i.value)
	}
}

//ToInclude asserts the given array contains the informed value
func (i info) ToInclude(theExpectedValue interface{}) {
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

//ToExclude asserts the given array not contains the informed value
func (i info) ToExclude(theExpectedValue interface{}) {
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

			if found {
				t.Errorf("Informed value %v found in slice %v", theExpectedValue, i.value)
			}
		}
	}
}

//Fail force fails
func (i info) Fail(message string) {
	if message != "" {
		t.Errorf(message)
	} else {
		t.Errorf("Explicitly forces failure")
	}
}

//ToBeAn check if the informed value is of some type
func (i info) ToBeAn(expectedType string) {
	valueType := reflect.TypeOf(i.value).Kind()
	checkType(valueType, expectedType)
}

//ToBeA check if the informed value is of some type
func (i info) ToBeA(expectedType string) {
	valueType := reflect.TypeOf(i.value).Kind()
	checkType(valueType, expectedType)
}

func checkType(valueType reflect.Kind, expectedType string) {
	var error = false

	switch expectedType {
	case "slice":
		if valueType != reflect.Slice {
			error = true
		}
	case "string":
		if valueType != reflect.String {
			error = true
		}
	case "int":
		if valueType != reflect.Int {
			error = true
		}
	case "bool":
		if valueType != reflect.Bool {
			error = true
		}
	case "struct":
		if valueType != reflect.Struct {
			error = true
		}
	case "interface":
		if valueType != reflect.Interface {
			error = true
		}
	default:
		t.Errorf("Expected type %s not exist", expectedType)
	}

	if error {
		t.Errorf("Expected the type %s but the type is %v", expectedType, valueType)
	}
}
