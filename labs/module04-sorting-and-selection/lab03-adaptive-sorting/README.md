# Lab 4.3: Adaptive Sorting

**Module**: Module 4 - Sorting & Selection  
**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

---

Lab 4.3: Adaptive Sorting

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Build sorts that detect and exploit pre-existing order
- Understand adaptivity and why it matters
- Implement Timsort-like optimizations

### Requirements

1. **Implement Adaptive Insertion Sort**:
   - Skip already-sorted regions
   - Use binary search for insertion position
   - Early termination when fully sorted

2. **Implement Natural Mergesort**:
   - Detect existing runs (ascending/descending sequences)
   - Merge natural runs instead of fixed-size subarrays
   - Avoid unnecessary comparisons in sorted regions

3. **Implement Timsort (simplified version)**:
   - Find runs (min run size)
   - Use insertion sort for small runs
   - Merge runs using galloping mode
   - Maintain run stack

4. **Measure adaptivity**:
   - Sorted array: Should be O(n)
   - Reverse sorted: Should detect and reverse runs
   - Nearly sorted: Should be faster than O(n log n)

### Implementation Details

```python
def natural_mergesort(array):
    runs = find_runs(array)
    
    while len(runs) > 1:
        merged_runs = []
        for i in range(0, len(runs), 2):
            if i + 1 < len(runs):
                merge_runs(array, runs[i], runs[i+1])
                merged_runs.append((runs[i][0], runs[i+1][1]))
            else:
                merged_runs.append(runs[i])
        runs = merged_runs
    
    return array

def find_runs(array):
    runs = []
    start = 0
    i = 1
    
    while i < len(array):
        # Detect ascending or descending run
        if i < len(array) and array[i] < array[i-1]:
            # Descending run
            while i < len(array) and array[i] < array[i-1]:
                i += 1
            reverse(array, start, i-1)
        else:
            # Ascending run
            while i < len(array) and array[i] >= array[i-1]:
                i += 1
        
        runs.append((start, i-1))
        start = i
        i += 1
    
    if start < len(array):
        runs.append((start, len(array)-1))
    
    return runs

class Timsort:
    MIN_MERGE = 32
    
    def __init__(self):
        self.min_gallop = 7
    
    def sort(self, array):
        n = len(array)
        min_run = self.compute_min_run(n)
        
        # Create initial runs
        for start in range(0, n, min_run):
            end = min(start + min_run - 1, n - 1)
            insertion_sort_range(array, start, end)
        
        # Merge runs
        size = min_run
        while size < n:
            for start in range(0, n, size * 2):
                mid = start + size - 1
                end = min(start + size * 2 - 1, n - 1)
                
                if mid < end:
                    self.merge_with_galloping(array, start, mid, end)
            
            size *= 2
    
    def compute_min_run(self, n):
        r = 0
        while n >= self.MIN_MERGE:
            r |= n & 1
            n >>= 1
        return n + r
    
    def merge_with_galloping(self, array, lo, mid, hi):
        # Timsort's galloping merge
        left = array[lo:mid+1]
        right = array[mid+1:hi+1]
        
        i = j = 0
        k = lo
        min_gallop = self.min_gallop
        
        while i < len(left) and j < len(right):
            # Regular merge until one side starts "winning"
            if left[i] <= right[j]:
                array[k] = left[i]
                i += 1
            else:
                array[k] = right[j]
                j += 1
            k += 1
            
            # Enter galloping mode if one side keeps winning
            # (simplified - full Timsort has more sophisticated galloping)
        
        while i < len(left):
            array[k] = left[i]
            i += 1
            k += 1
        
        while j < len(right):
            array[k] = right[j]
            j += 1
            k += 1

def adaptive_insertion_sort(array):
    # Binary search insertion sort
    for i in range(1, len(array)):
        key = array[i]
        
        # Binary search for insertion position
        left, right = 0, i - 1
        while left <= right:
            mid = (left + right) // 2
            if array[mid] > key:
                right = mid - 1
            else:
                left = mid + 1
        
        # Shift elements and insert
        for j in range(i, left, -1):
            array[j] = array[j-1]
        array[left] = key
```

### Test Cases
- Fully sorted array
- Reverse sorted array
- Partially sorted (e.g., sorted with 10% random elements)
- Data with multiple sorted runs
- Compare with non-adaptive sorts

### Deliverables
- `adaptive_insertion_sort.{py,go,zig}`
- `natural_mergesort.{py,go,zig}`
- `timsort.{py,go,zig}` - Simplified implementation
- `adaptivity_benchmarks.{py,go,zig}`
- `analysis.md` - When adaptivity helps
- `run_detection.{py,go,zig}` - Visualize detected runs

### Success Criteria
- Near-linear time on sorted/nearly-sorted data
- Understanding of why Python uses Timsort
- Can detect and measure adaptivity benefits
- Recognition of when adaptive sorts are worth the complexity

---

---

## Back to Module

[← Back to Module 4: Sorting & Selection](../module_04/README.md)

## Navigation

- [Previous Lab](./lab_4_2.md) (if exists)
- [Next Lab](./lab_4_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 4, Lab 3 of 7*
