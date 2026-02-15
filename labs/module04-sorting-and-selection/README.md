# Module 4: Sorting & Selection

**Duration**: 3 weeks  
**Difficulty**: Medium

## Overview

Sorting is one of the most studied problems in computer science. This module provides a comprehensive exploration of sorting algorithms, from classic comparison sorts to specialized linear-time algorithms. You'll implement multiple algorithms, understand their tradeoffs, and learn when each excels.

**Why this matters**: Sorting is everywhere. Understanding the landscape of sorting algorithms - their best cases, worst cases, and practical performance - is essential for choosing the right tool. Plus, many algorithmic techniques (divide-and-conquer, randomization, adaptivity) are best learned through sorting.

## Learning Objectives

- Implement major sorting algorithms from scratch
- Understand comparison sort lower bounds
- Master pivot selection and partition strategies
- Build adaptive sorting algorithms
- Implement external sorting for large datasets
- Use selection algorithms to find order statistics
- Create comprehensive performance comparisons

## Topics Covered

- Comparison sorts: quicksort, mergesort, heapsort
- Linear-time sorts: counting, radix, bucket sort
- Partition strategies and pivot selection
- Adaptive sorting and detecting pre-existing order
- External sorting for disk-based data
- Selection and order statistics
- Practical optimization techniques

---

## Lab 4.1: Implement Five Core Sorting Algorithms

**Duration**: 6-8 hours  
**Difficulty**: Medium

### Objectives
- Implement major sorting algorithms from scratch
- Understand implementation details and edge cases
- Compare theoretical and practical performance

### Requirements

1. **Implement the following sorts**:
   - **Quicksort**: With various pivot selection strategies
   - **Mergesort**: Both top-down and bottom-up
   - **Heapsort**: Using your heap from Module 3
   - **Insertion Sort**: For small arrays
   - **Selection Sort**: For comparison

2. **For each algorithm**:
   - Basic implementation
   - Optimized version
   - In-place variant where applicable
   - Stable variant where applicable

3. **Quicksort pivot strategies**:
   - First element
   - Last element
   - Middle element
   - Random element
   - Median-of-three
   - Median-of-medians (deterministic O(n log n))

### Implementation Details

```python
def quicksort(array, lo=0, hi=None, pivot_strategy="median_of_three"):
    if hi is None:
        hi = len(array) - 1
    
    if lo < hi:
        pivot_idx = partition(array, lo, hi, pivot_strategy)
        quicksort(array, lo, pivot_idx - 1, pivot_strategy)
        quicksort(array, pivot_idx + 1, hi, pivot_strategy)

def partition(array, lo, hi, strategy):
    # Choose pivot based on strategy
    pivot_idx = choose_pivot(array, lo, hi, strategy)
    array[pivot_idx], array[hi] = array[hi], array[pivot_idx]
    
    pivot = array[hi]
    i = lo - 1
    
    for j in range(lo, hi):
        if array[j] <= pivot:
            i += 1
            array[i], array[j] = array[j], array[i]
    
    array[i + 1], array[hi] = array[hi], array[i + 1]
    return i + 1

def mergesort_top_down(array):
    if len(array) <= 1:
        return array
    
    mid = len(array) // 2
    left = mergesort_top_down(array[:mid])
    right = mergesort_top_down(array[mid:])
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0
    
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    
    result.extend(left[i:])
    result.extend(right[j:])
    return result

def mergesort_bottom_up(array):
    n = len(array)
    size = 1
    
    while size < n:
        for start in range(0, n, size * 2):
            mid = min(start + size, n)
            end = min(start + size * 2, n)
            merge_in_place(array, start, mid, end)
        size *= 2

def heapsort(array):
    # Build max heap
    n = len(array)
    for i in range(n // 2 - 1, -1, -1):
        heapify_down(array, i, n)
    
    # Extract elements
    for i in range(n - 1, 0, -1):
        array[0], array[i] = array[i], array[0]
        heapify_down(array, 0, i)

def insertion_sort(array):
    for i in range(1, len(array)):
        key = array[i]
        j = i - 1
        while j >= 0 and array[j] > key:
            array[j + 1] = array[j]
            j -= 1
        array[j + 1] = key

def selection_sort(array):
    for i in range(len(array)):
        min_idx = i
        for j in range(i + 1, len(array)):
            if array[j] < array[min_idx]:
                min_idx = j
        array[i], array[min_idx] = array[min_idx], array[i]
```

