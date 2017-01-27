# go-expect

When you use go-expect, you write assertions similarly to how you would say them, e.g. "I expect this value to be equal to 3" or "I expect this array to contain 3". When you write assertions in this way, you don't need to remember the order of actual and expected arguments to functions like assert.equal, which helps you write better tests.

### toExist

> `expect.TheValue(someValue).toExist()`

Asserts the given `value` is not nil.

```js
expect.TheValue('something truthy').toExist()
```

### toNotExist

> `expect.TheValue(someValue).toNotExist()`

Asserts the given `value` is nil.

```js
expect.TheValue(nil).toNotExist()
```

### toBe

> `expect.TheValue(someValue).toBe(expectedValue)`

Asserts that `someValue` is strictly equal to `expectedValue`.

### toNotBe

> `expect(someValue).toNotBe(expectedValue)`

Asserts that `someValue` is not strictly equal to `expectedValue`.


### toBeLessThan

> `expect.TheValue(number).toBeLessThan(value)`

Asserts the given `number` is less than `value`.

```js
expect.TheValue(2).toBeLessThan(3)
```

### toBeLessThanOrEqualTo

> `expect.TheValue(number).toBeLessThanOrEqualTo(value)`

Asserts the given `number` is less than or equal to `value`.

```js
expect.TheValue(2).toBeLessThanOrEqualTo(3)
```

### toBeGreaterThan

> `expectTheValue(number).toBeGreaterThan(valu)`

Asserts the given `number` is greater than `value`.

```js
expect.TheValue(3).toBeGreaterThan(2)
```

### toBeGreaterThanOrEqualTo

> `expect.TheValue(number).toBeGreaterThanOrEqualTo(value)`

Asserts the given `number` is greater than or equal to `value`.

```js
expect.TheValue(3).toBeGreaterThanOrEqualTo(2)
```

![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)