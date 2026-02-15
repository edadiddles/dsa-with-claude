# Lab 4.1: Implement Five Core Sorting Algorithms

**Module**: Module 4 - Sorting & Selection  
**Duration**: 6-8 hours  
**Difficulty**: Medium

---

Lab 4.1: Implement Five Core Sorting Algorithms

**Duration**: 6-8 hours  
**Difficulty**: Medium

### Objectives
- Implement major sorting algorithms from scratch
- Understand implementation details and edge cases
- Compare theoretical and practical performance

### Requirements

1. **Implement the following sorts**:
   - **Quicksort**: With various pivot selection strategies
   - **Mergesort**: Both top-down and bottom-up
   - **Heapsort**: Using your heap from Module 3
   - **Insertion Sort**: For small arrays
   - **Selection Sort**: For comparison

2. **For each algorithm**:
   - Basic implementation
   - Optimized version
   - In-place variant where applicable
   - Stable variant where applicable

3. **Quicksort pivot strategies**:
   - First element
   - Last element
   - Middle element
   - Random element
   - Median-of-three
   - Median-of-medians (deterministic O(n log n))

### Implementation Details

```python
def quicksort(array, lo=0, hi=None, pivot_strategy="median_of_three"):
    if hi is None:
        hi = len(array) - 1
    
    if lo < hi:
        pivot_idx = partition(array, lo, hi, pivot_strategy)
        quicksort(array, lo, pivot_idx - 1, pivot_strategy)
        quicksort(array, pivot_idx + 1, hi, pivot_strategy)

def partition(array, lo, hi, strategy):
    # Choose pivot based on strategy
    pivot_idx = choose_pivot(array, lo, hi, strategy)
    array[pivot_idx], array[hi] = array[hi], array[pivot_idx]
    
    pivot = array[hi]
    i = lo - 1
    
    for j in range(lo, hi):
        if array[j] <= pivot:
            i += 1
            array[i], array[j] = array[j], array[i]
    
    array[i + 1], array[hi] = array[hi], array[i + 1]
    return i + 1

def mergesort_top_down(array):
    if len(array) <= 1:
        return array
    
    mid = len(array) // 2
    left = mergesort_top_down(array[:mid])
    right = mergesort_top_down(array[mid:])
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0
    
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    
    result.extend(left[i:])
    result.extend(right[j:])
    return result

def mergesort_bottom_up(array):
    n = len(array)
    size = 1
    
    while size < n:
        for start in range(0, n, size * 2):
            mid = min(start + size, n)
            end = min(start + size * 2, n)
            merge_in_place(array, start, mid, end)
        size *= 2

def heapsort(array):
    # Build max heap
    n = len(array)
    for i in range(n // 2 - 1, -1, -1):
        heapify_down(array, i, n)
    
    # Extract elements
    for i in range(n - 1, 0, -1):
        array[0], array[i] = array[i], array[0]
        heapify_down(array, 0, i)

def insertion_sort(array):
    for i in range(1, len(array)):
        key = array[i]
        j = i - 1
        while j >= 0 and array[j] > key:
            array[j + 1] = array[j]
            j -= 1
        array[j + 1] = key

def selection_sort(array):
    for i in range(len(array)):
        min_idx = i
        for j in range(i + 1, len(array)):
            if array[j] < array[min_idx]:
                min_idx = j
        array[i], array[min_idx] = array[min_idx], array[i]
```

### Test Cases
- Random arrays: [10, 100, 1K, 10K, 100K, 1M] elements
- Already sorted arrays (best case for some, worst for others)
- Reverse sorted arrays
- Arrays with many duplicates
- Nearly sorted arrays (important for real-world data)
- Arrays with specific patterns (sawtooth, organ-pipe, etc.)

### Deliverables
- `quicksort.{py,go,zig}` - Multiple pivot strategies
- `mergesort.{py,go,zig}` - Top-down and bottom-up
- `heapsort.{py,go,zig}` - Using heap from Module 3
- `insertion_sort.{py,go,zig}` - Optimized version
- `selection_sort.{py,go,zig}` - Basic implementation
- `sort_comparison.md` - Analysis of characteristics
- `benchmarks.{py,go,zig}` - Comprehensive performance tests
- `test_data_generator.{py,go,zig}` - Generate various input patterns

### Success Criteria
- All sorts produce correct output
- Understanding of stability, adaptivity, and space complexity
- Can explain when each algorithm is preferable
- Performance matches theoretical expectations

---

---

## Back to Module

[← Back to Module 4: Sorting & Selection](../module_04/README.md)

## Navigation

- [Previous Lab](./lab_4_0.md) (if exists)
- [Next Lab](./lab_4_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 4, Lab 1 of 7*
