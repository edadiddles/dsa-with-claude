# Module 10: Advanced Algorithm Techniques

**Duration**: 3 weeks  
**Difficulty**: Medium to Hard

## Overview

This module explores advanced algorithmic techniques that extend beyond the core paradigms of DP and greedy. You'll master divide-and-conquer, backtracking, randomized algorithms, and amortized analysis - essential tools for tackling complex problems that don't fit neatly into basic categories.

**Why this matters**: Real-world problems often require sophisticated techniques. Divide-and-conquer powers FFT and matrix multiplication. Backtracking solves constraint satisfaction problems. Randomized algorithms break through worst-case barriers. Amortized analysis justifies data structure performance. These techniques appear in compilers, databases, AI, and high-performance computing.

## Learning Objectives

- Master divide-and-conquer algorithm design
- Implement backtracking with pruning and optimization
- Understand and apply randomized algorithms
- Perform amortized analysis of data structures
- Recognize advanced algorithm design patterns
- Apply techniques to practical problems

## Topics Covered

- Divide-and-conquer recurrence analysis
- Advanced divide-and-conquer (Karatsuba, Strassen, FFT)
- Backtracking and branch-and-bound
- Constraint satisfaction problems
- Randomized algorithms and probabilistic analysis
- Las Vegas vs. Monte Carlo algorithms
- Amortized analysis techniques (aggregate, accounting, potential method)
- Algorithm design patterns and problem-solving strategies

---

## Lab 10.1: Divide and Conquer Deep Dive

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

## Lab 10.2: Backtracking & Branch-and-Bound

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Master backtracking technique
- Implement constraint satisfaction problems
- Apply branch-and-bound optimization
- Understand pruning strategies

### Requirements

1. **Implement classic backtracking problems**:
   - **N-Queens**
   - **Sudoku Solver**
   - **Graph Coloring**
   - **Subset Sum**
   - **Hamiltonian Path**
   - **Traveling Salesman (exact)**

2. **Optimizations**:
   - Constraint propagation
   - Intelligent ordering
   - Pruning strategies
   - Branch-and-bound for optimization

3. **Analysis**:
   - Measure search space reduction
   - Compare with/without pruning
   - Understand exponential complexity

### Implementation Details

