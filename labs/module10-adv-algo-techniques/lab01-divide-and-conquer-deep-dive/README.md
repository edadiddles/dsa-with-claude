# Lab 10.1: Divide and Conquer Deep Dive

**Module**: Module 10 - Advanced Algorithm Techniques  
**Duration**: 6-8 hours  
**Difficulty**: Hard

---

Lab 10.1: Divide and Conquer Deep Dive

**Duration**: 6-8 hours  
**Difficulty**: Hard

### Objectives
- Master divide-and-conquer paradigm
- Implement advanced divide-and-conquer algorithms
- Analyze recurrences using Master Theorem
- Optimize divide-and-conquer algorithms

### Requirements

1. **Implement classic divide-and-conquer algorithms**:
   - **Merge Sort** (review and optimize)
   - **Quick Sort** with optimizations
   - **Binary Search** (iterative and recursive)
   - **Closest Pair of Points** (2D geometric)
   - **Strassen's Matrix Multiplication**
   - **Karatsuba Multiplication**

2. **Recurrence analysis**:
   - Master Theorem application
   - Recursion tree method
   - Substitution method
   - Akra-Bazzi theorem (advanced)

3. **Optimizations**:
   - Base case tuning
   - Hybrid algorithms
   - In-place variants

### Implementation Details