### Test Cases
- Random arrays: [10, 100, 1K, 10K, 100K, 1M] elements
- Already sorted arrays (best case for some, worst for others)
- Reverse sorted arrays
- Arrays with many duplicates
- Nearly sorted arrays (important for real-world data)
- Arrays with specific patterns (sawtooth, organ-pipe, etc.)

### Deliverables
- `quicksort.{py,go,zig}` - Multiple pivot strategies
- `mergesort.{py,go,zig}` - Top-down and bottom-up
- `heapsort.{py,go,zig}` - Using heap from Module 3
- `insertion_sort.{py,go,zig}` - Optimized version
- `selection_sort.{py,go,zig}` - Basic implementation
- `sort_comparison.md` - Analysis of characteristics
- `benchmarks.{py,go,zig}` - Comprehensive performance tests
- `test_data_generator.{py,go,zig}` - Generate various input patterns

### Success Criteria
- All sorts produce correct output
- Understanding of stability, adaptivity, and space complexity
- Can explain when each algorithm is preferable
- Performance matches theoretical expectations

---

## Lab 4.2: Quicksort Partition Strategies

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

## Lab 4.3: Adaptive Sorting

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Build sorts that detect and exploit pre-existing order
- Understand adaptivity and why it matters
- Implement Timsort-like optimizations

### Requirements

1. **Implement Adaptive Insertion Sort**:
   - Skip already-sorted regions
   - Use binary search for insertion position
   - Early termination when fully sorted

2. **Implement Natural Mergesort**:
   - Detect existing runs (ascending/descending sequences)
   - Merge natural runs instead of fixed-size subarrays
   - Avoid unnecessary comparisons in sorted regions

3. **Implement Timsort (simplified version)**:
   - Find runs (min run size)
   - Use insertion sort for small runs
   - Merge runs using galloping mode
   - Maintain run stack

4. **Measure adaptivity**:
   - Sorted array: Should be O(n)
   - Reverse sorted: Should detect and reverse runs
   - Nearly sorted: Should be faster than O(n log n)

### Implementation Details

