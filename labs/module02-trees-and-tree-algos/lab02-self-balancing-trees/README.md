# Lab 2.2: Self-Balancing Trees (AVL or Red-Black)

**Module**: Module 2 - Trees & Tree Algorithms  
**Duration**: 8-12 hours  
**Difficulty**: Very Hard

---

Lab 2.2: Self-Balancing Trees (AVL or Red-Black)

**Duration**: 8-12 hours  
**Difficulty**: Very Hard

### Objectives
- Implement a complete self-balancing tree
- Understand rotation mechanics
- Guarantee O(log n) operations

### Requirements

1. **Choose AVL or Red-Black tree** (or implement both!)

2. **For AVL Tree**:
   - Track height in each node
   - Implement rotations: left, right, left-right, right-left
   - Rebalance after insertion and deletion
   - Maintain balance factor (-1, 0, +1)

3. **For Red-Black Tree**:
   - Track color in each node
   - Implement rotations and recoloring
   - Maintain RB properties after all operations
   - Handle all deletion cases

4. **Compare performance against unbalanced BST**:
   - Worst-case input (sorted sequence)
   - Random inputs
   - Height comparison

### Implementation Details (AVL)

```python
class AVLNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None
        self.height = 1
    
def get_balance(node):
    return get_height(node.left) - get_height(node.right)

def rotate_right(y):
    x = y.left
    T2 = x.right
    x.right = y
    y.left = T2
    y.height = 1 + max(get_height(y.left), get_height(y.right))
    x.height = 1 + max(get_height(x.left), get_height(x.right))
    return x
```

### Test Cases
- Insert sorted sequence: should remain balanced
- Insert reverse sorted: should remain balanced
- Random insertions and deletions
- Verify height is O(log n)
- Stress test with 1M operations

### Deliverables
- `avl_tree.{py,go,zig}` or `red_black_tree.{py,go,zig}` - Full implementation
- `rotation_visualizer.{py,go,zig}` - Show rotations step-by-step
- `balance_analysis.md` - Height analysis and proofs
- `performance_comparison.{py,go,zig}` - BST vs. balanced tree benchmarks
- `test_suite.{py,go,zig}` - Exhaustive tests

### Success Criteria
- Tree remains balanced after all operations
- Height is guaranteed O(log n)
- Can explain each rotation case
- Performance matches theoretical predictions

---

---

## Back to Module

[← Back to Module 2: Trees & Tree Algorithms](../module_02/README.md)

## Navigation

- [Previous Lab](./lab_2_1.md) (if exists)
- [Next Lab](./lab_2_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 2, Lab 2 of 6*
