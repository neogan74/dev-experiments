# [1115. Print FooBar Alternately](https://leetcode.com/problems/print-foobar-alternately/description/)

You are given the following skeleton implementation:

```java
class FooBar {
    public void foo() {
        for (int i = 0; i < n; i++) {
            print("foo");
        }
    }

    public void bar() {
        for (int i = 0; i < n; i++) {
            print("bar");
        }
    }
}
```

The same `FooBar` instance is shared by two threads:

- Thread A calls `foo()`
- Thread B calls `bar()`

Modify the program so that the combined output is `"foobar"` repeated `n` times.

## Examples

**Example 1**

- Input: `n = 1`
- Output: `"foobar"`
- Explanation: Two threads execute `foo()` and `bar()` asynchronously, producing one `"foobar"`.

**Example 2**

- Input: `n = 2`
- Output: `"foobarfoobar"`
- Explanation: The alternating output repeats twice.

## Constraints

- `1 <= n <= 1000`
