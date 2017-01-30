package test

import (
	"testing"

	"../expect"
)

func TestToBe(t *testing.T) {
	expect.TheValue(10).ToBe(10)(t)
	array := []int{1, 2, 3, 4}
	expect.TheValue(array).ToBe(array)(t)
}

func TestToNotBe(t *testing.T) {
	expect.TheValue(10).ToNotBe(9)(t)
}

func TestToBeLessThanOrEqualTo(t *testing.T) {
	expect.TheValue(10).ToBeLessThanOrEqualTo(10)(t)
	expect.TheValue(10).ToBeLessThanOrEqualTo(11)(t)
}

func TestToBeLessThan(t *testing.T) {
	expect.TheValue(10).ToBeLessThan(11)(t)
}

func TestToBeGreaterThanOrEqualTo(t *testing.T) {
	expect.TheValue(10).ToBeGreaterThanOrEqualTo(10)(t)
	expect.TheValue(10).ToBeGreaterThanOrEqualTo(9)(t)
}

func TestToBeGreaterThan(t *testing.T) {
	expect.TheValue(10).ToBeGreaterThan(9)(t)
}

func TestToExist(t *testing.T) {
	expect.TheValue(10).ToExist()(t)
}

func TestToNotExist(t *testing.T) {
	expect.TheValue(nil).ToNotExist()(t)
}
