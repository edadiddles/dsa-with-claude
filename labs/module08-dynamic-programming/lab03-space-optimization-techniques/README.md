# Lab 8.3: Space Optimization Techniques

**Module**: Module 8 - Dynamic Programming  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 8.3: Space Optimization Techniques

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

---

## Back to Module

[← Back to Module 8: Dynamic Programming](../module_08/README.md)

## Navigation

- [Previous Lab](./lab_8_2.md) (if exists)
- [Next Lab](./lab_8_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 8, Lab 3 of 7*
