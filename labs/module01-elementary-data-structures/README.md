# Module 1: Elementary Data Structures

**Duration**: 3 weeks  
**Difficulty**: Medium

## Overview

These are the building blocks of all algorithms. Understanding their implementation details and tradeoffs is crucial for everything that follows. You'll implement these from scratch to build deep intuition about how they work.

**Why this matters**: These data structures appear everywhere. The difference between a good programmer and a great one is knowing which structure to use when, and why.

## Learning Objectives

- Implement fundamental data structures from scratch
- Understand amortized analysis deeply
- Design and evaluate hash functions
- Measure actual memory usage and overhead
- Make informed decisions about data structure selection

## Topics Covered

- Dynamic arrays with amortized analysis
- Linked lists (singly, doubly, circular)
- Stacks and queues
- Hash tables (collision resolution, hash functions, load factors)
- Memory layout and cache effects

---

## Lab 1.1: Dynamic Array Implementation

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

## Lab 1.2: Hash Table Deep Dive

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

## Lab 1.3: Custom Hash Functions

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Design hash functions for specific data types
- Understand avalanche effect and distribution
- Test hash function quality

### Requirements

1. **Design custom hash functions for**:
   - Integers (including negative numbers)
   - Strings (ASCII and Unicode)
   - Floating-point numbers
   - Custom objects/structs
   - Composite keys (tuples, pairs)

2. **Implement hash function quality tests**:
   - **Distribution test**: Chi-square test for uniformity
   - **Avalanche test**: Small input changes cause large hash changes
   - **Collision test**: Measure collision rate on real datasets
   - **Performance test**: Hashing speed

3. **Compare your hash functions against standard library implementations**

### Implementation Details

```python
def hash_integer(x, table_size):
    # Multiply by prime, mix bits, modulo
    
def hash_string(s, table_size):
    # Polynomial rolling hash or other method
    
def test_distribution(hash_func, dataset, table_size):
    buckets = [0] * table_size
    for item in dataset:
        buckets[hash_func(item, table_size)] += 1
    return chi_square_statistic(buckets)
```

### Test Cases
- Sequential integers (1, 2, 3, ...)
- Random integers
- Real-world strings (dictionary words, URLs, etc.)
- Synthetic worst-case inputs

### Deliverables
- `hash_functions.{py,go,zig}` - Collection of custom hash functions
- `hash_quality_tests.{py,go,zig}` - Testing framework
- `analysis.md` - Results and insights on hash function quality
- `visualizations/` - Distribution histograms

### Success Criteria
- Custom hash functions pass distribution tests
- Understanding of why certain designs work better
- Can design hash function for novel data type

---

## Lab 1.4: Memory Layout Analysis

**Duration**: 4-5 hours  
**Difficulty**: Hard

### Objectives
- Understand actual memory usage of data structures
- Measure memory overhead and alignment
- Optimize for cache performance

### Requirements

1. **Measure actual memory consumption**:
   - Array overhead
   - Linked list node overhead
   - Hash table memory usage
   - Pointer size and alignment

2. **Analyze memory layout**:
   - Struct/class padding
   - Cache line alignment
   - False sharing in concurrent structures
   - Memory fragmentation

3. **Implement memory-optimized versions**:
   - Packed arrays (bit-level packing)
   - Intrusive linked lists (no separate node allocation)
   - Open-addressed hash tables (less pointer overhead)
   - SOA vs. AOS (Structure of Arrays vs. Array of Structures)

### Implementation Details

```python
def measure_memory_overhead():
    # Allocate structure, measure total memory
    # Calculate per-element overhead
    
def analyze_cache_line_usage(data_structure):
    # Measure how many cache lines accessed
    # Calculate cache efficiency
```

### Test Cases
- 1M element array: measure actual bytes used
- Linked list with small elements: calculate overhead percentage
- Compare SOA vs. AOS for struct arrays

### Deliverables
- `memory_profiler.{py,go,zig}` - Memory measurement tools
- `optimized_structures.{py,go,zig}` - Memory-efficient implementations
- `memory_analysis.md` - Detailed findings
- Comparison charts

### Success Criteria
- Accurate memory measurements within 5%
- Optimized structures show measurable improvement
- Understanding of memory vs. CPU tradeoffs

---

## Lab 1.5: When Does Linked List Win?

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Find scenarios where linked lists outperform arrays
- Understand practical complexity beyond big-O
- Design experiments to expose differences

### Requirements

1. **Implement both** singly-linked list and dynamic array with identical interfaces

2. **Design experiments for different workloads**:
   - Sequential access (iteration)
   - Random access (get by index)
   - Front insertion/deletion
   - Middle insertion/deletion
   - Back insertion/deletion
   - Mixed workloads (realistic usage patterns)

3. **Measure performance across**:
   - Element counts (10, 100, 1K, 10K, 100K, 1M)
   - Element sizes (1 byte, 8 bytes, 1KB, 1MB)
   - Different operation mixes

4. **Find the break-even points** where linked list becomes competitive

### Implementation Details

```python
def benchmark_insertions(structure, position, count):
    # position: "front", "middle", "back"
    # Measure time for count insertions
    
def mixed_workload_simulation(structure):
    # Realistic mix: 70% read, 20% insert, 10% delete
```

### Test Cases
- Insert 10K elements at front: linked list should win
- Random access 10K times: array should win
- Insert at random positions: depends on size

### Deliverables
- `linked_list.{py,go,zig}` - Full linked list implementation
- `array_list.{py,go,zig}` - Dynamic array implementation
- `benchmarks.{py,go,zig}` - Comprehensive benchmark suite
- `findings.md` - Analysis of when each wins
- Decision flowchart for choosing data structure

### Success Criteria
- Can identify specific scenarios where linked list wins
- Understanding of constant factors vs. asymptotic complexity
- Practical guidelines for data structure selection

---

## Module Resources

### CLRS References
- Chapter 10: Elementary Data Structures
- Chapter 11: Hash Tables

### Additional Reading
- "The Art of Computer Programming" Vol 3 (searching and sorting)
- Hash function papers (MurmurHash, CityHash, etc.)
- Memory allocation papers for your chosen language

### Key Takeaways

By the end of this module, you should:
1. Understand that implementation details matter as much as asymptotic complexity
2. Know how to measure and optimize for real hardware
3. Be able to choose the right data structure for a given problem
4. Have deep intuition about hash tables and collision resolution
5. Appreciate the difference between theory and practice

### Next Module

**Module 2: Trees & Tree Algorithms** - You'll apply everything learned here to hierarchical data structures, implementing BSTs, balanced trees, and more.

---

## Tips for Success

1. **Use your Module 0 tools** - Profile everything, measure everything
2. **Test edge cases thoroughly** - Empty structures, single elements, etc.
3. **Compare implementations** - Your implementations vs. standard library
4. **Document tradeoffs** - Why did you choose one approach over another?
5. **Build incrementally** - Get basic functionality working before optimizing

## Common Pitfalls

- **Forgetting to handle edge cases** - Empty list deletions, etc.
- **Memory leaks** - Especially in manual memory management languages
- **Off-by-one errors** - Indexing is tricky
- **Poor hash functions** - Don't use naive approaches in production
- **Ignoring load factor** - Hash tables degrade without resizing