```python
def natural_mergesort(array):
    runs = find_runs(array)
    
    while len(runs) > 1:
        merged_runs = []
        for i in range(0, len(runs), 2):
            if i + 1 < len(runs):
                merge_runs(array, runs[i], runs[i+1])
                merged_runs.append((runs[i][0], runs[i+1][1]))
            else:
                merged_runs.append(runs[i])
        runs = merged_runs
    
    return array

def find_runs(array):
    runs = []
    start = 0
    i = 1
    
    while i < len(array):
        # Detect ascending or descending run
        if i < len(array) and array[i] < array[i-1]:
            # Descending run
            while i < len(array) and array[i] < array[i-1]:
                i += 1
            reverse(array, start, i-1)
        else:
            # Ascending run
            while i < len(array) and array[i] >= array[i-1]:
                i += 1
        
        runs.append((start, i-1))
        start = i
        i += 1
    
    if start < len(array):
        runs.append((start, len(array)-1))
    
    return runs

class Timsort:
    MIN_MERGE = 32
    
    def __init__(self):
        self.min_gallop = 7
    
    def sort(self, array):
        n = len(array)
        min_run = self.compute_min_run(n)
        
        # Create initial runs
        for start in range(0, n, min_run):
            end = min(start + min_run - 1, n - 1)
            insertion_sort_range(array, start, end)
        
        # Merge runs
        size = min_run
        while size < n:
            for start in range(0, n, size * 2):
                mid = start + size - 1
                end = min(start + size * 2 - 1, n - 1)
                
                if mid < end:
                    self.merge_with_galloping(array, start, mid, end)
            
            size *= 2
    
    def compute_min_run(self, n):
        r = 0
        while n >= self.MIN_MERGE:
            r |= n & 1
            n >>= 1
        return n + r
    
    def merge_with_galloping(self, array, lo, mid, hi):
        # Timsort's galloping merge
        left = array[lo:mid+1]
        right = array[mid+1:hi+1]
        
        i = j = 0
        k = lo
        min_gallop = self.min_gallop
        
        while i < len(left) and j < len(right):
            # Regular merge until one side starts "winning"
            if left[i] <= right[j]:
                array[k] = left[i]
                i += 1
            else:
                array[k] = right[j]
                j += 1
            k += 1
            
            # Enter galloping mode if one side keeps winning
            # (simplified - full Timsort has more sophisticated galloping)
        
        while i < len(left):
            array[k] = left[i]
            i += 1
            k += 1
        
        while j < len(right):
            array[k] = right[j]
            j += 1
            k += 1

def adaptive_insertion_sort(array):
    # Binary search insertion sort
    for i in range(1, len(array)):
        key = array[i]
        
        # Binary search for insertion position
        left, right = 0, i - 1
        while left <= right:
            mid = (left + right) // 2
            if array[mid] > key:
                right = mid - 1
            else:
                left = mid + 1
        
        # Shift elements and insert
        for j in range(i, left, -1):
            array[j] = array[j-1]
        array[left] = key
```

### Test Cases
- Fully sorted array
- Reverse sorted array
- Partially sorted (e.g., sorted with 10% random elements)
- Data with multiple sorted runs
- Compare with non-adaptive sorts

### Deliverables
- `adaptive_insertion_sort.{py,go,zig}`
- `natural_mergesort.{py,go,zig}`
- `timsort.{py,go,zig}` - Simplified implementation
- `adaptivity_benchmarks.{py,go,zig}`
- `analysis.md` - When adaptivity helps
- `run_detection.{py,go,zig}` - Visualize detected runs

### Success Criteria
- Near-linear time on sorted/nearly-sorted data
- Understanding of why Python uses Timsort
- Can detect and measure adaptivity benefits
- Recognition of when adaptive sorts are worth the complexity

---

## Lab 4.4: External Sorting Simulation

**Duration**: 5-6 hours  
**Difficulty**: Hard

### Objectives
- Sort data larger than available memory
- Understand multi-pass external sorting
- Minimize disk I/O operations

### Requirements

1. **Simulate memory constraints**:
   - Limit usable memory to M elements
   - Track disk reads and writes
   - Implement external mergesort

2. **External sorting algorithm**:
   - Phase 1: Create sorted runs of size M
   - Phase 2: k-way merge with k = M/B (B = block size)
   - Minimize number of passes over data

3. **Optimizations**:
   - Replacement selection for longer initial runs
   - Polyphase merge
   - Overlapping I/O with computation

4. **Measure**:
   - Total disk I/O operations
   - Number of passes
   - Wall-clock time with simulated disk latency

### Implementation Details

