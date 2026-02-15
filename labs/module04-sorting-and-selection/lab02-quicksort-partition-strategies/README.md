# Lab 4.2: Quicksort Partition Strategies

**Module**: Module 4 - Sorting & Selection  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 4.2: Quicksort Partition Strategies

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Deep dive into partition schemes
- Understand how pivot selection affects performance
- Mitigate quicksort's worst-case behavior

### Requirements

1. **Implement partition schemes**:
   - **Lomuto partition**: Simple but less efficient
   - **Hoare partition**: More efficient, harder to get right
   - **3-way partition**: Dutch National Flag for duplicates
   - **Dual-pivot partition**: Modern approach (used in Java's sort)

2. **Test all pivot selection strategies**:
   - Fixed positions (first, last, middle)
   - Random pivot
   - Median-of-three
   - Median-of-medians (for guaranteed O(n log n))
   - Ninther (median-of-three medians-of-three)

3. **Measure impact**:
   - Number of comparisons
   - Number of swaps
   - Recursion depth
   - Running time

### Implementation Details

```python
def lomuto_partition(array, lo, hi):
    pivot = array[hi]
    i = lo - 1
    
    for j in range(lo, hi):
        if array[j] <= pivot:
            i += 1
            array[i], array[j] = array[j], array[i]
    
    array[i + 1], array[hi] = array[hi], array[i + 1]
    return i + 1

def hoare_partition(array, lo, hi):
    pivot = array[lo]
    i, j = lo - 1, hi + 1
    
    while True:
        i += 1
        while array[i] < pivot:
            i += 1
        
        j -= 1
        while array[j] > pivot:
            j -= 1
        
        if i >= j:
            return j
        
        array[i], array[j] = array[j], array[i]

def three_way_partition(array, lo, hi):
    # Returns (lt, gt) where:
    # array[lo:lt] < pivot
    # array[lt:gt+1] == pivot
    # array[gt+1:hi+1] > pivot
    
    pivot = array[lo]
    lt, i, gt = lo, lo, hi
    
    while i <= gt:
        if array[i] < pivot:
            array[lt], array[i] = array[i], array[lt]
            lt += 1
            i += 1
        elif array[i] > pivot:
            array[i], array[gt] = array[gt], array[i]
            gt -= 1
        else:
            i += 1
    
    return lt, gt

def dual_pivot_partition(array, lo, hi):
    # Use two pivots
    if array[lo] > array[hi]:
        array[lo], array[hi] = array[hi], array[lo]
    
    pivot1, pivot2 = array[lo], array[hi]
    lt, gt = lo + 1, hi - 1
    i = lt
    
    while i <= gt:
        if array[i] < pivot1:
            array[i], array[lt] = array[lt], array[i]
            lt += 1
        elif array[i] > pivot2:
            while array[gt] > pivot2 and i < gt:
                gt -= 1
            array[i], array[gt] = array[gt], array[i]
            gt -= 1
            if array[i] < pivot1:
                array[i], array[lt] = array[lt], array[i]
                lt += 1
        i += 1
    
    lt -= 1
    gt += 1
    array[lo], array[lt] = array[lt], array[lo]
    array[hi], array[gt] = array[gt], array[hi]
    
    return lt, gt

def median_of_three(array, lo, hi):
    mid = (lo + hi) // 2
    if array[lo] > array[mid]:
        array[lo], array[mid] = array[mid], array[lo]
    if array[lo] > array[hi]:
        array[lo], array[hi] = array[hi], array[lo]
    if array[mid] > array[hi]:
        array[mid], array[hi] = array[hi], array[mid]
    return mid

def median_of_medians(array, lo, hi):
    # Deterministic pivot selection for worst-case O(n log n)
    if hi - lo < 5:
        return (lo + hi) // 2
    
    # Divide into groups of 5
    medians = []
    for i in range(lo, hi + 1, 5):
        group_end = min(i + 4, hi)
        group_median = find_median_of_five(array, i, group_end)
        medians.append(array[group_median])
    
    # Recursively find median of medians
    median_array = medians[:]
    return select(median_array, len(medians) // 2)
```

### Test Cases
- Sorted array (worst case for naive pivot selection)
- Array with all equal elements (3-way partition shines)
- Array with few unique values
- Random arrays
- Adversarial inputs designed to trigger worst case

### Deliverables
- `partition_schemes.{py,go,zig}` - All partition implementations
- `pivot_strategies.{py,go,zig}` - All pivot selection methods
- `partition_analysis.md` - Detailed comparison
- `benchmarks.{py,go,zig}` - Performance measurements
- `adversarial_inputs.{py,go,zig}` - Generate worst-case inputs
- `comparison_charts/` - Visual comparisons

### Success Criteria
- Understanding of how partition scheme affects performance
- 3-way partition significantly faster on duplicates
- Random/median-of-three prevent worst-case on sorted arrays
- Can explain Hoare vs. Lomuto tradeoffs

---

---

## Back to Module

[← Back to Module 4: Sorting & Selection](../module_04/README.md)

## Navigation

- [Previous Lab](./lab_4_1.md) (if exists)
- [Next Lab](./lab_4_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 4, Lab 2 of 7*