```python
# N-QUEENS

def solve_n_queens(n):
    """
    Solve N-Queens problem using backtracking.
    Place n queens on n×n board so no two attack each other.
    """
    solutions = []
    board = [-1] * n  # board[row] = column of queen in that row
    
    def is_safe(row, col):
        """Check if queen can be placed at (row, col)"""
        for prev_row in range(row):
            prev_col = board[prev_row]
            
            # Check column conflict
            if prev_col == col:
                return False
            
            # Check diagonal conflict
            if abs(prev_row - row) == abs(prev_col - col):
                return False
        
        return True
    
    def backtrack(row):
        """Try placing queen in row"""
        if row == n:
            # Found complete solution
            solutions.append(board[:])
            return
        
        for col in range(n):
            if is_safe(row, col):
                board[row] = col
                backtrack(row + 1)
                board[row] = -1  # Backtrack
    
    backtrack(0)
    return solutions

def solve_n_queens_optimized(n):
    """
    Optimized N-Queens using bit manipulation.
    Tracks attacked columns and diagonals with bits.
    """
    solutions = []
    
    def backtrack(row, cols, diag1, diag2, board):
        if row == n:
            solutions.append(board[:])
            return
        
        # Try each column
        available = ((1 << n) - 1) & ~(cols | diag1 | diag2)
        
        while available:
            # Get rightmost available position
            col = available & -available
            col_idx = (col - 1).bit_length() - 1
            
            board.append(col_idx)
            
            backtrack(
                row + 1,
                cols | col,
                (diag1 | col) << 1,
                (diag2 | col) >> 1,
                board
            )
            
            board.pop()
            available &= available - 1
    
    backtrack(0, 0, 0, 0, [])
    return solutions

# SUDOKU SOLVER

def solve_sudoku(board):
    """
    Solve 9×9 Sudoku using backtracking.
    board is 9×9 list of lists, 0 represents empty cell.
    """
    def is_valid(row, col, num):
        """Check if num can be placed at (row, col)"""
        # Check row
        if num in board[row]:
            return False
        
        # Check column
        if num in [board[i][col] for i in range(9)]:
            return False
        
        # Check 3×3 box
        box_row, box_col = 3 * (row // 3), 3 * (col // 3)
        for i in range(box_row, box_row + 3):
            for j in range(box_col, box_col + 3):
                if board[i][j] == num:
                    return False
        
        return True
    
    def find_empty():
        """Find next empty cell"""
        for i in range(9):
            for j in range(9):
                if board[i][j] == 0:
                    return i, j
        return None
    
    def solve():
        """Recursive backtracking solver"""
        empty = find_empty()
        
        if empty is None:
            return True  # Solved
        
        row, col = empty
        
        for num in range(1, 10):
            if is_valid(row, col, num):
                board[row][col] = num
                
                if solve():
                    return True
                
                board[row][col] = 0  # Backtrack
        
        return False
    
    solve()
    return board

def solve_sudoku_optimized(board):
    """
    Optimized Sudoku with constraint propagation.
    """
    # Track possible values for each cell
    possible = [[set(range(1, 10)) if board[i][j] == 0 else set() 
                 for j in range(9)] for i in range(9)]
    
    # Initialize constraints
    for i in range(9):
        for j in range(9):
            if board[i][j] != 0:
                propagate_constraints(board, possible, i, j, board[i][j])
    
    def propagate_constraints(board, possible, row, col, num):
        """Remove num from possible values in row, col, box"""
        # Row and column
        for k in range(9):
            possible[row][k].discard(num)
            possible[k][col].discard(num)
        
        # Box
        box_row, box_col = 3 * (row // 3), 3 * (col // 3)
        for i in range(box_row, box_row + 3):
            for j in range(box_col, box_col + 3):
                possible[i][j].discard(num)
    
    def solve():
        # Find cell with minimum remaining values (MRV heuristic)
        min_choices = 10
        best_cell = None
        
        for i in range(9):
            for j in range(9):
                if board[i][j] == 0 and len(possible[i][j]) < min_choices:
                    min_choices = len(possible[i][j])
                    best_cell = (i, j)
        
        if best_cell is None:
            return True  # Solved
        
        if min_choices == 0:
            return False  # Dead end
        
        row, col = best_cell
        
        for num in list(possible[row][col]):
            board[row][col] = num
            old_possible = [row[:] for row in possible]
            
            propagate_constraints(board, possible, row, col, num)
            
            if solve():
                return True
            
            board[row][col] = 0
            possible[:] = old_possible
        
        return False
    
    solve()
    return board

# GRAPH COLORING

def graph_coloring(graph, k):
    """
    Color graph with k colors using backtracking.
    graph is adjacency list.
    """
    n = len(graph)
    colors = [-1] * n
    
    def is_safe(node, color):
        """Check if node can be colored with color"""
        for neighbor in graph[node]:
            if colors[neighbor] == color:
                return False
        return True
    
    def backtrack(node):
        if node == n:
            return True  # All nodes colored
        
        for color in range(k):
            if is_safe(node, color):
                colors[node] = color
                
                if backtrack(node + 1):
                    return True
                
                colors[node] = -1  # Backtrack
        
        return False
    
    if backtrack(0):
        return colors
    return None

# SUBSET SUM

def subset_sum_backtrack(nums, target):
    """
    Find subset that sums to target using backtracking.
    Returns all solutions.
    """
    solutions = []
    
    def backtrack(index, current_sum, subset):
        if current_sum == target:
            solutions.append(subset[:])
            return
        
        if index >= len(nums) or current_sum > target:
            return
        
        # Include current number
        subset.append(nums[index])
        backtrack(index + 1, current_sum + nums[index], subset)
        subset.pop()
        
        # Exclude current number
        backtrack(index + 1, current_sum, subset)
    
    backtrack(0, 0, [])
    return solutions

# TRAVELING SALESMAN (EXACT)

def tsp_backtracking(dist_matrix):
    """
    Exact TSP solution using backtracking with branch-and-bound.
    dist_matrix[i][j] = distance from city i to city j.
    """
    n = len(dist_matrix)
    min_cost = [float('inf')]
    best_path = []
    
    def bound(path, visited):
        """Lower bound on remaining cost"""
        # Minimum cost to complete tour from current state
        cost = 0
        
        # Add cost of path so far
        for i in range(len(path) - 1):
            cost += dist_matrix[path[i]][path[i+1]]
        
        # Add minimum outgoing edge from last city
        if path:
            last = path[-1]
            min_out = min(dist_matrix[last][j] for j in range(n) 
                         if j not in visited)
            cost += min_out
        
        # Add minimum edges for unvisited cities
        for city in range(n):
            if city not in visited:
                min_edge = min(dist_matrix[city])
                cost += min_edge
        
        return cost
    
    def backtrack(path, visited, cost):
        if len(path) == n:
            # Complete tour, return to start
            total = cost + dist_matrix[path[-1]][path[0]]
            if total < min_cost[0]:
                min_cost[0] = total
                best_path[:] = path + [path[0]]
            return
        
        # Branch and bound: prune if lower bound exceeds best
        if bound(path, visited) >= min_cost[0]:
            return
        
        last = path[-1] if path else 0
        
        for next_city in range(n):
            if next_city not in visited:
                new_cost = cost + dist_matrix[last][next_city]
                
                path.append(next_city)
                visited.add(next_city)
                
                backtrack(path, visited, new_cost)
                
                path.pop()
                visited.remove(next_city)
    
    backtrack([0], {0}, 0)
    return min_cost[0], best_path

# PRUNING STRATEGIES

class BacktrackingOptimizer:
    """Collection of pruning and optimization techniques"""
    
    @staticmethod
    def forward_checking(domain, constraints):
        """
        Forward checking: remove inconsistent values from future variables.
        """
        pass
    
    @staticmethod
    def arc_consistency(domain, constraints):
        """
        Maintain arc consistency during search.
        """
        pass
    
    @staticmethod
    def minimum_remaining_values(variables, domains):
        """
        MRV heuristic: choose variable with fewest remaining values.
        """
        return min(variables, key=lambda v: len(domains[v]))
    
    @staticmethod
    def degree_heuristic(variables, graph):
        """
        Degree heuristic: choose variable with most constraints.
        """
        return max(variables, key=lambda v: len(graph[v]))
    
    @staticmethod
    def least_constraining_value(variable, values, constraints):
        """
        LCV: choose value that rules out fewest choices for neighbors.
        """
        def count_conflicts(value):
            count = 0
            for neighbor in constraints[variable]:
                # Count how many values this eliminates for neighbor
                pass
            return count
        
        return sorted(values, key=count_conflicts)
```

