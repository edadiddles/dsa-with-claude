# Module 8: Dynamic Programming

**Duration**: 4 weeks  
**Difficulty**: Medium to Hard

## Overview

Dynamic Programming (DP) is one of the most powerful algorithmic paradigms. It's a method for solving complex problems by breaking them down into simpler subproblems, solving each subproblem once, and storing the results. This module provides extensive practice with DP through dozens of problems, from classic examples to advanced patterns.

**Why this matters**: DP appears everywhere - optimization problems, resource allocation, sequence analysis, path finding, game theory, and more. Mastering DP transforms seemingly intractable problems into solvable ones. It's also one of the most common interview topics.

## Learning Objectives

- Recognize problems with optimal substructure and overlapping subproblems
- Master both top-down (memoization) and bottom-up (tabulation) approaches
- Optimize DP solutions for space complexity
- Reconstruct solutions, not just compute optimal values
- Identify and apply common DP patterns
- Build intuition for when DP is the right approach

## Topics Covered

- DP fundamentals: optimal substructure, overlapping subproblems
- Memoization vs. tabulation
- Classic DP problems (LCS, edit distance, knapsack, matrix chain)
- Space optimization techniques
- Path/solution reconstruction
- Advanced patterns (bitmask DP, DP on trees, interval DP)
- State space design and transition functions

---

## Lab 8.1: DP Foundations & Classic Problems

**Duration**: 6-8 hours  
**Difficulty**: Medium

### Objectives
- Understand DP fundamentals through classic problems
- Implement both memoization and tabulation
- Recognize optimal substructure

### Requirements

1. **Implement these classic DP problems**:
   - **Fibonacci sequence** (warmup)
   - **Longest Common Subsequence (LCS)**
   - **Longest Increasing Subsequence (LIS)**
   - **Edit Distance (Levenshtein)**
   - **0/1 Knapsack**
   - **Coin Change** (minimum coins, number of ways)

2. **For each problem**:
   - Naive recursive solution
   - Top-down with memoization
   - Bottom-up tabulation
   - Analysis of time and space complexity

3. **Understand the structure**:
   - Define state
   - Identify recurrence relation
   - Determine base cases
   - Prove optimal substructure

### Implementation Details

