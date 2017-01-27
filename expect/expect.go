package expect

import "testing"

//Info store the value and the expect methods
type Info struct {
	value interface{}
}

//TheValue recive the data to use for assurance
func TheValue(i interface{}) Info {
	return Info{value: i}
}

//ToBe asserts that object is strictly equal to value.
func (i Info) ToBe(theExpectedValue interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		if i.value != theExpectedValue {
			t.Errorf("Expected the value %v to be %v", i.value, theExpectedValue)
		}
	}
}

//ToNotBe asserts that object is not strictly equal to value.
func (i Info) ToNotBe(theNotExpectedValue interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		if i.value == theNotExpectedValue {
			t.Errorf("Expected the value %v to not be %v", i.value, theNotExpectedValue)
		}
	}
}

//ToBeLessThanOrEqualTo asserts the given number is less than or equal to value.
func (i Info) ToBeLessThanOrEqualTo(theNotExpectedValue int) func(t *testing.T) {
	return func(t *testing.T) {
		value, ok := i.value.(int)
		if ok {
			if value > theNotExpectedValue {
				t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
			}
		} else {
			t.Errorf("The parameter %v is not of type int", i.value)
		}
	}
}

//ToBeLessThan asserts the given number is less than the value.
func (i Info) ToBeLessThan(theNotExpectedValue int) func(t *testing.T) {
	return func(t *testing.T) {
		value, ok := i.value.(int)
		if ok {
			if value >= theNotExpectedValue {
				t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
			}
		} else {
			t.Errorf("The parameter %v is not of type int", i.value)
		}
	}
}

//ToBeGreaterThanOrEqualTo asserts the given number is greater than or equal to value.
func (i Info) ToBeGreaterThanOrEqualTo(theNotExpectedValue int) func(t *testing.T) {
	return func(t *testing.T) {
		value, ok := i.value.(int)
		if ok {
			if value < theNotExpectedValue {
				t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
			}
		} else {
			t.Errorf("The parameter %v is not of type int", i.value)
		}
	}
}

//ToBeGreaterThan asserts the given number is greater than the value.
func (i Info) ToBeGreaterThan(theNotExpectedValue int) func(t *testing.T) {
	return func(t *testing.T) {
		value, ok := i.value.(int)
		if ok {
			if value <= theNotExpectedValue {
				t.Errorf("Expected the value %v to less than or equal to %v", i.value, theNotExpectedValue)
			}
		} else {
			t.Errorf("The parameter %v is not of type int", i.value)
		}
	}
}

//ToExist asserts the given object is not nil.
func (i Info) ToExist() func(t *testing.T) {
	return func(t *testing.T) {
		if i.value == nil {
			t.Errorf("Expected to exist but got %v", i.value)
		}
	}
}

//ToNotExist asserts the given object is nil.
func (i Info) ToNotExist() func(t *testing.T) {
	return func(t *testing.T) {
		if i.value != nil {
			t.Errorf("Expected to not exist but got %v", i.value)
		}
	}
}