### Test Cases
- N-Queens for various n (compare with/without optimization)
- Sudoku puzzles of varying difficulty
- Graph coloring instances
- Subset sum with different targets
- Small TSP instances (exact solution)

### Deliverables
- `n_queens.{py,go,zig}` - Both basic and optimized versions
- `sudoku_solver.{py,go,zig}` - With constraint propagation
- `graph_coloring.{py,go,zig}` - Backtracking coloring
- `tsp_exact.{py,go,zig}` - Branch-and-bound TSP
- `constraint_propagation.{py,go,zig}` - CSP utilities
- `pruning_analysis.md` - Effectiveness of pruning
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All backtracking algorithms find correct solutions
- Understanding of pruning effectiveness
- Branch-and-bound reduces search space significantly
- Can apply backtracking to new problems

---

## Lab 10.3: Randomized Algorithms

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Understand randomized algorithm paradigm
- Implement Monte Carlo and Las Vegas algorithms
- Analyze expected running time
- Apply randomization to practical problems

### Requirements

1. **Implement randomized algorithms**:
   - **Randomized Quicksort**
   - **Randomized Selection** (QuickSelect)
   - **Monte Carlo Primality Testing** (Miller-Rabin)
   - **Randomized Min-Cut** (Karger's algorithm)
   - **Skip List** (from Module 5, review)
   - **Bloom Filter** (from Module 5, review)

2. **Analysis**:
   - Expected time complexity
   - Probability of correctness
   - Worst-case vs. expected case

3. **Comparison**:
   - Deterministic vs. randomized
   - Las Vegas vs. Monte Carlo

### Implementation Details

```python
import random

# RANDOMIZED QUICKSORT

def randomized_quicksort(arr):
    """
    Quicksort with random pivot selection.
    Expected O(n log n), worst case O(n^2) but unlikely.
    """
    if len(arr) <= 1:
        return arr
    
    # Random pivot
    pivot_idx = random.randint(0, len(arr) - 1)
    pivot = arr[pivot_idx]
    
    # Partition
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    
    return randomized_quicksort(left) + middle + randomized_quicksort(right)

# RANDOMIZED SELECTION (QuickSelect)

def randomized_select(arr, k):
    """
    Find k-th smallest element using randomized selection.
    Expected O(n), worst case O(n^2).
    Las Vegas algorithm: always correct, randomized time.
    """
    if len(arr) == 1:
        return arr[0]
    
    # Random pivot
    pivot_idx = random.randint(0, len(arr) - 1)
    pivot = arr[pivot_idx]
    
    # Partition
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    
    if k < len(left):
        return randomized_select(left, k)
    elif k < len(left) + len(middle):
        return pivot
    else:
        return randomized_select(right, k - len(left) - len(middle))

# MILLER-RABIN PRIMALITY TEST

def miller_rabin(n, k=5):
    """
    Miller-Rabin primality test.
    Monte Carlo algorithm: randomized, small error probability.
    
    Parameters:
    - n: number to test
    - k: number of rounds (higher k = lower error probability)
    
    Error probability: at most 1/(4^k)
    """
    if n < 2:
        return False
    if n == 2 or n == 3:
        return True
    if n % 2 == 0:
        return False
    
    # Write n-1 as 2^r * d
    r, d = 0, n - 1
    while d % 2 == 0:
        r += 1
        d //= 2
    
    # Witness loop
    for _ in range(k):
        a = random.randint(2, n - 2)
        x = pow(a, d, n)
        
        if x == 1 or x == n - 1:
            continue
        
        for _ in range(r - 1):
            x = pow(x, 2, n)
            if x == n - 1:
                break
        else:
            return False
    
    return True

def generate_prime(bits):
    """Generate random prime with given bit length"""
    while True:
        candidate = random.getrandbits(bits)
        # Make odd
        candidate |= 1
        
        if miller_rabin(candidate, k=10):
            return candidate

# KARGER'S MIN-CUT ALGORITHM

def karger_min_cut(graph):
    """
    Karger's randomized min-cut algorithm.
    Find minimum cut in undirected graph.
    
    Monte Carlo: may give wrong answer, repeat for better probability.
    Running O(n^2 log n) times gives high probability of correctness.
    """
    # Graph representation: adjacency list with parallel edges
    # Copy graph to avoid modifying original
    import copy
    g = copy.deepcopy(graph)
    
    def contract_edge(g, u, v):
        """Contract edge (u,v) by merging v into u"""
        # Move all edges from v to u
        for neighbor in g[v]:
            if neighbor != u:
                g[u].append(neighbor)
                # Update neighbor's edge
                for i, n in enumerate(g[neighbor]):
                    if n == v:
                        g[neighbor][i] = u
        
        # Remove v
        del g[v]
        
        # Remove self-loops at u
        g[u] = [n for n in g[u] if n != u]
    
    # Contract until 2 vertices remain
    while len(g) > 2:
        # Choose random edge
        u = random.choice(list(g.keys()))
        v = random.choice(g[u])
        
        contract_edge(g, u, v)
    
    # Count edges between remaining vertices
    vertices = list(g.keys())
    return len(g[vertices[0]])

def karger_min_cut_repeated(graph, iterations=None):
    """
    Repeat Karger's algorithm multiple times.
    Choose minimum cut found.
    
    Running O(n^2 log n) times gives probability > 1 - 1/n of correctness.
    """
    n = len(graph)
    if iterations is None:
        iterations = n * n * int(math.log(n) + 1)
    
    min_cut = float('inf')
    
    for _ in range(iterations):
        cut_size = karger_min_cut(graph)
        min_cut = min(min_cut, cut_size)
    
    return min_cut

# RESERVOIR SAMPLING

def reservoir_sampling(stream, k):
    """
    Select k random elements from stream of unknown length.
    Each element has equal probability k/n of being selected.
    """
    reservoir = []
    
    for i, element in enumerate(stream):
        if i < k:
            reservoir.append(element)
        else:
            # Random index in [0, i]
            j = random.randint(0, i)
            if j < k:
                reservoir[j] = element
    
    return reservoir

# RANDOMIZED ROUNDING

def randomized_rounding(fractional_solution):
    """
    Convert fractional solution to integer solution randomly.
    Used in approximation algorithms.
    """
    integer_solution = []
    
    for x in fractional_solution:
        if random.random() < x:
            integer_solution.append(1)
        else:
            integer_solution.append(0)
    
    return integer_solution

# LAS VEGAS VS MONTE CARLO

class AlgorithmClassifier:
    """Classify randomized algorithms"""
    
    @staticmethod
    def is_las_vegas(algorithm):
        """
        Las Vegas: Always correct, randomized running time.
        Examples: Randomized Quicksort, QuickSelect
        """
        return {
            'always_correct': True,
            'randomized_time': True,
            'examples': ['Randomized Quicksort', 'QuickSelect']
        }
    
    @staticmethod
    def is_monte_carlo(algorithm):
        """
        Monte Carlo: Randomized correctness, fixed time.
        Examples: Miller-Rabin, Karger's Min-Cut
        """
        return {
            'randomized_correctness': True,
            'fixed_time': True,
            'error_probability': 'can be made arbitrarily small',
            'examples': ['Miller-Rabin', 'Karger Min-Cut']
        }

# PROBABILISTIC ANALYSIS

def analyze_expected_time(algorithm, inputs, trials=1000):
    """
    Empirically measure expected running time.
    """
    import time
    
    times = []
    
    for _ in range(trials):
        test_input = random.choice(inputs)
        
        start = time.time()
        algorithm(test_input)
        end = time.time()
        
        times.append(end - start)
    
    return {
        'mean': sum(times) / len(times),
        'min': min(times),
        'max': max(times),
        'median': sorted(times)[len(times)//2]
    }

def probability_of_correctness(monte_carlo_algorithm, test_cases, trials=1000):
    """
    Empirically measure probability of correctness.
    """
    correct = 0
    
    for test_input, expected_output in test_cases:
        for _ in range(trials):
            result = monte_carlo_algorithm(test_input)
            if result == expected_output:
                correct += 1
    
    return correct / (len(test_cases) * trials)
```

### Test Cases
- Randomized quicksort on various inputs
- QuickSelect for finding median
- Miller-Rabin on known primes and composites
- Karger's min-cut on small graphs
- Measure probability of correctness

### Deliverables
- `randomized_quicksort.{py,go,zig}` - Implementation
- `randomized_select.{py,go,zig}` - QuickSelect
- `miller_rabin.{py,go,zig}` - Primality testing
- `karger_mincut.{py,go,zig}` - Min-cut algorithm
- `reservoir_sampling.{py,go,zig}` - Streaming algorithm
- `probabilistic_analysis.{py,go,zig}` - Analysis tools
- `randomized_vs_deterministic.md` - Comparison
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Understanding of Las Vegas vs. Monte Carlo
- Can analyze expected running time
- Probability of correctness measurable
- Know when randomization helps

---

## Lab 10.4: Amortized Analysis Techniques

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

## Lab 10.5: Algorithm Design Patterns

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Recognize common algorithm design patterns
- Apply patterns to new problems
- Build problem-solving toolkit

### Requirements

1. **Master design patterns**:
   - Two pointers
   - Sliding window
   - Prefix sum / difference array
   - Monotonic stack/queue
   - Top-K elements
   - Union-Find applications

2. **Problem-solving framework**:
   - Pattern recognition
   - Template application
   - Adaptation to variants

3. **Practice problems**:
   - Solve 20+ problems using patterns
   - Categorize by pattern
   - Build intuition

### Implementation Details

```python
# TWO POINTERS

def two_sum_sorted(arr, target):
    """
    Find two numbers that sum to target in sorted array.
    Pattern: Two pointers (one from each end).
    """
    left, right = 0, len(arr) - 1
    
    while left < right:
        current_sum = arr[left] + arr[right]
        
        if current_sum == target:
            return [left, right]
        elif current_sum < target:
            left += 1
        else:
            right -= 1
    
    return None

def remove_duplicates(arr):
    """
    Remove duplicates from sorted array in-place.
    Pattern: Two pointers (fast and slow).
    """
    if not arr:
        return 0
    
    slow = 0
    
    for fast in range(1, len(arr)):
        if arr[fast] != arr[slow]:
            slow += 1
            arr[slow] = arr[fast]
    
    return slow + 1

# SLIDING WINDOW

def max_sum_subarray(arr, k):
    """
    Maximum sum of subarray of size k.
    Pattern: Sliding window.
    """
    if len(arr) < k:
        return None
    
    # Initial window
    window_sum = sum(arr[:k])
    max_sum = window_sum
    
    # Slide window
    for i in range(k, len(arr)):
        window_sum += arr[i] - arr[i - k]
        max_sum = max(max_sum, window_sum)
    
    return max_sum

def longest_substring_k_distinct(s, k):
    """
    Longest substring with at most k distinct characters.
    Pattern: Sliding window with hashmap.
    """
    from collections import defaultdict
    
    char_count = defaultdict(int)
    left = 0
    max_length = 0
    
    for right in range(len(s)):
        char_count[s[right]] += 1
        
        while len(char_count) > k:
            char_count[s[left]] -= 1
            if char_count[s[left]] == 0:
                del char_count[s[left]]
            left += 1
        
        max_length = max(max_length, right - left + 1)
    
    return max_length

# PREFIX SUM

def range_sum_query(arr):
    """
    Precompute prefix sums for O(1) range queries.
    Pattern: Prefix sum.
    """
    prefix = [0]
    for num in arr:
        prefix.append(prefix[-1] + num)
    
    def query(left, right):
        return prefix[right + 1] - prefix[left]
    
    return query

def subarray_sum_equals_k(arr, k):
    """
    Count subarrays with sum equal to k.
    Pattern: Prefix sum with hashmap.
    """
    from collections import defaultdict
    
    count = 0
    prefix_sum = 0
    sum_count = defaultdict(int)
    sum_count[0] = 1
    
    for num in arr:
        prefix_sum += num
        
        # Check if prefix_sum - k exists
        count += sum_count[prefix_sum - k]
        
        sum_count[prefix_sum] += 1
    
    return count

# MONOTONIC STACK

def next_greater_element(arr):
    """
    For each element, find next greater element.
    Pattern: Monotonic stack.
    """
    result = [-1] * len(arr)
    stack = []
    
    for i in range(len(arr)):
        while stack and arr[i] > arr[stack[-1]]:
            idx = stack.pop()
            result[idx] = arr[i]
        stack.append(i)
    
    return result

def largest_rectangle_histogram(heights):
    """
    Largest rectangle in histogram.
    Pattern: Monotonic stack.
    """
    stack = []
    max_area = 0
    
    for i in range(len(heights)):
        while stack and heights[i] < heights[stack[-1]]:
            h_idx = stack.pop()
            h = heights[h_idx]
            w = i if not stack else i - stack[-1] - 1
            max_area = max(max_area, h * w)
        
        stack.append(i)
    
    while stack:
        h_idx = stack.pop()
        h = heights[h_idx]
        w = len(heights) if not stack else len(heights) - stack[-1] - 1
        max_area = max(max_area, h * w)
    
    return max_area

# TOP K ELEMENTS

def top_k_frequent(nums, k):
    """
    Find k most frequent elements.
    Pattern: Heap / Bucket sort.
    """
    from collections import Counter
    import heapq
    
    count = Counter(nums)
    
    # Method 1: Min heap of size k
    heap = []
    for num, freq in count.items():
        heapq.heappush(heap, (freq, num))
        if len(heap) > k:
            heapq.heappop(heap)
    
    return [num for freq, num in heap]

def kth_largest(arr, k):
    """
    Find kth largest element.
    Pattern: QuickSelect or heap.
    """
    import heapq
    return heapq.nlargest(k, arr)[-1]

# PATTERN TEMPLATES

class AlgorithmPatterns:
    """Collection of algorithm pattern templates"""
    
    @staticmethod
    def two_pointer_template():
        """
        Template for two pointer problems:
        
        left, right = 0, len(arr) - 1
        
        while left < right:
            if condition:
                # Found answer
                return
            elif need_larger:
                left += 1
            else:
                right -= 1
        """
        pass
    
    @staticmethod
    def sliding_window_template():
        """
        Template for sliding window:
        
        left = 0
        for right in range(len(arr)):
            # Add arr[right] to window
            
            while window_invalid:
                # Remove arr[left] from window
                left += 1
            
            # Update result
        """
        pass
    
    @staticmethod
    def monotonic_stack_template():
        """
        Template for monotonic stack:
        
        stack = []
        for i in range(len(arr)):
            while stack and condition(arr[i], arr[stack[-1]]):
                # Process stack.pop()
            
            stack.append(i)
        """
        pass

# PROBLEM RECOGNITION GUIDE

def identify_pattern(problem_description):
    """
    Identify which pattern to use:
    
    Two Pointers:
    - Sorted array
    - Find pair/triplet
    - Opposite ends moving toward middle
    
    Sliding Window:
    - Contiguous subarray/substring
    - All subarrays of size k
    - Longest/shortest with property
    
    Prefix Sum:
    - Range sum queries
    - Subarray sum equals k
    - Multiple queries on static array
    
    Monotonic Stack:
    - Next greater/smaller element
    - Largest rectangle
    - Stock span problem
    
    Top K:
    - K largest/smallest
    - K most frequent
    - Heap or QuickSelect
    """
    pass
```

### Test Cases
- Problems for each pattern
- Variations requiring adaptation
- Combined patterns

### Deliverables
- `two_pointers.{py,go,zig}` - 10+ problems
- `sliding_window.{py,go,zig}` - 10+ problems
- `prefix_sum.{py,go,zig}` - 10+ problems
- `monotonic_stack.{py,go,zig}` - 10+ problems
- `top_k.{py,go,zig}` - 10+ problems
- `pattern_templates.{py,go,zig}` - Reusable templates
- `pattern_recognition_guide.md` - How to identify patterns
- `problem_catalog.md` - 50+ problems categorized

### Success Criteria
- Can recognize patterns in new problems
- Fluent with pattern templates
- Can combine multiple patterns
- Build problem-solving intuition

---

## Module Resources

### CLRS References
- Chapter 4: Divide-and-Conquer
- Chapter 5: Probabilistic Analysis and Randomized Algorithms
- Chapter 17: Amortized Analysis

### Additional Reading
- "Algorithm Design Manual" by Skiena
- Randomized algorithms textbooks
- Amortized analysis papers

### Key Takeaways

By the end of this module, you should:
1. Master divide-and-conquer and recurrence analysis
2. Implement backtracking with effective pruning
3. Understand randomized algorithms and probabilistic analysis
4. Perform amortized analysis using multiple methods
5. Recognize and apply common algorithm patterns
6. Have expanded problem-solving toolkit

### Next Module

**Module 11: String Algorithms** - Master string matching, suffix structures, and pattern matching algorithms.

---

## Tips for Success

1. **Practice recurrence analysis** - Master Theorem is your friend
2. **Visualize backtracking trees** - Understand search space
3. **Run randomized algorithms multiple times** - See probability in action
4. **Do amortized analysis carefully** - Track potential/credits rigorously
5. **Build pattern library** - Recognize templates quickly

## Common Pitfalls

- **Incorrect base cases** in divide-and-conquer
- **Insufficient pruning** in backtracking
- **Confusing Las Vegas and Monte Carlo**
- **Mistaking amortized for average-case**
- **Not adapting patterns** to specific problems
- **Premature optimization** - Get correctness first!
