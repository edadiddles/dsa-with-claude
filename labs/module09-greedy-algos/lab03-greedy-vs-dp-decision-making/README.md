# Lab 9.3: Greedy vs. DP Decision Making

**Module**: Module 9 - Greedy Algorithms  
**Duration**: 3-4 hours  
**Difficulty**: Medium

---

Lab 9.3: Greedy vs. DP Decision Making

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Distinguish greedy-solvable from DP-required problems
- Build decision framework
- Practice on ambiguous problems

### Requirements

1. **Analyze problem characteristics**:
   - Greedy choice property presence
   - Optimal substructure verification
   - Local vs. global optimization

2. **Implement both approaches** for comparison:
   - Coin change (greedy fails, DP works)
   - Weighted job scheduling (greedy fails, DP works)
   - Activity selection (greedy works)
   - Minimum spanning tree (greedy works)

3. **Decision framework**:
   - Checklist for identifying greedy problems
   - Red flags indicating DP needed
   - Testing approach validity

### Implementation Details

```python
# DECISION FRAMEWORK

def is_greedy_applicable(problem_description):
    """
    Checklist for greedy algorithms:
    
    ✓ Greedy Choice Property:
      - Can we make a locally optimal choice?
      - Does this choice lead to global optimum?
      - Can we prove it with exchange argument?
    
    ✓ Optimal Substructure:
      - After greedy choice, is remaining problem similar?
      - Can we solve recursively?
    
    ✓ No Reconsideration:
      - Once choice is made, never revisit?
      - No need to explore alternatives?
    
    ✗ Red flags (use DP instead):
      - Multiple interdependent choices
      - Need to consider all possibilities
      - Greedy choice depends on future choices
      - Can construct counterexample
    """
    pass

# EXAMPLE 1: Coin Change

def coin_change_comparison(coins, amount):
    """
    Compare greedy vs. DP for coin change.
    
    Greedy approach: Always take largest coin.
    DP approach: Try all possibilities.
    """
    # Greedy
    def greedy(coins, amount):
        coins = sorted(coins, reverse=True)
        count = 0
        for coin in coins:
            while amount >= coin:
                amount -= coin
                count += 1
        return count if amount == 0 else -1
    
    # DP
    def dp(coins, amount):
        dp_table = [float('inf')] * (amount + 1)
        dp_table[0] = 0
        
        for i in range(1, amount + 1):
            for coin in coins:
                if coin <= i:
                    dp_table[i] = min(dp_table[i], 1 + dp_table[i - coin])
        
        return dp_table[amount] if dp_table[amount] != float('inf') else -1
    
    greedy_result = greedy(coins, amount)
    dp_result = dp(coins, amount)
    
    return {
        'greedy': greedy_result,
        'dp': dp_result,
        'greedy_optimal': greedy_result == dp_result
    }

def when_greedy_coin_change_works():
    """
    Greedy works for canonical coin systems:
    - US coins: [1, 5, 10, 25]
    - Euro coins: [1, 2, 5, 10, 20, 50]
    
    Canonical system: greedy always gives optimal solution.
    
    Non-canonical example: [1, 3, 4]
    - For amount=6: greedy gives 3 coins (4+1+1)
    - DP gives 2 coins (3+3)
    """
    # US coins - greedy works
    result1 = coin_change_comparison([1, 5, 10, 25], 41)
    print(f"US coins, amount=41: {result1}")
    
    # Non-canonical - greedy fails
    result2 = coin_change_comparison([1, 3, 4], 6)
    print(f"[1,3,4], amount=6: {result2}")

# EXAMPLE 2: Weighted Interval Scheduling

def weighted_interval_comparison(intervals):
    """
    intervals: [(start, end, weight), ...]
    
    Greedy approaches (all fail):
    1. Earliest finish time
    2. Shortest interval
    3. Highest weight
    
    DP approach (works):
    Sort by finish time, DP on whether to include each interval.
    """
    # Greedy by weight (fails)
    def greedy_by_weight(intervals):
        sorted_intervals = sorted(intervals, key=lambda x: x[2], reverse=True)
        
        selected = []
        for interval in sorted_intervals:
            start, end, weight = interval
            
            # Check if compatible with selected
            compatible = True
            for s, e, w in selected:
                if not (end <= s or start >= e):
                    compatible = False
                    break
            
            if compatible:
                selected.append(interval)
        
        return sum(w for s, e, w in selected)
    
    # DP (correct)
    def dp_weighted_intervals(intervals):
        # Sort by finish time
        intervals = sorted(intervals, key=lambda x: x[1])
        n = len(intervals)
        
        # Find latest non-overlapping interval for each
        prev = [-1] * n
        for i in range(n):
            for j in range(i - 1, -1, -1):
                if intervals[j][1] <= intervals[i][0]:
                    prev[i] = j
                    break
        
        # DP
        dp = [0] * (n + 1)
        for i in range(1, n + 1):
            # Don't include interval i-1
            without = dp[i-1]
            
            # Include interval i-1
            weight = intervals[i-1][2]
            with_interval = weight + (dp[prev[i-1] + 1] if prev[i-1] != -1 else 0)
            
            dp[i] = max(without, with_interval)
        
        return dp[n]
    
    greedy_result = greedy_by_weight(intervals)
    dp_result = dp_weighted_intervals(intervals)
    
    return {
        'greedy': greedy_result,
        'dp': dp_result,
        'greedy_optimal': greedy_result == dp_result
    }

# TESTING FRAMEWORK

class GreedyValidator:
    """Framework for testing greedy algorithms"""
    
    def __init__(self, greedy_func, brute_force_func):
        self.greedy = greedy_func
        self.brute_force = brute_force_func
    
    def validate(self, test_cases):
        """Test greedy against brute force"""
        results = []
        
        for test_input in test_cases:
            greedy_output = self.greedy(test_input)
            optimal_output = self.brute_force(test_input)
            
            is_optimal = (greedy_output == optimal_output)
            
            results.append({
                'input': test_input,
                'greedy': greedy_output,
                'optimal': optimal_output,
                'correct': is_optimal
            })
        
        return results
    
    def find_counterexample(self, input_generator, max_tries=1000):
        """Try to find counterexample where greedy fails"""
        for _ in range(max_tries):
            test_input = input_generator()
            
            greedy_output = self.greedy(test_input)
            optimal_output = self.brute_force(test_input)
            
            if greedy_output != optimal_output:
                return {
                    'found': True,
                    'input': test_input,
                    'greedy': greedy_output,
                    'optimal': optimal_output
                }
        
        return {'found': False}

# PROBLEM CLASSIFICATION

def classify_problem(problem_characteristics):
    """
    Classify whether problem is greedy or DP based on characteristics.
    
    Greedy indicators:
    - Sorting helps
    - Local optimal leads to global optimal
    - Proof by exchange argument exists
    - Never need to reconsider choices
    
    DP indicators:
    - Need to explore all possibilities
    - Optimal depends on solving subproblems optimally
    - Can't determine best choice without looking ahead
    - Overlapping subproblems exist
    """
    score_greedy = 0
    score_dp = 0
    
    if problem_characteristics.get('sorting_helps'):
        score_greedy += 2
    
    if problem_characteristics.get('local_optimal_works'):
        score_greedy += 3
    
    if problem_characteristics.get('exchange_argument_exists'):
        score_greedy += 3
    
    if problem_characteristics.get('overlapping_subproblems'):
        score_dp += 3
    
    if problem_characteristics.get('need_all_possibilities'):
        score_dp += 3
    
    if problem_characteristics.get('future_dependent'):
        score_dp += 2
    
    if score_greedy > score_dp:
        return 'greedy'
    elif score_dp > score_greedy:
        return 'dp'
    else:
        return 'unclear'

# EXAMPLES OF CLASSIFICATION

def problem_examples():
    """
    Clear greedy problems:
    1. Activity selection: Sort by finish, take earliest
    2. Huffman coding: Always merge two smallest
    3. MST (Kruskal): Always add cheapest edge
    4. Fractional knapsack: Take by value/weight ratio
    
    Clear DP problems:
    1. 0/1 knapsack: Can't decide without seeing future
    2. Longest common subsequence: Need all possibilities
    3. Edit distance: Multiple ways to transform
    4. Matrix chain multiplication: Order matters
    
    Ambiguous (need analysis):
    1. Coin change: Depends on coin system
    2. Interval scheduling: Depends on whether weighted
    3. Subset sum: Generally DP, but special cases greedy
    """
    pass
```

### Test Cases
- Problems with known greedy solutions
- Problems with known DP solutions
- Ambiguous problems requiring analysis
- Generate counterexamples systematically

### Deliverables
- `greedy_vs_dp_framework.{py,go,zig}` - Decision framework
- `coin_change_analysis.{py,go,zig}` - Greedy vs. DP comparison
- `interval_scheduling_analysis.{py,go,zig}` - Weighted vs. unweighted
- `greedy_validator.{py,go,zig}` - Testing framework
- `problem_classifier.{py,go,zig}` - Automatic classification
- `decision_guide.md` - How to choose approach
- `examples_catalog.md` - Categorized problem examples

### Success Criteria
- Can distinguish greedy from DP problems
- Understanding of why certain problems need DP
- Can construct counterexamples
- Know how to validate greedy algorithms

---

---

## Back to Module

[← Back to Module 9: Greedy Algorithms](../module_09/README.md)

## Navigation

- [Previous Lab](./lab_9_2.md) (if exists)
- [Next Lab](./lab_9_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 9, Lab 3 of 5*
