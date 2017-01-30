# go-expect

When you use go-expect, you write assertions similarly to how you would say them, e.g. "I expect this value to be equal to 3" or "I expect this array to contain 3". When you write assertions in this way, you don't need to remember the order of actual and expected arguments to functions like assert.equal, which helps you write better tests.

### toExist

> `expect.TheValue(someValue).toExist()(t)`

Asserts the given `value` is not nil.

```js
expect.TheValue('something truthy').toExist()(t)
```

### toNotExist

> `expect.TheValue(someValue).toNotExist()(t)`

Asserts the given `value` is nil.

```js
expect.TheValue(nil).toNotExist()(t)
```

### toBe

> `expect.TheValue(someValue).toBe(expectedValue)(t)`

Asserts that `someValue` is strictly equal to `expectedValue`.

### toNotBe

> `expect(someValue).toNotBe(expectedValue)(t)`

Asserts that `someValue` is not strictly equal to `expectedValue`.


### toBeLessThan

> `expect.TheValue(number).toBeLessThan(value)(t)`

Asserts the given `number` is less than `value`.

```js
expect.TheValue(2).toBeLessThan(3)(t)
```

### toBeLessThanOrEqualTo

> `expect.TheValue(number).toBeLessThanOrEqualTo(value)(t)`

Asserts the given `number` is less than or equal to `value`.

```js
expect.TheValue(2).toBeLessThanOrEqualTo(3)(t)
```

### toBeGreaterThan

> `expectTheValue(number).toBeGreaterThan(valu)(t)`

Asserts the given `number` is greater than `value`.

```js
expect.TheValue(3).toBeGreaterThan(2)(t)
```

### toBeGreaterThanOrEqualTo

> `expect.TheValue(number).toBeGreaterThanOrEqualTo(value)(t)`

Asserts the given `number` is greater than or equal to `value`.

```js
expect.TheValue(3).toBeGreaterThanOrEqualTo(2)(t)
```

![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)