<!-- $theme: gaia -->



# Algorithm Patterns: Part 1



#### Algorithm Pattern:  a method, strategy, or technique of solving a problem.

---

### square a list

#### specific solution

```go
for i := 0; i < len(arr)-1; i++ {
  arr[i] *= 2
}
```

---


#### general solution

Apply some function to each thing


```haskell
fmap :: (a -> b)-> f a -> f b

fmap list (*2)
fmap list (\x -> x*x)
fmap...
```

---

## Pattern 1: Two Pointers

### One Liner

An efficient techniquefor for searching pairs in sorted arrays.

---
## Pattern 1: Two Pointers


### General Idea

Use two pointers, representing the first and last element, and add the values together. If the result matches the predicate then you done. If the sum is to high, decrement the end pointer, if the sum is to low, increment the start pointer.

---
## Pattern 1: Two Pointers


### Identifying

Sorted array/linked list

Calculate something of a given size



---

## Pattern 2: Sliding Window

### One Liner

A nested loop can be re-written as a single loop in some problems

---

## Pattern 2: Sliding Window
### General Idea

---

## Pattern 2: Sliding Window

### Identifying

You have an ordered container (array, string, et cetera)

You need to findsome specificy type of sub sequence that satisfies some condition

A O(n^2) is available

----

## Pattern 3: Fast and Slow Pointers

### One Liner

Use two pointers which move through some sequential  data structure at different speeds.

 
----

## Pattern 3: Fast and Slow Pointers

### General Idea

Use two pointers which move through some sequential  data structure at different speeds.

[not much else to say]

----

## Pattern 3: Fast and Slow Pointers
### Identifying

Handling cycles in a linked list or arrays

---

## Pattern 4: Merge Intervals

### One Liner


---

## Pattern 4: Merge Intervals

### General Idea

---

## Pattern 4: Merge Intervals
### Identifying

Overlapping intervals



----

## Six Cases

----

![merge_intervals_1](/Users/jason/go/src/algorithm-patterns/img/merge_intervals_1.png)

---

![merge_intervals_2](/Users/jason/go/src/algorithm-patterns/img/merge_intervals_2.png)

---

![merge_intervals_3](/Users/jason/go/src/algorithm-patterns/img/merge_intervals_3.png)

---

![merge_intervals_4](/Users/jason/go/src/algorithm-patterns/img/merge_intervals_4.png)

---

![merge_intervals_5](/Users/jason/go/src/algorithm-patterns/img/merge_intervals_5.png)

---



![merge_intervals_6](/Users/jason/go/src/algorithm-patterns/img/merge_intervals_6.png)

---

