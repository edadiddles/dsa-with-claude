# Lab 4.5: Selection Algorithms

**Module**: Module 4 - Sorting & Selection  
**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

---

Lab 4.5: Selection Algorithms

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Find kth smallest element without full sort
- Implement both randomized and deterministic algorithms
- Understand expected vs. worst-case guarantees

### Requirements

1. **Implement Randomized Quickselect**:
   - Expected O(n) time
   - Use random pivot
   - Partition and recurse on one side only

2. **Implement Median of Medians**:
   - Worst-case O(n) time
   - Deterministic pivot selection
   - Guaranteed linear time

3. **Additional selection algorithms**:
   - Heap-based selection (for small k)
   - Partial quicksort
   - Iterative vs. recursive versions

4. **Applications**:
   - Find median
   - Find k largest/smallest elements
   - Percentile calculation

### Implementation Details

```python
import random

def randomized_select(array, k):
    """Find kth smallest element (0-indexed)"""
    if len(array) == 1:
        return array[0]
    
    pivot_idx = random.randint(0, len(array) - 1)
    pivot = array[pivot_idx]
    
    lows = [x for x in array if x < pivot]
    highs = [x for x in array if x > pivot]
    pivots = [x for x in array if x == pivot]
    
    if k < len(lows):
        return randomized_select(lows, k)
    elif k < len(lows) + len(pivots):
        return pivot
    else:
        return randomized_select(highs, k - len(lows) - len(pivots))

def median_of_medians(array, k):
    """Deterministic O(n) selection"""
    if len(array) <= 5:
        return sorted(array)[k]
    
    # Divide into groups of 5
    medians = []
    for i in range(0, len(array), 5):
        group = array[i:i+5]
        medians.append(sorted(group)[len(group)//2])
    
    # Find median of medians recursively
    pivot = median_of_medians(medians, len(medians)//2)
    
    # Partition around pivot
    lows = [x for x in array if x < pivot]
    highs = [x for x in array if x > pivot]
    pivots = [x for x in array if x == pivot]
    
    if k < len(lows):
        return median_of_medians(lows, k)
    elif k < len(lows) + len(pivots):
        return pivot
    else:
        return median_of_medians(highs, k - len(lows) - len(pivots))

def heap_select(array, k):
    """Use min-heap to find k smallest elements"""
    import heapq
    
    # For k smallest: use max heap of size k
    # For k largest: use min heap of size k
    heap = []
    
    for num in array:
        if len(heap) < k:
            heapq.heappush(heap, -num)  # Negative for max heap
        elif num < -heap[0]:
            heapq.heapreplace(heap, -num)
    
    return -heap[0]  # kth smallest

def find_median(array):
    """Find median using selection"""
    n = len(array)
    if n % 2 == 1:
        return randomized_select(array, n // 2)
    else:
        return (randomized_select(array, n // 2 - 1) + 
                randomized_select(array, n // 2)) / 2.0

def find_percentile(array, p):
    """Find pth percentile (p in [0, 100])"""
    k = int(len(array) * p / 100)
    return randomized_select(array, k)

def partition_in_place(array, lo, hi, pivot_idx):
    """Partition for in-place quickselect"""
    pivot = array[pivot_idx]
    array[pivot_idx], array[hi] = array[hi], array[pivot_idx]
    
    store_idx = lo
    for i in range(lo, hi):
        if array[i] < pivot:
            array[i], array[store_idx] = array[store_idx], array[i]
            store_idx += 1
    
    array[store_idx], array[hi] = array[hi], array[store_idx]
    return store_idx

def quickselect_in_place(array, lo, hi, k):
    """In-place quickselect"""
    if lo == hi:
        return array[lo]
    
    pivot_idx = random.randint(lo, hi)
    pivot_idx = partition_in_place(array, lo, hi, pivot_idx)
    
    if k == pivot_idx:
        return array[k]
    elif k < pivot_idx:
        return quickselect_in_place(array, lo, pivot_idx - 1, k)
    else:
        return quickselect_in_place(array, pivot_idx + 1, hi, k)
```

### Test Cases
- Find median in random array
- Find min (k=0) and max (k=n-1)
- Find quartiles
- Arrays with duplicates
- Adversarial inputs
- Large arrays (verify O(n) performance)

### Deliverables
- `randomized_select.{py,go,zig}`
- `median_of_medians.{py,go,zig}`
- `heap_select.{py,go,zig}` - For small k
- `benchmarks.{py,go,zig}` - Compare approaches
- `analysis.md` - When to use each algorithm
- `applications.{py,go,zig}` - Median, percentiles, etc.

### Success Criteria
- Both algorithms find correct kth element
- Randomized is faster in practice
- Understanding of worst-case guarantees
- Can explain when to use heap-based selection (small k)

---

---

## Back to Module

[← Back to Module 4: Sorting & Selection](../module_04/README.md)

## Navigation

- [Previous Lab](./lab_4_4.md) (if exists)
- [Next Lab](./lab_4_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 4, Lab 5 of 7*