```python
# FIBONACCI - Warmup Problem

def fibonacci_recursive(n):
    """Naive recursive - O(2^n)"""
    if n <= 1:
        return n
    return fibonacci_recursive(n-1) + fibonacci_recursive(n-2)

def fibonacci_memoized(n, memo=None):
    """Top-down with memoization - O(n)"""
    if memo is None:
        memo = {}
    
    if n in memo:
        return memo[n]
    
    if n <= 1:
        return n
    
    memo[n] = fibonacci_memoized(n-1, memo) + fibonacci_memoized(n-2, memo)
    return memo[n]

def fibonacci_tabulation(n):
    """Bottom-up tabulation - O(n)"""
    if n <= 1:
        return n
    
    dp = [0] * (n + 1)
    dp[1] = 1
    
    for i in range(2, n + 1):
        dp[i] = dp[i-1] + dp[i-2]
    
    return dp[n]

def fibonacci_optimized(n):
    """Space-optimized - O(1) space"""
    if n <= 1:
        return n
    
    prev2, prev1 = 0, 1
    
    for i in range(2, n + 1):
        current = prev1 + prev2
        prev2, prev1 = prev1, current
    
    return prev1

# LONGEST COMMON SUBSEQUENCE

def lcs_recursive(s1, s2, i, j):
    """Naive recursive - O(2^(m+n))"""
    if i == 0 or j == 0:
        return 0
    
    if s1[i-1] == s2[j-1]:
        return 1 + lcs_recursive(s1, s2, i-1, j-1)
    else:
        return max(lcs_recursive(s1, s2, i-1, j), 
                   lcs_recursive(s1, s2, i, j-1))

def lcs_memoized(s1, s2):
    """Top-down with memoization - O(m*n)"""
    memo = {}
    
    def helper(i, j):
        if i == 0 or j == 0:
            return 0
        
        if (i, j) in memo:
            return memo[(i, j)]
        
        if s1[i-1] == s2[j-1]:
            result = 1 + helper(i-1, j-1)
        else:
            result = max(helper(i-1, j), helper(i, j-1))
        
        memo[(i, j)] = result
        return result
    
    return helper(len(s1), len(s2))

def lcs_tabulation(s1, s2):
    """Bottom-up tabulation - O(m*n)"""
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = 1 + dp[i-1][j-1]
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    
    return dp[m][n]

def lcs_with_sequence(s1, s2):
    """LCS with actual sequence reconstruction"""
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = 1 + dp[i-1][j-1]
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    
    # Reconstruct sequence
    lcs = []
    i, j = m, n
    while i > 0 and j > 0:
        if s1[i-1] == s2[j-1]:
            lcs.append(s1[i-1])
            i -= 1
            j -= 1
        elif dp[i-1][j] > dp[i][j-1]:
            i -= 1
        else:
            j -= 1
    
    return dp[m][n], ''.join(reversed(lcs))

# LONGEST INCREASING SUBSEQUENCE

def lis_dp(arr):
    """LIS using DP - O(n^2)"""
    if not arr:
        return 0
    
    n = len(arr)
    dp = [1] * n  # dp[i] = length of LIS ending at i
    
    for i in range(1, n):
        for j in range(i):
            if arr[j] < arr[i]:
                dp[i] = max(dp[i], dp[j] + 1)
    
    return max(dp)

def lis_binary_search(arr):
    """LIS using binary search - O(n log n)"""
    if not arr:
        return 0
    
    import bisect
    tails = []  # tails[i] = smallest tail of LIS of length i+1
    
    for num in arr:
        pos = bisect.bisect_left(tails, num)
        if pos == len(tails):
            tails.append(num)
        else:
            tails[pos] = num
    
    return len(tails)

# EDIT DISTANCE

def edit_distance(s1, s2):
    """
    Minimum edit distance (Levenshtein distance).
    Operations: insert, delete, replace.
    """
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Base cases
    for i in range(m + 1):
        dp[i][0] = i  # Delete all characters
    for j in range(n + 1):
        dp[0][j] = j  # Insert all characters
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = dp[i-1][j-1]  # No operation needed
            else:
                dp[i][j] = 1 + min(
                    dp[i-1][j],    # Delete
                    dp[i][j-1],    # Insert
                    dp[i-1][j-1]   # Replace
                )
    
    return dp[m][n]

def edit_distance_with_operations(s1, s2):
    """Edit distance with operation reconstruction"""
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    ops = [[None] * (n + 1) for _ in range(m + 1)]
    
    # Base cases
    for i in range(m + 1):
        dp[i][0] = i
        if i > 0:
            ops[i][0] = ('delete', s1[i-1])
    for j in range(n + 1):
        dp[0][j] = j
        if j > 0:
            ops[0][j] = ('insert', s2[j-1])
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = dp[i-1][j-1]
                ops[i][j] = ('match', s1[i-1])
            else:
                delete_cost = dp[i-1][j]
                insert_cost = dp[i][j-1]
                replace_cost = dp[i-1][j-1]
                
                min_cost = min(delete_cost, insert_cost, replace_cost)
                dp[i][j] = 1 + min_cost
                
                if min_cost == replace_cost:
                    ops[i][j] = ('replace', s1[i-1], s2[j-1])
                elif min_cost == delete_cost:
                    ops[i][j] = ('delete', s1[i-1])
                else:
                    ops[i][j] = ('insert', s2[j-1])
    
    # Reconstruct operations
    operations = []
    i, j = m, n
    while i > 0 or j > 0:
        op = ops[i][j]
        if op[0] == 'match':
            i -= 1
            j -= 1
        elif op[0] == 'replace':
            operations.append(op)
            i -= 1
            j -= 1
        elif op[0] == 'delete':
            operations.append(op)
            i -= 1
        else:  # insert
            operations.append(op)
            j -= 1
    
    return dp[m][n], list(reversed(operations))

# 0/1 KNAPSACK

def knapsack_01(weights, values, capacity):
    """
    0/1 Knapsack problem.
    Each item can be taken at most once.
    """
    n = len(weights)
    dp = [[0] * (capacity + 1) for _ in range(n + 1)]
    
    for i in range(1, n + 1):
        for w in range(capacity + 1):
            # Don't take item i-1
            dp[i][w] = dp[i-1][w]
            
            # Take item i-1 if it fits
            if weights[i-1] <= w:
                dp[i][w] = max(dp[i][w], 
                              values[i-1] + dp[i-1][w - weights[i-1]])
    
    return dp[n][capacity]

def knapsack_01_with_items(weights, values, capacity):
    """0/1 Knapsack with item reconstruction"""
    n = len(weights)
    dp = [[0] * (capacity + 1) for _ in range(n + 1)]
    
    # Fill DP table
    for i in range(1, n + 1):
        for w in range(capacity + 1):
            dp[i][w] = dp[i-1][w]
            if weights[i-1] <= w:
                dp[i][w] = max(dp[i][w], 
                              values[i-1] + dp[i-1][w - weights[i-1]])
    
    # Reconstruct items
    items = []
    w = capacity
    for i in range(n, 0, -1):
        if dp[i][w] != dp[i-1][w]:
            items.append(i-1)
            w -= weights[i-1]
    
    return dp[n][capacity], list(reversed(items))

# COIN CHANGE

def coin_change_min(coins, amount):
    """
    Minimum number of coins to make amount.
    Returns -1 if impossible.
    """
    dp = [float('inf')] * (amount + 1)
    dp[0] = 0
    
    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i:
                dp[i] = min(dp[i], 1 + dp[i - coin])
    
    return dp[amount] if dp[amount] != float('inf') else -1

def coin_change_ways(coins, amount):
    """
    Number of ways to make amount using coins.
    """
    dp = [0] * (amount + 1)
    dp[0] = 1  # One way to make 0: use no coins
    
    for coin in coins:
        for i in range(coin, amount + 1):
            dp[i] += dp[i - coin]
    
    return dp[amount]

def coin_change_with_coins(coins, amount):
    """Coin change with actual coins used"""
    dp = [float('inf')] * (amount + 1)
    parent = [-1] * (amount + 1)
    dp[0] = 0
    
    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i and 1 + dp[i - coin] < dp[i]:
                dp[i] = 1 + dp[i - coin]
                parent[i] = coin
    
    if dp[amount] == float('inf'):
        return -1, []
    
    # Reconstruct coins
    coins_used = []
    current = amount
    while current > 0:
        coin = parent[current]
        coins_used.append(coin)
        current -= coin
    
    return dp[amount], coins_used
```

### Test Cases
- Small inputs (verify by hand)
- Edge cases (empty strings, zero capacity, etc.)
- Large inputs (performance testing)
- Compare memoization vs. tabulation performance
- Verify solution reconstruction

### Deliverables
- `classic_dp.{py,go,zig}` - All classic problems
- `memoization_vs_tabulation.{py,go,zig}` - Comparison framework
- `solution_reconstruction.{py,go,zig}` - Path/solution recovery
- `dp_visualizer.{py,go,zig}` - Visualize DP tables
- `complexity_analysis.md` - Time/space analysis for each
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All problems solved correctly
- Understanding of recurrence relations
- Can explain optimal substructure for each
- Memoization and tabulation both work
- Solution reconstruction works

---

## Lab 8.2: Memoization vs. Tabulation

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Deeply understand top-down vs. bottom-up approaches
- Know when to use each
- Master conversion between approaches

### Requirements

1. **Implement same problems both ways**:
   - Choose 5+ DP problems
   - Implement with memoization
   - Implement with tabulation
   - Compare performance

2. **Conversion practice**:
   - Given memoized solution, convert to tabulation
   - Given tabulated solution, convert to memoization
   - Understand order of computation

3. **Analysis**:
   - Stack space usage (recursion depth)
   - Cache performance
   - Code complexity
   - When each is preferable

