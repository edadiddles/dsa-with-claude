# Lab 13.3: Approximation Algorithms

**Module**: Module 13 - NP-Completeness & Approximation Algorithms  
**Duration**: 6-8 hours  
**Difficulty**: Hard

---

Lab 13.3: Approximation Algorithms

**Duration**: 6-8 hours  
**Difficulty**: Hard

### Objectives
- Design approximation algorithms
- Prove approximation ratios
- Understand performance guarantees
- Apply to practical problems

### Requirements

1. **Implement approximation algorithms**:
   - **Vertex Cover** (2-approximation)
   - **Set Cover** (ln n approximation)
   - **TSP** (metric TSP 2-approximation)
   - **Bin Packing** (First Fit, Best Fit)
   - **Load Balancing** (greedy 2-approximation)
   - **Max Cut** (randomized 2-approximation)

2. **Prove approximation ratios**:
   - Upper bound on algorithm performance
   - Lower bound examples
   - Tightness of analysis

3. **Compare with optimal**:
   - Measure actual approximation ratio
   - Understand when approximation is good enough

### Implementation Details

```python
import random

# VERTEX COVER APPROXIMATION

def vertex_cover_approx(graph):
    """
    2-approximation for Vertex Cover.
    
    Algorithm:
    1. While edges remain:
       - Pick arbitrary edge (u, v)
       - Add both u and v to cover
       - Remove all edges incident to u or v
    
    Approximation ratio: 2
    Proof: Every edge forces at least one endpoint in optimal cover.
           We add both endpoints, so at most 2× optimal.
    
    Time: O(E)
    """
    cover = set()
    remaining_edges = list(graph)
    
    while remaining_edges:
        # Pick arbitrary edge
        u, v = remaining_edges[0]
        
        # Add both endpoints
        cover.add(u)
        cover.add(v)
        
        # Remove all edges incident to u or v
        remaining_edges = [
            (a, b) for a, b in remaining_edges
            if a not in {u, v} and b not in {u, v}
        ]
    
    return list(cover)

def vertex_cover_approx_analysis():
    """
    Approximation analysis:
    
    Let OPT = size of optimal vertex cover
    Let ALG = size of our cover
    
    Claim: ALG ≤ 2 * OPT
    
    Proof:
    - Let E' = edges selected by algorithm
    - |E'| edges were selected
    - Both endpoints added to cover, so ALG = 2|E'|
    - Each edge in E' must have at least one endpoint in OPT
    - No two edges in E' share endpoints
    - Therefore |E'| ≤ OPT
    - Therefore ALG = 2|E'| ≤ 2 * OPT
    
    This is tight: Consider matching graph.
    """
    pass

# SET COVER APPROXIMATION

def set_cover_greedy(universe, sets):
    """
    Greedy approximation for Set Cover.
    
    Algorithm:
    1. While uncovered elements remain:
       - Pick set covering most uncovered elements
       - Add to cover
    
    Approximation ratio: ln n (where n = |universe|)
    
    Time: O(|sets| * |universe|)
    """
    cover = []
    uncovered = set(universe)
    
    while uncovered:
        # Find set covering most uncovered elements
        best_set = None
        best_coverage = 0
        
        for s in sets:
            coverage = len(s & uncovered)
            if coverage > best_coverage:
                best_coverage = coverage
                best_set = s
        
        if best_set is None:
            break
        
        cover.append(best_set)
        uncovered -= best_set
    
    return cover

# TSP APPROXIMATION (METRIC TSP)

def tsp_approx_mst(graph, vertices):
    """
    2-approximation for Metric TSP using MST.
    
    Algorithm:
    1. Find MST of graph
    2. Do DFS on MST to get tour
    3. Shortcut tour (using triangle inequality)
    
    Approximation ratio: 2 (for metric TSP)
    
    Time: O(V^2) or O(E log V) with better MST
    """
    # Find MST (using Prim's or Kruskal's)
    mst = find_mst(graph, vertices)
    
    # DFS on MST
    visited = set()
    tour = []
    
    def dfs(v):
        visited.add(v)
        tour.append(v)
        
        for u in mst.get(v, []):
            if u not in visited:
                dfs(u)
    
    start = vertices[0]
    dfs(start)
    tour.append(start)  # Complete cycle
    
    return tour

def find_mst(graph, vertices):
    """Find MST using Prim's algorithm"""
    # Simplified - use implementation from Module 6
    import heapq
    
    mst = {v: [] for v in vertices}
    visited = set()
    edges = []
    
    start = vertices[0]
    visited.add(start)
    
    # Add edges from start
    for v in vertices:
        if v != start and (start, v) in graph:
            heapq.heappush(edges, (graph[(start, v)], start, v))
    
    while edges and len(visited) < len(vertices):
        weight, u, v = heapq.heappop(edges)
        
        if v in visited:
            continue
        
        visited.add(v)
        mst[u].append(v)
        mst[v].append(u)
        
        # Add edges from v
        for w in vertices:
            if w not in visited and (v, w) in graph:
                heapq.heappush(edges, (graph[(v, w)], v, w))
    
    return mst

def tsp_christofides(graph, vertices):
    """
    Christofides algorithm: 1.5-approximation for Metric TSP.
    
    Algorithm:
    1. Find MST
    2. Find minimum matching on odd-degree vertices
    3. Combine to get Eulerian graph
    4. Find Eulerian tour
    5. Convert to Hamiltonian tour (shortcut)
    
    Approximation ratio: 1.5
    
    This is more complex - simplified here.
    """
    # Complex algorithm - placeholder
    pass

# BIN PACKING APPROXIMATIONS

def bin_packing_first_fit(items, bin_capacity):
    """
    First Fit approximation for Bin Packing.
    
    Algorithm:
    For each item, place in first bin that fits.
    If no bin fits, open new bin.
    
    Approximation ratio: 1.7 * OPT
    
    Time: O(n * m) where m = number of bins
    """
    bins = []
    
    for item in items:
        # Find first bin that fits
        placed = False
        for bin in bins:
            if sum(bin) + item <= bin_capacity:
                bin.append(item)
                placed = True
                break
        
        if not placed:
            bins.append([item])
    
    return bins

def bin_packing_best_fit(items, bin_capacity):
    """
    Best Fit approximation for Bin Packing.
    
    Algorithm:
    For each item, place in bin with least remaining space that fits.
    
    Approximation ratio: 1.7 * OPT
    
    Time: O(n * m)
    """
    bins = []
    
    for item in items:
        # Find best bin (minimum remaining space after placing item)
        best_bin = None
        min_remaining = bin_capacity + 1
        
        for bin in bins:
            remaining = bin_capacity - sum(bin) - item
            if remaining >= 0 and remaining < min_remaining:
                min_remaining = remaining
                best_bin = bin
        
        if best_bin is not None:
            best_bin.append(item)
        else:
            bins.append([item])
    
    return bins

def bin_packing_first_fit_decreasing(items, bin_capacity):
    """
    First Fit Decreasing (FFD) for Bin Packing.
    
    Algorithm:
    1. Sort items in decreasing order
    2. Apply First Fit
    
    Approximation ratio: 11/9 * OPT + 6/9
    Better than First Fit!
    
    Time: O(n log n + n * m)
    """
    sorted_items = sorted(items, reverse=True)
    return bin_packing_first_fit(sorted_items, bin_capacity)

# LOAD BALANCING

def load_balancing_greedy(jobs, m):
    """
    Greedy approximation for Load Balancing.
    
    Problem: Assign n jobs to m machines, minimize maximum load.
    
    Algorithm:
    Assign each job to machine with current minimum load.
    
    Approximation ratio: 2 - 1/m
    
    Time: O(n log m) with heap
    """
    import heapq
    
    # Min heap of (load, machine_id)
    machines = [(0, i) for i in range(m)]
    heapq.heapify(machines)
    
    assignment = [[] for _ in range(m)]
    
    for job in jobs:
        # Assign to machine with minimum load
        load, machine = heapq.heappop(machines)
        
        assignment[machine].append(job)
        
        heapq.heappush(machines, (load + job, machine))
    
    max_load = max(sum(jobs) for jobs in assignment)
    
    return assignment, max_load

# MAX CUT APPROXIMATION

def max_cut_randomized(graph):
    """
    Randomized 2-approximation for Max Cut.
    
    Algorithm:
    Randomly partition vertices into two sets.
    
    Expected approximation ratio: 2
    (Expected cut size ≥ OPT/2)
    
    Time: O(V)
    """
    vertices = set()
    for u, v in graph:
        vertices.add(u)
        vertices.add(v)
    
    vertices = list(vertices)
    
    # Random partition
    partition = {v: random.choice([0, 1]) for v in vertices}
    
    # Count edges crossing partition
    cut_size = sum(
        1 for u, v in graph
        if partition[u] != partition[v]
    )
    
    return partition, cut_size

def max_cut_greedy(graph):
    """
    Greedy algorithm for Max Cut.
    Also 2-approximation, deterministic.
    
    Algorithm:
    Build sets incrementally, placing each vertex in set
    that maximizes current cut.
    """
    vertices = set()
    for u, v in graph:
        vertices.add(u)
        vertices.add(v)
    
    set_A = set()
    set_B = set()
    
    for v in vertices:
        # Count edges to each set
        edges_to_A = sum(1 for u, w in graph if (u == v and w in set_A) or (w == v and u in set_A))
        edges_to_B = sum(1 for u, w in graph if (u == v and w in set_B) or (w == v and u in set_B))
        
        # Place in set with more edges
        if edges_to_A > edges_to_B:
            set_B.add(v)
        else:
            set_A.add(v)
    
    cut_size = sum(
        1 for u, v in graph
        if (u in set_A and v in set_B) or (u in set_B and v in set_A)
    )
    
    return set_A, set_B, cut_size

# APPROXIMATION RATIO MEASUREMENT

def measure_approximation_ratio(approx_algo, optimal_algo, test_cases):
    """
    Empirically measure approximation ratio.
    """
    ratios = []
    
    for instance in test_cases:
        approx_cost = approx_algo(instance)
        optimal_cost = optimal_algo(instance)
        
        ratio = approx_cost / optimal_cost if optimal_cost > 0 else 0
        ratios.append(ratio)
    
    return {
        'max_ratio': max(ratios),
        'avg_ratio': sum(ratios) / len(ratios),
        'min_ratio': min(ratios)
    }

# PTAS (POLYNOMIAL TIME APPROXIMATION SCHEME)

def knapsack_fptas(weights, values, capacity, epsilon):
    """
    FPTAS for Knapsack: (1 + ε)-approximation in polynomial time.
    
    Runs in O(n^3 / ε) time.
    
    For any ε > 0, gets solution within (1 + ε) of optimal.
    """
    # Scale values
    max_value = max(values)
    K = epsilon * max_value / len(values)
    
    scaled_values = [int(v / K) for v in values]
    
    # Run DP on scaled values
    # (Implementation similar to knapsack DP from Module 8)
    
    # Return actual values
    pass
```

### Test Cases
- Compare approximation to optimal (small instances)
- Measure approximation ratio
- Large instances (optimal infeasible)
- Worst-case examples

### Deliverables
- `vertex_cover_approx.{py,go,zig}` - 2-approximation
- `set_cover_greedy.{py,go,zig}` - ln n approximation
- `tsp_approx.{py,go,zig}` - MST and Christofides
- `bin_packing_approx.{py,go,zig}` - All variants
- `load_balancing.{py,go,zig}` - Greedy algorithm
- `max_cut.{py,go,zig}` - Randomized and greedy
- `approximation_analysis.md` - Proofs of ratios
- `empirical_evaluation.{py,go,zig}` - Measure actual performance
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All approximation algorithms implemented
- Understanding of approximation ratios
- Can prove approximation guarantees
- Know when approximation is acceptable

---

---

## Back to Module

[← Back to Module 13: NP-Completeness & Approximation Algorithms](../module_13/README.md)

## Navigation

- [Previous Lab](./lab_13_2.md) (if exists)
- [Next Lab](./lab_13_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 13, Lab 3 of 5*
