# [1114. Print in Order](https://leetcode.com/problems/print-in-order/description/)

Suppose we have a class:

```java
public class Foo {
    public void first() { print("first"); }
    public void second() { print("second"); }
    public void third() { print("third"); }
}
```

The same instance of `Foo` is passed to three different threads. Thread A calls `first()`, thread B calls `second()`, and thread C calls `third()`. Design a mechanism and modify the program so that `second()` executes after `first()`, and `third()` executes after `second()`.

> **Note:** The thread scheduling order is unknown. The input format simply ensures that the tests cover all permutations.

**Example 1**

```
Input: nums = [1,2,3]
Output: "firstsecondthird"
Explanation: Threads call `first()`, `second()`, and `third()` respectively. The correct output is "firstsecondthird".
```

**Example 2**

```
Input: nums = [1,3,2]
Output: "firstsecondthird"
Explanation: Threads call `first()`, `third()`, and `second()` respectively. The synchronization mechanism still enforces "firstsecondthird".
```

**Constraints**

- `nums` is a permutation of `[1, 2, 3]`.
