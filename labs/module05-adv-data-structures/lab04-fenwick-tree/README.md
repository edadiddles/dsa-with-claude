# Lab 5.4: Fenwick Tree (Binary Indexed Tree)

**Module**: Module 5 - Advanced Data Structures  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 5.4: Fenwick Tree (Binary Indexed Tree)

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Implement Fenwick tree for prefix sums
- Understand bit manipulation tricks
- Compare with segment tree

### Requirements

1. **Implement Fenwick tree with**:
   - `build(array)` - O(n) construction
   - `prefix_sum(index)` - Sum of elements [0, index]
   - `range_sum(left, right)` - Sum of elements [left, right]
   - `update(index, delta)` - Add delta to element at index

2. **Understand the bit manipulation**:
   - `i & (-i)` gives lowest set bit
   - Parent: `i + (i & -i)`
   - Next: `i - (i & -i)`

3. **Extensions**:
   - 2D Fenwick tree (for matrix prefix sums)
   - Range update, point query variant
   - Binary search on Fenwick tree

### Implementation Details

```python
class FenwickTree:
    def __init__(self, n):
        self.n = n
        self.tree = [0] * (n + 1)  # 1-indexed
    
    def update(self, index, delta):
        """Add delta to element at index (0-indexed input)"""
        index += 1  # Convert to 1-indexed
        while index <= self.n:
            self.tree[index] += delta
            index += index & (-index)  # Move to parent
    
    def prefix_sum(self, index):
        """Sum of elements [0, index] (0-indexed input)"""
        index += 1  # Convert to 1-indexed
        result = 0
        while index > 0:
            result += self.tree[index]
            index -= index & (-index)  # Move to next node
        return result
    
    def range_sum(self, left, right):
        """Sum of elements [left, right] (0-indexed)"""
        if left == 0:
            return self.prefix_sum(right)
        return self.prefix_sum(right) - self.prefix_sum(left - 1)
    
    def build(self, array):
        """Build from array in O(n)"""
        for i, val in enumerate(array):
            self.update(i, val)
    
    def build_optimized(self, array):
        """O(n) build by propagating upward"""
        self.tree = [0] + array[:]  # 1-indexed with values
        
        for i in range(1, self.n + 1):
            parent = i + (i & -i)
            if parent <= self.n:
                self.tree[parent] += self.tree[i]

class FenwickTree2D:
    def __init__(self, rows, cols):
        self.rows = rows
        self.cols = cols
        self.tree = [[0] * (cols + 1) for _ in range(rows + 1)]
    
    def update(self, row, col, delta):
        """Add delta to element at (row, col)"""
        row += 1
        col += 1
        
        i = row
        while i <= self.rows:
            j = col
            while j <= self.cols:
                self.tree[i][j] += delta
                j += j & (-j)
            i += i & (-i)
    
    def prefix_sum(self, row, col):
        """Sum of rectangle from (0,0) to (row, col)"""
        row += 1
        col += 1
        result = 0
        
        i = row
        while i > 0:
            j = col
            while j > 0:
                result += self.tree[i][j]
                j -= j & (-j)
            i -= i & (-i)
        
        return result
    
    def range_sum(self, r1, c1, r2, c2):
        """Sum of rectangle from (r1,c1) to (r2,c2)"""
        return (self.prefix_sum(r2, c2) - 
                self.prefix_sum(r1-1, c2) - 
                self.prefix_sum(r2, c1-1) + 
                self.prefix_sum(r1-1, c1-1))

class RangeUpdateFenwick:
    """Fenwick tree for range updates and point queries"""
    def __init__(self, n):
        self.n = n
        self.tree = [0] * (n + 1)
    
    def range_update(self, left, right, delta):
        """Add delta to all elements in [left, right]"""
        self._update(left, delta)
        self._update(right + 1, -delta)
    
    def _update(self, index, delta):
        index += 1
        while index <= self.n:
            self.tree[index] += delta
            index += index & (-index)
    
    def point_query(self, index):
        """Get value at index"""
        index += 1
        result = 0
        while index > 0:
            result += self.tree[index]
            index -= index & (-index)
        return result

def binary_search_fenwick(fenwick, target):
    """Find smallest index where prefix_sum >= target"""
    # Works if all elements are non-negative
    index = 0
    power = 1
    
    # Find highest power of 2 <= n
    while power * 2 <= fenwick.n:
        power *= 2
    
    current_sum = 0
    
    while power > 0:
        if index + power <= fenwick.n and current_sum + fenwick.tree[index + power] < target:
            index += power
            current_sum += fenwick.tree[index]
        power //= 2
    
    return index  # 0-indexed result
```

### Test Cases
- Build from array, verify prefix sums
- Random updates and queries
- Compare with naive O(n) prefix sum
- 2D matrix operations
- Verify correctness against segment tree
- Performance benchmarks

### Deliverables
- `fenwick_tree.{py,go,zig}` - 1D implementation
- `fenwick_tree_2d.{py,go,zig}` - 2D implementation
- `range_update_fenwick.{py,go,zig}` - Range update variant
- `bit_manipulation_explanation.md` - Explain the magic
- `fenwick_vs_segment.md` - Comparison
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Correct prefix sum and update operations
- Understanding of bit manipulation
- Knowledge of when Fenwick tree is preferable to segment tree (simpler, less memory)
- Can explain the `i & (-i)` trick

---

---

## Back to Module

[← Back to Module 5: Advanced Data Structures](../module_05/README.md)

## Navigation

- [Previous Lab](./lab_5_3.md) (if exists)
- [Next Lab](./lab_5_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 5, Lab 4 of 6*