### Implementation Details

```python
class DPConverter:
    """Tools for converting between memoization and tabulation"""
    
    @staticmethod
    def memoize_function(func):
        """Generic memoization decorator"""
        cache = {}
        
        def memoized(*args):
            if args in cache:
                return cache[args]
            result = func(*args)
            cache[args] = result
            return result
        
        memoized.cache = cache
        return memoized
    
    @staticmethod
    def analyze_recursion_depth(func, *args):
        """Measure maximum recursion depth"""
        max_depth = [0]
        current_depth = [0]
        
        original_func = func
        
        def wrapper(*args):
            current_depth[0] += 1
            max_depth[0] = max(max_depth[0], current_depth[0])
            result = original_func(*args)
            current_depth[0] -= 1
            return result
        
        # Monkey patch temporarily
        import sys
        old_limit = sys.getrecursionlimit()
        sys.setrecursionlimit(10000)
        
        try:
            wrapper(*args)
        finally:
            sys.setrecursionlimit(old_limit)
        
        return max_depth[0]

# MATRIX CHAIN MULTIPLICATION

def matrix_chain_memoized(dims):
    """
    Matrix chain multiplication (top-down).
    dims[i] gives dimensions of matrix i: dims[i-1] x dims[i]
    """
    n = len(dims) - 1
    memo = {}
    
    def mcm(i, j):
        if i == j:
            return 0
        
        if (i, j) in memo:
            return memo[(i, j)]
        
        min_cost = float('inf')
        for k in range(i, j):
            cost = (mcm(i, k) + 
                   mcm(k+1, j) + 
                   dims[i-1] * dims[k] * dims[j])
            min_cost = min(min_cost, cost)
        
        memo[(i, j)] = min_cost
        return min_cost
    
    return mcm(1, n)

def matrix_chain_tabulation(dims):
    """Matrix chain multiplication (bottom-up)"""
    n = len(dims) - 1
    dp = [[0] * (n + 1) for _ in range(n + 1)]
    
    # length is chain length
    for length in range(2, n + 1):
        for i in range(1, n - length + 2):
            j = i + length - 1
            dp[i][j] = float('inf')
            
            for k in range(i, j):
                cost = (dp[i][k] + 
                       dp[k+1][j] + 
                       dims[i-1] * dims[k] * dims[j])
                dp[i][j] = min(dp[i][j], cost)
    
    return dp[1][n]

# ROD CUTTING

def rod_cutting_memoized(prices, n):
    """
    Rod cutting problem (top-down).
    prices[i] is price of rod of length i+1.
    """
    memo = {}
    
    def cut(length):
        if length == 0:
            return 0
        
        if length in memo:
            return memo[length]
        
        max_revenue = 0
        for i in range(min(length, len(prices))):
            revenue = prices[i] + cut(length - i - 1)
            max_revenue = max(max_revenue, revenue)
        
        memo[length] = max_revenue
        return max_revenue
    
    return cut(n)

def rod_cutting_tabulation(prices, n):
    """Rod cutting problem (bottom-up)"""
    dp = [0] * (n + 1)
    
    for length in range(1, n + 1):
        max_revenue = 0
        for i in range(min(length, len(prices))):
            revenue = prices[i] + dp[length - i - 1]
            max_revenue = max(max_revenue, revenue)
        dp[length] = max_revenue
    
    return dp[n]

# HOUSE ROBBER

def house_robber_memoized(nums):
    """
    House robber (top-down).
    Cannot rob adjacent houses.
    """
    memo = {}
    
    def rob(i):
        if i >= len(nums):
            return 0
        
        if i in memo:
            return memo[i]
        
        # Either rob this house or skip it
        memo[i] = max(nums[i] + rob(i + 2),  # Rob current
                      rob(i + 1))              # Skip current
        return memo[i]
    
    return rob(0)

def house_robber_tabulation(nums):
    """House robber (bottom-up)"""
    if not nums:
        return 0
    if len(nums) == 1:
        return nums[0]
    
    n = len(nums)
    dp = [0] * n
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])
    
    for i in range(2, n):
        dp[i] = max(nums[i] + dp[i-2],  # Rob current
                    dp[i-1])              # Skip current
    
    return dp[n-1]

# PERFORMANCE COMPARISON

def compare_approaches(problem_name, memoized_func, tabulated_func, *args):
    """Compare memoization vs tabulation"""
    import time
    
    # Benchmark memoization
    start = time.time()
    memo_result = memoized_func(*args)
    memo_time = time.time() - start
    
    # Benchmark tabulation
    start = time.time()
    tab_result = tabulated_func(*args)
    tab_time = time.time() - start
    
    # Check correctness
    assert memo_result == tab_result, "Results don't match!"
    
    print(f"\n{problem_name}:")
    print(f"  Memoization: {memo_time:.6f}s")
    print(f"  Tabulation:  {tab_time:.6f}s")
    print(f"  Result: {memo_result}")
    
    if memo_time < tab_time:
        print(f"  Winner: Memoization ({tab_time/memo_time:.2f}x faster)")
    else:
        print(f"  Winner: Tabulation ({memo_time/tab_time:.2f}x faster)")
```

### Test Cases
- Same inputs for both approaches (verify same results)
- Large inputs (stress test)
- Measure stack depth for memoization
- Profile cache hit rates

### Deliverables
- `memoization_examples.{py,go,zig}` - 5+ problems with memoization
- `tabulation_examples.{py,go,zig}` - Same 5+ problems with tabulation
- `conversion_guide.md` - How to convert between approaches
- `performance_comparison.{py,go,zig}` - Benchmark both
- `when_to_use_what.md` - Decision guide

### Success Criteria
- Can convert any problem between approaches
- Understanding of order of computation
- Knowledge of when each is preferable
- Performance analysis shows expected results

---

## Lab 8.3: Space Optimization Techniques

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Reduce space complexity of DP solutions
- Understand space-time tradeoffs
- Recognize optimization opportunities

### Requirements

