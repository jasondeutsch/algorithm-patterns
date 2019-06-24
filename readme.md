<!-- $theme: gaia -->



# Algorithm Patterns: Part 1

---


#### Who hates leetcode?
![](img/dogs_cat.jpg)

---

![](img/intro-2.png)


<sup> source: https://www.nomachetejuggling.com/2014/06/24/the-worst-programming-interview-question/ </sup>

---


#### Algorithm Pattern:  a method, strategy, or technique of solving a problem.

---

# we are conditioned to look for complicated solutions

---

## Patterns
* two pointers
* sliding window
* fast and slow pointers
* merge intervals
* in place linked list traversal
* cyclic sort
* tree breadth first search
* tree depth first searh
* two heaps

---

## Patterns continued
* subsets	
* modified binary search
* top elements
* knapsack
* k way marge
* topological sort
* a few others...(string balancing, etc)


---

## What we'll cover (hopefully!)
* two pointers
* sliding window
* fast and slow pointers
* merge intervals
---

## Personal story

---

## Pattern 1: Two Pointers

### One Liner

An efficient technique for searching pairs in sorted arrays.

---
## Pattern 1: Two Pointers


### General Idea

Use two pointers, representing the first and last element and add the values together. If the result matches the predicate then you done. If the sum is to high, decrement the end pointer, if the sum is to low, increment the start pointer.

---
## Pattern 1: Two Pointers


### Identifying

Sorted array

Calculate something of a given size


---

## Pattern 2: Sliding Window
---

## Pattern 2: Sliding Window

### One Liner

Dynamicaly expand the end and contract the the start of a subsequence during iteration, while never evaluating the index twice. 


---
### One Liner

Dynamicaly expand the end and contract the the start of a subsequence during iteration, while never evaluating the index twice. 

(not as complex at sounds)

---

## Pattern 2: Sliding Window

### General Idea
start with two pointers (index variables) to the start of an array or list)
Consider these as "start and end"
increment end variable until some condition met.
Process values as you go.
When condition met
* add result to result set
* remove the start index's value
* increment start


---

### What is a window?
A contiguouss subsequence of a sequental structure.

There are different kinds of sliding windows, we're going to look fixed window.

---

#### What is a window
![](img/slide_window_1.png)


---


##### Key Insight
![](img/sliding_window_2.png)



---

## Pattern 2: Sliding Window

### Identifying

You have an ordered container (array, string, et cetera)

You need to find some specific type of sub sequence that satisfies some condition

A O(n^2) is available

A nested loop can be re-written as a single loop in some problems

----


## Pattern 3: Fast and Slow Pointers


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

Handling cycles in a linked list or array

---
#### A cycle in a linked list

<img src="./img/linked-list-cycle.png" width="1000"/>

image:	 hackkerrank.com

---

## Pattern 4: Merge Intervals

### One Liner


---

## Pattern 4: Merge Intervals

### General Idea (double check)

1. Sort the Intervals based on the increasing order of start time

2. Push the first interval on to a result set

3. If the current interval does not overlap with the stack top, push it
If the current element overlaps with stack top and ending time of current interval is more     than that of stack top, update stack top with the ending time of current interval
---

## Pattern 4: Merge Intervals
### Identifying

Overlapping intervals


----

## Six Cases

----	

![merge_intervals_1](img/merge_intervals_1.png)

---

![merge_intervals_2](img/merge_intervals_2.png)

---

![merge_intervals_3](img/merge_intervals_3.png)

---

![merge_intervals_4](img/merge_intervals_4.png)

---

![merge_intervals_5](img/merge_intervals_5.png)

---



![merge_intervals_6](img/merge_intervals_6.png)

---

## Find The Pattern