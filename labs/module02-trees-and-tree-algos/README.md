# Module 2: Trees & Tree Algorithms

**Duration**: 4 weeks  
**Difficulty**: Medium to Very Hard

## Overview

Tree structures model hierarchical relationships and enable efficient search, insertion, and deletion operations. This module takes you from basic binary search trees to self-balancing trees and advanced augmented structures.

**Why this matters**: Trees are everywhere - filesystems, databases, compilers, game trees, decision trees. Understanding tree algorithms deeply is essential for systems programming and algorithm design.

## Learning Objectives

- Implement binary search trees with all standard operations
- Master tree traversals (recursive and iterative)
- Build self-balancing trees (AVL or Red-Black)
- Understand augmented data structures
- Recognize when tree structures are optimal

## Topics Covered

- Binary search trees (BST)
- Tree traversals and their applications
- Self-balancing trees (AVL, Red-Black)
- B-trees for disk-based storage
- Augmented data structures (order statistics, intervals)

---

## Lab 2.1: Binary Search Tree Fundamentals

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

## Lab 2.2: Self-Balancing Trees (AVL or Red-Black)

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

## Lab 2.3: Augmented Data Structures

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

## Lab 2.4: B-Tree for Disk Storage

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

## Lab 2.5: Tree Visualizer

**Duration**: 4-6 hours  
**Difficulty**: Medium

### Objectives
- Build tools to visualize tree structures
- Aid debugging and understanding
- Support multiple tree types

### Requirements

1. **Create visualizations for**:
   - BST, AVL, Red-Black trees
   - B-trees
   - Expression trees
   - Any custom tree structure

2. **Visualization features**:
   - Text-based ASCII art
   - Graphical rendering (DOT/Graphviz or similar)
   - Step-by-step animation of operations
   - Highlight nodes during operations

3. **Interactive features**:
   - Click to insert/delete nodes
   - Show rotation/rebalancing steps
   - Display node metadata (height, color, size, etc.)

### Implementation Details

```python
def visualize_tree(root, format="ascii"):
    if format == "ascii":
        return generate_ascii_tree(root)
    elif format == "graphviz":
        return generate_dot_graph(root)
    elif format == "interactive":
        launch_gui(root)

def animate_operation(tree, operation, *args):
    # Show before state
    # Execute operation step-by-step
    # Show after state
```

### Test Cases
- Visualize small trees (easy to verify by hand)
- Visualize large trees (stress test layout algorithm)
- Animate insertion into AVL tree (show rotations)

### Deliverables
- `tree_visualizer.{py,go,zig}` - Core visualization library
- `cli_visualizer.{py,go,zig}` - Command-line interface
- `gui_visualizer.{py,go,zig}` - Graphical interface (optional)
- `examples/` - Visualizations of various tree types

### Success Criteria
- Clear, readable visualizations
- Supports all tree types implemented in module
- Useful debugging tool for subsequent labs

---

## Lab 2.6: Comparative Tree Performance

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Benchmark all tree implementations
- Identify when each tree type is optimal
- Understand practical vs. theoretical performance

### Requirements

1. **Benchmark all tree types** (BST, AVL/RB, B-tree):
   - Sequential insertions
   - Random insertions
   - Search performance
   - Deletion performance
   - Mixed workloads

2. **Measure**:
   - Operation throughput (ops/sec)
   - Height of resulting tree
   - Memory usage
   - Cache efficiency

3. **Create decision matrix**:
   - When to use unbalanced BST
   - When to use AVL vs. Red-Black
   - When to use B-tree

### Deliverables
- `tree_benchmarks.{py,go,zig}` - Comprehensive benchmark suite
- `performance_report.md` - Detailed analysis
- `decision_guide.md` - Guidelines for choosing tree type
- Performance graphs and tables

### Success Criteria
- Clear performance differences demonstrated
- Understanding of tradeoffs
- Can justify tree choice for given scenario

---

## Module Resources

### CLRS References
- Chapter 12: Binary Search Trees
- Chapter 13: Red-Black Trees
- Chapter 14: Augmenting Data Structures
- Chapter 18: B-Trees

### Additional Reading
- "Introduction to Algorithms" sections on AVL trees
- Database systems textbooks for B-tree implementations
- Persistent data structures papers

### Key Takeaways

By the end of this module, you should:
1. Understand tree rotations and rebalancing deeply
2. Be able to augment any tree with additional information
3. Know when self-balancing is worth the complexity
4. Recognize tree-based solutions to problems
5. Have built robust visualization and testing tools

### Next Module

**Module 3: Heaps & Priority Queues** - You'll implement heap-based structures and understand their applications in sorting and graph algorithms.

---

## Tips for Success

1. **Visualize everything** - Trees are visual; draw them out
2. **Test rotations extensively** - They're easy to get wrong
3. **Start with BST** - Get comfortable before tackling balanced trees
4. **Use your visualizer** - It'll catch bugs quickly
5. **Prove invariants** - Formally verify balance properties

## Common Pitfalls

- **Incorrect rotation logic** - Off by one in pointer updates
- **Forgetting to update augmented data** - Size, height, etc.
- **Deletion edge cases** - Two-child deletion is tricky
- **Memory leaks** - Especially during tree restructuring
- **Not testing degenerate cases** - Single node, two nodes, etc.