1. **Space optimization patterns**:
   - Reduce 2D DP to 1D
   - Sliding window approach
   - Alternating arrays
   - State compression

2. **Implement optimized versions**:
   - LCS with O(min(m,n)) space
   - Knapsack with O(capacity) space
   - Edit distance with O(min(m,n)) space

3. **Trade-off analysis**:
   - When optimization is possible
   - Impact on solution reconstruction
   - Code complexity increase

### Implementation Details

```python
# LCS SPACE OPTIMIZATIONS

def lcs_space_optimized(s1, s2):
    """
    LCS with O(min(m,n)) space instead of O(m*n).
    Can only compute length, not reconstruct sequence.
    """
    # Make s1 the shorter string
    if len(s1) > len(s2):
        s1, s2 = s2, s1
    
    m, n = len(s1), len(s2)
    prev = [0] * (m + 1)
    curr = [0] * (m + 1)
    
    for j in range(1, n + 1):
        for i in range(1, m + 1):
            if s1[i-1] == s2[j-1]:
                curr[i] = 1 + prev[i-1]
            else:
                curr[i] = max(prev[i], curr[i-1])
        
        prev, curr = curr, prev
    
    return prev[m]

def lcs_hirschberg(s1, s2):
    """
    Hirschberg's algorithm: LCS with O(min(m,n)) space
    AND sequence reconstruction.
    Uses divide-and-conquer.
    """
    def lcs_length(s1, s2):
        """Compute LCS length row by row"""
        m, n = len(s1), len(s2)
        prev = [0] * (n + 1)
        
        for i in range(1, m + 1):
            curr = [0] * (n + 1)
            for j in range(1, n + 1):
                if s1[i-1] == s2[j-1]:
                    curr[j] = 1 + prev[j-1]
                else:
                    curr[j] = max(prev[j], curr[j-1])
            prev = curr
        
        return prev
    
    def hirschberg_recursive(s1, s2):
        if not s1:
            return ""
        elif len(s1) == 1:
            return s1 if s1[0] in s2 else ""
        
        mid = len(s1) // 2
        
        # Compute LCS lengths for first half and second half
        left_lengths = lcs_length(s1[:mid], s2)
        right_lengths = lcs_length(s1[mid:][::-1], s2[::-1])[::-1]
        
        # Find optimal split point
        k = 0
        max_val = left_lengths[0] + right_lengths[0]
        for j in range(len(s2) + 1):
            if left_lengths[j] + right_lengths[j] > max_val:
                max_val = left_lengths[j] + right_lengths[j]
                k = j
        
        # Recursively solve subproblems
        return (hirschberg_recursive(s1[:mid], s2[:k]) +
                hirschberg_recursive(s1[mid:], s2[k:]))
    
    return hirschberg_recursive(s1, s2)

# KNAPSACK SPACE OPTIMIZATIONS

def knapsack_space_optimized(weights, values, capacity):
    """
    0/1 Knapsack with O(capacity) space instead of O(n*capacity).
    Process items in reverse to avoid overwriting needed values.
    """
    dp = [0] * (capacity + 1)
    
    for i in range(len(weights)):
        # Process in reverse to avoid using updated values
        for w in range(capacity, weights[i] - 1, -1):
            dp[w] = max(dp[w], values[i] + dp[w - weights[i]])
    
    return dp[capacity]

def knapsack_two_arrays(weights, values, capacity):
    """
    Knapsack with two alternating arrays.
    Clearer than single array approach.
    """
    prev = [0] * (capacity + 1)
    curr = [0] * (capacity + 1)
    
    for i in range(len(weights)):
        for w in range(capacity + 1):
            curr[w] = prev[w]  # Don't take item
            if weights[i] <= w:
                curr[w] = max(curr[w], values[i] + prev[w - weights[i]])
        
        prev, curr = curr, prev
    
    return prev[capacity]

# EDIT DISTANCE SPACE OPTIMIZATION

def edit_distance_space_optimized(s1, s2):
    """
    Edit distance with O(min(m,n)) space.
    """
    # Make s2 the shorter string (use fewer columns)
    if len(s1) < len(s2):
        s1, s2 = s2, s1
    
    m, n = len(s1), len(s2)
    prev = list(range(n + 1))
    
    for i in range(1, m + 1):
        curr = [i]  # First column is always i
        
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                curr.append(prev[j-1])
            else:
                curr.append(1 + min(prev[j],      # Delete
                                    curr[j-1],     # Insert
                                    prev[j-1]))    # Replace
        
        prev = curr
    
    return prev[n]

# PALINDROME PARTITIONING

def palindrome_partition_space_optimized(s):
    """
    Minimum cuts to partition string into palindromes.
    Optimize using palindrome lookup table.
    """
    n = len(s)
    
    # Precompute palindrome table: O(n^2) space but saves recomputation
    is_palindrome = [[False] * n for _ in range(n)]
    for i in range(n):
        is_palindrome[i][i] = True
    
    for length in range(2, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            if s[i] == s[j]:
                is_palindrome[i][j] = (length == 2 or is_palindrome[i+1][j-1])
    
    # DP for minimum cuts: O(n) space
    dp = [0] * n
    
    for i in range(n):
        if is_palindrome[0][i]:
            dp[i] = 0  # Entire prefix is palindrome
        else:
            dp[i] = float('inf')
            for j in range(i):
                if is_palindrome[j+1][i]:
                    dp[i] = min(dp[i], 1 + dp[j])
    
    return dp[n-1]

# STATE COMPRESSION (BITMASK DP)

def traveling_salesman_bitmask(dist):
    """
    TSP using bitmask for space optimization.
    State: (current_city, visited_set)
    Instead of 2^n * n array, use dict with (city, mask) as key.
    """
    n = len(dist)
    memo = {}
    
    def tsp(pos, visited):
        # visited is a bitmask
        if visited == (1 << n) - 1:  # All cities visited
            return dist[pos][0]  # Return to start
        
        if (pos, visited) in memo:
            return memo[(pos, visited)]
        
        min_cost = float('inf')
        for next_city in range(n):
            if not (visited & (1 << next_city)):  # Not visited
                new_visited = visited | (1 << next_city)
                cost = dist[pos][next_city] + tsp(next_city, new_visited)
                min_cost = min(min_cost, cost)
        
        memo[(pos, visited)] = min_cost
        return min_cost
    
    return tsp(0, 1)  # Start at city 0
```

