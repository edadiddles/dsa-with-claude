# Lab 13.1: Understanding NP-Completeness

**Module**: Module 13 - NP-Completeness & Approximation Algorithms  
**Duration**: 5-6 hours  
**Difficulty**: Hard

---

Lab 13.1: Understanding NP-Completeness

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

---

## Back to Module

[← Back to Module 13: NP-Completeness & Approximation Algorithms](../module_13/README.md)

## Navigation

- [Previous Lab](./lab_13_0.md) (if exists)
- [Next Lab](./lab_13_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 13, Lab 1 of 5*
