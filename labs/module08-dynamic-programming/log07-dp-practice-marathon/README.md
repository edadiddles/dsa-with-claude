# Lab 8.7: DP Practice Marathon

**Module**: Module 8 - Dynamic Programming  
**Duration**: 8-10 hours  
**Difficulty**: Medium-Hard

---

Lab 8.7: DP Practice Marathon

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

---

## Back to Module

[← Back to Module 8: Dynamic Programming](../module_08/README.md)

## Navigation

- [Previous Lab](./lab_8_6.md) (if exists)
- [Next Lab](./lab_8_8.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 8, Lab 7 of 7*
