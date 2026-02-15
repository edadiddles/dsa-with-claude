# Lab 1.2: Hash Table Deep Dive

**Module**: Module 1 - Elementary Data Structures  
**Duration**: 6-8 hours  
**Difficulty**: Hard

---

Lab 1.2: Hash Table Deep Dive

**Duration**: 6-8 hours  
**Difficulty**: Hard

### Objectives
- Implement hash tables with different collision resolution strategies
- Design and test hash functions
- Understand load factor and rehashing

### Requirements

1. **Implement three collision resolution strategies**:
   - **Separate Chaining**: Using linked lists or dynamic arrays
   - **Linear Probing**: Open addressing with linear probe sequence
   - **Double Hashing**: Open addressing with secondary hash function

2. **Implement multiple hash functions**:
   - Division method: `h(k) = k mod m`
   - Multiplication method: `h(k) = floor(m * (k * A mod 1))`
   - Universal hashing: Random hash from family
   - For strings: Rolling hash, FNV-1a, MurmurHash

3. **Add dynamic resizing**:
   - Monitor load factor (α = n/m)
   - Resize when α > threshold (e.g., 0.75)
   - Rehash all elements
   - Experiment with different thresholds

4. **Performance analysis**:
   - Measure collision rates for different hash functions
   - Compare lookup performance across strategies
   - Analyze worst-case vs. average-case behavior
   - Study impact of load factor on performance

### Implementation Details

```python
class HashTable:
    def __init__(self, collision_strategy="chaining", hash_function="division"):
        self.size = 0
        self.capacity = 16
        self.load_factor_threshold = 0.75
        self.table = allocate(self.capacity)
        
    def insert(self, key, value):
        index = self.hash(key)
        # Handle collision according to strategy
        self.size += 1
        if self.load_factor() > self.threshold:
            self.rehash()
    
    def search(self, key):
        index = self.hash(key)
        # Resolve collision and return value
    
    def delete(self, key):
        # Special handling for open addressing (tombstones)
```

### Test Cases
- Insert/search/delete 100K random keys
- Test with poor hash function (many collisions)
- Test with perfect hash function (no collisions)
- Pathological cases (all keys hash to same bucket)
- String keys with various distributions

### Deliverables
- `hash_table_chaining.{py,go,zig}` - Separate chaining implementation
- `hash_table_linear.{py,go,zig}` - Linear probing implementation
- `hash_table_double.{py,go,zig}` - Double hashing implementation
- `hash_functions.{py,go,zig}` - Collection of hash function implementations
- `collision_analysis.md` - Analysis of collision rates
- `performance_comparison.md` - Detailed performance study
- `benchmarks/` - Comprehensive benchmark suite
- `tests/` - Unit tests for all implementations

### Success Criteria
- All three strategies implemented correctly
- Load factor tracking and automatic resizing works
- Can explain when each strategy is preferable
- Hash function quality correlates with measured collision rates
- Performance matches theoretical expectations

---

---

## Back to Module

[← Back to Module 1: Elementary Data Structures](../module_01/README.md)

## Navigation

- [Previous Lab](./lab_1_1.md) (if exists)
- [Next Lab](./lab_1_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 1, Lab 2 of 5*
