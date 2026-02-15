# Module 3: Heaps & Priority Queues

**Duration**: 2 weeks  
**Difficulty**: Medium

## Overview

Priority queues are fundamental to many algorithms, from graph algorithms to task scheduling. Heaps provide an efficient implementation with guaranteed logarithmic operations. This module explores heap structures, heap sort, and practical applications.

**Why this matters**: Priority queues appear everywhere - operating system schedulers, network routing, compression algorithms, event simulation, and more. Understanding heaps deeply gives you a powerful tool for solving ordering problems efficiently.

## Learning Objectives

- Implement binary heaps from scratch with array representation
- Understand heap property and heapify operations
- Master heap sort algorithm
- Build priority queue applications
- Explore d-ary heaps and their tradeoffs
- Compare heap-based vs. tree-based priority queues

## Topics Covered

- Binary heaps (min and max)
- Heap operations and complexity
- Heap sort and variants
- Priority queue operations
- D-ary heaps
- Fibonacci heaps (theoretical analysis)
- Applications: scheduling, graph algorithms, streaming data

---

## Lab 3.1: Binary Heap Implementation

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement min-heap and max-heap from scratch
- Understand heap property and heapify operations
- Use array-based representation efficiently

### Requirements

1. **Implement binary heap using array**:
   - Parent at index i, children at 2i+1 and 2i+2
   - `insert(value)` - O(log n)
   - `extract_min()` or `extract_max()` - O(log n)
   - `peek()` - O(1)
   - `decrease_key(index, new_value)` - O(log n)
   - `delete(index)` - O(log n)

2. **Implement heapify operations**:
   - `heapify_up(index)` - Bubble up
   - `heapify_down(index)` - Bubble down
   - `build_heap(array)` - O(n) construction from unsorted array

3. **Implement both min-heap and max-heap** (or parameterized comparator)

### Implementation Details

```python
class BinaryHeap:
    def __init__(self, is_min_heap=True):
        self.heap = []
        self.comparator = (lambda a, b: a < b) if is_min_heap else (lambda a, b: a > b)
    
    def insert(self, value):
        self.heap.append(value)
        self._heapify_up(len(self.heap) - 1)
    
    def _heapify_up(self, index):
        parent = (index - 1) // 2
        if parent >= 0 and self.comparator(self.heap[index], self.heap[parent]):
            self.heap[index], self.heap[parent] = self.heap[parent], self.heap[index]
            self._heapify_up(parent)
    
    def extract_min(self):
        if len(self.heap) == 0:
            return None
        if len(self.heap) == 1:
            return self.heap.pop()
        
        root = self.heap[0]
        self.heap[0] = self.heap.pop()
        self._heapify_down(0)
        return root
    
    def _heapify_down(self, index):
        smallest = index
        left = 2 * index + 1
        right = 2 * index + 2
        
        if left < len(self.heap) and self.comparator(self.heap[left], self.heap[smallest]):
            smallest = left
        if right < len(self.heap) and self.comparator(self.heap[right], self.heap[smallest]):
            smallest = right
        
        if smallest != index:
            self.heap[index], self.heap[smallest] = self.heap[smallest], self.heap[index]
            self._heapify_down(smallest)
    
    def build_heap(self, array):
        self.heap = array[:]
        # Start from last non-leaf node and heapify down
        for i in range(len(self.heap) // 2 - 1, -1, -1):
            self._heapify_down(i)
```

### Test Cases
- Insert random elements, verify heap property after each insertion
- Extract all elements, verify sorted order
- Build heap from unsorted array, verify O(n) time
- Test decrease_key and delete operations
- Edge cases: empty heap, single element, duplicate values

### Deliverables
- `binary_heap.{py,go,zig}` - Complete heap implementation
- `heap_visualizer.{py,go,zig}` - Visualize heap as tree and array
- `build_heap_analysis.md` - Prove O(n) build time
- `test_suite.{py,go,zig}` - Comprehensive tests
- `benchmarks.{py,go,zig}` - Performance measurements