```python
class DiskSimulator:
    def __init__(self, read_latency_ms=10, write_latency_ms=10):
        self.read_latency = read_latency_ms / 1000.0
        self.write_latency = write_latency_ms / 1000.0
        self.reads = 0
        self.writes = 0
        self.disk_storage = {}
    
    def read_block(self, filename, block_id):
        time.sleep(self.read_latency)
        self.reads += 1
        return self.disk_storage.get((filename, block_id), [])
    
    def write_block(self, filename, block_id, data):
        time.sleep(self.write_latency)
        self.writes += 1
        self.disk_storage[(filename, block_id)] = data[:]

def external_sort(input_file, output_file, memory_size, block_size, disk):
    # Phase 1: Create sorted runs
    runs = []
    chunk_id = 0
    
    while has_more_data(input_file):
        chunk = read_chunk(input_file, memory_size, disk)
        chunk.sort()  # In-memory sort
        
        run_file = f"run_{chunk_id}.tmp"
        write_run_to_disk(chunk, run_file, block_size, disk)
        runs.append(run_file)
        chunk_id += 1
    
    # Phase 2: k-way merge
    k = memory_size // block_size
    
    while len(runs) > 1:
        new_runs = []
        
        for i in range(0, len(runs), k):
            batch = runs[i:i+k]
            merged_file = f"merged_{len(new_runs)}.tmp"
            k_way_merge_external(batch, merged_file, memory_size, block_size, disk)
            new_runs.append(merged_file)
        
        # Clean up old runs
        for run in runs:
            delete_file(run, disk)
        
        runs = new_runs
    
    # Final run is the sorted file
    rename_file(runs[0], output_file, disk)
    
    return disk.reads, disk.writes

def k_way_merge_external(run_files, output_file, memory_size, block_size, disk):
    import heapq
    
    # Allocate buffers
    k = len(run_files)
    buffer_size = memory_size // (k + 1)  # k input buffers + 1 output buffer
    
    # Initialize input buffers
    input_buffers = []
    heap = []
    
    for i, run_file in enumerate(run_files):
        buffer = read_block_from_run(run_file, 0, buffer_size, disk)
        if buffer:
            input_buffers.append({
                'run_file': run_file,
                'buffer': buffer,
                'block_id': 0,
                'position': 0
            })
            heapq.heappush(heap, (buffer[0], i, 0))
    
    # Output buffer
    output_buffer = []
    output_block_id = 0
    
    # Merge
    while heap:
        value, run_idx, pos = heapq.heappop(heap)
        output_buffer.append(value)
        
        # Flush output buffer if full
        if len(output_buffer) >= buffer_size:
            write_block_to_file(output_file, output_block_id, output_buffer, disk)
            output_buffer = []
            output_block_id += 1
        
        # Refill input buffer if needed
        buf = input_buffers[run_idx]
        buf['position'] += 1
        
        if buf['position'] >= len(buf['buffer']):
            # Read next block
            buf['block_id'] += 1
            new_buffer = read_block_from_run(
                buf['run_file'], 
                buf['block_id'], 
                buffer_size, 
                disk
            )
            
            if new_buffer:
                buf['buffer'] = new_buffer
                buf['position'] = 0
                heapq.heappush(heap, (new_buffer[0], run_idx, 0))
        else:
            heapq.heappush(heap, (buf['buffer'][buf['position']], run_idx, buf['position']))
    
    # Flush remaining output
    if output_buffer:
        write_block_to_file(output_file, output_block_id, output_buffer, disk)

def replacement_selection(input_file, memory_size, disk):
    # Generate longer initial runs using a heap
    import heapq
    
    heap = []
    runs = []
    current_run = []
    run_id = 0
    
    # Fill initial heap
    for _ in range(memory_size):
        value = read_next_value(input_file, disk)
        if value is not None:
            heapq.heappush(heap, (value, run_id))
    
    while heap:
        value, run_tag = heapq.heappop(heap)
        
        if run_tag == run_id:
            current_run.append(value)
            
            # Read next value
            next_value = read_next_value(input_file, disk)
            if next_value is not None:
                if next_value >= value:
                    heapq.heappush(heap, (next_value, run_id))
                else:
                    heapq.heappush(heap, (next_value, run_id + 1))
        else:
            # Start new run
            write_run_to_disk(current_run, f"run_{run_id}.tmp", disk)
            runs.append(f"run_{run_id}.tmp")
            run_id += 1
            current_run = [value]
    
    if current_run:
        write_run_to_disk(current_run, f"run_{run_id}.tmp", disk)
        runs.append(f"run_{run_id}.tmp")
    
    return runs
```

### Test Cases
- Data size 10x memory size
- Data size 100x memory size
- Different block sizes
- Varying number of runs
- Compare replacement selection vs. simple run creation

