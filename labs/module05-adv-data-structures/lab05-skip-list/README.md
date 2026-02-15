# Lab 5.5: Skip List

**Module**: Module 5 - Advanced Data Structures  
**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

---

Lab 5.5: Skip List

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Implement probabilistic balanced structure
- Understand randomized data structures
- Compare with balanced BST

### Requirements

1. **Implement skip list with**:
   - `insert(value)` - O(log n) expected
   - `search(value)` - O(log n) expected
   - `delete(value)` - O(log n) expected
   - Random level generation (p = 0.5 typically)

2. **Features**:
   - Configurable probability p
   - Maximum level limit
   - Iterator support (in-order traversal)

3. **Analysis**:
   - Measure actual height distribution
   - Compare search cost with BST
   - Study effect of different p values

### Implementation Details

```python
import random

class SkipNode:
    def __init__(self, value, level):
        self.value = value
        self.forward = [None] * (level + 1)

class SkipList:
    def __init__(self, max_level=16, p=0.5):
        self.max_level = max_level
        self.p = p
        self.header = SkipNode(None, max_level)
        self.level = 0
    
    def random_level(self):
        level = 0
        while random.random() < self.p and level < self.max_level:
            level += 1
        return level
    
    def search(self, value):
        current = self.header
        
        for i in range(self.level, -1, -1):
            while current.forward[i] and current.forward[i].value < value:
                current = current.forward[i]
        
        current = current.forward[0]
        return current and current.value == value
    
    def insert(self, value):
        update = [None] * (self.max_level + 1)
        current = self.header
        
        # Find position to insert
        for i in range(self.level, -1, -1):
            while current.forward[i] and current.forward[i].value < value:
                current = current.forward[i]
            update[i] = current
        
        current = current.forward[0]
        
        # Check if already exists
        if current and current.value == value:
            return  # Duplicate
        
        # Generate random level for new node
        new_level = self.random_level()
        
        if new_level > self.level:
            for i in range(self.level + 1, new_level + 1):
                update[i] = self.header
            self.level = new_level
        
        new_node = SkipNode(value, new_level)
        
        # Insert node
        for i in range(new_level + 1):
            new_node.forward[i] = update[i].forward[i]
            update[i].forward[i] = new_node
    
    def delete(self, value):
        update = [None] * (self.max_level + 1)
        current = self.header
        
        # Find position
        for i in range(self.level, -1, -1):
            while current.forward[i] and current.forward[i].value < value:
                current = current.forward[i]
            update[i] = current
        
        current = current.forward[0]
        
        if current and current.value == value:
            # Remove node
            for i in range(self.level + 1):
                if update[i].forward[i] != current:
                    break
                update[i].forward[i] = current.forward[i]
            
            # Update level
            while self.level > 0 and self.header.forward[self.level] is None:
                self.level -= 1
    
    def __iter__(self):
        current = self.header.forward[0]
        while current:
            yield current.value
            current = current.forward[0]
    
    def range_query(self, start, end):
        """Find all values in [start, end]"""
        result = []
        current = self.header
        
        # Navigate to start
        for i in range(self.level, -1, -1):
            while current.forward[i] and current.forward[i].value < start:
                current = current.forward[i]
        
        current = current.forward[0]
        
        # Collect values in range
        while current and current.value <= end:
            result.append(current.value)
            current = current.forward[0]
        
        return result
    
    def get_height(self):
        return self.level + 1
    
    def measure_level_distribution(self):
        """Measure how many nodes at each level"""
        counts = [0] * (self.level + 1)
        current = self.header.forward[0]
        
        while current:
            for i in range(len(current.forward)):
                if current.forward[i] is not None or i == 0:
                    counts[i] += 1
            current = current.forward[0]
        
        return counts

def analyze_skip_list_performance(p_values, sizes):
    """Analyze performance with different p values"""
    results = {}
    
    for p in p_values:
        results[p] = {}
        for size in sizes:
            skip_list = SkipList(p=p)
            values = list(range(size))
            random.shuffle(values)
            
            # Measure insertion time
            import time
            start = time.time()
            for v in values:
                skip_list.insert(v)
            insert_time = time.time() - start
            
            # Measure search time
            start = time.time()
            for v in values:
                skip_list.search(v)
            search_time = time.time() - start
            
            results[p][size] = {
                'height': skip_list.get_height(),
                'insert_time': insert_time,
                'search_time': search_time,
                'level_dist': skip_list.measure_level_distribution()
            }
    
    return results
```

### Test Cases
- Insert random values, verify search
- Delete values, verify structure
- Measure level distribution
- Compare with AVL/Red-Black tree
- Test different p values (0.25, 0.5, 0.75)

### Deliverables
- `skip_list.{py,go,zig}` - Full implementation
- `probability_analysis.md` - Study of p values
- `height_distribution.{py,go,zig}` - Measure actual heights
- `comparison.md` - Skip list vs. BST
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Correct operations
- Understanding of probabilistic guarantees
- Knowledge of skip list advantages (simpler than balanced BST, good cache locality)
- Can explain expected height formula

---

---

## Back to Module

[← Back to Module 5: Advanced Data Structures](../module_05/README.md)

## Navigation

- [Previous Lab](./lab_5_4.md) (if exists)
- [Next Lab](./lab_5_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 5, Lab 5 of 6*
