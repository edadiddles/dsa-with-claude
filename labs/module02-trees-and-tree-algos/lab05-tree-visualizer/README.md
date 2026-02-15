# Lab 2.5: Tree Visualizer

**Module**: Module 2 - Trees & Tree Algorithms  
**Duration**: 4-6 hours  
**Difficulty**: Medium

---

Lab 2.5: Tree Visualizer

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

---

## Back to Module

[← Back to Module 2: Trees & Tree Algorithms](../module_02/README.md)

## Navigation

- [Previous Lab](./lab_2_4.md) (if exists)
- [Next Lab](./lab_2_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 2, Lab 5 of 6*
