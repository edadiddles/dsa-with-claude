# Module 13: NP-Completeness & Approximation Algorithms

**Duration**: 2 weeks (Optional Advanced Module)  
**Difficulty**: Hard

## Overview

This module explores the theory of computational complexity, NP-completeness, and practical approaches to intractable problems. You'll learn to recognize NP-complete problems, prove hardness through reductions, and design approximation algorithms when exact solutions are infeasible.

**Why this matters**: Many real-world optimization problems are NP-complete - from scheduling to circuit design to protein folding. Understanding computational complexity theory helps you recognize when problems are fundamentally hard and know when to use approximation algorithms, heuristics, or accept exponential solutions. This knowledge is crucial for research, algorithm design, and practical engineering decisions.

## Learning Objectives

- Understand P, NP, and NP-completeness theory
- Prove problems are NP-complete using reductions
- Design approximation algorithms with guaranteed bounds
- Recognize common NP-complete problems
- Apply heuristics and randomized approaches
- Understand when exact solutions are infeasible
- Make informed engineering tradeoffs

## Topics Covered

- Complexity classes: P, NP, NP-complete, NP-hard
- Cook-Levin theorem and SAT
- Polynomial-time reductions
- Classic NP-complete problems (SAT, 3-SAT, Vertex Cover, TSP, etc.)
- Approximation algorithms and approximation ratios
- Randomized approximation
- Heuristics and metaheuristics
- Parameterized complexity
- Practical approaches to hard problems

---

## Lab 13.1: Understanding NP-Completeness

**Duration**: 5-6 hours  
**Difficulty**: Hard

### Objectives
- Understand complexity classes P and NP
- Learn the concept of NP-completeness
- Recognize NP-complete problems
- Understand polynomial-time reductions

### Requirements

1. **Understand complexity theory**:
   - **P**: Problems solvable in polynomial time
   - **NP**: Problems verifiable in polynomial time
   - **NP-complete**: Hardest problems in NP
   - **NP-hard**: At least as hard as NP-complete
   - **Relationship between classes**

2. **Study classic NP-complete problems**:
   - Boolean Satisfiability (SAT)
   - 3-SAT
   - Vertex Cover
   - Independent Set
   - Clique
   - Hamiltonian Cycle
   - Traveling Salesman Problem
   - Graph Coloring
   - Subset Sum
   - Partition Problem

3. **Implement verification algorithms**:
   - Certificate verification
   - Demonstrate problems are in NP
   - Brute force solutions for small instances

### Implementation Details

