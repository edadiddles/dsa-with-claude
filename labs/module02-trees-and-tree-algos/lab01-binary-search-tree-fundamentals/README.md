# Lab 2.1: Binary Search Tree Fundamentals

**Module**: Module 2 - Trees & Tree Algorithms  
**Duration**: 5-6 hours  
**Difficulty**: Medium

---

Lab 2.1: Binary Search Tree Fundamentals

**Duration**: 5-6 hours  
**Difficulty**: Medium

### Objectives
- Implement complete BST with all operations
- Master tree traversals (inorder, preorder, postorder, level-order)
- Understand BST property maintenance

### Requirements

1. **Implement BST with operations**:
   - `insert(value)` - O(h)
   - `search(value)` - O(h)
   - `delete(value)` - O(h) (including successor/predecessor)
   - `minimum()` / `maximum()` - O(h)
   - `height()` - O(n)
   - `size()` - O(n) or O(1) with augmentation

2. **Implement all traversals**:
   - Recursive inorder, preorder, postorder
   - Iterative versions using explicit stack
   - Level-order using queue
   - Morris traversal (O(1) space)

3. **Additional operations**:
   - `floor(value)` - Largest element ≤ value
   - `ceiling(value)` - Smallest element ≥ value
   - `range_query(min, max)` - All elements in range
   - `kth_smallest(k)` - Find kth element

4. **Validate BST property** with comprehensive tests

### Implementation Details

```python
class BSTNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

class BST:
    def __init__(self):
        self.root = None
    
    def insert(self, value):
        self.root = self._insert_recursive(self.root, value)
    
    def _insert_recursive(self, node, value):
        if node is None:
            return BSTNode(value)
        if value < node.value:
            node.left = self._insert_recursive(node.left, value)
        else:
            node.right = self._insert_recursive(node.right, value)
        return node
```

### Test Cases
- Insert sorted sequence: should create degenerate tree
- Insert random sequence: should create balanced-ish tree
- Delete nodes with 0, 1, 2 children
- Verify BST property after all operations
- Test with duplicates

### Deliverables
- `bst.{py,go,zig}` - Complete BST implementation
- `traversals.{py,go,zig}` - All traversal implementations
- `visualizer.{py,go,zig}` - Tree visualization tool
- `test_suite.{py,go,zig}` - Comprehensive tests
- `complexity_analysis.md` - Analysis of height impact on performance

### Success Criteria
- All operations maintain BST property
- Can explain deletion algorithm thoroughly
- Morris traversal works correctly
- Visualization helps debugging

---

---

## Back to Module

[← Back to Module 2: Trees & Tree Algorithms](../module_02/README.md)

## Navigation

- [Previous Lab](./lab_2_0.md) (if exists)
- [Next Lab](./lab_2_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 2, Lab 1 of 6*
