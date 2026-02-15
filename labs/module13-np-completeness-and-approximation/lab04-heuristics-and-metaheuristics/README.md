# Lab 13.4: Heuristics & Metaheuristics

**Module**: Module 13 - NP-Completeness & Approximation Algorithms  
**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

---

Lab 13.4: Heuristics & Metaheuristics

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

---

## Back to Module

[← Back to Module 13: NP-Completeness & Approximation Algorithms](../module_13/README.md)

## Navigation

- [Previous Lab](./lab_13_3.md) (if exists)
- [Next Lab](./lab_13_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 13, Lab 4 of 5*