```python
# COMPLEXITY CLASSES DEMONSTRATION

def is_in_P(problem_instance):
    """
    P: Problems solvable in polynomial time.
    Examples: Sorting, Shortest Path, MST, etc.
    """
    # These problems have polynomial-time algorithms
    pass

def is_in_NP(problem_instance):
    """
    NP: Problems where solutions can be verified in polynomial time.
    
    A problem is in NP if:
    1. Given a solution (certificate)
    2. We can verify it's correct in polynomial time
    
    Note: All P problems are also in NP.
    """
    pass

# SAT (BOOLEAN SATISFIABILITY)

def evaluate_boolean_formula(formula, assignment):
    """
    Evaluate a boolean formula given variable assignment.
    
    formula: CNF formula as list of clauses
    assignment: dict mapping variables to True/False
    
    This is polynomial time - SAT is in NP because we can verify solutions.
    """
    for clause in formula:
        clause_satisfied = False
        for literal in clause:
            if literal.startswith('~'):
                var = literal[1:]
                if not assignment.get(var, False):
                    clause_satisfied = True
                    break
            else:
                if assignment.get(literal, False):
                    clause_satisfied = True
                    break
        
        if not clause_satisfied:
            return False
    
    return True

def verify_SAT_solution(formula, assignment):
    """
    Verify a SAT solution in polynomial time.
    This proves SAT is in NP.
    
    Time: O(mn) where m = clauses, n = variables
    """
    return evaluate_boolean_formula(formula, assignment)

def solve_SAT_bruteforce(formula, variables):
    """
    Solve SAT by trying all assignments.
    Time: O(2^n * mn) - exponential!
    
    This is why SAT is hard - no known polynomial algorithm.
    """
    from itertools import product
    
    for assignment_values in product([False, True], repeat=len(variables)):
        assignment = dict(zip(variables, assignment_values))
        if evaluate_boolean_formula(formula, assignment):
            return True, assignment
    
    return False, None

# 3-SAT

def verify_3SAT_solution(formula, assignment):
    """
    Verify 3-SAT solution.
    3-SAT: Each clause has exactly 3 literals.
    
    3-SAT is NP-complete (SAT is also NP-complete).
    """
    for clause in formula:
        if len(clause) != 3:
            raise ValueError("Not a 3-SAT formula")
        
        clause_satisfied = any(
            (not lit.startswith('~') and assignment.get(lit, False)) or
            (lit.startswith('~') and not assignment.get(lit[1:], False))
            for lit in clause
        )
        
        if not clause_satisfied:
            return False
    
    return True

# VERTEX COVER

def verify_vertex_cover(graph, cover, k):
    """
    Verify that cover is a vertex cover of size <= k.
    
    Vertex Cover: Select at most k vertices such that every edge
    has at least one endpoint in the cover.
    
    Time: O(E) - polynomial verification, so in NP
    """
    if len(cover) > k:
        return False
    
    cover_set = set(cover)
    
    # Check every edge is covered
    for u, v in graph:
        if u not in cover_set and v not in cover_set:
            return False
    
    return True

def solve_vertex_cover_bruteforce(graph, k):
    """
    Solve vertex cover by trying all subsets.
    Time: O(n choose k * E) - exponential
    """
    from itertools import combinations
    
    vertices = set()
    for u, v in graph:
        vertices.add(u)
        vertices.add(v)
    
    vertices = list(vertices)
    
    for size in range(k + 1):
        for cover in combinations(vertices, size):
            if verify_vertex_cover(graph, cover, k):
                return True, cover
    
    return False, None

# INDEPENDENT SET

def verify_independent_set(graph, ind_set, k):
    """
    Verify that ind_set is an independent set of size >= k.
    
    Independent Set: Select at least k vertices with no edges between them.
    
    Note: Complement of vertex cover!
    """
    if len(ind_set) < k:
        return False
    
    ind_set = set(ind_set)
    
    # Check no edge has both endpoints in set
    for u, v in graph:
        if u in ind_set and v in ind_set:
            return False
    
    return True

# CLIQUE

def verify_clique(graph, clique, k):
    """
    Verify that clique is a clique of size >= k.
    
    Clique: Select at least k vertices with edges between all pairs.
    """
    if len(clique) < k:
        return False
    
    clique = list(clique)
    
    # Check all pairs are connected
    for i in range(len(clique)):
        for j in range(i + 1, len(clique)):
            edge = (clique[i], clique[j])
            reverse_edge = (clique[j], clique[i])
            if edge not in graph and reverse_edge not in graph:
                return False
    
    return True

# HAMILTONIAN CYCLE

def verify_hamiltonian_cycle(graph, cycle):
    """
    Verify that cycle is a Hamiltonian cycle.
    
    Hamiltonian Cycle: Visit every vertex exactly once and return to start.
    """
    vertices = set()
    for u, v in graph:
        vertices.add(u)
        vertices.add(v)
    
    # Check all vertices visited
    if set(cycle[:-1]) != vertices:
        return False
    
    # Check cycle is valid
    if cycle[0] != cycle[-1]:
        return False
    
    # Check all edges exist
    for i in range(len(cycle) - 1):
        edge = (cycle[i], cycle[i + 1])
        reverse_edge = (cycle[i + 1], cycle[i])
        if edge not in graph and reverse_edge not in graph:
            return False
    
    return True

# TRAVELING SALESMAN PROBLEM (DECISION VERSION)

def verify_TSP_solution(distances, tour, max_cost):
    """
    Verify TSP solution has cost <= max_cost.
    
    TSP (decision): Is there a tour with cost <= max_cost?
    """
    n = len(tour) - 1
    
    # Check valid tour (visits all cities once)
    if len(set(tour[:-1])) != n or tour[0] != tour[-1]:
        return False
    
    # Compute cost
    total_cost = 0
    for i in range(len(tour) - 1):
        u, v = tour[i], tour[i + 1]
        if (u, v) not in distances:
            return False
        total_cost += distances[(u, v)]
    
    return total_cost <= max_cost

# GRAPH COLORING

def verify_graph_coloring(graph, coloring, k):
    """
    Verify graph can be colored with k colors.
    
    Graph k-Coloring: Assign one of k colors to each vertex
    such that adjacent vertices have different colors.
    """
    # Check uses at most k colors
    if max(coloring.values()) >= k:
        return False
    
    # Check adjacent vertices have different colors
    for u, v in graph:
        if coloring[u] == coloring[v]:
            return False
    
    return True

# SUBSET SUM

def verify_subset_sum(numbers, subset, target):
    """
    Verify subset sums to target.
    
    Subset Sum: Is there a subset that sums to exactly target?
    """
    return sum(numbers[i] for i in subset) == target

def solve_subset_sum_bruteforce(numbers, target):
    """
    Solve subset sum by trying all subsets.
    Time: O(2^n) - exponential
    """
    from itertools import combinations
    
    n = len(numbers)
    
    for size in range(n + 1):
        for subset in combinations(range(n), size):
            if sum(numbers[i] for i in subset) == target:
                return True, subset
    
    return False, None

# PARTITION PROBLEM

def verify_partition(numbers, partition1):
    """
    Verify numbers can be partitioned into two equal-sum subsets.
    
    Partition: Divide numbers into two sets with equal sum.
    Special case of Subset Sum.
    """
    total = sum(numbers)
    
    if total % 2 != 0:
        return False
    
    sum1 = sum(numbers[i] for i in partition1)
    
    return sum1 == total // 2

# COMPLEXITY CLASS RELATIONSHIPS

def demonstrate_complexity_classes():
    """
    Demonstrate relationships between complexity classes:
    
    P ⊆ NP
    - Every problem solvable in polynomial time
      is also verifiable in polynomial time
    
    NP-complete ⊆ NP
    - NP-complete problems are the hardest in NP
    
    NP-hard
    - At least as hard as NP-complete
    - May not be in NP (decision vs. optimization)
    
    Unknown: P = NP ?
    - Million dollar question!
    - If P = NP, polynomial algorithms exist for all NP problems
    - If P ≠ NP, some problems are inherently hard
    
    Practical impact:
    - P = NP: Cryptography breaks, optimization becomes easy
    - P ≠ NP (believed): Some problems need approximations/heuristics
    """
    pass

# RECOGNITION GUIDE

def is_problem_likely_np_complete(problem_description):
    """
    Heuristics for recognizing NP-complete problems:
    
    Warning signs:
    1. "Find the minimum/maximum ..." (optimization)
    2. "Is there a subset/selection that ..." (combinatorial)
    3. "Visit all ..." or "cover all ..." (covering problems)
    4. Problem involves graphs with global properties
    5. Problem is about partitioning or packing
    6. Small changes make problem easy/hard (3-SAT vs 2-SAT)
    7. No known polynomial algorithm despite extensive research
    
    Common NP-complete patterns:
    - Graph problems: Coloring, Clique, Independent Set, Hamiltonian Cycle
    - Covering: Vertex Cover, Set Cover
    - Packing: Bin Packing, Knapsack
    - Scheduling: Job Shop Scheduling
    - Logic: SAT, 3-SAT
    """
    pass
```

