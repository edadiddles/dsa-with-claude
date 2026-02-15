# Lab 5.3: Segment Tree

**Module**: Module 5 - Advanced Data Structures  
**Duration**: 5-6 hours  
**Difficulty**: Hard

---

Lab 5.3: Segment Tree

**Duration**: 5-6 hours  
**Difficulty**: Hard

### Objectives
- Implement segment tree for range queries
- Support point updates and range updates
- Understand lazy propagation

### Requirements

1. **Implement segment tree supporting**:
   - `build(array)` - O(n) construction
   - `query(left, right)` - O(log n) range query
   - `update(index, value)` - O(log n) point update
   - Operations: sum, min, max, GCD, etc.

2. **Implement lazy propagation for**:
   - Range updates: add value to range
   - Range set: set all values in range
   - More complex updates

3. **Additional features**:
   - Persistent segment tree
   - 2D segment tree (for matrix range queries)
   - Dynamic segment tree (for large coordinate spaces)

### Implementation Details

```python
class SegmentTree:
    def __init__(self, array, operation='sum'):
        self.n = len(array)
        self.tree = [0] * (4 * self.n)
        self.operation = operation
        self.identity = self._get_identity()
        self._build(array, 0, 0, self.n - 1)
    
    def _get_identity(self):
        if self.operation == 'sum':
            return 0
        elif self.operation == 'min':
            return float('inf')
        elif self.operation == 'max':
            return float('-inf')
        elif self.operation == 'gcd':
            return 0
    
    def _combine(self, left_val, right_val):
        if self.operation == 'sum':
            return left_val + right_val
        elif self.operation == 'min':
            return min(left_val, right_val)
        elif self.operation == 'max':
            return max(left_val, right_val)
        elif self.operation == 'gcd':
            import math
            return math.gcd(left_val, right_val)
    
    def _build(self, array, node, start, end):
        if start == end:
            self.tree[node] = array[start]
        else:
            mid = (start + end) // 2
            left_child = 2 * node + 1
            right_child = 2 * node + 2
            
            self._build(array, left_child, start, mid)
            self._build(array, right_child, mid + 1, end)
            
            self.tree[node] = self._combine(
                self.tree[left_child],
                self.tree[right_child]
            )
    
    def query(self, left, right):
        return self._query(0, 0, self.n - 1, left, right)
    
    def _query(self, node, start, end, left, right):
        # No overlap
        if start > right or end < left:
            return self.identity
        
        # Complete overlap
        if start >= left and end <= right:
            return self.tree[node]
        
        # Partial overlap
        mid = (start + end) // 2
        left_child = 2 * node + 1
        right_child = 2 * node + 2
        
        left_result = self._query(left_child, start, mid, left, right)
        right_result = self._query(right_child, mid + 1, end, left, right)
        
        return self._combine(left_result, right_result)
    
    def update(self, index, value):
        self._update(0, 0, self.n - 1, index, value)
    
    def _update(self, node, start, end, index, value):
        if start == end:
            self.tree[node] = value
        else:
            mid = (start + end) // 2
            left_child = 2 * node + 1
            right_child = 2 * node + 2
            
            if index <= mid:
                self._update(left_child, start, mid, index, value)
            else:
                self._update(right_child, mid + 1, end, index, value)
            
            self.tree[node] = self._combine(
                self.tree[left_child],
                self.tree[right_child]
            )

class LazySegmentTree:
    """Segment tree with lazy propagation for range updates"""
    def __init__(self, array):
        self.n = len(array)
        self.tree = [0] * (4 * self.n)
        self.lazy = [0] * (4 * self.n)
        self._build(array, 0, 0, self.n - 1)
    
    def _build(self, array, node, start, end):
        if start == end:
            self.tree[node] = array[start]
        else:
            mid = (start + end) // 2
            self._build(array, 2*node+1, start, mid)
            self._build(array, 2*node+2, mid+1, end)
            self.tree[node] = self.tree[2*node+1] + self.tree[2*node+2]
    
    def _push(self, node, start, end):
        """Push lazy updates down"""
        if self.lazy[node] != 0:
            self.tree[node] += (end - start + 1) * self.lazy[node]
            
            if start != end:
                self.lazy[2*node+1] += self.lazy[node]
                self.lazy[2*node+2] += self.lazy[node]
            
            self.lazy[node] = 0
    
    def range_update(self, left, right, value):
        """Add value to all elements in [left, right]"""
        self._range_update(0, 0, self.n - 1, left, right, value)
    
    def _range_update(self, node, start, end, left, right, value):
        self._push(node, start, end)
        
        if start > right or end < left:
            return
        
        if start >= left and end <= right:
            self.lazy[node] += value
            self._push(node, start, end)
            return
        
        mid = (start + end) // 2
        self._range_update(2*node+1, start, mid, left, right, value)
        self._range_update(2*node+2, mid+1, end, left, right, value)
        
        self._push(2*node+1, start, mid)
        self._push(2*node+2, mid+1, end)
        self.tree[node] = self.tree[2*node+1] + self.tree[2*node+2]
    
    def query(self, left, right):
        return self._query(0, 0, self.n - 1, left, right)
    
    def _query(self, node, start, end, left, right):
        if start > right or end < left:
            return 0
        
        self._push(node, start, end)
        
        if start >= left and end <= right:
            return self.tree[node]
        
        mid = (start + end) // 2
        return (self._query(2*node+1, start, mid, left, right) +
                self._query(2*node+2, mid+1, end, left, right))

class SegmentTree2D:
    """2D segment tree for matrix range queries"""
    def __init__(self, matrix):
        self.rows = len(matrix)
        self.cols = len(matrix[0])
        self.tree = [[0] * (4 * self.cols) for _ in range(4 * self.rows)]
        self._build_rows(matrix, 0, 0, self.rows - 1)
    
    def _build_rows(self, matrix, node, start, end):
        if start == end:
            self._build_cols(matrix[start], node, 0, 0, self.cols - 1)
        else:
            mid = (start + end) // 2
            self._build_rows(matrix, 2*node+1, start, mid)
            self._build_rows(matrix, 2*node+2, mid+1, end)
            
            # Merge column segment trees
            for col_node in range(4 * self.cols):
                self.tree[node][col_node] = (
                    self.tree[2*node+1][col_node] + 
                    self.tree[2*node+2][col_node]
                )
    
    def _build_cols(self, row, row_node, col_node, start, end):
        if start == end:
            self.tree[row_node][col_node] = row[start]
        else:
            mid = (start + end) // 2
            self._build_cols(row, row_node, 2*col_node+1, start, mid)
            self._build_cols(row, row_node, 2*col_node+2, mid+1, end)
            self.tree[row_node][col_node] = (
                self.tree[row_node][2*col_node+1] + 
                self.tree[row_node][2*col_node+2]
            )
```

### Test Cases
- Range sum queries on random arrays
- Verify against naive O(n) queries
- Point updates and re-query
- Range updates with lazy propagation
- 2D queries on matrices
- Stress test with 100K operations

### Deliverables
- `segment_tree.{py,go,zig}` - Basic implementation
- `lazy_segment_tree.{py,go,zig}` - With lazy propagation
- `segment_tree_2d.{py,go,zig}` - 2D version
- `applications.{py,go,zig}` - Range queries in practice
- `complexity_analysis.md` - Proof of O(log n) operations
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Correct range queries and updates
- Lazy propagation works efficiently
- Understanding of when segment trees are needed
- Can extend to other operations (min, max, GCD)

---

---

## Back to Module

[← Back to Module 5: Advanced Data Structures](../module_05/README.md)

## Navigation

- [Previous Lab](./lab_5_2.md) (if exists)
- [Next Lab](./lab_5_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 5, Lab 3 of 6*
