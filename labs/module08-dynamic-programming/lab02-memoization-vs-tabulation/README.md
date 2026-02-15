# Lab 8.2: Memoization vs. Tabulation

**Module**: Module 8 - Dynamic Programming  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 8.2: Memoization vs. Tabulation

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

---

## Back to Module

[← Back to Module 8: Dynamic Programming](../module_08/README.md)

## Navigation

- [Previous Lab](./lab_8_1.md) (if exists)
- [Next Lab](./lab_8_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 8, Lab 2 of 7*
