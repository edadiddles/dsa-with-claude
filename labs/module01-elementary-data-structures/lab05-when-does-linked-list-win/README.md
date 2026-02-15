# Lab 1.5: When Does Linked List Win?

**Module**: Module 1 - Elementary Data Structures  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 1.5: When Does Linked List Win?

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

---

## Back to Module

[← Back to Module 1: Elementary Data Structures](../module_01/README.md)

## Navigation

- [Previous Lab](./lab_1_4.md) (if exists)
- [Next Lab](./lab_1_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 1, Lab 5 of 5*