### Deliverables
- `external_sort.{py,go,zig}` - Complete external sort
- `disk_simulator.{py,go,zig}` - Simulated disk I/O
- `run_generator.{py,go,zig}` - Create sorted runs
- `replacement_selection.{py,go,zig}` - Advanced run creation
- `analysis.md` - I/O analysis and optimizations
- Graphs showing I/O vs. data size

### Success Criteria
- Correctly sorts data larger than memory
- I/O count matches theoretical predictions
- Understanding of external sorting tradeoffs
- Replacement selection produces longer runs

---

## Lab 4.5: Selection Algorithms

**Duration**: 4-5 hours  
**Difficulty**: Medium-Hard

### Objectives
- Find kth smallest element without full sort
- Implement both randomized and deterministic algorithms
- Understand expected vs. worst-case guarantees

### Requirements

1. **Implement Randomized Quickselect**:
   - Expected O(n) time
   - Use random pivot
   - Partition and recurse on one side only

2. **Implement Median of Medians**:
   - Worst-case O(n) time
   - Deterministic pivot selection
   - Guaranteed linear time

3. **Additional selection algorithms**:
   - Heap-based selection (for small k)
   - Partial quicksort
   - Iterative vs. recursive versions

4. **Applications**:
   - Find median
   - Find k largest/smallest elements
   - Percentile calculation

### Implementation Details

```python
import random

def randomized_select(array, k):
    """Find kth smallest element (0-indexed)"""
    if len(array) == 1:
        return array[0]
    
    pivot_idx = random.randint(0, len(array) - 1)
    pivot = array[pivot_idx]
    
    lows = [x for x in array if x < pivot]
    highs = [x for x in array if x > pivot]
    pivots = [x for x in array if x == pivot]
    
    if k < len(lows):
        return randomized_select(lows, k)
    elif k < len(lows) + len(pivots):
        return pivot
    else:
        return randomized_select(highs, k - len(lows) - len(pivots))

def median_of_medians(array, k):
    """Deterministic O(n) selection"""
    if len(array) <= 5:
        return sorted(array)[k]
    
    # Divide into groups of 5
    medians = []
    for i in range(0, len(array), 5):
        group = array[i:i+5]
        medians.append(sorted(group)[len(group)//2])
    
    # Find median of medians recursively
    pivot = median_of_medians(medians, len(medians)//2)
    
    # Partition around pivot
    lows = [x for x in array if x < pivot]
    highs = [x for x in array if x > pivot]
    pivots = [x for x in array if x == pivot]
    
    if k < len(lows):
        return median_of_medians(lows, k)
    elif k < len(lows) + len(pivots):
        return pivot
    else:
        return median_of_medians(highs, k - len(lows) - len(pivots))

def heap_select(array, k):
    """Use min-heap to find k smallest elements"""
    import heapq
    
    # For k smallest: use max heap of size k
    # For k largest: use min heap of size k
    heap = []
    
    for num in array:
        if len(heap) < k:
            heapq.heappush(heap, -num)  # Negative for max heap
        elif num < -heap[0]:
            heapq.heapreplace(heap, -num)
    
    return -heap[0]  # kth smallest

def find_median(array):
    """Find median using selection"""
    n = len(array)
    if n % 2 == 1:
        return randomized_select(array, n // 2)
    else:
        return (randomized_select(array, n // 2 - 1) + 
                randomized_select(array, n // 2)) / 2.0

def find_percentile(array, p):
    """Find pth percentile (p in [0, 100])"""
    k = int(len(array) * p / 100)
    return randomized_select(array, k)

def partition_in_place(array, lo, hi, pivot_idx):
    """Partition for in-place quickselect"""
    pivot = array[pivot_idx]
    array[pivot_idx], array[hi] = array[hi], array[pivot_idx]
    
    store_idx = lo
    for i in range(lo, hi):
        if array[i] < pivot:
            array[i], array[store_idx] = array[store_idx], array[i]
            store_idx += 1
    
    array[store_idx], array[hi] = array[hi], array[store_idx]
    return store_idx

def quickselect_in_place(array, lo, hi, k):
    """In-place quickselect"""
    if lo == hi:
        return array[lo]
    
    pivot_idx = random.randint(lo, hi)
    pivot_idx = partition_in_place(array, lo, hi, pivot_idx)
    
    if k == pivot_idx:
        return array[k]
    elif k < pivot_idx:
        return quickselect_in_place(array, lo, pivot_idx - 1, k)
    else:
        return quickselect_in_place(array, pivot_idx + 1, hi, k)
```

