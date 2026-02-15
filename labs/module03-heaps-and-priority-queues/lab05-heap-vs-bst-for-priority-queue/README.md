# Lab 3.5: Heap vs. BST for Priority Queue

**Module**: Module 3 - Heaps & Priority Queues  
**Duration**: 2-3 hours  
**Difficulty**: Easy-Medium

---

Lab 3.5: Heap vs. BST for Priority Queue

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

---

## Back to Module

[← Back to Module 3: Heaps & Priority Queues](../module_03/README.md)

## Navigation

- [Previous Lab](./lab_3_4.md) (if exists)
- [Next Lab](./lab_3_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 3, Lab 5 of 5*
