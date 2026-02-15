# Lab 3.3: Priority Queue Applications

**Module**: Module 3 - Heaps & Priority Queues  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 3.3: Priority Queue Applications

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

---

## Back to Module

[← Back to Module 3: Heaps & Priority Queues](../module_03/README.md)

## Navigation

- [Previous Lab](./lab_3_2.md) (if exists)
- [Next Lab](./lab_3_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 3, Lab 3 of 5*