### Test Cases
- Find median in random array
- Find min (k=0) and max (k=n-1)
- Find quartiles
- Arrays with duplicates
- Adversarial inputs
- Large arrays (verify O(n) performance)

### Deliverables
- `randomized_select.{py,go,zig}`
- `median_of_medians.{py,go,zig}`
- `heap_select.{py,go,zig}` - For small k
- `benchmarks.{py,go,zig}` - Compare approaches
- `analysis.md` - When to use each algorithm
- `applications.{py,go,zig}` - Median, percentiles, etc.

### Success Criteria
- Both algorithms find correct kth element
- Randomized is faster in practice
- Understanding of worst-case guarantees
- Can explain when to use heap-based selection (small k)

---

## Lab 4.6: Linear-Time Sorts

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement non-comparison sorts
- Understand when O(n) sorting is possible
- Recognize problem constraints enabling linear time

### Requirements

1. **Implement Counting Sort**:
   - For integers in known range [0, k]
   - Time: O(n + k), Space: O(k)
   - Stable version

2. **Implement Radix Sort**:
   - LSD (Least Significant Digit) version
   - MSD (Most Significant Digit) version
   - Works for integers and strings
   - Time: O(d(n + k)) where d = number of digits

3. **Implement Bucket Sort**:
   - For uniformly distributed data
   - Expected O(n) time
   - Choose appropriate number of buckets

4. **Analysis**:
   - When can these be used?
   - What constraints must hold?
   - Comparison with O(n log n) sorts

### Implementation Details

```python
def counting_sort(array, max_val):
    """Stable counting sort for integers in [0, max_val]"""
    count = [0] * (max_val + 1)
    
    # Count occurrences
    for x in array:
        count[x] += 1
    
    # Cumulative counts for stable sort
    for i in range(1, len(count)):
        count[i] += count[i-1]
    
    # Build output array
    output = [0] * len(array)
    for x in reversed(array):  # Reverse to maintain stability
        output[count[x] - 1] = x
        count[x] -= 1
    
    return output

def counting_sort_by_digit(array, digit, base=10):
    """Helper for radix sort - sort by specific digit"""
    count = [0] * base
    
    for x in array:
        digit_value = (x // (base ** digit)) % base
        count[digit_value] += 1
    
    for i in range(1, base):
        count[i] += count[i-1]
    
    output = [0] * len(array)
    for x in reversed(array):
        digit_value = (x // (base ** digit)) % base
        output[count[digit_value] - 1] = x
        count[digit_value] -= 1
    
    return output

def radix_sort_lsd(array, num_digits=None, base=10):
    """Least Significant Digit radix sort"""
    if not array:
        return array
    
    if num_digits is None:
        num_digits = len(str(max(array)))
    
    for digit in range(num_digits):
        array = counting_sort_by_digit(array, digit, base)
    
    return array

def radix_sort_msd(array, digit=None, base=10):
    """Most Significant Digit radix sort"""
    if not array or len(array) <= 1:
        return array
    
    if digit is None:
        digit = len(str(max(array))) - 1
    
    if digit < 0:
        return array
    
    # Partition by digit
    buckets = [[] for _ in range(base)]
    for x in array:
        digit_value = (x // (base ** digit)) % base
        buckets[digit_value].append(x)
    
    # Recursively sort each bucket
    result = []
    for bucket in buckets:
        result.extend(radix_sort_msd(bucket, digit - 1, base))
    
    return result

def bucket_sort(array, num_buckets=None):
    """Bucket sort for uniformly distributed data"""
    if not array:
        return array
    
    min_val, max_val = min(array), max(array)
    
    if num_buckets is None:
        num_buckets = len(array)
    
    bucket_range = (max_val - min_val) / num_buckets
    if bucket_range == 0:
        bucket_range = 1
    
    # Create buckets
    buckets = [[] for _ in range(num_buckets)]
    
    # Distribute elements
    for x in array:
        bucket_idx = int((x - min_val) / bucket_range)
        if bucket_idx == num_buckets:  # Handle max_val
            bucket_idx -= 1
        buckets[bucket_idx].append(x)
    
    # Sort each bucket and concatenate
    result = []
    for bucket in buckets:
        # Use insertion sort for small buckets
        insertion_sort(bucket)
        result.extend(bucket)
    
    return result

def radix_sort_strings(strings):
    """Radix sort for strings (MSD)"""
    def get_char(s, pos):
        return ord(s[pos]) if pos < len(s) else -1
    
    def radix_msd_helper(strings, pos, result):
        if not strings or pos >= max(len(s) for s in strings):
            result.extend(strings)
            return
        
        # Bucket by character at position
        buckets = {}
        for s in strings:
            char = get_char(s, pos)
            if char not in buckets:
                buckets[char] = []
            buckets[char].append(s)
        
        # Process buckets in order
        for char in sorted(buckets.keys()):
            if char == -1:
                result.extend(buckets[char])
            else:
                radix_msd_helper(buckets[char], pos + 1, result)
    
    result = []
    radix_msd_helper(strings, 0, result)
    return result
```

