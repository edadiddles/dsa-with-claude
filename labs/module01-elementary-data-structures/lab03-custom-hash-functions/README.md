# Lab 1.3: Custom Hash Functions

**Module**: Module 1 - Elementary Data Structures  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 1.3: Custom Hash Functions

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

---

## Back to Module

[← Back to Module 1: Elementary Data Structures](../module_01/README.md)

## Navigation

- [Previous Lab](./lab_1_2.md) (if exists)
- [Next Lab](./lab_1_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 1, Lab 3 of 5*