### Success Criteria
- Heap property maintained after all operations
- Build heap is provably O(n)
- Understanding of parent-child index relationships
- Can explain why bottom-up build is O(n)

---

## Lab 3.2: Heap Sort

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

## Lab 3.3: Priority Queue Applications

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Build practical applications using priority queues
- Understand when priority queues are the right tool
- Implement various priority queue variants

### Requirements

1. **Implement Task Scheduler**:
   - Tasks have priority and deadline
   - Schedule highest priority task first
   - Support adding tasks dynamically
   - Handle task preemption

2. **Implement K-way Merge**:
   - Merge k sorted arrays into one sorted array
   - Use min-heap of size k
   - Optimal O(n log k) algorithm

3. **Implement Top-K Elements**:
   - Find k largest/smallest elements in stream
   - Use heap of size k
   - Compare with sorting approach

4. **Implement Median Maintenance**:
   - Maintain median of elements seen so far
   - Use two heaps (max-heap and min-heap)
   - Support insertions and median queries

### Implementation Details

```python
class TaskScheduler:
    def __init__(self):
        self.task_queue = BinaryHeap(is_min_heap=False)  # Max heap by priority
    
    def add_task(self, task, priority, deadline):
        self.task_queue.insert((priority, deadline, task))
    
    def get_next_task(self):
        if len(self.task_queue.heap) == 0:
            return None
        return self.task_queue.extract_min()[2]  # Return task only
    
    def peek_next(self):
        if len(self.task_queue.heap) == 0:
            return None
        return self.task_queue.peek()[2]

def k_way_merge(arrays):
    heap = BinaryHeap(is_min_heap=True)
    result = []
    
    # Initialize heap with first element from each array
    for i, arr in enumerate(arrays):
        if arr:
            heap.insert((arr[0], i, 0))  # (value, array_index, element_index)
    
    while len(heap.heap) > 0:
        value, arr_idx, elem_idx = heap.extract_min()
        result.append(value)
        
        # Add next element from same array
        if elem_idx + 1 < len(arrays[arr_idx]):
            next_val = arrays[arr_idx][elem_idx + 1]
            heap.insert((next_val, arr_idx, elem_idx + 1))
    
    return result

class TopKFinder:
    def __init__(self, k):
        self.k = k
        self.heap = BinaryHeap(is_min_heap=True)  # Min heap for top k largest
    
    def add(self, value):
        if len(self.heap.heap) < self.k:
            self.heap.insert(value)
        elif value > self.heap.peek():
            self.heap.extract_min()
            self.heap.insert(value)
    
    def get_top_k(self):
        return sorted(self.heap.heap, reverse=True)

class MedianMaintainer:
    def __init__(self):
        self.max_heap = BinaryHeap(is_min_heap=False)  # Lower half
        self.min_heap = BinaryHeap(is_min_heap=True)   # Upper half
    
    def insert(self, value):
        # Add to max heap first
        if len(self.max_heap.heap) == 0 or value <= self.max_heap.peek():
            self.max_heap.insert(value)
        else:
            self.min_heap.insert(value)
        
        # Rebalance heaps
        if len(self.max_heap.heap) > len(self.min_heap.heap) + 1:
            self.min_heap.insert(self.max_heap.extract_min())
        elif len(self.min_heap.heap) > len(self.max_heap.heap):
            self.max_heap.insert(self.min_heap.extract_min())
    
    def get_median(self):
        if len(self.max_heap.heap) == 0:
            return None
        if len(self.max_heap.heap) > len(self.min_heap.heap):
            return self.max_heap.peek()
        return (self.max_heap.peek() + self.min_heap.peek()) / 2.0
```

### Test Cases
- Task scheduler: Verify tasks executed in priority order
- K-way merge: Test with k=2, k=10, k=100
- Top-K: Verify correct k elements found
- Median: Test with odd/even number of elements, streaming data