```python
import math

# MERGE SORT - Optimized

def merge_sort(arr, left=0, right=None):
    """
    Optimized merge sort with several improvements.
    """
    if right is None:
        right = len(arr) - 1
    
    # Optimization 1: Use insertion sort for small subarrays
    if right - left < 10:
        insertion_sort_range(arr, left, right)
        return
    
    if left < right:
        mid = (left + right) // 2
        
        merge_sort(arr, left, mid)
        merge_sort(arr, mid + 1, right)
        
        # Optimization 2: Skip merge if already sorted
        if arr[mid] <= arr[mid + 1]:
            return
        
        merge(arr, left, mid, right)

def merge(arr, left, mid, right):
    """Merge two sorted subarrays"""
    # Create temporary arrays
    left_arr = arr[left:mid+1]
    right_arr = arr[mid+1:right+1]
    
    i = j = 0
    k = left
    
    while i < len(left_arr) and j < len(right_arr):
        if left_arr[i] <= right_arr[j]:
            arr[k] = left_arr[i]
            i += 1
        else:
            arr[k] = right_arr[j]
            j += 1
        k += 1
    
    # Copy remaining elements
    while i < len(left_arr):
        arr[k] = left_arr[i]
        i += 1
        k += 1
    
    while j < len(right_arr):
        arr[k] = right_arr[j]
        j += 1
        k += 1

def insertion_sort_range(arr, left, right):
    """Insertion sort for range [left, right]"""
    for i in range(left + 1, right + 1):
        key = arr[i]
        j = i - 1
        while j >= left and arr[j] > key:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key

# QUICK SORT - Optimized

def quick_sort(arr, left=0, right=None):
    """
    Optimized quicksort with median-of-three pivot and tail recursion.
    """
    if right is None:
        right = len(arr) - 1
    
    while left < right:
        # Use insertion sort for small subarrays
        if right - left < 10:
            insertion_sort_range(arr, left, right)
            return
        
        # Median-of-three pivot selection
        pivot_idx = median_of_three(arr, left, right)
        
        # Partition
        pivot_idx = partition(arr, left, right, pivot_idx)
        
        # Tail recursion optimization: recurse on smaller partition
        if pivot_idx - left < right - pivot_idx:
            quick_sort(arr, left, pivot_idx - 1)
            left = pivot_idx + 1
        else:
            quick_sort(arr, pivot_idx + 1, right)
            right = pivot_idx - 1

def median_of_three(arr, left, right):
    """Select median of first, middle, last as pivot"""
    mid = (left + right) // 2
    
    # Sort three elements
    if arr[left] > arr[mid]:
        arr[left], arr[mid] = arr[mid], arr[left]
    if arr[left] > arr[right]:
        arr[left], arr[right] = arr[right], arr[left]
    if arr[mid] > arr[right]:
        arr[mid], arr[right] = arr[right], arr[mid]
    
    return mid

def partition(arr, left, right, pivot_idx):
    """Partition array around pivot"""
    pivot = arr[pivot_idx]
    
    # Move pivot to end
    arr[pivot_idx], arr[right] = arr[right], arr[pivot_idx]
    
    store_idx = left
    for i in range(left, right):
        if arr[i] < pivot:
            arr[i], arr[store_idx] = arr[store_idx], arr[i]
            store_idx += 1
    
    # Move pivot to final position
    arr[store_idx], arr[right] = arr[right], arr[store_idx]
    
    return store_idx

# CLOSEST PAIR OF POINTS

def closest_pair(points):
    """
    Find closest pair of points in 2D.
    Divide-and-conquer: O(n log n)
    """
    # Sort points by x-coordinate
    points_x = sorted(points, key=lambda p: p[0])
    points_y = sorted(points, key=lambda p: p[1])
    
    return closest_pair_recursive(points_x, points_y)

def closest_pair_recursive(px, py):
    """Recursive helper for closest pair"""
    n = len(px)
    
    # Base case: brute force for small n
    if n <= 3:
        return brute_force_closest(px)
    
    # Divide
    mid = n // 2
    midpoint = px[mid]
    
    pyl = [p for p in py if p[0] <= midpoint[0]]
    pyr = [p for p in py if p[0] > midpoint[0]]
    
    # Conquer
    dl = closest_pair_recursive(px[:mid], pyl)
    dr = closest_pair_recursive(px[mid:], pyr)
    
    # Find minimum
    d = min(dl, dr)
    
    # Check strip
    strip = [p for p in py if abs(p[0] - midpoint[0]) < d]
    
    for i in range(len(strip)):
        j = i + 1
        while j < len(strip) and (strip[j][1] - strip[i][1]) < d:
            dist = distance(strip[i], strip[j])
            d = min(d, dist)
            j += 1
    
    return d

def brute_force_closest(points):
    """Brute force closest pair for small inputs"""
    min_dist = float('inf')
    
    for i in range(len(points)):
        for j in range(i + 1, len(points)):
            dist = distance(points[i], points[j])
            min_dist = min(min_dist, dist)
    
    return min_dist

def distance(p1, p2):
    """Euclidean distance between two points"""
    return math.sqrt((p1[0] - p2[0])**2 + (p1[1] - p2[1])**2)

# STRASSEN'S MATRIX MULTIPLICATION

def strassen_multiply(A, B):
    """
    Strassen's algorithm for matrix multiplication.
    O(n^log2(7)) ≈ O(n^2.807) vs O(n^3) for naive.
    """
    n = len(A)
    
    # Base case: use standard multiplication for small matrices
    if n <= 64:
        return standard_matrix_multiply(A, B)
    
    # Pad to power of 2 if needed
    # (Simplified implementation assumes square matrices)
    
    # Divide matrices into quarters
    mid = n // 2
    
    A11 = [[A[i][j] for j in range(mid)] for i in range(mid)]
    A12 = [[A[i][j] for j in range(mid, n)] for i in range(mid)]
    A21 = [[A[i][j] for j in range(mid)] for i in range(mid, n)]
    A22 = [[A[i][j] for j in range(mid, n)] for i in range(mid, n)]
    
    B11 = [[B[i][j] for j in range(mid)] for i in range(mid)]
    B12 = [[B[i][j] for j in range(mid, n)] for i in range(mid)]
    B21 = [[B[i][j] for j in range(mid)] for i in range(mid, n)]
    B22 = [[B[i][j] for j in range(mid, n)] for i in range(mid, n)]
    
    # Compute 7 products
    M1 = strassen_multiply(
        matrix_add(A11, A22),
        matrix_add(B11, B22)
    )
    M2 = strassen_multiply(
        matrix_add(A21, A22),
        B11
    )
    M3 = strassen_multiply(
        A11,
        matrix_subtract(B12, B22)
    )
    M4 = strassen_multiply(
        A22,
        matrix_subtract(B21, B11)
    )
    M5 = strassen_multiply(
        matrix_add(A11, A12),
        B22
    )
    M6 = strassen_multiply(
        matrix_subtract(A21, A11),
        matrix_add(B11, B12)
    )
    M7 = strassen_multiply(
        matrix_subtract(A12, A22),
        matrix_add(B21, B22)
    )
    
    # Combine results
    C11 = matrix_add(matrix_subtract(matrix_add(M1, M4), M5), M7)
    C12 = matrix_add(M3, M5)
    C21 = matrix_add(M2, M4)
    C22 = matrix_add(matrix_subtract(matrix_add(M1, M3), M2), M6)
    
    # Combine quarters
    C = [[0] * n for _ in range(n)]
    for i in range(mid):
        for j in range(mid):
            C[i][j] = C11[i][j]
            C[i][j + mid] = C12[i][j]
            C[i + mid][j] = C21[i][j]
            C[i + mid][j + mid] = C22[i][j]
    
    return C

def matrix_add(A, B):
    """Add two matrices"""
    n = len(A)
    return [[A[i][j] + B[i][j] for j in range(n)] for i in range(n)]

def matrix_subtract(A, B):
    """Subtract two matrices"""
    n = len(A)
    return [[A[i][j] - B[i][j] for j in range(n)] for i in range(n)]

def standard_matrix_multiply(A, B):
    """Standard O(n^3) matrix multiplication"""
    n = len(A)
    C = [[0] * n for _ in range(n)]
    
    for i in range(n):
        for j in range(n):
            for k in range(n):
                C[i][j] += A[i][k] * B[k][j]
    
    return C

# KARATSUBA MULTIPLICATION

def karatsuba(x, y):
    """
    Karatsuba algorithm for fast integer multiplication.
    O(n^log2(3)) ≈ O(n^1.585) vs O(n^2) for grade school.
    """
    # Base case
    if x < 10 or y < 10:
        return x * y
    
    # Calculate number of digits
    n = max(len(str(x)), len(str(y)))
    m = n // 2
    
    # Split numbers
    high1, low1 = divmod(x, 10**m)
    high2, low2 = divmod(y, 10**m)
    
    # Three recursive multiplications
    z0 = karatsuba(low1, low2)
    z1 = karatsuba(low1 + high1, low2 + high2)
    z2 = karatsuba(high1, high2)
    
    # Combine results
    return z2 * 10**(2*m) + (z1 - z2 - z0) * 10**m + z0

# RECURRENCE ANALYSIS

def master_theorem(a, b, d):
    """
    Apply Master Theorem to recurrence T(n) = a*T(n/b) + O(n^d)
    
    Returns time complexity class:
    - If a > b^d: O(n^log_b(a))
    - If a = b^d: O(n^d log n)
    - If a < b^d: O(n^d)
    """
    log_b_a = math.log(a) / math.log(b)
    
    if a > b**d:
        return f"O(n^{log_b_a:.3f})"
    elif abs(a - b**d) < 1e-9:
        return f"O(n^{d} log n)"
    else:
        return f"O(n^{d})"

def analyze_merge_sort():
    """
    Merge sort: T(n) = 2*T(n/2) + O(n)
    a=2, b=2, d=1
    a = b^d → O(n log n)
    """
    return master_theorem(2, 2, 1)

def analyze_strassen():
    """
    Strassen: T(n) = 7*T(n/2) + O(n^2)
    a=7, b=2, d=2
    a > b^d → O(n^log_2(7)) ≈ O(n^2.807)
    """
    return master_theorem(7, 2, 2)

def analyze_karatsuba():
    """
    Karatsuba: T(n) = 3*T(n/2) + O(n)
    a=3, b=2, d=1
    a > b^d → O(n^log_2(3)) ≈ O(n^1.585)
    """
    return master_theorem(3, 2, 1)
```

### Test Cases
- Sort random arrays with all algorithms
- Compare performance (merge sort vs quick sort)
- Test closest pair with random points
- Verify matrix multiplication correctness
- Test integer multiplication with large numbers
- Analyze recurrences for all algorithms

### Deliverables
- `merge_sort_optimized.{py,go,zig}` - Production-ready merge sort
- `quick_sort_optimized.{py,go,zig}` - Optimized quicksort
- `closest_pair.{py,go,zig}` - 2D closest pair algorithm
- `strassen.{py,go,zig}` - Strassen's matrix multiplication
- `karatsuba.{py,go,zig}` - Karatsuba multiplication
- `recurrence_analyzer.{py,go,zig}` - Master theorem application
- `performance_comparison.md` - Benchmarks and analysis
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All divide-and-conquer algorithms work correctly
- Understanding of Master Theorem
- Can optimize divide-and-conquer algorithms
- Performance improvements measurable

---

---

## Back to Module

[← Back to Module 10: Advanced Algorithm Techniques](../module_10/README.md)

## Navigation

- [Previous Lab](./lab_10_0.md) (if exists)
- [Next Lab](./lab_10_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 10, Lab 1 of 5*