### Test Cases
- Small problem instances (brute force solvable)
- Verify polynomial-time verification
- Compare verification time vs. solution time
- Test on known NP-complete problems

### Deliverables
- `complexity_classes.{py,go,zig}` - P, NP definitions and examples
- `sat_solver.{py,go,zig}` - SAT and 3-SAT verification and brute force
- `graph_problems.{py,go,zig}` - Vertex Cover, Clique, Independent Set
- `verification_framework.{py,go,zig}` - Generic NP verification
- `np_complete_examples.md` - Catalog of NP-complete problems
- `complexity_theory.md` - Explanation of P vs NP
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Understanding of P, NP, NP-complete
- Can verify solutions in polynomial time
- Recognition of NP-complete problems
- Understanding why these problems are hard

---

## Lab 13.2: Polynomial-Time Reductions

**Duration**: 6-8 hours  
**Difficulty**: Hard

### Objectives
- Understand polynomial-time reductions
- Prove NP-completeness through reductions
- Master reduction techniques
- Build reduction chains

### Requirements

1. **Implement classic reductions**:
   - **3-SAT ≤_p Vertex Cover**
   - **3-SAT ≤_p Independent Set**
   - **3-SAT ≤_p Clique**
   - **Independent Set ≤_p Clique** (via complement graph)
   - **Vertex Cover ≤_p Independent Set**
   - **Subset Sum ≤_p Partition**

2. **Understand reduction technique**:
   - Transform instance A to instance B
   - Solution to B implies solution to A
   - Transformation is polynomial time
   - Correctness proof

3. **Build reduction chains**:
   - SAT → 3-SAT → Vertex Cover → ...
   - Understand Cook-Levin theorem

### Implementation Details