### Deliverables
- `task_scheduler.{py,go,zig}` - Complete scheduler implementation
- `k_way_merge.{py,go,zig}` - K-way merge algorithm
- `top_k.{py,go,zig}` - Top-K elements finder
- `median_maintainer.{py,go,zig}` - Running median
- `applications.md` - Discussion of use cases
- `benchmarks.{py,go,zig}` - Performance analysis

### Success Criteria
- All applications work correctly
- Understanding of why priority queue is optimal for each problem
- Can identify new problems suited for priority queues
- Performance matches theoretical complexity

---

## Lab 3.4: D-ary Heaps

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Implement heaps with branching factor d
- Find optimal d for different workloads
- Understand cache effects and tree height tradeoffs

### Requirements

1. **Implement parameterized d-ary heap**:
   - Binary heap (d=2) as special case
   - Ternary heap (d=3)
   - Arbitrary d

2. **Index calculations**:
   - Parent: `(i-1) / d`
   - kth child: `d*i + k + 1` for k ∈ [0, d-1]

3. **Find optimal d for**:
   - Insert-heavy workloads
   - Extract-heavy workloads
   - Mixed workloads
   - Different element sizes

4. **Analyze tradeoffs**:
   - Tree height: decreases with larger d
   - Heapify cost: increases with larger d (more children to compare)
   - Cache performance: larger d may improve locality

### Implementation Details

```python
class DaryHeap:
    def __init__(self, d=2, is_min_heap=True):
        self.d = d
        self.heap = []
        self.comparator = (lambda a, b: a < b) if is_min_heap else (lambda a, b: a > b)
    
    def parent(self, i):
        return (i - 1) // self.d
    
    def kth_child(self, i, k):
        return self.d * i + k + 1
    
    def insert(self, value):
        self.heap.append(value)
        self._heapify_up(len(self.heap) - 1)
    
    def _heapify_up(self, index):
        if index == 0:
            return
        
        parent_idx = self.parent(index)
        if self.comparator(self.heap[index], self.heap[parent_idx]):
            self.heap[index], self.heap[parent_idx] = self.heap[parent_idx], self.heap[index]
            self._heapify_up(parent_idx)
    
    def extract_min(self):
        if len(self.heap) == 0:
            return None
        if len(self.heap) == 1:
            return self.heap.pop()
        
        root = self.heap[0]
        self.heap[0] = self.heap.pop()
        self._heapify_down(0)
        return root
    
    def _heapify_down(self, index):
        extreme_child = self._find_extreme_child(index)
        if extreme_child != -1 and self.comparator(self.heap[extreme_child], self.heap[index]):
            self.heap[index], self.heap[extreme_child] = self.heap[extreme_child], self.heap[index]
            self._heapify_down(extreme_child)
    
    def _find_extreme_child(self, index):
        # Find min/max among d children
        first_child = self.kth_child(index, 0)
        if first_child >= len(self.heap):
            return -1
        
        extreme = first_child
        for k in range(1, self.d):
            child = self.kth_child(index, k)
            if child >= len(self.heap):
                break
            if self.comparator(self.heap[child], self.heap[extreme]):
                extreme = child
        
        return extreme
    
    def height(self):
        import math
        if len(self.heap) == 0:
            return 0
        return math.ceil(math.log(len(self.heap) * (self.d - 1) + 1, self.d))
```

### Test Cases
- Test d=2,3,4,8,16 for correctness
- Benchmark insert/extract for each d
- Measure tree height for different d
- Cache performance profiling

### Deliverables
- `d_ary_heap.{py,go,zig}` - Parameterized heap implementation
- `optimal_d_analysis.md` - Analysis of optimal branching factor
- `benchmarks.{py,go,zig}` - Performance across different d values
- `height_analysis.{py,go,zig}` - Measure actual heights
- Graphs showing tradeoffs

### Success Criteria
- Correct operation for any d ≥ 2
- Clear understanding of height vs. comparison tradeoff
- Can recommend optimal d for given workload
- Understanding of cache effects

---

## Lab 3.5: Heap vs. BST for Priority Queue

**Duration**: 2-3 hours  
**Difficulty**: Easy-Medium