### Test Cases
- Verify optimized versions match unoptimized
- Large inputs (demonstrate space savings)
- Memory profiling
- Benchmark time impact

### Deliverables
- `space_optimized_dp.{py,go,zig}` - All optimizations
- `hirschberg.{py,go,zig}` - Hirschberg's algorithm
- `bitmask_dp.{py,go,zig}` - State compression examples
- `optimization_guide.md` - When and how to optimize
- `memory_profiler.{py,go,zig}` - Measure space usage

### Success Criteria
- Space complexity reduced as expected
- Understanding of optimization techniques
- Can identify optimization opportunities
- Knowledge of trade-offs

---

## Lab 8.4: Path Reconstruction

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

## Lab 8.5: Advanced DP Patterns

**Duration**: 6-8 hours  
**Difficulty**: Hard

### Objectives
- Master advanced DP patterns
- Solve complex multi-dimensional DP problems
- Recognize pattern applications

### Requirements

1. **Advanced patterns**:
   - **Bitmask DP**: TSP, assignment problems
   - **DP on trees**: Tree diameter, vertex cover
   - **Interval DP**: Burst balloons, matrix chain
   - **Digit DP**: Count numbers with properties
   - **DP with states**: Finite automata patterns

2. **Implement challenging problems**:
   - Traveling Salesman Problem
   - Burst Balloons
   - Regular Expression Matching
   - Wildcard Matching
   - Egg Drop Problem

### Implementation Details

```python
# BITMASK DP - TSP

def tsp_dp(dist):
    """
    Traveling Salesman Problem using bitmask DP.
    Time: O(n^2 * 2^n), Space: O(n * 2^n)
    """
    n = len(dist)
    dp = [[float('inf')] * n for _ in range(1 << n)]
    dp[1][0] = 0  # Start at city 0
    
    for mask in range(1 << n):
        for u in range(n):
            if not (mask & (1 << u)):
                continue
            
            for v in range(n):
                if mask & (1 << v):
                    continue
                
                new_mask = mask | (1 << v)
                dp[new_mask][v] = min(dp[new_mask][v],
                                     dp[mask][u] + dist[u][v])
    
    # Return to start
    full_mask = (1 << n) - 1
    min_cost = float('inf')
    for u in range(n):
        min_cost = min(min_cost, dp[full_mask][u] + dist[u][0])
    
    return min_cost

# DP ON TREES

def tree_diameter(adj_list, root=0):
    """
    Find diameter of tree using DP.
    Diameter = longest path between any two nodes.
    """
    diameter = [0]
    
    def dfs(node, parent):
        # Returns longest path from this node downward
        max1, max2 = 0, 0  # Two longest paths from this node
        
        for child in adj_list[node]:
            if child != parent:
                child_depth = dfs(child, node)
                if child_depth > max1:
                    max2, max1 = max1, child_depth
                elif child_depth > max2:
                    max2 = child_depth
        
        # Update diameter (path through this node)
        diameter[0] = max(diameter[0], max1 + max2)
        
        return max1 + 1
    
    dfs(root, -1)
    return diameter[0]

def min_vertex_cover_tree(adj_list, root=0):
    """
    Minimum vertex cover on tree.
    dp[node][0] = min cover if node not included
    dp[node][1] = min cover if node included
    """
    dp = {}
    
    def dfs(node, parent):
        # Not include this node: must include all children
        not_include = 0
        # Include this node: can choose for children
        include = 1
        
        for child in adj_list[node]:
            if child != parent:
                child_dp = dfs(child, node)
                not_include += child_dp[1]  # Must include child
                include += min(child_dp[0], child_dp[1])
        
        dp[node] = (not_include, include)
        return dp[node]
    
    result = dfs(root, -1)
    return min(result)

# INTERVAL DP

def burst_balloons(nums):
    """
    Burst balloons to maximize coins.
    Classic interval DP problem.
    """
    # Add boundary balloons
    nums = [1] + nums + [1]
    n = len(nums)
    dp = [[0] * n for _ in range(n)]
    
    # length is the interval length
    for length in range(2, n):
        for left in range(n - length):
            right = left + length
            
            # Try bursting each balloon last in this interval
            for i in range(left + 1, right):
                coins = (nums[left] * nums[i] * nums[right] +
                        dp[left][i] + dp[i][right])
                dp[left][right] = max(dp[left][right], coins)
    
    return dp[0][n-1]

def optimal_bst(keys, freq):
    """
    Optimal Binary Search Tree.
    Minimize search cost given search frequencies.
    """
    n = len(keys)
    dp = [[0] * n for _ in range(n)]
    sum_freq = [[0] * n for _ in range(n)]
    
    # Precompute frequency sums
    for i in range(n):
        sum_freq[i][i] = freq[i]
        for j in range(i + 1, n):
            sum_freq[i][j] = sum_freq[i][j-1] + freq[j]
    
    # Single key trees
    for i in range(n):
        dp[i][i] = freq[i]
    
    # Larger intervals
    for length in range(2, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            dp[i][j] = float('inf')
            
            # Try each key as root
            for r in range(i, j + 1):
                cost = (dp[i][r-1] if r > i else 0) + \
                       (dp[r+1][j] if r < j else 0) + \
                       sum_freq[i][j]
                dp[i][j] = min(dp[i][j], cost)
    
    return dp[0][n-1]

# DIGIT DP

def count_numbers_with_digit_sum(n, target_sum):
    """
    Count numbers from 1 to n with digit sum = target_sum.
    Classic digit DP.
    """
    digits = [int(d) for d in str(n)]
    memo = {}
    
    def dp(pos, current_sum, tight):
        # pos: current digit position
        # current_sum: sum of digits so far
        # tight: whether we're still bounded by n
        
        if current_sum > target_sum:
            return 0
        
        if pos == len(digits):
            return 1 if current_sum == target_sum else 0
        
        if (pos, current_sum, tight) in memo:
            return memo[(pos, current_sum, tight)]
        
        limit = digits[pos] if tight else 9
        result = 0
        
        for digit in range(0, limit + 1):
            new_tight = tight and (digit == limit)
            result += dp(pos + 1, current_sum + digit, new_tight)
        
        memo[(pos, current_sum, tight)] = result
        return result
    
    return dp(0, 0, True)

# REGULAR EXPRESSION MATCHING

def regex_match(s, p):
    """
    Regular expression matching with . and *.
    . matches any single character
    * matches zero or more of preceding element
    """
    m, n = len(s), len(p)
    dp = [[False] * (n + 1) for _ in range(m + 1)]
    dp[0][0] = True
    
    # Handle patterns like a*, a*b*, etc. that match empty string
    for j in range(2, n + 1, 2):
        if p[j-1] == '*':
            dp[0][j] = dp[0][j-2]
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if p[j-1] == '*':
                # Match zero occurrences
                dp[i][j] = dp[i][j-2]
                
                # Match one or more occurrences
                if p[j-2] == '.' or p[j-2] == s[i-1]:
                    dp[i][j] = dp[i][j] or dp[i-1][j]
            elif p[j-1] == '.' or p[j-1] == s[i-1]:
                dp[i][j] = dp[i-1][j-1]
    
    return dp[m][n]

# EGG DROP

def egg_drop(eggs, floors):
    """
    Egg drop problem: find minimum trials needed
    to find critical floor with given number of eggs.
    """
    dp = [[float('inf')] * (eggs + 1) for _ in range(floors + 1)]
    
    # Base cases
    for e in range(eggs + 1):
        dp[0][e] = 0  # No floors = no trials
        dp[1][e] = 1  # One floor = one trial
    
    for f in range(floors + 1):
        dp[f][1] = f  # One egg = f trials (linear search)
    
    for f in range(2, floors + 1):
        for e in range(2, eggs + 1):
            # Try dropping from each floor
            for x in range(1, f + 1):
                # Egg breaks: check floors below with e-1 eggs
                # Egg doesn't break: check floors above with e eggs
                worst = 1 + max(dp[x-1][e-1], dp[f-x][e])
                dp[f][e] = min(dp[f][e], worst)
    
    return dp[floors][eggs]

def egg_drop_optimized(eggs, floors):
    """
    Egg drop with binary search optimization.
    Time: O(eggs * floors * log(floors))
    """
    dp = [[0] * (eggs + 1) for _ in range(floors + 1)]
    
    for f in range(1, floors + 1):
        dp[f][1] = f
    
    for e in range(1, eggs + 1):
        dp[1][e] = 1
    
    for f in range(2, floors + 1):
        for e in range(2, eggs + 1):
            # Binary search for optimal floor to drop from
            left, right = 1, f
            
            while left + 1 < right:
                mid = (left + right) // 2
                
                # dp[mid-1][e-1] increases as mid increases (egg breaks)
                # dp[f-mid][e] decreases as mid increases (egg doesn't break)
                breaks = dp[mid-1][e-1]
                no_break = dp[f-mid][e]
                
                if breaks > no_break:
                    right = mid
                else:
                    left = mid
            
            dp[f][e] = 1 + min(max(dp[x-1][e-1], dp[f-x][e]) 
                              for x in [left, right])
    
    return dp[floors][eggs]
```

