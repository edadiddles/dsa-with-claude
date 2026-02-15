# Lab 3.1: Binary Heap Implementation

**Module**: Module 3 - Heaps & Priority Queues  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 3.1: Binary Heap Implementation

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

---

## Back to Module

[← Back to Module 3: Heaps & Priority Queues](../module_03/README.md)

## Navigation

- [Previous Lab](./lab_3_0.md) (if exists)
- [Next Lab](./lab_3_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 3, Lab 1 of 5*