### Objectives
- Compare heap and BST as priority queue implementations
- Understand operation-specific tradeoffs
- Make informed implementation choices

### Requirements

1. **Implement priority queue using both**:
   - Binary heap
   - Balanced BST (AVL or Red-Black from Module 2)

2. **Benchmark operations**:
   - Insert
   - Extract min/max
   - Find min/max (peek)
   - Decrease key
   - Delete arbitrary element
   - Build from array

3. **Test different workloads**:
   - Insert-only
   - Extract-only
   - Alternating insert/extract
   - Rare decrease-key operations
   - Frequent arbitrary deletions

### Implementation Details

```python
class HeapPriorityQueue:
    def __init__(self):
        self.heap = BinaryHeap(is_min_heap=True)
    
    def insert(self, value):
        self.heap.insert(value)
    
    def extract_min(self):
        return self.heap.extract_min()
    
    def peek(self):
        return self.heap.peek()
    
    def decrease_key(self, old_value, new_value):
        # Requires index tracking - expensive for heap
        pass
    
    def delete(self, value):
        # Requires search - O(n) for heap
        pass

class BSTPriorityQueue:
    def __init__(self):
        self.bst = AVLTree()  # Or RedBlackTree
    
    def insert(self, value):
        self.bst.insert(value)
    
    def extract_min(self):
        min_val = self.bst.minimum()
        self.bst.delete(min_val)
        return min_val
    
    def peek(self):
        return self.bst.minimum()
    
    def decrease_key(self, old_value, new_value):
        self.bst.delete(old_value)
        self.bst.insert(new_value)
    
    def delete(self, value):
        self.bst.delete(value)

def benchmark_priority_queue_operations(pq_type, operations):
    if pq_type == "heap":
        pq = HeapPriorityQueue()
    else:
        pq = BSTPriorityQueue()
    
    times = {op: 0 for op in ["insert", "extract", "peek", "decrease_key", "delete"]}
    
    # Benchmark each operation type
    # ... timing code ...
    
    return times
```

### Test Cases
- 100K insertions followed by 100K extractions
- Interleaved operations
- Decrease-key heavy workload
- Random deletions

### Deliverables
- `priority_queue_heap.{py,go,zig}` - Heap-based PQ
- `priority_queue_bst.{py,go,zig}` - BST-based PQ
- `comparison.md` - Detailed analysis
- `benchmarks.{py,go,zig}` - Performance comparison
- `decision_matrix.md` - Guidelines for choosing implementation

### Success Criteria
- Both implementations correct
- Understanding that heap is typically better for standard PQ operations
- Knowledge of when BST might be preferable (e.g., need sorted iteration)
- Clear performance data supporting conclusions

---

## Module Resources

### CLRS References
- Chapter 6: Heapsort
- Chapter 19: Fibonacci Heaps (theoretical)

### Additional Reading
- "Priority Queues and Heaps" by Mehlhorn and Sanders
- Fibonacci heap papers (advanced)
- Cache-oblivious heaps research

### Key Takeaways

By the end of this module, you should:
1. Have deep understanding of heap property and operations
2. Know when to use heaps vs. other structures
3. Understand heap sort's place among sorting algorithms
4. Be able to implement priority queues efficiently
5. Recognize priority queue applications in algorithms

### Next Module

**Module 4: Sorting & Selection** - You'll implement and analyze multiple sorting algorithms, understanding when each excels and how to optimize for real-world data.

---

## Tips for Success

1. **Visualize the heap** - Draw the tree representation
2. **Trace heapify operations** - Step through manually
3. **Test edge cases** - Empty heap, single element
4. **Measure everything** - Use your Module 0 tools
5. **Compare implementations** - Heap vs BST, different d values

## Common Pitfalls

- **Array indexing errors** - Parent/child calculations are easy to get wrong
- **Forgetting to maintain heap property** - After every modification
- **Inefficient decrease_key in heap** - Requires index tracking
- **Not considering cache effects** - D-ary heap performance depends on it
- **Incorrect heap sort** - Remember to heapify entire array first
