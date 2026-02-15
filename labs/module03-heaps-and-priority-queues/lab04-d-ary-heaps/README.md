# Lab 3.4: D-ary Heaps

**Module**: Module 3 - Heaps & Priority Queues  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 3.4: D-ary Heaps

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

---

## Back to Module

[← Back to Module 3: Heaps & Priority Queues](../module_03/README.md)

## Navigation

- [Previous Lab](./lab_3_3.md) (if exists)
- [Next Lab](./lab_3_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 3, Lab 4 of 5*