```python
# POLYNOMIAL-TIME REDUCTIONS

def reduction_framework(problem_A_instance, transform, solve_B, back_transform):
    """
    Generic reduction framework:
    1. Transform instance of A to instance of B
    2. Solve B
    3. Transform B's solution back to A's solution
    
    If this works in polynomial time, A ≤_p B
    """
    # Transform A to B
    problem_B_instance = transform(problem_A_instance)
    
    # Solve B
    solution_B = solve_B(problem_B_instance)
    
    # Transform back
    solution_A = back_transform(solution_B)
    
    return solution_A

# REDUCTION: 3-SAT to VERTEX COVER

def reduce_3SAT_to_vertex_cover(formula):
    """
    Reduce 3-SAT to Vertex Cover.
    
    Construction:
    1. For each clause (a ∨ b ∨ c), create triangle of 3 vertices
    2. For each variable x, create edge between x and ¬x across clauses
    3. Set k = number of clauses + number of variables
    
    Claim: Formula is satisfiable iff graph has vertex cover of size k.
    
    Time: O(mn) where m = clauses, n = variables
    """
    graph = []
    k = 0
    
    # Track variable occurrences
    var_occurrences = {}  # var -> list of (clause_id, literal_id, negated)
    
    # Create triangle for each clause
    for clause_id, clause in enumerate(formula):
        if len(clause) != 3:
            raise ValueError("Not a 3-SAT formula")
        
        # Add triangle edges
        graph.append((f"c{clause_id}_0", f"c{clause_id}_1"))
        graph.append((f"c{clause_id}_1", f"c{clause_id}_2"))
        graph.append((f"c{clause_id}_2", f"c{clause_id}_0"))
        
        # Track variables
        for lit_id, literal in enumerate(clause):
            negated = literal.startswith('~')
            var = literal[1:] if negated else literal
            
            if var not in var_occurrences:
                var_occurrences[var] = []
            var_occurrences[var].append((clause_id, lit_id, negated))
        
        k += 1  # Need 1 vertex from each triangle (covers 2 edges)
    
    # Add edges between opposite literals
    for var, occurrences in var_occurrences.items():
        positive = [(c, l) for c, l, neg in occurrences if not neg]
        negative = [(c, l) for c, l, neg in occurrences if neg]
        
        for c1, l1 in positive:
            for c2, l2 in negative:
                graph.append((f"c{c1}_{l1}", f"c{c2}_{l2}"))
        
        k += 1  # Need 1 vertex for each variable
    
    return graph, k

def vertex_cover_to_3SAT_solution(formula, graph, vertex_cover):
    """
    Given vertex cover, construct SAT assignment.
    """
    assignment = {}
    
    # Extract variable assignments from vertex cover
    # (Implementation depends on how vertices are labeled)
    
    return assignment

# REDUCTION: INDEPENDENT SET to CLIQUE

def reduce_independent_set_to_clique(graph, k):
    """
    Reduce Independent Set to Clique via complement graph.
    
    Construction:
    1. Create complement graph G' where (u,v) ∈ G' iff (u,v) ∉ G
    2. Independent set of size k in G = clique of size k in G'
    
    Time: O(V^2)
    """
    vertices = set()
    for u, v in graph:
        vertices.add(u)
        vertices.add(v)
    
    # Create complement graph
    complement = []
    for u in vertices:
        for v in vertices:
            if u != v:
                edge = (u, v)
                reverse = (v, u)
                # Add edge if NOT in original graph
                if edge not in graph and reverse not in graph:
                    complement.append((u, v))
    
    return complement, k

# REDUCTION: VERTEX COVER to INDEPENDENT SET

def reduce_vertex_cover_to_independent_set(graph, k):
    """
    Reduce Vertex Cover to Independent Set.
    
    Construction:
    Vertex cover of size k in G = independent set of size n-k in G.
    (Where n = number of vertices)
    
    Proof:
    - S is vertex cover → V-S is independent set
    - If V-S had edge, that edge wouldn't be covered by S
    
    Time: O(V)
    """
    vertices = set()
    for u, v in graph:
        vertices.add(u)
        vertices.add(v)
    
    n = len(vertices)
    k_independent = n - k
    
    return graph, k_independent  # Same graph, different k

# REDUCTION: SUBSET SUM to PARTITION

def reduce_subset_sum_to_partition(numbers, target):
    """
    Reduce Subset Sum to Partition.
    
    Construction:
    Given Subset Sum instance (numbers, target):
    1. Let S = sum(numbers)
    2. Add two elements: S - 2*target and target
    3. Partition problem: Can we split into two equal sums?
    
    Time: O(n)
    """
    S = sum(numbers)
    
    # Create Partition instance
    partition_numbers = numbers + [S - 2 * target, target]
    
    # If Partition has solution, one part sums to (S + (S - 2*target) + target) / 2 = S
    # This means Subset Sum has solution
    
    return partition_numbers

# REDUCTION: 3-SAT to HAMILTONIAN CYCLE

def reduce_3SAT_to_hamiltonian_cycle(formula):
    """
    Reduce 3-SAT to Hamiltonian Cycle.
    (Classic but complex reduction)
    
    Construction sketch:
    1. Create gadgets for each variable
    2. Create gadgets for each clause
    3. Connect gadgets to enforce constraints
    
    This is a complex reduction - simplified here.
    """
    # Simplified placeholder
    graph = []
    
    # Variable gadgets: Paths that can be traversed left-to-right or right-to-left
    # Choice represents true/false assignment
    
    # Clause gadgets: Ensure at least one literal in each clause is true
    
    return graph

# REDUCTION VERIFICATION

def verify_reduction(reduction_func, problem_A_verify, problem_B_verify,
                     test_cases):
    """
    Verify a reduction is correct:
    1. Transformation is polynomial time
    2. YES instance of A → YES instance of B
    3. NO instance of A → NO instance of B
    """
    import time
    
    for instance_A, expected_result in test_cases:
        # Time the transformation
        start = time.time()
        instance_B = reduction_func(instance_A)
        transform_time = time.time() - start
        
        # Check polynomial time (heuristic)
        # Should verify asymptotic behavior more rigorously
        
        # Verify correctness
        result_A = problem_A_verify(instance_A)
        result_B = problem_B_verify(instance_B)
        
        assert result_A == result_B, f"Reduction incorrect: {result_A} != {result_B}"
        assert result_A == expected_result

# REDUCTION CHAINS

def demonstrate_reduction_chain():
    """
    Classic reduction chain showing NP-completeness:
    
    SAT (Cook-Levin: first NP-complete problem)
    ↓
    3-SAT (SAT ≤_p 3-SAT)
    ↓
    Vertex Cover (3-SAT ≤_p Vertex Cover)
    ↓
    Independent Set (Vertex Cover ≤_p Independent Set)
    ↓
    Clique (Independent Set ≤_p Clique)
    
    To prove problem X is NP-complete:
    1. Show X is in NP (polynomial verification)
    2. Reduce known NP-complete problem to X
    
    This proves X is as hard as all NP problems.
    """
    pass

# PROOF TEMPLATE

def np_completeness_proof_template():
    """
    Template for NP-completeness proof:
    
    Theorem: Problem X is NP-complete.
    
    Proof:
    1. Show X ∈ NP:
       - Given certificate (solution)
       - Verify in polynomial time
       - [Provide verification algorithm]
    
    2. Show X is NP-hard:
       - Choose known NP-complete problem Y
       - Construct reduction Y ≤_p X
       - Describe transformation
       - Prove transformation is polynomial time
       - Prove correctness:
         a) YES instance of Y → YES instance of X
         b) NO instance of Y → NO instance of X
    
    Therefore X is NP-complete. QED.
    """
    pass
```

### Test Cases
- Small instances where both problems solvable
- Verify reduction preserves satisfiability
- Test transformation time (polynomial)
- Verify correctness of reduction

### Deliverables
- `reductions.{py,go,zig}` - All classic reductions
- `reduction_framework.{py,go,zig}` - Generic reduction template
- `reduction_verifier.{py,go,zig}` - Test reduction correctness
- `np_completeness_proofs.md` - Formal proofs
- `reduction_chains.md` - Visualization of reduction relationships
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Can construct polynomial-time reductions
- Understanding of reduction correctness
- Can prove NP-completeness
- Recognition of reduction patterns

---

## Lab 13.3: Approximation Algorithms

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

## Lab 13.4: Heuristics & Metaheuristics

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Implement practical heuristics for NP-hard problems
- Understand local search techniques
- Apply metaheuristics
- Compare heuristic performance

### Requirements

1. **Implement heuristics**:
   - **Local Search** (hill climbing)
   - **Simulated Annealing**
   - **Genetic Algorithms**
   - **Tabu Search**
   - **Greedy Randomized Adaptive Search (GRASP)**

