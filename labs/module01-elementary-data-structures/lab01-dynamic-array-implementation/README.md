# Lab 1.1: Dynamic Array Implementation

**Module**: Module 1 - Elementary Data Structures  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 1.1: Dynamic Array Implementation

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement a dynamic array from scratch
- Understand amortized analysis deeply
- Explore growth factor tradeoffs

### Requirements

1. **Implement a complete dynamic array** with:
   - `push(value)` - O(1) amortized
   - `pop()` - O(1)
   - `get(index)` - O(1)
   - `set(index, value)` - O(1)
   - `insert(index, value)` - O(n)
   - `delete(index)` - O(n)
   - `size()` - O(1)
   - `capacity()` - O(1)

2. **Experiment with different growth strategies**:
   - Factor of 2 (double when full)
   - Factor of 1.5
   - Factor of 3
   - Additive growth (add constant amount)
   - Hybrid strategies

3. **Implement shrinking strategies**:
   - Shrink when 1/4 full (to avoid thrashing)
   - Shrink when 1/2 full
   - Never shrink
   - Compare space-time tradeoffs

4. **Perform amortized analysis**:
   - Accounting method
   - Potential method
   - Aggregate analysis
   - Prove O(1) amortized insertion

### Implementation Details

```python
class DynamicArray:
    def __init__(self, growth_factor=2.0):
        self.data = allocate(initial_capacity)
        self.size = 0
        self.capacity = initial_capacity
        self.growth_factor = growth_factor
    
    def push(self, value):
        if self.size == self.capacity:
            self._resize(self.capacity * self.growth_factor)
        self.data[self.size] = value
        self.size += 1
```

### Test Cases
- Insert 1M elements, verify O(n) total time
- Test all edge cases (empty, one element, full capacity)
- Memory leak detection
- Stress test with random operations

### Deliverables
- `dynamic_array.{py,go,zig}` - Full implementation with multiple strategies
- `amortized_analysis.md` - Mathematical proofs of amortized bounds
- `benchmarks.{py,go,zig}` - Performance comparison of growth factors
- `visualization.{py,go,zig}` - Visualize capacity changes over time
- `test_suite.{py,go,zig}` - Comprehensive unit tests

### Success Criteria
- All operations have correct time complexity
- Growth factor of 2 shows best amortized performance
- Shrinking strategy prevents thrashing
- Can explain amortized analysis using all three methods

---

---

## Back to Module

[← Back to Module 1: Elementary Data Structures](../module_01/README.md)

## Navigation

- [Previous Lab](./lab_1_0.md) (if exists)
- [Next Lab](./lab_1_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 1, Lab 1 of 5*
