# go-expect

When you use go-expect, you write assertions similarly to how you would say them, e.g. "I expect this value to be equal to 3" or "I expect this array to contain 3". When you write assertions in this way, you don't need to remember the order of actual and expected arguments to functions like assert.equal, which helps you write better tests.

## Installing:

```sh
$ go get github.com/viniciusmarson/go-expect/expect
```

## Features: 

### ToExist

> `expect.TheValue(someValue).ToExist()(t)`

Asserts the given `value` is not nil.

```js
expect.TheValue('something truthy').ToExist()(t)
```

### ToNotExist

> `expect.TheValue(someValue).ToNotExist()(t)`

Asserts the given `value` is nil.

```js
expect.TheValue(nil).ToNotExist()(t)
```

### ToBe

> `expect.TheValue(someValue).ToBe(expectedValue)(t)`

Asserts that `someValue` is strictly equal to `expectedValue`.

```js
expect.TheValue([]int{ 1, 2, 3, 4 }).toBe([]int{ 1, 2, 3, 4 })(t)
```


### ToNotBe

> `expect(someValue).ToNotBe(expectedValue)(t)`

Asserts that `someValue` is not strictly equal to `expectedValue`.

```js
expect.TheValue([]int{ 1, 2, 3, 4 }).toBe([]int{ 4, 3, 2, 1 })(t)
```


### ToBeLessThan

> `expect.TheValue(number).ToBeLessThan(value)(t)`

Asserts the given `number` is less than `value`.

```js
expect.TheValue(2).ToBeLessThan(3)(t)
```


### ToBeLessThanOrEqualTo

> `expect.TheValue(number).ToBeLessThanOrEqualTo(value)(t)`

Asserts the given `number` is less than or equal to `value`.

```js
expect.TheValue(2).ToBeLessThanOrEqualTo(3)(t)
```


### ToBeGreaterThan

> `expectTheValue(number).ToBeGreaterThan(valu)(t)`

Asserts the given `number` is greater than `value`.

```js
expect.TheValue(3).ToBeGreaterThan(2)(t)
```


### ToBeGreaterThanOrEqualTo

> `expect.TheValue(number).ToBeGreaterThanOrEqualTo(value)(t)`

Asserts the given `number` is greater than or equal to `value`.

```js
expect.TheValue(3).ToBeGreaterThanOrEqualTo(2)(t)
```


### ToBeGreaterThanOrEqualTo

> `expect.TheValue(number).ToBeGreaterThanOrEqualTo(value)(t)`

Asserts the given `number` is greater than or equal to `value`.

```js
expect.TheValue(3).ToBeGreaterThanOrEqualTo(2)(t)
```


### Contains

> `expect.TheValue(slice).Contains(value)(t)`

Asserts the given `slice` contains the `value`.

```js
expect.TheValue([]int{ 10, 9 , 8 }).Contains(9)(t)
```

## Example:

```js
func TestContains(t *testing.T) {
	TheValue([]int{10,9,8}).Contains(10)(t)
}
```


![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)