### Test Cases
- Counting sort: Small range (k << n), large range (k >> n)
- Radix sort: Integers, strings, floating-point (with tricks)
- Bucket sort: Uniform distribution, skewed distribution
- Edge cases: empty arrays, single elements, all duplicates

### Deliverables
- `counting_sort.{py,go,zig}`
- `radix_sort.{py,go,zig}` - Both LSD and MSD
- `bucket_sort.{py,go,zig}`
- `applicability_guide.md` - When to use each
- `benchmarks.{py,go,zig}` - Compare with comparison sorts
- `string_radix_sort.{py,go,zig}` - For string sorting

### Success Criteria
- All sorts correct and stable where required
- Understanding of problem constraints for linear time
- Knowledge of tradeoffs (time vs. space, range vs. size)
- Can identify when linear-time sorting is possible

---

## Lab 4.7: Comprehensive Sorting Comparison

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Create definitive comparison of all sorting algorithms
- Build intuition for algorithm selection
- Understand practical vs. theoretical performance

### Requirements

1. **Benchmark all implemented sorts on**:
   - Random data
   - Sorted data
   - Reverse sorted data
   - Nearly sorted data (with various noise levels)
   - Data with many duplicates
   - Data with specific distributions (Gaussian, uniform, Zipfian)

2. **Measure**:
   - Running time
   - Number of comparisons
   - Number of swaps/moves
   - Memory usage
   - Cache performance

3. **Create decision tree/matrix**:
   - Input size ranges
   - Data characteristics
   - Memory constraints
   - Stability requirements
   - Recommended algorithm

### Implementation Details