2. **Apply to problems**:
   - TSP
   - Graph Coloring
   - SAT
   - Job Scheduling

3. **Evaluate performance**:
   - Solution quality
   - Running time
   - Consistency
   - Parameter sensitivity

### Implementation Details

```python
import random
import math

# LOCAL SEARCH (HILL CLIMBING)

def local_search_tsp(tour, distances):
    """
    Local search for TSP using 2-opt moves.
    
    Algorithm:
    1. Start with initial tour
    2. Repeatedly apply 2-opt improvements
    3. Stop when no improvement found
    
    Not guaranteed to find global optimum (gets stuck in local optimum).
    """
    n = len(tour)
    improved = True
    
    while improved:
        improved = False
        
        for i in range(n - 1):
            for j in range(i + 2, n):
                # Try reversing tour[i:j]
                new_tour = tour[:i] + tour[i:j][::-1] + tour[j:]
                
                if tour_cost(new_tour, distances) < tour_cost(tour, distances):
                    tour = new_tour
                    improved = True
                    break
            
            if improved:
                break
    
    return tour

def tour_cost(tour, distances):
    """Calculate total tour cost"""
    cost = 0
    for i in range(len(tour) - 1):
        cost += distances.get((tour[i], tour[i+1]), float('inf'))
    return cost

# SIMULATED ANNEALING

def simulated_annealing_tsp(tour, distances, initial_temp=1000, cooling_rate=0.995, iterations=10000):
    """
    Simulated Annealing for TSP.
    
    Algorithm:
    1. Start with initial solution
    2. At each iteration:
       - Generate neighbor solution
       - Accept if better, or with probability e^(-Δ/T) if worse
    3. Decrease temperature
    
    Allows escaping local optima!
    """
    current = tour
    current_cost = tour_cost(current, distances)
    best = current
    best_cost = current_cost
    
    temp = initial_temp
    
    for iteration in range(iterations):
        # Generate neighbor (2-opt move)
        i, j = sorted(random.sample(range(len(current)), 2))
        neighbor = current[:i] + current[i:j][::-1] + current[j:]
        neighbor_cost = tour_cost(neighbor, distances)
        
        # Accept or reject
        delta = neighbor_cost - current_cost
        
        if delta < 0 or random.random() < math.exp(-delta / temp):
            current = neighbor
            current_cost = neighbor_cost
            
            if current_cost < best_cost:
                best = current
                best_cost = current_cost
        
        # Cool down
        temp *= cooling_rate
    
    return best, best_cost

# GENETIC ALGORITHM

def genetic_algorithm_tsp(cities, distances, population_size=100, generations=500, mutation_rate=0.01):
    """
    Genetic Algorithm for TSP.
    
    Algorithm:
    1. Initialize population of random tours
    2. For each generation:
       - Select parents (fitness-based)
       - Crossover to create offspring
       - Mutate offspring
       - Replace population
    """
    # Initialize population
    population = [random.sample(cities, len(cities)) for _ in range(population_size)]
    
    for gen in range(generations):
        # Evaluate fitness
        fitness = [1 / tour_cost(tour + [tour[0]], distances) for tour in population]
        
        # Selection (tournament)
        new_population = []
        
        for _ in range(population_size):
            # Tournament selection
            parent1 = tournament_select(population, fitness)
            parent2 = tournament_select(population, fitness)
            
            # Crossover
            child = order_crossover(parent1, parent2)
            
            # Mutation
            if random.random() < mutation_rate:
                child = swap_mutate(child)
            
            new_population.append(child)
        
        population = new_population
    
    # Return best tour
    best_tour = min(population, key=lambda t: tour_cost(t + [t[0]], distances))
    return best_tour, tour_cost(best_tour + [best_tour[0]], distances)

def tournament_select(population, fitness, tournament_size=5):
    """Tournament selection"""
    tournament = random.sample(list(zip(population, fitness)), tournament_size)
    return max(tournament, key=lambda x: x[1])[0]

def order_crossover(parent1, parent2):
    """Order crossover for TSP"""
    size = len(parent1)
    start, end = sorted(random.sample(range(size), 2))
    
    child = [None] * size
    child[start:end] = parent1[start:end]
    
    pos = end
    for city in parent2[end:] + parent2[:end]:
        if city not in child:
            if pos >= size:
                pos = 0
            child[pos] = city
            pos += 1
    
    return child

def swap_mutate(tour):
    """Swap mutation"""
    tour = tour[:]
    i, j = random.sample(range(len(tour)), 2)
    tour[i], tour[j] = tour[j], tour[i]
    return tour

# TABU SEARCH

def tabu_search_tsp(tour, distances, max_iterations=1000, tabu_tenure=20):
    """
    Tabu Search for TSP.
    
    Algorithm:
    1. Maintain tabu list of recently visited solutions
    2. At each iteration:
       - Explore neighborhood
       - Select best non-tabu move
       - Update tabu list
    
    Prevents cycling back to recent solutions.
    """
    current = tour
    current_cost = tour_cost(current, distances)
    best = current
    best_cost = current_cost
    
    tabu_list = []
    
    for iteration in range(max_iterations):
        # Generate all 2-opt neighbors
        neighbors = []
        
        for i in range(len(current) - 1):
            for j in range(i + 2, len(current)):
                neighbor = current[:i] + current[i:j][::-1] + current[j:]
                cost = tour_cost(neighbor, distances)
                neighbors.append((neighbor, cost, (i, j)))
        
        # Find best non-tabu neighbor
        neighbors.sort(key=lambda x: x[1])
        
        for neighbor, cost, move in neighbors:
            if move not in tabu_list:
                current = neighbor
                current_cost = cost
                
                # Update tabu list
                tabu_list.append(move)
                if len(tabu_list) > tabu_tenure:
                    tabu_list.pop(0)
                
                # Update best
                if cost < best_cost:
                    best = neighbor
                    best_cost = cost
                
                break
    
    return best, best_cost

# GRASP (Greedy Randomized Adaptive Search Procedure)

def grasp_tsp(cities, distances, iterations=100):
    """
    GRASP for TSP.
    
    Algorithm:
    1. Construction phase: Build solution greedily with randomization
    2. Local search phase: Improve solution
    3. Repeat and keep best
    """
    best = None
    best_cost = float('inf')
    
    for _ in range(iterations):
        # Construction phase
        tour = greedy_randomized_construction(cities, distances)
        
        # Local search phase
        tour = local_search_tsp(tour + [tour[0]], distances)[:-1]
        
        cost = tour_cost(tour + [tour[0]], distances)
        
        if cost < best_cost:
            best = tour
            best_cost = cost
    
    return best, best_cost

def greedy_randomized_construction(cities, distances, alpha=0.3):
    """
    Build TSP tour greedily with randomization.
    
    alpha controls greediness (0 = fully random, 1 = fully greedy)
    """
    tour = [random.choice(cities)]
    remaining = set(cities) - {tour[0]}
    
    while remaining:
        current = tour[-1]
        
        # Calculate distances to remaining cities
        dists = [(city, distances.get((current, city), float('inf'))) 
                for city in remaining]
        dists.sort(key=lambda x: x[1])
        
        # Restricted candidate list (RCL)
        min_dist = dists[0][1]
        max_dist = dists[-1][1]
        threshold = min_dist + alpha * (max_dist - min_dist)
        
        rcl = [city for city, dist in dists if dist <= threshold]
        
        # Random selection from RCL
        next_city = random.choice(rcl)
        tour.append(next_city)
        remaining.remove(next_city)
    
    return tour

# HEURISTIC COMPARISON

def compare_heuristics(cities, distances):
    """Compare all heuristics on same instance"""
    import time
    
    results = {}
    
    # Initial random tour
    initial_tour = random.sample(cities, len(cities))
    
    # Local Search
    start = time.time()
    ls_tour, ls_cost = local_search_tsp(initial_tour + [initial_tour[0]], distances)
    ls_time = time.time() - start
    results['Local Search'] = {'cost': ls_cost, 'time': ls_time}
    
    # Simulated Annealing
    start = time.time()
    sa_tour, sa_cost = simulated_annealing_tsp(initial_tour + [initial_tour[0]], distances)
    sa_time = time.time() - start
    results['Simulated Annealing'] = {'cost': sa_cost, 'time': sa_time}
    
    # Genetic Algorithm
    start = time.time()
    ga_tour, ga_cost = genetic_algorithm_tsp(cities, distances)
    ga_time = time.time() - start
    results['Genetic Algorithm'] = {'cost': ga_cost, 'time': ga_time}
    
    # Tabu Search
    start = time.time()
    ts_tour, ts_cost = tabu_search_tsp(initial_tour + [initial_tour[0]], distances)
    ts_time = time.time() - start
    results['Tabu Search'] = {'cost': ts_cost, 'time': ts_time}
    
    # GRASP
    start = time.time()
    grasp_tour, grasp_cost = grasp_tsp(cities, distances)
    grasp_time = time.time() - start
    results['GRASP'] = {'cost': grasp_cost, 'time': grasp_time}
    
    return results
```