### Test Cases
- TSP on small graphs (verify by brute force)
- Tree problems on various tree structures
- Interval DP with different inputs
- Digit DP with various constraints
- Regex matching with complex patterns

### Deliverables
- `bitmask_dp.{py,go,zig}` - TSP and related problems
- `tree_dp.{py,go,zig}` - DP on trees
- `interval_dp.{py,go,zig}` - Interval-based problems
- `digit_dp.{py,go,zig}` - Digit DP examples
- `regex_matching.{py,go,zig}` - Pattern matching problems
- `advanced_patterns.md` - Pattern recognition guide
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All advanced problems solved correctly
- Understanding of when each pattern applies
- Can recognize patterns in new problems
- Performance matches expected complexity

---

## Lab 8.6: DP Visualization Tools

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Build tools to visualize DP computation
- Aid understanding and debugging
- Create educational animations

### Requirements

1. **Visualization features**:
   - Display DP table/state space
   - Highlight computation order
   - Show dependencies between states
   - Animate filling of DP table

2. **Support for**:
   - 1D DP problems
   - 2D DP problems
   - State space visualization
   - Recurrence visualization

3. **Output formats**:
   - Terminal-based (ASCII)
   - HTML/JavaScript interactive
   - Step-by-step animation

### Implementation Details