```python
class SortBenchmark:
    def __init__(self):
        self.algorithms = {
            'quicksort': quicksort,
            'mergesort': mergesort_top_down,
            'heapsort': heapsort,
            'timsort': timsort,
            'counting': counting_sort,
            'radix': radix_sort_lsd,
            'bucket': bucket_sort,
        }
    
    def benchmark_all(self, test_data, runs=10):
        results = {}
        
        for name, algorithm in self.algorithms.items():
            times = []
            comparisons = []
            swaps = []
            
            for _ in range(runs):
                data = test_data[:]
                
                start = time.time()
                algorithm(data)
                elapsed = time.time() - start
                
                times.append(elapsed)
            
            results[name] = {
                'mean_time': statistics.mean(times),
                'std_time': statistics.stdev(times) if len(times) > 1 else 0,
                'min_time': min(times),
                'max_time': max(times),
            }
        
        return results
    
    def test_all_patterns(self, size):
        patterns = {
            'random': generate_random(size),
            'sorted': generate_sorted(size),
            'reverse': generate_reverse_sorted(size),
            'nearly_sorted': generate_nearly_sorted(size, noise=0.1),
            'many_duplicates': generate_many_duplicates(size),
            'gaussian': generate_gaussian(size),
        }
        
        all_results = {}
        for pattern_name, data in patterns.items():
            all_results[pattern_name] = self.benchmark_all(data)
        
        return all_results

def generate_test_data(size, pattern):
    if pattern == 'random':
        return [random.randint(0, size) for _ in range(size)]
    elif pattern == 'sorted':
        return list(range(size))
    elif pattern == 'reverse':
        return list(range(size, 0, -1))
    elif pattern == 'nearly_sorted':
        data = list(range(size))
        # Add 10% noise
        for _ in range(size // 10):
            i, j = random.randint(0, size-1), random.randint(0, size-1)
            data[i], data[j] = data[j], data[i]
        return data
    elif pattern == 'many_duplicates':
        return [random.randint(0, size // 10) for _ in range(size)]
    elif pattern == 'gaussian':
        return [int(random.gauss(size/2, size/6)) for _ in range(size)]
    elif pattern == 'zipfian':
        # Implement Zipfian distribution
        pass

def create_decision_matrix():
    """
    Generate decision matrix for choosing sorting algorithm
    """
    decision_matrix = {
        'small_n_random': 'insertion_sort',
        'small_n_nearly_sorted': 'insertion_sort',
        'medium_n_random': 'quicksort',
        'medium_n_nearly_sorted': 'timsort',
        'large_n_random': 'quicksort',
        'large_n_nearly_sorted': 'timsort',
        'stability_required': 'mergesort or timsort',
        'limited_memory': 'heapsort',
        'known_range_integers': 'counting_sort or radix_sort',
        'uniform_distribution': 'bucket_sort',
        'external_data': 'external_mergesort',
    }
    
    return decision_matrix
```

### Deliverables
- `sort_comparison.{py,go,zig}` - Comprehensive benchmark suite
- `test_data_generator.{py,go,zig}` - Generate various input patterns
- `results_visualizer.{py,go,zig}` - Create comparison charts
- `sorting_guide.md` - Decision guide for choosing sort
- `performance_report.md` - Detailed analysis with graphs
- `decision_matrix.md` - Quick reference guide

### Success Criteria
- Clear performance patterns emerge
- Decision guide is actionable
- Understanding of why certain sorts excel in certain scenarios
- Can make informed sorting algorithm choices

---

## Module Resources

### CLRS References
- Chapter 7: Quicksort
- Chapter 8: Sorting in Linear Time
- Chapter 9: Medians and Order Statistics

### Additional Reading
- "Engineering a Sort Function" by Bentley & McIlroy
- Timsort description by Tim Peters
- "Introspective Sorting and Selection" (introsort paper)

### Key Takeaways

By the end of this module, you should:
1. Have implemented all major sorting algorithms
2. Understand the O(n log n) comparison sort lower bound
3. Know when linear-time sorting is possible
4. Be able to choose the right sort for any scenario
5. Understand adaptivity and its practical importance

### Next Module

**Module 5: Advanced Data Structures** - You'll implement specialized structures like Union-Find, tries, segment trees, and more.

---

## Tips for Success

1. **Test on real data** - Not just random arrays
2. **Measure everything** - Use your Module 0 tools
3. **Understand edge cases** - All equal, two elements, etc.
4. **Compare implementations** - Yours vs. standard library
5. **Document tradeoffs** - No one sort is best for everything

## Common Pitfalls

- **Not handling duplicates** - Especially in quicksort
- **Incorrect stability** - Test with duplicate values
- **Poor pivot selection** - Can make quicksort O(n²)
- **Ignoring adaptivity** - Real data is often partially sorted
- **Overengineering** - Simple sorts work fine for small data
