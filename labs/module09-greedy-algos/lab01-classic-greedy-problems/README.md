# Lab 9.1: Classic Greedy Problems

**Module**: Module 9 - Greedy Algorithms  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 9.1: Classic Greedy Problems

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement fundamental greedy algorithms
- Understand greedy choice property through examples
- Build intuition for when greedy works

### Requirements

1. **Implement classic greedy algorithms**:
   - **Activity Selection** (interval scheduling)
   - **Fractional Knapsack**
   - **Coin Change** (greedy version)
   - **Job Sequencing with Deadlines**
   - **Minimum Platforms** (train scheduling)
   - **Gas Station** (circular tour)

2. **For each problem**:
   - Prove greedy choice property
   - Show optimal substructure
   - Analyze time complexity
   - Compare with non-greedy alternatives

3. **Demonstrate greedy fails**:
   - Show when greedy doesn't work
   - Provide counterexamples
   - Compare with DP solutions

### Implementation Details

```python
# ACTIVITY SELECTION - Classic Greedy

def activity_selection(activities):
    """
    Select maximum number of non-overlapping activities.
    Activities: list of (start, finish) tuples.
    Greedy: always pick activity that finishes earliest.
    """
    if not activities:
        return []
    
    # Sort by finish time
    activities = sorted(activities, key=lambda x: x[1])
    
    selected = [activities[0]]
    last_finish = activities[0][1]
    
    for start, finish in activities[1:]:
        if start >= last_finish:
            selected.append((start, finish))
            last_finish = finish
    
    return selected

def activity_selection_proof():
    """
    Proof of correctness using exchange argument:
    
    Let A = optimal solution
    Let G = greedy solution
    
    Claim: |G| = |A|
    
    Proof:
    1. Let a1 be the first activity in G (finishes earliest)
    2. Let b1 be the first activity in A
    3. If a1 = b1, done
    4. If a1 != b1, then finish(a1) <= finish(b1)
    5. Replace b1 with a1 in A → still valid solution
    6. Repeat for remaining activities
    7. Therefore G is optimal
    """
    pass

# FRACTIONAL KNAPSACK

def fractional_knapsack(weights, values, capacity):
    """
    Fractional knapsack: can take fractions of items.
    Greedy: sort by value-to-weight ratio, take greedily.
    """
    n = len(weights)
    
    # Compute value-to-weight ratios
    items = [(values[i] / weights[i], weights[i], values[i]) 
             for i in range(n)]
    
    # Sort by ratio (descending)
    items.sort(reverse=True)
    
    total_value = 0
    remaining_capacity = capacity
    taken = []
    
    for ratio, weight, value in items:
        if remaining_capacity == 0:
            break
        
        if weight <= remaining_capacity:
            # Take whole item
            taken.append((weight, value, 1.0))
            total_value += value
            remaining_capacity -= weight
        else:
            # Take fraction
            fraction = remaining_capacity / weight
            taken.append((weight, value, fraction))
            total_value += value * fraction
            remaining_capacity = 0
    
    return total_value, taken

def fractional_vs_01_knapsack():
    """
    Demonstration that greedy works for fractional but not 0/1:
    
    Example:
    weights = [10, 20, 30]
    values = [60, 100, 120]
    capacity = 50
    
    Greedy (by ratio):
    - Item 0: ratio = 6.0, take all → 60
    - Item 1: ratio = 5.0, take all → 100
    - Total: 160
    
    Optimal for 0/1:
    - Items 1, 2: 100 + 120 = 220 (better!)
    
    Greedy fails for 0/1 knapsack!
    """
    pass

# COIN CHANGE - Greedy Version

def coin_change_greedy(coins, amount):
    """
    Coin change using greedy approach.
    WARNING: Only works for certain coin systems!
    
    Works for: US coins (1, 5, 10, 25)
    Fails for: (1, 3, 4) with amount=6
    """
    coins = sorted(coins, reverse=True)
    
    count = 0
    used = []
    
    for coin in coins:
        while amount >= coin:
            amount -= coin
            count += 1
            used.append(coin)
    
    if amount == 0:
        return count, used
    else:
        return -1, []  # Impossible

def coin_change_greedy_counterexample():
    """
    Greedy fails for coins = [1, 3, 4], amount = 6:
    
    Greedy: 4 + 1 + 1 = 3 coins
    Optimal: 3 + 3 = 2 coins
    
    Must use DP for general coin change!
    """
    coins = [1, 3, 4]
    amount = 6
    
    greedy_count, _ = coin_change_greedy(coins, amount)
    # greedy_count = 3
    
    # DP gives optimal = 2
    print("Greedy fails! Need DP for coin change.")

# JOB SEQUENCING WITH DEADLINES

def job_sequencing(jobs):
    """
    Jobs with deadlines and profits.
    Each job takes 1 unit time.
    Maximize profit.
    
    jobs: list of (id, deadline, profit)
    Greedy: sort by profit (descending), schedule as late as possible.
    """
    # Sort by profit (descending)
    jobs = sorted(jobs, key=lambda x: x[2], reverse=True)
    
    max_deadline = max(job[1] for job in jobs)
    
    # Track which time slots are filled
    slots = [-1] * max_deadline
    total_profit = 0
    scheduled = []
    
    for job_id, deadline, profit in jobs:
        # Try to schedule in latest possible slot before deadline
        for t in range(min(deadline, max_deadline) - 1, -1, -1):
            if slots[t] == -1:
                slots[t] = job_id
                total_profit += profit
                scheduled.append((job_id, t + 1, profit))
                break
    
    return total_profit, scheduled

# MINIMUM PLATFORMS

def min_platforms(arrivals, departures):
    """
    Minimum platforms needed for trains.
    arrivals[i] = arrival time of train i
    departures[i] = departure time of train i
    
    Greedy: track overlapping intervals.
    """
    events = []
    
    for arr in arrivals:
        events.append((arr, 'arrival'))
    for dep in departures:
        events.append((dep, 'departure'))
    
    # Sort events by time
    events.sort()
    
    platforms_needed = 0
    max_platforms = 0
    
    for time, event_type in events:
        if event_type == 'arrival':
            platforms_needed += 1
            max_platforms = max(max_platforms, platforms_needed)
        else:
            platforms_needed -= 1
    
    return max_platforms

# GAS STATION (Circular Tour)

def can_complete_circuit(gas, cost):
    """
    Circular route of gas stations.
    gas[i] = gas at station i
    cost[i] = gas needed to go from i to i+1
    
    Return starting station index, or -1 if impossible.
    
    Greedy insight: if total gas >= total cost, solution exists.
    Start from first station where we can begin accumulating surplus.
    """
    total_gas = sum(gas)
    total_cost = sum(cost)
    
    if total_gas < total_cost:
        return -1
    
    tank = 0
    start = 0
    
    for i in range(len(gas)):
        tank += gas[i] - cost[i]
        
        if tank < 0:
            # Can't reach next station, try starting from next
            start = i + 1
            tank = 0
    
    return start

def gas_station_proof():
    """
    Proof that greedy works:
    
    1. If total_gas < total_cost, impossible
    2. If total_gas >= total_cost, solution exists
    3. If we fail starting from station i, all stations 
       from i to current are invalid starts
    4. Start from first station after failure point
    5. If we complete loop, this is the answer
    
    Key insight: We only need to find ONE valid start.
    """
    pass

# GREEDY INTERVAL SCHEDULING VARIANTS

def weighted_interval_scheduling(intervals):
    """
    Activity selection with weights.
    This is NOT solvable by greedy - need DP!
    
    Counterexample:
    intervals = [(0, 10, 100), (1, 2, 10), (3, 4, 10)]
    
    Greedy (by finish time): (1, 2, 10), (3, 4, 10) → 20
    Optimal: (0, 10, 100) → 100
    
    Must use DP!
    """
    # This is a counterexample showing greedy fails
    pass

def interval_partitioning(intervals):
    """
    Partition intervals into minimum number of groups
    such that no intervals in same group overlap.
    
    Greedy: sort by start time, use earliest finishing group.
    """
    if not intervals:
        return []
    
    intervals = sorted(intervals)
    
    groups = []
    
    for interval in intervals:
        start, end = interval
        
        # Find group where this interval fits
        placed = False
        for group in groups:
            # Check if compatible with this group
            if start >= group[-1][1]:  # No overlap
                group.append(interval)
                placed = True
                break
        
        if not placed:
            # Need new group
            groups.append([interval])
    
    return groups

# COMPARISON: Greedy vs DP

def compare_greedy_dp():
    """
    Problems where greedy works:
    1. Activity selection
    2. Fractional knapsack
    3. Huffman coding
    4. MST (Kruskal, Prim)
    5. Dijkstra's shortest path
    
    Problems where greedy fails (need DP):
    1. 0/1 knapsack
    2. Weighted interval scheduling
    3. Longest common subsequence
    4. Edit distance
    5. General coin change
    
    Key difference: Greedy has NO reconsideration.
    DP explores all possibilities.
    """
    pass
```

### Test Cases
- Activity selection with various overlaps
- Fractional knapsack vs. 0/1 comparison
- Coin change with canonical and non-canonical systems
- Job sequencing with tight deadlines
- Platform problems with complex schedules
- Gas station with various configurations

### Deliverables
- `activity_selection.{py,go,zig}` - Complete implementation with proof
- `fractional_knapsack.{py,go,zig}` - With comparison to 0/1
- `coin_change_greedy.{py,go,zig}` - With counterexamples
- `job_sequencing.{py,go,zig}` - Job scheduling implementation
- `platform_scheduling.{py,go,zig}` - Train platform problem
- `gas_station.{py,go,zig}` - Circular tour problem
- `greedy_proofs.md` - Correctness proofs for each algorithm
- `counterexamples.md` - When greedy fails
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All greedy algorithms work correctly
- Can prove correctness using exchange arguments
- Understanding of greedy choice property
- Can identify when greedy fails
- Know difference between greedy and DP problems

---

---

## Back to Module

[← Back to Module 9: Greedy Algorithms](../module_09/README.md)

## Navigation

- [Previous Lab](./lab_9_0.md) (if exists)
- [Next Lab](./lab_9_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 9, Lab 1 of 5*
