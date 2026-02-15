# Lab 8.1: DP Foundations & Classic Problems

**Module**: Module 8 - Dynamic Programming  
**Duration**: 6-8 hours  
**Difficulty**: Medium

---

Lab 8.1: DP Foundations & Classic Problems

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

---

## Back to Module

[← Back to Module 8: Dynamic Programming](../module_08/README.md)

## Navigation

- [Previous Lab](./lab_8_0.md) (if exists)
- [Next Lab](./lab_8_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 8, Lab 1 of 7*