### Test Cases
- TSP instances of various sizes
- Compare solution quality
- Measure running time
- Parameter sensitivity analysis

### Deliverables
- `local_search.{py,go,zig}` - Hill climbing variants
- `simulated_annealing.{py,go,zig}` - SA implementation
- `genetic_algorithm.{py,go,zig}` - GA for TSP
- `tabu_search.{py,go,zig}` - Tabu search
- `grasp.{py,go,zig}` - GRASP algorithm
- `heuristic_comparison.{py,go,zig}` - Benchmark all methods
- `parameter_tuning.{py,go,zig}` - Optimize parameters
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- All heuristics implemented
- Understanding of metaheuristic concepts
- Can tune parameters effectively
- Know tradeoffs between methods

---

## Lab 13.5: Practical Approaches to Hard Problems

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Apply multiple techniques to real problems
- Make engineering tradeoffs
- Build practical solvers
- Understand when to use each approach

### Requirements

1. **Build complete solvers**:
   - **SAT Solver** (DPLL with heuristics)
   - **TSP Solver** (combining exact and approximate)
   - **Scheduling System** (with constraints)
   - **Resource Allocation** (practical optimization)

2. **Hybrid approaches**:
   - Exact for small subproblems
   - Approximation for large instances
   - Branch and bound with approximation
   - Preprocessing and reduction

3. **Real-world considerations**:
   - Time budgets
   - Solution quality requirements
   - Scalability
   - Robustness

### Implementation Details

