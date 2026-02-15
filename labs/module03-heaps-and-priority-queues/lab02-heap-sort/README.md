# Lab 3.2: Heap Sort

**Module**: Module 3 - Heaps & Priority Queues  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 3.2: Heap Sort

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Implement heap sort algorithm
- Compare in-place vs. non-in-place variants
- Understand heap sort properties (unstable, in-place, O(n log n))

### Requirements

1. **Implement heap sort**:
   - In-place version (modifying input array)
   - Non-in-place version (using separate heap)
   - Both ascending and descending order

2. **Optimization**:
   - Bottom-up heapify for better constants
   - Early termination for partially sorted arrays
   - Hybrid approaches (switch to insertion sort for small subarrays)

3. **Analysis**:
   - Count comparisons and swaps
   - Measure cache performance
   - Compare with quicksort and mergesort

### Implementation Details

```python
def heap_sort_in_place(array):
    n = len(array)
    
    # Build max heap
    for i in range(n // 2 - 1, -1, -1):
        heapify_down(array, i, n)
    
    # Extract elements one by one
    for i in range(n - 1, 0, -1):
        array[0], array[i] = array[i], array[0]
        heapify_down(array, 0, i)

def heapify_down(array, index, heap_size):
    largest = index
    left = 2 * index + 1
    right = 2 * index + 2
    
    if left < heap_size and array[left] > array[largest]:
        largest = left
    if right < heap_size and array[right] > array[largest]:
        largest = right
    
    if largest != index:
        array[index], array[largest] = array[largest], array[index]
        heapify_down(array, largest, heap_size)

def heap_sort_non_in_place(array):
    heap = BinaryHeap(is_min_heap=True)
    heap.build_heap(array)
    
    result = []
    while len(heap.heap) > 0:
        result.append(heap.extract_min())
    
    return result

def hybrid_heap_sort(array, threshold=10):
    # Use heap sort for large arrays, insertion sort for small
    if len(array) <= threshold:
        return insertion_sort(array)
    return heap_sort_in_place(array)
```

### Test Cases
- Random arrays of various sizes (10, 100, 1K, 10K, 100K)
- Already sorted arrays
- Reverse sorted arrays
- Arrays with duplicates
- Nearly sorted arrays

### Deliverables
- `heap_sort.{py,go,zig}` - Multiple heap sort implementations
- `comparison_analysis.md` - Compare with other O(n log n) sorts
- `benchmarks.{py,go,zig}` - Performance measurements
- `visualization.{py,go,zig}` - Visualize sorting process

### Success Criteria
- Correct sorting for all test cases
- In-place version uses O(1) extra space
- Understanding of why heap sort isn't typically fastest in practice
- Can explain stability issues

---

---

## Back to Module

[← Back to Module 3: Heaps & Priority Queues](../module_03/README.md)

## Navigation

- [Previous Lab](./lab_3_1.md) (if exists)
- [Next Lab](./lab_3_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 3, Lab 2 of 5*
