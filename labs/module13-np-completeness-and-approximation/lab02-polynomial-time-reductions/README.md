# Lab 13.2: Polynomial-Time Reductions

**Module**: Module 13 - NP-Completeness & Approximation Algorithms  
**Duration**: 6-8 hours  
**Difficulty**: Hard

---

Lab 13.2: Polynomial-Time Reductions

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

---

## Back to Module

[← Back to Module 13: NP-Completeness & Approximation Algorithms](../module_13/README.md)

## Navigation

- [Previous Lab](./lab_13_1.md) (if exists)
- [Next Lab](./lab_13_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 13, Lab 2 of 5*
