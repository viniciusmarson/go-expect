package expect

import (
	"testing"
)

func TestToBe(t *testing.T) {
	TheValue(10).ToBe(10)(t)
	x := []int{10, 20, 30, 4, 5, 6, 7, 8, 9, 10}
	y := []int{10, 20, 30, 4, 5, 6, 7, 8, 9, 10}
	TheValue(x).ToBe(y)(t)
}

func TestToNotBe(t *testing.T) {
	TheValue(10).ToNotBe(9)(t)
	x := []int{10, 2, 30, 4, 5, 6, 7, 8, 9, 10}
	y := []int{10, 20, 30, 4, 5, 6, 7, 8, 9, 10}
	TheValue(x).ToNotBe(y)(t)
}

func TestToBeLessThanOrEqualTo(t *testing.T) {
	TheValue(10).ToBeLessThanOrEqualTo(10)(t)
	TheValue(10).ToBeLessThanOrEqualTo(11)(t)
}

func TestToBeLessThan(t *testing.T) {
	TheValue(10).ToBeLessThan(11)(t)
}

func TestToBeGreaterThanOrEqualTo(t *testing.T) {
	TheValue(10).ToBeGreaterThanOrEqualTo(10)(t)
	TheValue(10).ToBeGreaterThanOrEqualTo(9)(t)
}

func TestToBeGreaterThan(t *testing.T) {
	TheValue(10).ToBeGreaterThan(9)(t)
}

func TestToExist(t *testing.T) {
	TheValue(10).ToExist()(t)
}

func TestToNotExist(t *testing.T) {
	TheValue(nil).ToNotExist()(t)
}
