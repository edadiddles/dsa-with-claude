# Lab 2.3: Augmented Data Structures

**Module**: Module 2 - Trees & Tree Algorithms  
**Duration**: 4-6 hours  
**Difficulty**: Medium-Hard

---

Lab 2.3: Augmented Data Structures

**Duration**: 4-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Extend trees with additional information
- Maintain augmented data during updates
- Solve new problems efficiently with augmentation

### Requirements

1. **Implement Order Statistic Tree**:
   - Augment each node with subtree size
   - `select(i)` - Find ith smallest element in O(log n)
   - `rank(value)` - Find rank of value in O(log n)
   - Maintain size during insertions/deletions

2. **Implement Interval Tree**:
   - Store intervals [low, high] in nodes
   - Augment with max endpoint in subtree
   - `search_overlap(interval)` - Find overlapping interval
   - `all_overlaps(interval)` - Find all overlapping intervals

3. **Implement Range Tree for range sum queries**:
   - Augment with subtree sum
   - `range_sum(min, max)` - Sum of values in range

### Implementation Details

```python
class OrderStatNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None
        self.size = 1  # Size of subtree
    
def select(node, i):
    # Find ith smallest element
    left_size = get_size(node.left)
    if i == left_size:
        return node.value
    elif i < left_size:
        return select(node.left, i)
    else:
        return select(node.right, i - left_size - 1)
```

### Test Cases
- Order statistic tree: Verify select and rank are inverses
- Interval tree: Test with overlapping and non-overlapping intervals
- Range tree: Verify sum against naive implementation

### Deliverables
- `order_statistic_tree.{py,go,zig}` - OS tree implementation
- `interval_tree.{py,go,zig}` - Interval tree implementation
- `range_tree.{py,go,zig}` - Range sum tree
- `augmentation_guide.md` - How to augment data structures
- `applications.md` - Real-world use cases

### Success Criteria
- All augmented operations work correctly
- Augmented data maintained during updates
- Understanding of how to design new augmentations

---

---

## Back to Module

[← Back to Module 2: Trees & Tree Algorithms](../module_02/README.md)

## Navigation

- [Previous Lab](./lab_2_2.md) (if exists)
- [Next Lab](./lab_2_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 2, Lab 3 of 6*
