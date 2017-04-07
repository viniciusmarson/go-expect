# go-expect

[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/derekparker/delve/master/LICENSE) [![GoDoc](https://godoc.org/github.com/derekparker/delve?status.svg)](https://godoc.org/github.com/derekparker/delve) [![Build Status](https://travis-ci.org/derekparker/delve.svg?branch=travis-ci)](https://travis-ci.org/derekparker/delve)


When you use go-expect, you write assertions similarly to how you would say them, e.g. "I expect this value to be equal to 3" or "I expect this array to contain 3". When you write assertions in this way, you don't need to remember the order of actual and expected arguments to functions like assert.equal, which helps you write better tests.



## Installing:

```sh
$ go get github.com/viniciusmarson/go-expect/expect
```

## Features: 

### ToExist

> `expect.This(value).ToExist()(t)`

Asserts the given `value` is not nil.

```go
expect.This('something truthy').ToExist()(t)
```

### ToNotExist

> `expect.This(value).ToNotExist()(t)`

Asserts the given `value` is nil.

```go
expect.This(nil).ToNotExist()(t)
```

### ToBe

> `expect.This(value).ToBe(expectedValue)(t)`

Asserts that `value` is strictly equal to `expectedValue`.

```go
expect.This([]int{ 1, 2, 3, 4 }).ToBe([]int{ 1, 2, 3, 4 })(t)
```


### ToNotBe

> `expect.This(value).ToNotBe(expectedValue)(t)`

Asserts that `value` is not strictly equal to `expectedValue`.

```go
expect.This([]int{ 1, 2, 3, 4 }).ToNotBe([]int{ 4, 3, 2, 1 })(t)
```


### ToBeLessThan

> `expect.This(number).ToBeLessThan(value)(t)`

Asserts the given `number` is less than `value`.

```go
expect.This(2).ToBeLessThan(3)(t)
```


### ToBeLessThanOrEqualTo

> `expect.This(number).ToBeLessThanOrEqualTo(value)(t)`

Asserts the given `number` is less than or equal to `value`.

```go
expect.This(2).ToBeLessThanOrEqualTo(3)(t)
```


### ToBeGreaterThan

> `expect.This(number).ToBeGreaterThan(valu)(t)`

Asserts the given `number` is greater than `value`.

```go
expect.This(3).ToBeGreaterThan(2)(t)
```


### ToBeGreaterThanOrEqualTo

> `expect.This(number).ToBeGreaterThanOrEqualTo(value)(t)`

Asserts the given `number` is greater than or equal to `value`.

```go
expect.This(3).ToBeGreaterThanOrEqualTo(2)(t)
```


### ToBeGreaterThanOrEqualTo

> `expect.This(number).ToBeGreaterThanOrEqualTo(value)(t)`

Asserts the given `number` is greater than or equal to `value`.

```go
expect.This(3).ToBeGreaterThanOrEqualTo(2)(t)
```


### ToInclude

> `expect.This(slice).ToInclude(value)(t)`

Asserts the given `slice` contains the `value`.

```go
expect.This([]int{ 10, 9 , 8 }).ToInclude(9)(t)
```

## Example:

```go
package main

import (
  "testing"
  "github.com/viniciusmarson/go-expect/expect"
)

func TestToInclude(t *testing.T) {
   expect.This([]int{10,9,8}).ToInclude(10)(t)
}

func TestBubbleSort(t *testing.T) {
	theExpectedResponse := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	response := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(response)
	expect.This(response).ToBe(theExpectedResponse)(t)
}
```


![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)
