![license](http://img.shields.io/badge/license-MIT-blue.svg)

![Golang testing](https://cdn-images-1.medium.com/max/800/1*8jnAiQdv4Vxh5AH3-2uswA.jpeg)

When you use go-expect, you write beautiful assertions as if you were writing a text. When you write assertions in this way, you don't need to remember the order of actual and expected arguments to functions like assert.equal, which helps you write better tests.


## Installing:

```sh
$ go get github.com/viniciusmarson/go-expect/expect
```


## Example:

```go
package main

import (
  "testing"
  "github.com/viniciusmarson/go-expect/expect"
)

func TestToInclude(t *testing.T) {
   expect := expect.New(t)
   expect([]int{10,9,8}).ToInclude(10)
}

func TestBubbleSort(t *testing.T) {
	expect := expect.New(t)
	theExpectedResponse := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	response := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(response)
	expect(response).ToBe(theExpectedResponse)
}
```


## Features: 


### ToExist

> `expect(value).ToExist()`

Asserts the given `value` is not nil.

```go
expect("something truthy").ToExist()
```


### ToNotExist

> `expect(value).ToNotExist()`

Asserts the given `value` is nil.

```go
expect(nil).ToNotExist()
```


### ToBe

> `expect(value).ToBe(expectedValue)`

Asserts that `value` is strictly equal to `expectedValue`.

```go
expect([]int{ 1, 2, 3, 4 }).ToBe([]int{ 1, 2, 3, 4 })
```


### ToNotBe

> `expect(value).ToNotBe(expectedValue)`

Asserts that `value` is not strictly equal to `expectedValue`.

```go
expect([]int{ 1, 2, 3, 4 }).ToNotBe([]int{ 4, 3, 2, 1 })
```


### ToBeLessThan

> `expect(number).ToBeLessThan(value)`

Asserts the given `number` is less than `value`.

```go
expect(2).ToBeLessThan(3)
```


### ToBeLessThanOrEqualTo

> `expect(number).ToBeLessThanOrEqualTo(value)`

Asserts the given `number` is less than or equal to `value`.

```go
expect(2).ToBeLessThanOrEqualTo(3)
```


### ToBeGreaterThan

> `expect(number).ToBeGreaterThan(valu)`

Asserts the given `number` is greater than `value`.

```go
expect(3).ToBeGreaterThan(2)
```


### ToBeGreaterThanOrEqualTo

> `expect(number).ToBeGreaterThanOrEqualTo(value)`

Asserts the given `number` is greater than or equal to `value`.

```go
expect(3).ToBeGreaterThanOrEqualTo(2)
```


### ToInclude

> `expect(slice).ToInclude(value)`

Asserts the given `slice` contains the `value`.

```go
expect([]int{ 10, 9 , 8 }).ToInclude(9)
```

## TODO 

* ToExclude
* Contains
* NotContains
* ToBeTrue
* ToBeFalse


![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)
