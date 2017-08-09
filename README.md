![license](http://img.shields.io/badge/license-MIT-blue.svg)

![Golang testing](https://cdn-images-1.medium.com/max/800/1*8jnAiQdv4Vxh5AH3-2uswA.jpeg)

When you use go-expect, you write beautiful assertions as if you were writing a text. When you write assertions in this way, you don't need to remember the order of actual and expected arguments to functions like assert.equal, which helps you write better tests.


## Installing:

```sh
$ go get github.com/viniciusmarson/go-expect/expect
```

&nbsp;

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

&nbsp;

## Features: 


### ToExist

> `expect(value).ToExist()`

Asserts the given `value` is not nil.

```go
expect("something truthy").ToExist()
```

&nbsp;
### ToNotExist

> `expect(value).ToNotExist()`

Asserts the given `value` is nil.

```go
expect(nil).ToNotExist()
```

&nbsp;
### ToBe

> `expect(value).ToBe(expectedValue)`

Asserts that `value` is strictly equal to `expectedValue`.

```go
expect([]int{ 1, 2, 3, 4 }).ToBe([]int{ 1, 2, 3, 4 })

expect("test").ToBe("test")

expect(11).ToBe(11)
```

&nbsp;
### ToNotBe

> `expect(value).ToNotBe(expectedValue)`

Asserts that `value` is not strictly equal to `expectedValue`.

```go
expect([]int{ 1, 2, 3, 4 }).ToNotBe([]int{ 4, 3, 2, 1 })

expect("test").ToNotBe("tset")

expect(10).ToNotBe(11)
```

&nbsp;
### ToBeAn and ToBeA

> `expect(value).ToBeAn(expectedType)`

Asserts that `value` is of the type `expectedType`.

```go
expect("test").ToBeA("string")
expect(true).ToBeA("bool")
expect(10).ToBeAn("int")
expect([]int{}).ToBeA("slice")
expect(`some interface{}`).ToBeAn("interface")
expect(`some struct`).ToBeA("struct")
```


&nbsp;
### ToBeTrue

> `expect(bool).ToBeTrue()`

Asserts that `bool` is true.

```go
expect(true).ToBeTrue()
```

&nbsp;
### ToBeFalse

> `expect(bool).ToBeFalse()`

Asserts that `bool` is false.

```go
expect(false).ToBeFalse()
```

&nbsp;
### Contains

> `expect(string).Contains(expectedString)`

Asserts that `string` contains `expectedString`.

```go
expect("banana").Contains("nana")
```

&nbsp;
### NotContains

> `expect(string).NotContains(notExpectedString)`

Asserts that `string` not contains `notExpectedString`.

```go
expect("banana").NotContains("haha")
```

&nbsp;
### ToBeLessThan

> `expect(number).ToBeLessThan(value)`

Asserts the given `number` is less than `value`.

```go
expect(2).ToBeLessThan(3)
```

&nbsp;
### ToBeLessThanOrEqualTo

> `expect(number).ToBeLessThanOrEqualTo(value)`

Asserts the given `number` is less than or equal to `value`.

```go
expect(2).ToBeLessThanOrEqualTo(3)
```

&nbsp;
### ToBeGreaterThan

> `expect(number).ToBeGreaterThan(valu)`

Asserts the given `number` is greater than `value`.

```go
expect(3).ToBeGreaterThan(2)
```

&nbsp;
### ToBeGreaterThanOrEqualTo

> `expect(number).ToBeGreaterThanOrEqualTo(value)`

Asserts the given `number` is greater than or equal to `value`.

```go
expect(3).ToBeGreaterThanOrEqualTo(2)
```

&nbsp;
### ToInclude

> `expect(slice).ToInclude(value)`

Asserts the given `slice` contains the `value`.

```go
expect([]int{ 10, 9 , 8 }).ToInclude(9)
```

&nbsp;
### ToExclude

> `expect(slice).ToExclude(value)`

Asserts the given `slice` not contains the `value`.

```go
expect([]int{ 10, 9 , 8 }).ToExclude(2)
```


&nbsp;
### Fail

> `expect().Fail()`

Explicitly forces failure.

```go
expect().Fail()
expect().Fail("Custom message")
```


&nbsp;
## TODO 



![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)
