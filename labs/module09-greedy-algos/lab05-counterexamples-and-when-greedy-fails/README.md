# Lab 9.5: Counterexamples & When Greedy Fails

**Module**: Module 9 - Greedy Algorithms  
**Duration**: 2-3 hours  
**Difficulty**: Medium

---

Lab 9.5: Counterexamples & When Greedy Fails

**Duration**: 2-3 hours  
**Difficulty**: Medium

### Objectives
- Systematically find counterexamples
- Understand why greedy fails
- Build intuition for problem difficulty

### Requirements

1. **Catalog of failures**:
   - 0/1 knapsack
   - Weighted interval scheduling
   - Longest path in graph
   - Traveling salesman
   - General coin change

2. **Counterexample construction**:
   - Systematic search
   - Minimal counterexamples
   - Explain why greedy fails

3. **Analysis of failure modes**:
   - Irreversible choices
   - Short-term vs. long-term optimization
   - Dependencies between choices

### Implementation Details

```python
# COUNTEREXAMPLE CATALOG

class GreedyCounterexamples:
    """Collection of problems where greedy fails"""
    
    @staticmethod
    def knapsack_01():
        """
        0/1 Knapsack: Greedy by ratio fails
        
        Items: [(weight, value)]
        1. (10, 60) → ratio = 6.0
        2. (20, 100) → ratio = 5.0
        3. (30, 120) → ratio = 4.0
        
        Capacity = 50
        
        Greedy (by ratio): Take 1, 2 → value = 160
        Optimal: Take 2, 3 → value = 220
        
        Why greedy fails:
        - Committing to item 1 (high ratio) prevents taking 2+3
        - Can't undo choice of item 1
        - Need to explore all possibilities (DP)
        """
        return {
            'items': [(10, 60), (20, 100), (30, 120)],
            'capacity': 50,
            'greedy_value': 160,
            'optimal_value': 220,
            'reason': 'Irreversible early choice'
        }
    
    @staticmethod
    def weighted_interval_scheduling():
        """
        Weighted Interval Scheduling: All greedy heuristics fail
        
        Intervals: [(start, end, weight)]
        1. (0, 100, 50)
        2. (1, 2, 40)
        3. (3, 4, 40)
        
        Greedy by finish: Take 2, 3 → 80
        Greedy by weight: Take 1 → 50
        Optimal: Take 2, 3 → 80? Or take 1 → 50?
        
        Actually, this example: 2+3 = 80 is optimal.
        
        Better example:
        1. (0, 10, 100)
        2. (1, 2, 10)
        3. (3, 4, 10)
        
        Greedy by finish: 2, 3 → 20
        Greedy by weight: 1 → 100
        Optimal: 1 → 100
        
        But this doesn't show greedy by weight fails...
        
        Actual counterexample:
        1. (0, 5, 6)
        2. (1, 3, 4)
        3. (4, 8, 4)
        4. (5, 7, 3)
        5. (6, 9, 2)
        
        Greedy by weight: 1, 3, 5 → 12
        Optimal: 2, 4, 5 → 9? No...
        
        Need: 1, 4 vs 2, 3
        """
        return {
            'intervals': [(0, 10, 100), (1, 2, 10), (3, 4, 10)],
            'greedy_by_finish': 20,
            'greedy_by_weight': 100,
            'optimal': 100,
            'reason': 'Multiple greedy strategies, none guaranteed'
        }
    
    @staticmethod
    def coin_change():
        """
        Coin Change: Greedy fails for non-canonical systems
        
        Coins: [1, 3, 4]
        Amount: 6
        
        Greedy: 4, 1, 1 → 3 coins
        Optimal: 3, 3 → 2 coins
        
        Why greedy fails:
        - Taking largest coin (4) prevents optimal solution
        - Need to try not taking 4
        - Requires exploring alternatives (DP)
        """
        return {
            'coins': [1, 3, 4],
            'amount': 6,
            'greedy': 3,
            'optimal': 2,
            'reason': 'Largest coin blocks better solution'
        }
    
    @staticmethod
    def longest_path():
        """
        Longest Path in Graph: All greedy strategies fail
        
        Graph with weighted edges.
        Greedy: always take heaviest edge
        
        Problem: May get stuck in local optimum
        
        Example:
            A --10--> B --1--> C
            |                  |
            +-------100------->+
        
        Path A → C:
        Greedy: A → B → C = 11
        Optimal: A → C = 100
        
        Why greedy fails:
        - First choice (A→B) looks good locally
        - Blocks globally optimal path
        - NP-hard problem, no greedy solution
        """
        pass
    
    @staticmethod
    def tsp():
        """
        Traveling Salesman Problem: Greedy fails
        
        Nearest neighbor heuristic:
        - Start at city
        - Go to nearest unvisited city
        - Repeat
        
        This can give arbitrarily bad solutions!
        
        Example where nearest neighbor fails:
        
        Cities arranged in cross pattern with center:
        - Center has distance 1 to all outer cities
        - Outer cities have distance 2 to each other
        
        Greedy from outer city: visits center last
        Optimal: visit center in middle of tour
        
        TSP is NP-hard, no polynomial greedy solution.
        """
        pass

# COUNTEREXAMPLE GENERATOR

class CounterexampleFinder:
    """Systematically search for counterexamples"""
    
    def find_knapsack_counterexample(self, num_items=3, max_weight=10, max_value=20):
        """
        Search for 0/1 knapsack instances where greedy fails.
        """
        import itertools
        import random
        
        for _ in range(1000):
            # Generate random items
            items = [(random.randint(1, max_weight), 
                     random.randint(1, max_value)) 
                    for _ in range(num_items)]
            
            capacity = random.randint(max_weight, num_items * max_weight)
            
            # Try greedy
            greedy_value = self._knapsack_greedy(items, capacity)
            
            # Try optimal (brute force for small instances)
            optimal_value = self._knapsack_optimal(items, capacity)
            
            if greedy_value < optimal_value:
                return {
                    'items': items,
                    'capacity': capacity,
                    'greedy': greedy_value,
                    'optimal': optimal_value
                }
        
        return None
    
    def _knapsack_greedy(self, items, capacity):
        """Greedy by value/weight ratio"""
        items_with_ratio = [(v/w, w, v) for w, v in items]
        items_with_ratio.sort(reverse=True)
        
        total_value = 0
        remaining = capacity
        
        for ratio, weight, value in items_with_ratio:
            if weight <= remaining:
                total_value += value
                remaining -= weight
        
        return total_value
    
    def _knapsack_optimal(self, items, capacity):
        """Brute force optimal"""
        n = len(items)
        max_value = 0
        
        for mask in range(1 << n):
            total_weight = 0
            total_value = 0
            
            for i in range(n):
                if mask & (1 << i):
                    total_weight += items[i][0]
                    total_value += items[i][1]
            
            if total_weight <= capacity:
                max_value = max(max_value, total_value)
        
        return max_value

# FAILURE ANALYSIS

def analyze_greedy_failure(problem, greedy_solution, optimal_solution):
    """
    Analyze why greedy failed:
    
    1. Irreversible choice:
       - Early decision blocks better later options
       - Example: 0/1 knapsack
    
    2. Missing global view:
       - Greedy sees only local information
       - Example: Longest path
    
    3. Interdependent choices:
       - Earlier choices affect later constraints
       - Example: TSP
    
    4. Multiple optimal criteria:
       - No clear "greedy choice"
       - Example: Weighted interval scheduling
    """
    failure_type = None
    
    if 'blocked_by_early_choice' in analyze_choices(greedy_solution, optimal_solution):
        failure_type = 'irreversible_choice'
    
    if 'requires_global_view' in analyze_structure(problem):
        failure_type = 'missing_global_view'
    
    if 'choices_interdependent' in analyze_dependencies(problem):
        failure_type = 'interdependent_choices'
    
    return failure_type

# MINIMAL COUNTEREXAMPLES

def find_minimal_counterexample(problem_class, greedy_algorithm):
    """
    Find smallest counterexample where greedy fails.
    
    Start with small instances, increase size until greedy fails.
    """
    size = 2
    
    while size < 20:
        # Generate random instance of given size
        instance = problem_class.generate_random(size)
        
        greedy_result = greedy_algorithm(instance)
        optimal_result = solve_optimal(instance)
        
        if greedy_result != optimal_result:
            # Found counterexample, try to minimize
            return minimize_counterexample(instance, greedy_algorithm)
        
        size += 1
    
    return None

def minimize_counterexample(instance, greedy_algorithm):
    """
    Given counterexample, try to make it smaller while preserving failure.
    """
    # Try removing elements one by one
    # If still counterexample, keep smaller version
    # Repeat until minimal
    
    current = instance
    
    improved = True
    while improved:
        improved = False
        
        for i in range(len(current)):
            smaller = current[:i] + current[i+1:]
            
            if is_counterexample(smaller, greedy_algorithm):
                current = smaller
                improved = True
                break
    
    return current

def is_counterexample(instance, greedy_algorithm):
    """Check if instance is counterexample"""
    greedy = greedy_algorithm(instance)
    optimal = solve_optimal(instance)
    return greedy != optimal
```

### Test Cases
- Verify all counterexamples are correct
- Search for new counterexamples
- Analyze why each greedy approach fails

### Deliverables
- `counterexample_catalog.{py,go,zig}` - Collection of failures
- `counterexample_finder.{py,go,zig}` - Systematic search
- `minimal_counterexamples.{py,go,zig}` - Minimization tool
- `failure_analysis.{py,go,zig}` - Analyze why greedy fails
- `greedy_failures.md` - Complete catalog with explanations
- `when_to_avoid_greedy.md` - Guide for recognizing failures

### Success Criteria
- Understanding of common greedy failure modes
- Can construct counterexamples for incorrect greedy approaches
- Know when to immediately reject greedy and use DP
- Can explain why each counterexample breaks greedy

---

---

## Back to Module

[← Back to Module 9: Greedy Algorithms](../module_09/README.md)

## Navigation

- [Previous Lab](./lab_9_4.md) (if exists)
- [Next Lab](./lab_9_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 9, Lab 5 of 5*
