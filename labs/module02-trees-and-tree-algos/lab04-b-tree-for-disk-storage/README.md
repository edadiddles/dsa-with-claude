# Lab 2.4: B-Tree for Disk Storage

**Module**: Module 2 - Trees & Tree Algorithms  
**Duration**: 6-8 hours  
**Difficulty**: Hard

---

Lab 2.4: B-Tree for Disk Storage

**Duration**: 6-8 hours  
**Difficulty**: Hard

### Objectives
- Implement B-tree optimized for disk I/O
- Understand multi-level indexing
- Simulate disk access patterns

### Requirements

1. **Implement B-tree with configurable order** (minimum degree t):
   - Each node has ≤ 2t-1 keys
   - Each internal node has ≤ 2t children
   - Root has ≥ 1 key
   - All other nodes have ≥ t-1 keys

2. **Operations**:
   - `search(key)` - O(log_t n) disk accesses
   - `insert(key)` - O(log_t n) disk accesses
   - `delete(key)` - O(log_t n) disk accesses
   - Split and merge nodes as needed

3. **Simulate disk I/O**:
   - Track disk reads and writes
   - Measure total I/O cost
   - Compare with BST simulation

4. **Find optimal t for different scenarios**:
   - Small keys (integers)
   - Large keys (strings)
   - Different page sizes

### Implementation Details

```python
class BTreeNode:
    def __init__(self, t, leaf=True):
        self.t = t  # Minimum degree
        self.keys = []
        self.children = []
        self.leaf = leaf
        self.n = 0  # Number of keys
    
    def split_child(self, i, y):
        # Split full child y of node at index i
```

### Test Cases
- Insert sequential keys
- Insert random keys
- Delete keys requiring merges
- Verify all B-tree properties
- Measure disk I/O for different t values

### Deliverables
- `btree.{py,go,zig}` - Complete B-tree implementation
- `disk_simulator.{py,go,zig}` - Simulated disk I/O tracking
- `optimal_order_analysis.md` - Finding optimal t
- `comparison.md` - B-tree vs. BST I/O comparison

### Success Criteria
- B-tree properties maintained
- I/O cost is O(log_t n)
- Understanding of why B-trees are disk-friendly

---

---

## Back to Module

[← Back to Module 2: Trees & Tree Algorithms](../module_02/README.md)

## Navigation

- [Previous Lab](./lab_2_3.md) (if exists)
- [Next Lab](./lab_2_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 2, Lab 4 of 6*
