# Lab 10.4: Amortized Analysis Techniques

**Module**: Module 10 - Advanced Algorithm Techniques  
**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

---

Lab 10.4: Amortized Analysis Techniques

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Master amortized analysis techniques
- Apply aggregate, accounting, and potential methods
- Analyze data structures with amortized complexity
- Understand difference from average-case analysis

### Requirements

1. **Learn three methods**:
   - **Aggregate analysis**
   - **Accounting method**
   - **Potential method**

2. **Analyze data structures**:
   - Dynamic array resizing
   - Stack with MultiPop
   - Binary counter
   - Splay tree (simplified)
   - Union-Find with path compression

3. **Applications**:
   - Justify real-world data structure performance
   - Design new amortized structures

### Implementation Details

```python
# DYNAMIC ARRAY - Amortized Analysis

class DynamicArray:
    """
    Dynamic array with amortized O(1) append.
    
    Amortized analysis (aggregate method):
    - n appends cause O(log n) resizes
    - Each resize copies O(n) elements
    - Total work: n + n/2 + n/4 + ... + 1 = O(n)
    - Amortized cost per append: O(n)/n = O(1)
    """
    def __init__(self):
        self.array = [None] * 1
        self.size = 0
        self.capacity = 1
        self.resize_count = 0
        self.copy_count = 0
    
    def append(self, value):
        if self.size == self.capacity:
            self._resize()
        
        self.array[self.size] = value
        self.size += 1
    
    def _resize(self):
        """Double capacity"""
        self.capacity *= 2
        new_array = [None] * self.capacity
        
        for i in range(self.size):
            new_array[i] = self.array[i]
            self.copy_count += 1
        
        self.array = new_array
        self.resize_count += 1
    
    def amortized_cost_analysis(self):
        """
        Return amortized analysis metrics.
        """
        total_operations = self.size
        total_copies = self.copy_count
        
        return {
            'appends': self.size,
            'resizes': self.resize_count,
            'copies': self.copy_count,
            'amortized_cost': total_copies / total_operations if total_operations > 0 else 0
        }

# STACK WITH MULTIPOP

class StackWithMultiPop:
    """
    Stack supporting MultiPop(k) - pop k elements.
    
    Amortized analysis (accounting method):
    - Charge $2 per Push: $1 for push, $1 credit for future Pop
    - Pop uses stored credit: $0 amortized
    - MultiPop(k) uses k stored credits: $0 amortized
    - Amortized cost per operation: O(1)
    """
    def __init__(self):
        self.stack = []
        self.credits = 0  # Track credits for accounting method
        self.total_charge = 0
    
    def push(self, value):
        self.stack.append(value)
        self.credits += 1  # Store one credit per element
        self.total_charge += 2  # Charge 2 for accounting
    
    def pop(self):
        if self.stack:
            self.credits -= 1  # Use stored credit
            return self.stack.pop()
        return None
    
    def multi_pop(self, k):
        result = []
        for _ in range(min(k, len(self.stack))):
            result.append(self.pop())
        return result
    
    def amortized_analysis(self):
        return {
            'total_charged': self.total_charge,
            'operations': len(self.stack) + self.total_charge // 2,
            'amortized_per_op': self.total_charge / (len(self.stack) + self.total_charge // 2) if self.total_charge > 0 else 0
        }

# BINARY COUNTER

class BinaryCounter:
    """
    Binary counter with increment operation.
    
    Amortized analysis (aggregate method):
    - n increments flip O(n) bits total
    - Bit i flips every 2^i increments
    - Total flips: n + n/2 + n/4 + ... = 2n
    - Amortized cost per increment: O(1)
    """
    def __init__(self, bits=32):
        self.counter = [0] * bits
        self.flip_count = 0
    
    def increment(self):
        i = 0
        
        while i < len(self.counter) and self.counter[i] == 1:
            self.counter[i] = 0
            self.flip_count += 1
            i += 1
        
        if i < len(self.counter):
            self.counter[i] = 1
            self.flip_count += 1
    
    def value(self):
        """Get counter value"""
        result = 0
        for i in range(len(self.counter)):
            result += self.counter[i] * (2 ** i)
        return result
    
    def amortized_analysis(self, increments):
        """Analyze n increments"""
        return {
            'increments': increments,
            'total_flips': self.flip_count,
            'amortized_flips': self.flip_count / increments if increments > 0 else 0
        }

# POTENTIAL METHOD

class AmortizedAnalyzer:
    """Tools for potential method analysis"""
    
    @staticmethod
    def potential_dynamic_array(array):
        """
        Potential function for dynamic array.
        Φ(D) = 2*size - capacity
        
        This ensures enough potential to pay for resize.
        """
        return 2 * array.size - array.capacity
    
    @staticmethod
    def amortized_cost_append(array, potential_before):
        """
        Calculate amortized cost of append using potential method.
        
        Amortized cost = actual cost + ΔΦ
        
        Without resize: actual = 1, ΔΦ = 2, amortized = 3
        With resize: actual = 1 + size, ΔΦ = -size + 2, amortized = 3
        """
        actual_cost = 1
        if array.size == array.capacity:
            actual_cost += array.size  # Resize cost
        
        array.append(None)  # Hypothetical append
        
        potential_after = AmortizedAnalyzer.potential_dynamic_array(array)
        delta_potential = potential_after - potential_before
        
        amortized_cost = actual_cost + delta_potential
        
        return {
            'actual': actual_cost,
            'delta_potential': delta_potential,
            'amortized': amortized_cost
        }
    
    @staticmethod
    def potential_binary_counter(counter):
        """
        Potential function for binary counter.
        Φ(D) = number of 1s in counter
        
        Each increment: ΔΦ = -k + 1, where k = number of 1s flipped to 0
        Actual cost = k + 1 (flip k zeros and one zero to one)
        Amortized = (k+1) + (-k+1) = 2
        """
        return sum(counter.counter)

# SPLAY TREE SIMPLIFIED

class SplayTreeNode:
    def __init__(self, key):
        self.key = key
        self.left = None
        self.right = None

class SplayTree:
    """
    Simplified splay tree for amortized analysis demonstration.
    
    Amortized analysis (potential method):
    - Potential = sum of log(size of subtrees)
    - Splay operation: O(log n) amortized
    - Sequence of m operations: O(m log n)
    """
    def __init__(self):
        self.root = None
    
    def splay(self, node):
        """Splay node to root (simplified)"""
        # Rotate node to root using rotations
        # Each rotation changes potential
        # Amortized cost: O(log n)
        pass
    
    def potential(self, node):
        """Potential of tree"""
        if node is None:
            return 0
        
        size = self._size(node)
        return math.log2(size) if size > 0 else 0
    
    def _size(self, node):
        """Size of subtree"""
        if node is None:
            return 0
        return 1 + self._size(node.left) + self._size(node.right)

# COMPARISON OF METHODS

def compare_analysis_methods():
    """
    Compare three amortized analysis methods:
    
    1. Aggregate Analysis:
       - Compute total cost of n operations
       - Divide by n
       - Pros: Simple, direct
       - Cons: Need to analyze entire sequence
    
    2. Accounting Method:
       - Assign "charge" to each operation
       - Some operations pay for future operations
       - Maintain credit invariant
       - Pros: Intuitive, per-operation view
       - Cons: Need to design good charging scheme
    
    3. Potential Method:
       - Define potential function Φ
       - Amortized cost = actual cost + ΔΦ
       - Pros: Mathematical, rigorous
       - Cons: Need to design good potential function
    
    All three give same amortized bound!
    """
    pass

# AMORTIZED VS AVERAGE CASE

def amortized_vs_average():
    """
    Key differences:
    
    Amortized Analysis:
    - Worst-case analysis over sequence of operations
    - No probability involved
    - Guarantees: worst sequence has this cost
    
    Average-Case Analysis:
    - Expected cost for random input
    - Probability distribution over inputs
    - Guarantees: expected cost over random inputs
    
    Example: Dynamic array
    - Amortized: O(1) per append (worst sequence)
    - Average: O(1) per append (random order)
    - Both happen to be same, but for different reasons!
    """
    pass
```

### Test Cases
- Dynamic array with many appends
- Stack with mixed Push/Pop/MultiPop
- Binary counter with many increments
- Verify amortized bounds empirically

### Deliverables
- `dynamic_array_analysis.{py,go,zig}` - With amortized metrics
- `stack_multipop.{py,go,zig}` - Accounting method
- `binary_counter.{py,go,zig}` - Aggregate analysis
- `potential_method.{py,go,zig}` - Potential function tools
- `amortized_analysis_guide.md` - When to use each method
- `comparison.md` - Three methods compared
- `test_suite.{py,go,zig}` - Verify bounds

### Success Criteria
- Understanding of all three methods
- Can apply appropriate method to new problems
- Distinction from average-case analysis clear
- Can design potential functions

---

---

## Back to Module

[← Back to Module 10: Advanced Algorithm Techniques](../module_10/README.md)

## Navigation

- [Previous Lab](./lab_10_3.md) (if exists)
- [Next Lab](./lab_10_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 10, Lab 4 of 5*
