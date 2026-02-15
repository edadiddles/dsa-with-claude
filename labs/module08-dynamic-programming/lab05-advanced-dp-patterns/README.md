# Lab 8.5: Advanced DP Patterns

**Module**: Module 8 - Dynamic Programming  
**Duration**: 6-8 hours  
**Difficulty**: Hard

---

Lab 8.5: Advanced DP Patterns

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

---

## Back to Module

[← Back to Module 8: Dynamic Programming](../module_08/README.md)

## Navigation

- [Previous Lab](./lab_8_4.md) (if exists)
- [Next Lab](./lab_8_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 8, Lab 5 of 7*