```python
# SAT SOLVER (DPLL Algorithm)

def dpll_sat_solver(formula, assignment=None):
    """
    DPLL algorithm for SAT solving.
    
    More sophisticated than brute force:
    1. Unit propagation
    2. Pure literal elimination
    3. Intelligent branching
    
    Still exponential worst-case, but much better in practice.
    """
    if assignment is None:
        assignment = {}
    
    # Simplify formula
    formula = simplify_formula(formula, assignment)
    
    # Base cases
    if not formula:
        return True, assignment  # All clauses satisfied
    
    if [] in formula:
        return False, None  # Empty clause (unsatisfiable)
    
    # Unit propagation
    unit = find_unit_clause(formula)
    if unit:
        var, value = unit
        assignment[var] = value
        return dpll_sat_solver(formula, assignment)
    
    # Pure literal elimination
    pure = find_pure_literal(formula)
    if pure:
        var, value = pure
        assignment[var] = value
        return dpll_sat_solver(formula, assignment)
    
    # Branching
    var = choose_variable(formula)
    
    # Try true
    assignment[var] = True
    sat, result = dpll_sat_solver(formula, assignment.copy())
    if sat:
        return True, result
    
    # Try false
    assignment[var] = False
    return dpll_sat_solver(formula, assignment.copy())

def simplify_formula(formula, assignment):
    """Simplify formula given partial assignment"""
    simplified = []
    
    for clause in formula:
        # Check if clause is satisfied
        satisfied = False
        new_clause = []
        
        for lit in clause:
            negated = lit.startswith('~')
            var = lit[1:] if negated else lit
            
            if var in assignment:
                if assignment[var] != negated:
                    satisfied = True
                    break
            else:
                new_clause.append(lit)
        
        if not satisfied:
            simplified.append(new_clause)
    
    return simplified

def find_unit_clause(formula):
    """Find clause with single literal"""
    for clause in formula:
        if len(clause) == 1:
            lit = clause[0]
            negated = lit.startswith('~')
            var = lit[1:] if negated else lit
            return (var, not negated)
    return None

def find_pure_literal(formula):
    """Find variable that appears only positive or only negative"""
    literals = {}
    
    for clause in formula:
        for lit in clause:
            negated = lit.startswith('~')
            var = lit[1:] if negated else lit
            
            if var not in literals:
                literals[var] = set()
            literals[var].add(negated)
    
    for var, polarities in literals.items():
        if len(polarities) == 1:
            return (var, not list(polarities)[0])
    
    return None

def choose_variable(formula):
    """Choose variable for branching (heuristic)"""
    # Simple heuristic: choose most frequent variable
    counts = {}
    
    for clause in formula:
        for lit in clause:
            var = lit[1:] if lit.startswith('~') else lit
            counts[var] = counts.get(var, 0) + 1
    
    return max(counts, key=counts.get)

# HYBRID TSP SOLVER

def hybrid_tsp_solver(cities, distances, time_budget=10.0):
    """
    Hybrid TSP solver with time budget.
    
    Strategy:
    1. If small (n < 12): exact (branch and bound)
    2. If medium (12 ≤ n < 20): Christofides + local search
    3. If large (n ≥ 20): Best heuristic
    """
    import time
    
    n = len(cities)
    start_time = time.time()
    
    if n < 12:
        # Exact solution
        return branch_and_bound_tsp(cities, distances, time_budget)
    
    elif n < 20:
        # Approximation + improvement
        tour, cost = tsp_approx_mst(distances, cities)
        
        remaining_time = time_budget - (time.time() - start_time)
        if remaining_time > 0:
            tour = local_search_tsp(tour, distances)
        
        return tour, tour_cost(tour, distances)
    
    else:
        # Best heuristic within time budget
        remaining_time = time_budget - (time.time() - start_time)
        return simulated_annealing_tsp(
            random.sample(cities, len(cities)) + [cities[0]],
            distances,
            iterations=int(remaining_time * 1000)
        )

def branch_and_bound_tsp(cities, distances, time_budget):
    """
    Exact TSP using branch and bound.
    Uses MST lower bound.
    """
    import time
    
    best_tour = None
    best_cost = float('inf')
    start_time = time.time()
    
    def bound(partial_tour, unvisited):
        """Lower bound on completion cost"""
        # Cost so far
        cost = sum(distances.get((partial_tour[i], partial_tour[i+1]), float('inf'))
                  for i in range(len(partial_tour) - 1))
        
        # Add MST of unvisited vertices
        if unvisited:
            # Simplified: use minimum edge to unvisited
            min_edge = min(
                distances.get((partial_tour[-1], v), float('inf'))
                for v in unvisited
            )
            cost += min_edge
        
        return cost
    
    def branch(partial_tour, unvisited):
        nonlocal best_tour, best_cost
        
        # Check time budget
        if time.time() - start_time > time_budget:
            return
        
        if not unvisited:
            # Complete tour
            total_cost = tour_cost(partial_tour + [partial_tour[0]], distances)
            if total_cost < best_cost:
                best_cost = total_cost
                best_tour = partial_tour
            return
        
        # Prune if bound exceeds best
        if bound(partial_tour, unvisited) >= best_cost:
            return
        
        # Branch on each unvisited city
        for city in sorted(unvisited, 
                          key=lambda c: distances.get((partial_tour[-1], c), float('inf'))):
            branch(partial_tour + [city], unvisited - {city})
    
    # Start from first city
    branch([cities[0]], set(cities) - {cities[0]})
    
    return best_tour, best_cost

# PRACTICAL SCHEDULING SOLVER

def schedule_jobs_with_constraints(jobs, constraints, time_budget=5.0):
    """
    Practical job scheduling with constraints.
    
    jobs: list of (duration, deadline, priority)
    constraints: list of (job1, job2) where job1 before job2
    
    Uses combination of:
    1. Greedy initial solution
    2. Local search
    3. Constraint propagation
    """
    import time
    
    start_time = time.time()
    
    # Initial greedy solution
    schedule = greedy_schedule(jobs, constraints)
    
    # Local search with time budget
    while time.time() - start_time < time_budget:
        # Try swapping two jobs
        improved = False
        
        for i in range(len(schedule)):
            for j in range(i + 1, len(schedule)):
                # Check if swap is valid
                new_schedule = schedule[:]
                new_schedule[i], new_schedule[j] = new_schedule[j], new_schedule[i]
                
                if is_valid_schedule(new_schedule, constraints):
                    if schedule_cost(new_schedule, jobs) < schedule_cost(schedule, jobs):
                        schedule = new_schedule
                        improved = True
                        break
            
            if improved:
                break
    
    return schedule

def greedy_schedule(jobs, constraints):
    """Greedy scheduling respecting constraints"""
    # Build dependency graph
    dependencies = {i: set() for i in range(len(jobs))}
    
    for j1, j2 in constraints:
        dependencies[j2].add(j1)
    
    # Topological sort with priority
    scheduled = []
    available = [i for i in range(len(jobs)) if not dependencies[i]]
    
    while available:
        # Choose job with highest priority/deadline urgency
        next_job = max(available, key=lambda j: jobs[j][2])  # priority
        scheduled.append(next_job)
        
        available.remove(next_job)
        
        # Update available
        for job in range(len(jobs)):
            if job not in scheduled and next_job in dependencies[job]:
                dependencies[job].remove(next_job)
                if not dependencies[job]:
                    available.append(job)
    
    return scheduled

def is_valid_schedule(schedule, constraints):
    """Check if schedule respects constraints"""
    position = {job: i for i, job in enumerate(schedule)}
    
    for j1, j2 in constraints:
        if position[j1] >= position[j2]:
            return False
    
    return True

def schedule_cost(schedule, jobs):
    """Calculate schedule cost (lateness, priority)"""
    cost = 0
    time = 0
    
    for job_id in schedule:
        duration, deadline, priority = jobs[job_id]
        completion = time + duration
        
        # Lateness penalty
        if completion > deadline:
            cost += (completion - deadline) * priority
        
        time = completion
    
    return cost

# ALGORITHM SELECTION FRAMEWORK

def select_algorithm(problem_type, instance_size, time_budget, quality_requirement):
    """
    Choose algorithm based on problem characteristics.
    
    Returns: algorithm name and expected performance
    """
    if problem_type == "TSP":
        if instance_size < 12:
            return "Exact (Branch & Bound)", "optimal"
        elif instance_size < 20 and time_budget > 1.0:
            return "Christofides + Local Search", "1.5-approx + improvement"
        elif quality_requirement == "high":
            return "Simulated Annealing", "good quality, no guarantee"
        else:
            return "Greedy + 2-opt", "fast, reasonable quality"
    
    elif problem_type == "SAT":
        if instance_size < 50:
            return "DPLL", "exact but may be slow"
        else:
            return "WalkSAT (local search)", "incomplete but fast"
    
    elif problem_type == "Vertex Cover":
        if quality_requirement == "exact":
            return "Branch & Bound", "optimal, exponential time"
        elif quality_requirement == "good":
            return "2-approximation", "guaranteed 2× optimal, fast"
        else:
            return "Greedy", "fast, no guarantee"
    
    return "Unknown", "unknown"
```