```python
import matplotlib.pyplot as plt
import matplotlib.animation as animation
import numpy as np

class DPVisualizer:
    """Visualize DP computation"""
    
    def visualize_table_2d(self, dp, title="DP Table", labels=None):
        """Visualize 2D DP table as heatmap"""
        fig, ax = plt.subplots(figsize=(10, 8))
        
        # Create heatmap
        im = ax.imshow(dp, cmap='YlOrRd', aspect='auto')
        
        # Add colorbar
        plt.colorbar(im, ax=ax)
        
        # Add labels
        if labels:
            ax.set_xlabel(labels.get('x', 'Column'))
            ax.set_ylabel(labels.get('y', 'Row'))
        
        # Add grid
        ax.set_xticks(np.arange(len(dp[0])))
        ax.set_yticks(np.arange(len(dp)))
        ax.grid(which='both', color='black', linewidth=0.5)
        
        # Add values as text
        for i in range(len(dp)):
            for j in range(len(dp[0])):
                val = dp[i][j]
                if val != float('inf'):
                    text = ax.text(j, i, f'{val:.0f}',
                                 ha="center", va="center", color="black")
        
        ax.set_title(title)
        plt.tight_layout()
        return fig
    
    def animate_filling(self, problem_func, *args):
        """
        Animate the filling of DP table.
        problem_func should yield (i, j, value) for each update.
        """
        updates = list(problem_func(*args))
        
        # Determine table size
        max_i = max(u[0] for u in updates) + 1
        max_j = max(u[1] for u in updates) + 1
        
        dp = [[0] * max_j for _ in range(max_i)]
        
        fig, ax = plt.subplots()
        im = ax.imshow(dp, cmap='YlOrRd', vmin=0, 
                      vmax=max(u[2] for u in updates))
        
        def update(frame):
            i, j, val = updates[frame]
            dp[i][j] = val
            im.set_array(dp)
            ax.set_title(f'Step {frame + 1}/{len(updates)}')
            return [im]
        
        ani = animation.FuncAnimation(fig, update, frames=len(updates),
                                     interval=200, blit=True)
        return ani
    
    def print_table_ascii(self, dp, title="DP Table"):
        """Print DP table in ASCII format"""
        print(f"\n{title}")
        print("=" * 50)
        
        if isinstance(dp[0], list):
            # 2D table
            for row in dp:
                print(" ".join(f"{val:5.0f}" if val != float('inf') else "  inf" 
                             for val in row))
        else:
            # 1D table
            print(" ".join(f"{val:5.0f}" if val != float('inf') else "  inf" 
                         for val in dp))
        print()

def visualize_lcs_computation(s1, s2):
    """Visualize LCS computation step by step"""
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    steps = []
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = 1 + dp[i-1][j-1]
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            
            steps.append({
                'i': i,
                'j': j,
                'value': dp[i][j],
                'table': [row[:] for row in dp],
                'match': s1[i-1] == s2[j-1]
            })
    
    return steps

def generate_dp_html(problem_name, steps):
    """Generate interactive HTML visualization"""
    html = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>{problem_name} - DP Visualization</title>
        <style>
            .dp-table {{ border-collapse: collapse; margin: 20px; }}
            .dp-cell {{ 
                border: 1px solid #ccc;
                padding: 10px;
                width: 40px;
                height: 40px;
                text-align: center;
            }}
            .current {{ background-color: #ffeb3b; }}
            .computed {{ background-color: #c8e6c9; }}
            .controls {{ margin: 20px; }}
        </style>
    </head>
    <body>
        <h1>{problem_name}</h1>
        <div class="controls">
            <button onclick="prevStep()">Previous</button>
            <button onclick="nextStep()">Next</button>
            <span id="step-info"></span>
        </div>
        <div id="visualization"></div>
        
        <script>
            const steps = {steps};
            let currentStep = 0;
            
            function renderStep() {{
                // Render logic here
            }}
            
            function nextStep() {{
                if (currentStep < steps.length - 1) {{
                    currentStep++;
                    renderStep();
                }}
            }}
            
            function prevStep() {{
                if (currentStep > 0) {{
                    currentStep--;
                    renderStep();
                }}
            }}
            
            renderStep();
        </script>
    </body>
    </html>
    """
    return html
```

### Test Cases
- Visualize classic DP problems
- Test animation on different problems
- Export to various formats
- Interactive controls work

### Deliverables
- `dp_visualizer.{py,go,zig}` - Core visualization tools
- `ascii_viz.{py,go,zig}` - Terminal visualization
- `matplotlib_viz.{py,go,zig}` - Graphical visualization
- `html_export.{py,go,zig}` - Interactive HTML export
- `animation_examples/` - Sample animations
- `visualization_guide.md` - How to use tools

### Success Criteria
- Clear, understandable visualizations
- Helps with debugging DP solutions
- Useful for learning and teaching
- Works for various DP problems

---

## Lab 8.7: DP Practice Marathon

**Duration**: 8-10 hours  
**Difficulty**: Medium-Hard

### Objectives
- Solve 20+ DP problems of varying difficulty
- Build problem-solving speed
- Recognize patterns quickly

### Requirements

1. **Solve problems across categories**:
   - Fibonacci-style (5 problems)
   - Knapsack variants (5 problems)
   - String/sequence problems (5 problems)
   - Grid/matrix problems (5 problems)
   - Advanced problems (5+ problems)

2. **For each problem**:
   - Identify the DP pattern
   - Define state and recurrence
   - Implement solution
   - Test and verify

3. **Track metrics**:
   - Time to solution
   - Bugs encountered
   - Pattern recognition

### Problem List

