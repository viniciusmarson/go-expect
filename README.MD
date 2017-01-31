# go-expect

When you use go-expect, you write assertions similarly to how you would say them, e.g. "I expect this value to be equal to 3" or "I expect this array to contain 3". When you write assertions in this way, you don't need to remember the order of actual and expected arguments to functions like assert.equal, which helps you write better tests.

## Installing:

```sh
$ go get github.com/viniciusmarson/go-expect/expect
```

## Features: 

### ToExist

> `expect.This(value).ToExist()(t)`

Asserts the given `value` is not nil.

```js
expect.This('something truthy').ToExist()(t)
```

### ToNotExist

> `expect.This(value).ToNotExist()(t)`

Asserts the given `value` is nil.

```js
expect.This(nil).ToNotExist()(t)
```

### ToBe

> `expect.This(value).ToBe(expectedValue)(t)`

Asserts that `value` is strictly equal to `expectedValue`.

```js
expect.This([]int{ 1, 2, 3, 4 }).toBe([]int{ 1, 2, 3, 4 })(t)
```


### ToNotBe

> `expect(value).ToNotBe(expectedValue)(t)`

Asserts that `value` is not strictly equal to `expectedValue`.

```js
expect.This([]int{ 1, 2, 3, 4 }).toBe([]int{ 4, 3, 2, 1 })(t)
```


### ToBeLessThan

> `expect.This(number).ToBeLessThan(value)(t)`

Asserts the given `number` is less than `value`.

```js
expect.This(2).ToBeLessThan(3)(t)
```


### ToBeLessThanOrEqualTo

> `expect.This(number).ToBeLessThanOrEqualTo(value)(t)`

Asserts the given `number` is less than or equal to `value`.

```js
expect.This(2).ToBeLessThanOrEqualTo(3)(t)
```


### ToBeGreaterThan

> `expectThis(number).ToBeGreaterThan(valu)(t)`

Asserts the given `number` is greater than `value`.

```js
expect.This(3).ToBeGreaterThan(2)(t)
```


### ToBeGreaterThanOrEqualTo

> `expect.This(number).ToBeGreaterThanOrEqualTo(value)(t)`

Asserts the given `number` is greater than or equal to `value`.

```js
expect.This(3).ToBeGreaterThanOrEqualTo(2)(t)
```


### ToBeGreaterThanOrEqualTo

> `expect.This(number).ToBeGreaterThanOrEqualTo(value)(t)`

Asserts the given `number` is greater than or equal to `value`.

```js
expect.This(3).ToBeGreaterThanOrEqualTo(2)(t)
```


### ToInclude

> `expect.This(slice).ToInclude(value)(t)`

Asserts the given `slice` contains the `value`.

```js
expect.This([]int{ 10, 9 , 8 }).ToInclude(9)(t)
```

## Example:

```js
package main

import (
  "testing"
  "github.com/viniciusmarson/go-expect/expect"
)

func TestToInclude(t *testing.T) {
   expect.This([]int{10,9,8}).ToInclude(10)(t)
}
```


![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)
