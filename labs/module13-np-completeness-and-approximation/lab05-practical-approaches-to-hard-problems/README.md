# Lab 13.5: Practical Approaches to Hard Problems

**Module**: Module 13 - NP-Completeness & Approximation Algorithms  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 13.5: Practical Approaches to Hard Problems

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

---

## Back to Module

[← Back to Module 13: NP-Completeness & Approximation Algorithms](../module_13/README.md)

## Navigation

- [Previous Lab](./lab_13_4.md) (if exists)
- [Next Lab](./lab_13_6.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 13, Lab 5 of 5*