```python
# FIBONACCI-STYLE PROBLEMS

def climbing_stairs(n):
    """
    Ways to climb n stairs (1 or 2 steps at a time).
    Classic Fibonacci variant.
    """
    if n <= 2:
        return n
    
    prev2, prev1 = 1, 2
    for i in range(3, n + 1):
        current = prev1 + prev2
        prev2, prev1 = prev1, current
    
    return prev1

def house_robber_circular(nums):
    """
    House robber with houses in a circle.
    Cannot rob first and last together.
    """
    if not nums:
        return 0
    if len(nums) == 1:
        return nums[0]
    
    def rob_linear(houses):
        if not houses:
            return 0
        if len(houses) == 1:
            return houses[0]
        
        prev2, prev1 = houses[0], max(houses[0], houses[1])
        for i in range(2, len(houses)):
            current = max(houses[i] + prev2, prev1)
            prev2, prev1 = prev1, current
        
        return prev1
    
    # Either rob first house (can't rob last) or rob last (can't rob first)
    return max(rob_linear(nums[:-1]), rob_linear(nums[1:]))

# KNAPSACK VARIANTS

def partition_equal_subset(nums):
    """Can partition array into two subsets with equal sum?"""
    total = sum(nums)
    if total % 2 != 0:
        return False
    
    target = total // 2
    dp = [False] * (target + 1)
    dp[0] = True
    
    for num in nums:
        for s in range(target, num - 1, -1):
            dp[s] = dp[s] or dp[s - num]
    
    return dp[target]

def target_sum(nums, target):
    """
    Count ways to assign +/- to nums to sum to target.
    """
    total = sum(nums)
    if (total + target) % 2 != 0 or total < abs(target):
        return 0
    
    # Transform to subset sum problem
    subset_sum = (total + target) // 2
    
    dp = [0] * (subset_sum + 1)
    dp[0] = 1
    
    for num in nums:
        for s in range(subset_sum, num - 1, -1):
            dp[s] += dp[s - num]
    
    return dp[subset_sum]

# STRING PROBLEMS

def palindromic_substrings(s):
    """Count palindromic substrings"""
    n = len(s)
    dp = [[False] * n for _ in range(n)]
    count = 0
    
    # Single characters
    for i in range(n):
        dp[i][i] = True
        count += 1
    
    # Two characters
    for i in range(n - 1):
        if s[i] == s[i+1]:
            dp[i][i+1] = True
            count += 1
    
    # Longer substrings
    for length in range(3, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            if s[i] == s[j] and dp[i+1][j-1]:
                dp[i][j] = True
                count += 1
    
    return count

def distinct_subsequences(s, t):
    """
    Count distinct subsequences of s that equal t.
    """
    m, n = len(s), len(t)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    for i in range(m + 1):
        dp[i][0] = 1  # Empty string is subsequence
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            dp[i][j] = dp[i-1][j]  # Don't use s[i-1]
            if s[i-1] == t[j-1]:
                dp[i][j] += dp[i-1][j-1]  # Use s[i-1]
    
    return dp[m][n]

# GRID PROBLEMS

def unique_paths_with_obstacles(grid):
    """Unique paths in grid with obstacles"""
    if not grid or grid[0][0] == 1:
        return 0
    
    m, n = len(grid), len(grid[0])
    dp = [[0] * n for _ in range(m)]
    dp[0][0] = 1
    
    for i in range(m):
        for j in range(n):
            if grid[i][j] == 1:
                dp[i][j] = 0
            else:
                if i > 0:
                    dp[i][j] += dp[i-1][j]
                if j > 0:
                    dp[i][j] += dp[i][j-1]
    
    return dp[m-1][n-1]

def min_path_sum(grid):
    """Minimum path sum from top-left to bottom-right"""
    m, n = len(grid), len(grid[0])
    dp = [[0] * n for _ in range(m)]
    dp[0][0] = grid[0][0]
    
    for i in range(1, m):
        dp[i][0] = dp[i-1][0] + grid[i][0]
    
    for j in range(1, n):
        dp[0][j] = dp[0][j-1] + grid[0][j]
    
    for i in range(1, m):
        for j in range(1, n):
            dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
    
    return dp[m-1][n-1]

# ADVANCED PROBLEMS

def longest_valid_parentheses(s):
    """Length of longest valid parentheses substring"""
    n = len(s)
    if n < 2:
        return 0
    
    dp = [0] * n
    max_len = 0
    
    for i in range(1, n):
        if s[i] == ')':
            if s[i-1] == '(':
                dp[i] = (dp[i-2] if i >= 2 else 0) + 2
            elif i - dp[i-1] > 0 and s[i - dp[i-1] - 1] == '(':
                dp[i] = dp[i-1] + 2 + (dp[i - dp[i-1] - 2] if i - dp[i-1] >= 2 else 0)
            
            max_len = max(max_len, dp[i])
    
    return max_len

def maximal_rectangle(matrix):
    """Largest rectangle in binary matrix"""
    if not matrix:
        return 0
    
    m, n = len(matrix), len(matrix[0])
    heights = [0] * n
    max_area = 0
    
    def largest_rectangle_histogram(heights):
        stack = []
        max_area = 0
        index = 0
        
        while index < len(heights):
            if not stack or heights[index] >= heights[stack[-1]]:
                stack.append(index)
                index += 1
            else:
                top = stack.pop()
                width = index if not stack else index - stack[-1] - 1
                max_area = max(max_area, heights[top] * width)
        
        while stack:
            top = stack.pop()
            width = index if not stack else index - stack[-1] - 1
            max_area = max(max_area, heights[top] * width)
        
        return max_area
    
    for i in range(m):
        for j in range(n):
            if matrix[i][j] == '1':
                heights[j] += 1
            else:
                heights[j] = 0
        
        max_area = max(max_area, largest_rectangle_histogram(heights))
    
    return max_area
```

### Deliverables
- `dp_problems.{py,go,zig}` - All 20+ problems solved
- `problem_patterns.md` - Pattern recognition guide
- `solving_strategies.md` - Approach for each category
- `performance_log.md` - Time and bug tracking
- `test_suite.{py,go,zig}` - Tests for all problems

### Success Criteria
- All 20+ problems solved correctly
- Can identify patterns quickly
- Average solving time improves over practice
- Understanding of DP deepens

---

## Module Resources

### CLRS References
- Chapter 15: Dynamic Programming

### Additional Reading
- "Introduction to Algorithms" DP chapter
- Competitive programming DP tutorials
- LeetCode DP problem set
- Dynamic Programming for Coding Interviews

### Key Takeaways

By the end of this module, you should:
1. Recognize DP problems instantly
2. Define state and recurrence systematically
3. Implement both memoization and tabulation fluently
4. Optimize space complexity when possible
5. Reconstruct solutions from DP tables
6. Apply advanced DP patterns confidently
7. Solve DP problems quickly and accurately

### Next Module

**Module 9: Greedy Algorithms** - You'll learn when greedy works, how to prove correctness, and identify greedy vs. DP problems.

---

## Tips for Success

1. **Draw the recurrence tree** - Visualize overlapping subproblems
2. **Start with recursion** - Then add memoization, then convert to tabulation
3. **Define state carefully** - This is the hardest part
4. **Verify with small examples** - Test by hand first
5. **Practice, practice, practice** - DP requires lots of exposure

## Common Pitfalls

- **Wrong state definition** - Most common mistake
- **Incorrect base cases** - Check edge conditions carefully
- **Off-by-one errors** - Especially in array indexing
- **Missing dependencies** - Ensure all needed subproblems computed first
- **Space optimization too early** - Get it working first, optimize later
- **Not practicing enough** - DP needs repetition to master
