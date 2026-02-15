# Lab 8.4: Path Reconstruction

**Module**: Module 8 - Dynamic Programming  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 8.4: Path Reconstruction

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Reconstruct optimal solutions, not just values
- Master backtracking through DP tables
- Handle multiple optimal solutions

### Requirements

1. **Reconstruction techniques**:
   - Store parent pointers
   - Backtrack through DP table
   - Iterative vs recursive reconstruction

2. **Implement reconstruction for**:
   - LCS (actual subsequence)
   - Edit distance (sequence of operations)
   - Knapsack (items selected)
   - Shortest path problems
   - Matrix chain (parenthesization)

3. **Multiple solutions**:
   - Find all optimal solutions
   - Count optimal solutions
   - Generate k-best solutions

### Implementation Details

```python
# MATRIX CHAIN PARENTHESIZATION

def matrix_chain_with_parenthesization(dims):
    """
    Matrix chain with optimal parenthesization.
    """
    n = len(dims) - 1
    dp = [[0] * (n + 1) for _ in range(n + 1)]
    split = [[0] * (n + 1) for _ in range(n + 1)]
    
    for length in range(2, n + 1):
        for i in range(1, n - length + 2):
            j = i + length - 1
            dp[i][j] = float('inf')
            
            for k in range(i, j):
                cost = (dp[i][k] + dp[k+1][j] + 
                       dims[i-1] * dims[k] * dims[j])
                if cost < dp[i][j]:
                    dp[i][j] = cost
                    split[i][j] = k
    
    def print_parens(i, j):
        if i == j:
            return f"A{i}"
        k = split[i][j]
        return f"({print_parens(i, k)} x {print_parens(k+1, j)})"
    
    return dp[1][n], print_parens(1, n)

# SUBSET SUM WITH RECONSTRUCTION

def subset_sum_with_elements(nums, target):
    """
    Find subset that sums to target, return the subset.
    """
    n = len(nums)
    dp = [[False] * (target + 1) for _ in range(n + 1)]
    dp[0][0] = True
    
    for i in range(1, n + 1):
        for s in range(target + 1):
            # Don't include nums[i-1]
            dp[i][s] = dp[i-1][s]
            
            # Include nums[i-1]
            if s >= nums[i-1]:
                dp[i][s] = dp[i][s] or dp[i-1][s - nums[i-1]]
    
    if not dp[n][target]:
        return None  # No solution
    
    # Reconstruct subset
    subset = []
    i, s = n, target
    while i > 0 and s > 0:
        if not dp[i-1][s]:  # Must have included nums[i-1]
            subset.append(nums[i-1])
            s -= nums[i-1]
        i -= 1
    
    return subset

# ALL LCS SEQUENCES

def all_lcs(s1, s2):
    """Find all longest common subsequences"""
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = 1 + dp[i-1][j-1]
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    
    # Find all LCS by backtracking
    def backtrack(i, j, current):
        if i == 0 or j == 0:
            result.add(current[::-1])
            return
        
        if s1[i-1] == s2[j-1]:
            backtrack(i-1, j-1, current + s1[i-1])
        else:
            if dp[i-1][j] == dp[i][j]:
                backtrack(i-1, j, current)
            if dp[i][j-1] == dp[i][j]:
                backtrack(i, j-1, current)
    
    result = set()
    backtrack(m, n, "")
    return list(result)

# COUNT PATHS IN DP

def count_paths_in_grid(grid):
    """
    Count paths from top-left to bottom-right.
    Can move right or down. 0 = obstacle, 1 = free.
    """
    if not grid or grid[0][0] == 0:
        return 0
    
    m, n = len(grid), len(grid[0])
    dp = [[0] * n for _ in range(m)]
    dp[0][0] = 1
    
    for i in range(m):
        for j in range(n):
            if grid[i][j] == 0:
                continue
            
            if i > 0:
                dp[i][j] += dp[i-1][j]
            if j > 0:
                dp[i][j] += dp[i][j-1]
    
    return dp[m-1][n-1]

def all_paths_in_grid(grid):
    """
    Return all paths from top-left to bottom-right.
    """
    if not grid or grid[0][0] == 0:
        return []
    
    m, n = len(grid), len(grid[0])
    paths = []
    
    def dfs(i, j, path):
        if i == m - 1 and j == n - 1:
            paths.append(path[:])
            return
        
        if i + 1 < m and grid[i+1][j] == 1:
            path.append('D')
            dfs(i+1, j, path)
            path.pop()
        
        if j + 1 < n and grid[i][j+1] == 1:
            path.append('R')
            dfs(i, j+1, path)
            path.pop()
    
    dfs(0, 0, [])
    return paths

# K-BEST SOLUTIONS

def k_best_subsets(nums, target, k):
    """
    Find k best subsets that sum closest to target.
    """
    import heapq
    
    n = len(nums)
    # dp[i][s] = list of (subset, sum) achieving sum s using first i elements
    dp = [[[] for _ in range(target + 1)] for _ in range(n + 1)]
    dp[0][0] = [([], 0)]
    
    for i in range(1, n + 1):
        for s in range(target + 1):
            # Don't include nums[i-1]
            dp[i][s] = dp[i-1][s][:]
            
            # Include nums[i-1]
            if s >= nums[i-1]:
                for subset, sum_val in dp[i-1][s - nums[i-1]]:
                    new_subset = subset + [nums[i-1]]
                    new_sum = sum_val + nums[i-1]
                    dp[i][s].append((new_subset, new_sum))
            
            # Keep only k best for this state
            dp[i][s].sort(key=lambda x: abs(target - x[1]))
            dp[i][s] = dp[i][s][:k]
    
    # Find k best overall
    all_solutions = []
    for s in range(target + 1):
        all_solutions.extend(dp[n][s])
    
    all_solutions.sort(key=lambda x: abs(target - x[1]))
    return all_solutions[:k]
```

### Test Cases
- Problems with unique optimal solution
- Problems with multiple optimal solutions
- Verify reconstructed solutions are optimal
- Test k-best generation

### Deliverables
- `solution_reconstruction.{py,go,zig}` - All reconstruction techniques
- `multiple_solutions.{py,go,zig}` - Find all optimal solutions
- `k_best.{py,go,zig}` - Generate k-best solutions
- `reconstruction_patterns.md` - Common patterns
- `test_suite.{py,go,zig}` - Verify reconstructions

### Success Criteria
- All reconstructions produce optimal solutions
- Can find multiple solutions when they exist
- Understanding of backtracking through DP tables
- k-best generation works correctly

---

---

## Back to Module

[← Back to Module 8: Dynamic Programming](../module_08/README.md)

## Navigation

- [Previous Lab](./lab_8_3.md) (if exists)
- [Next Lab](./lab_8_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 8, Lab 4 of 7*
