package expect

import (
	"testing"
)

func TestToBe(t *testing.T) {
	expect := New(t)
	expect(10).ToBe(10)
	x := []int{10, 20, 30, 4, 5, 6, 7, 8, 9, 10}
	y := []int{10, 20, 30, 4, 5, 6, 7, 8, 9, 10}
	expect(x).ToBe(y)
}

func TestToNotBe(t *testing.T) {
	expect := New(t)
	expect(10).ToNotBe(9)
	x := []int{10, 2, 30, 4, 5, 6, 7, 8, 9, 10}
	y := []int{10, 20, 30, 4, 5, 6, 7, 8, 9, 10}
	expect(x).ToNotBe(y)
}

func TestToBeTrue(t *testing.T) {
	expect := New(t)
	expect(true).ToBeTrue()
}

func TestToBeFalse(t *testing.T) {
	expect := New(t)
	expect(false).ToBeFalse()
}

func TestContains(t *testing.T) {
	expect := New(t)
	expect("banana").Contains("nana")
}

func TestNotContains(t *testing.T) {
	expect := New(t)
	expect("banana").NotContains("haha")
}

func TestToBeLessThanOrEqualTo(t *testing.T) {
	expect := New(t)
	expect(10).ToBeLessThanOrEqualTo(10)
	expect(10).ToBeLessThanOrEqualTo(11)
}

func TestToBeLessThan(t *testing.T) {
	expect := New(t)
	expect(10).ToBeLessThan(11)
}

func TestToBeGreaterThanOrEqualTo(t *testing.T) {
	expect := New(t)
	expect(10).ToBeGreaterThanOrEqualTo(10)
	expect(10).ToBeGreaterThanOrEqualTo(9)
}

func TestToBeGreaterThan(t *testing.T) {
	expect := New(t)
	expect(10).ToBeGreaterThan(9)
}

func TestToExist(t *testing.T) {
	expect := New(t)
	expect(10).ToExist()
}

func TestToNotExist(t *testing.T) {
	expect := New(t)
	expect(nil).ToNotExist()
}

func TestToInclude(t *testing.T) {
	expect := New(t)
	expect([]int{10, 9, 8}).ToInclude(10)
}

func TestToExclude(t *testing.T) {
	expect := New(t)
	expect([]int{10, 9, 8}).ToExclude(2)
}

func TestToBeAn(t *testing.T) {
	expect := New(t)
	expect("test").ToBeAn("string")
}