### Test Cases
- Real-world problem instances
- Compare exact vs. approximate
- Measure quality vs. time tradeoffs
- Test with time budgets

### Deliverables
- `sat_solver.{py,go,zig}` - DPLL implementation
- `hybrid_tsp_solver.{py,go,zig}` - Adaptive TSP solver
- `scheduling_system.{py,go,zig}` - Practical scheduler
- `algorithm_selector.{py,go,zig}` - Choose best approach
- `tradeoff_analyzer.{py,go,zig}` - Quality vs. time analysis
- `case_studies.md` - Real-world applications
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Can solve real problem instances
- Understanding of engineering tradeoffs
- Appropriate algorithm selection
- Practical solutions within constraints

---

## Module Resources

### CLRS References
- Chapter 34: NP-Completeness
- Chapter 35: Approximation Algorithms

### Additional Reading
- "Computers and Intractability" by Garey & Johnson (NP-completeness bible)
- "Approximation Algorithms" by Vazirani
- "The Design of Approximation Algorithms" by Williamson & Shmoys
- SAT solver research papers

### Key Takeaways

By the end of this module, you should:
1. Understand P, NP, and NP-completeness
2. Recognize NP-complete problems
3. Prove NP-completeness via reductions
4. Design approximation algorithms
5. Apply heuristics effectively
6. Make informed engineering decisions
7. Know when exact solutions are infeasible

### Completion

**Congratulations!** This completes the entire DSA curriculum. You now have:
- Complete understanding of data structures and algorithms
- Theoretical foundations (complexity theory)
- Practical problem-solving skills
- Knowledge of 13 comprehensive modules
- 75+ labs with detailed implementations

---

## Tips for Success

1. **Don't expect polynomial solutions** - Accept intractability
2. **Know your reductions** - Build from SAT/3-SAT
3. **Approximation is powerful** - Often good enough
4. **Combine techniques** - Hybrid approaches work best
5. **Measure real performance** - Theory meets practice

## Common Pitfalls

- **Trying to find polynomial algorithms for NP-complete problems**
- **Not verifying NP membership before proving NP-completeness**
- **Incorrect reduction directions** (A ≤_p B, not B ≤_p A)
- **Confusing approximation ratio with error**
- **Not testing heuristics thoroughly**
- **Ignoring time budgets in real applications**
