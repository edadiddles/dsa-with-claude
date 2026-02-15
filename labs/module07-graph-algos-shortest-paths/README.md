# Module 7: Graph Algorithms II - Shortest Paths

**Duration**: 2 weeks  
**Difficulty**: Medium to Hard

## Overview

Finding shortest paths in weighted graphs is one of the most practical problems in computer science. This module covers the essential shortest path algorithms - from Dijkstra's for non-negative weights to Bellman-Ford for general graphs, and Floyd-Warshall for all-pairs shortest paths.

**Why this matters**: Shortest path algorithms power GPS navigation, network routing protocols, game AI pathfinding, flight booking systems, and countless other real-world applications. Understanding when and how to use each algorithm is crucial for building efficient systems.

## Learning Objectives

- Implement Dijkstra's algorithm with different priority queue strategies
- Handle negative weights with Bellman-Ford
- Compute all-pairs shortest paths with Floyd-Warshall
- Exploit DAG structure for linear-time shortest paths
- Build a practical multi-criteria route planning system
- Understand when to use which shortest path algorithm

## Topics Covered

- Single-source shortest paths (Dijkstra's, Bellman-Ford)
- All-pairs shortest paths (Floyd-Warshall, Johnson's)
- Shortest paths in DAGs
- Negative cycle detection
- Path reconstruction
- A* search and heuristics
- Practical route planning with constraints

---

## Lab 7.1: Dijkstra's Algorithm

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Implement Dijkstra's shortest path algorithm
- Use different priority queue implementations
- Understand limitations (no negative weights)
- Implement optimizations (A*, bidirectional search)

### Requirements

1. **Implement Dijkstra's with**:
   - Binary heap priority queue
   - Array-based (for dense graphs)
   - Analysis of Fibonacci heap (theoretical)

2. **Path reconstruction**:
   - Store parent pointers
   - Reconstruct shortest path
   - Find all shortest paths (multiple paths with same length)

3. **Optimizations**:
   - Bidirectional search
   - A* search with heuristics
   - Early termination when target found

### Implementation Details

```python
import heapq

def dijkstra(graph, source):
    """
    Dijkstra's algorithm for single-source shortest paths.
    Returns distances and parent pointers for path reconstruction.
    """
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    dist[source] = 0
    
    # Min-heap: (distance, vertex)
    pq = [(0, source)]
    visited = set()
    
    while pq:
        d, u = heapq.heappop(pq)
        
        if u in visited:
            continue
        
        visited.add(u)
        
        # Relax all edges from u
        for v in graph.get_neighbors(u):
            weight = graph.get_weight(u, v)
            if dist[u] + weight < dist[v]:
                dist[v] = dist[u] + weight
                parent[v] = u
                heapq.heappush(pq, (dist[v], v))
    
    return dist, parent

def reconstruct_path(parent, source, target):
    """Reconstruct shortest path from parent pointers"""
    if parent[target] == -1 and target != source:
        return None  # No path exists
    
    path = []
    current = target
    
    while current != -1:
        path.append(current)
        if current == source:
            break
        current = parent[current]
    
    return list(reversed(path))

def dijkstra_with_path(graph, source, target):
    """Dijkstra's that returns path and distance to specific target"""
    dist, parent = dijkstra(graph, source)
    
    if dist[target] == float('inf'):
        return None, float('inf')
    
    path = reconstruct_path(parent, source, target)
    return path, dist[target]

def dijkstra_early_termination(graph, source, target):
    """Dijkstra's with early termination when target is reached"""
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    dist[source] = 0
    
    pq = [(0, source)]
    visited = set()
    
    while pq:
        d, u = heapq.heappop(pq)
        
        if u == target:
            # Target reached, reconstruct path and return
            path = reconstruct_path(parent, source, target)
            return path, dist[target]
        
        if u in visited:
            continue
        
        visited.add(u)
        
        for v in graph.get_neighbors(u):
            weight = graph.get_weight(u, v)
            if dist[u] + weight < dist[v]:
                dist[v] = dist[u] + weight
                parent[v] = u
                heapq.heappush(pq, (dist[v], v))
    
    return None, float('inf')  # Target not reachable

def dijkstra_array_based(graph, source):
    """
    Dijkstra's using array-based priority queue.
    Better for dense graphs: O(V²) instead of O((V+E) log V).
    """
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    visited = [False] * graph.num_vertices
    dist[source] = 0
    
    for _ in range(graph.num_vertices):
        # Find minimum distance vertex not yet visited
        u = -1
        min_dist = float('inf')
        for v in range(graph.num_vertices):
            if not visited[v] and dist[v] < min_dist:
                min_dist = dist[v]
                u = v
        
        if u == -1:
            break  # All reachable vertices processed
        
        visited[u] = True
        
        # Relax edges
        for v in graph.get_neighbors(u):
            weight = graph.get_weight(u, v)
            if dist[u] + weight < dist[v]:
                dist[v] = dist[u] + weight
                parent[v] = u
    
    return dist, parent

def bidirectional_dijkstra(graph, source, target):
    """
    Bidirectional Dijkstra's search.
    Search from both source and target simultaneously.
    """
    dist_forward = [float('inf')] * graph.num_vertices
    dist_backward = [float('inf')] * graph.num_vertices
    parent_forward = [-1] * graph.num_vertices
    parent_backward = [-1] * graph.num_vertices
    
    dist_forward[source] = 0
    dist_backward[target] = 0
    
    pq_forward = [(0, source)]
    pq_backward = [(0, target)]
    
    visited_forward = set()
    visited_backward = set()
    
    best_distance = float('inf')
    meeting_vertex = -1
    
    # Need reverse graph for backward search
    graph_reverse = transpose(graph)
    
    while pq_forward or pq_backward:
        # Expand from forward direction
        if pq_forward:
            d, u = heapq.heappop(pq_forward)
            
            if u not in visited_forward:
                visited_forward.add(u)
                
                # Check if we've met the backward search
                if u in visited_backward:
                    total_dist = dist_forward[u] + dist_backward[u]
                    if total_dist < best_distance:
                        best_distance = total_dist
                        meeting_vertex = u
                
                # Relax edges
                for v in graph.get_neighbors(u):
                    weight = graph.get_weight(u, v)
                    if dist_forward[u] + weight < dist_forward[v]:
                        dist_forward[v] = dist_forward[u] + weight
                        parent_forward[v] = u
                        heapq.heappush(pq_forward, (dist_forward[v], v))
        
        # Expand from backward direction
        if pq_backward:
            d, u = heapq.heappop(pq_backward)
            
            if u not in visited_backward:
                visited_backward.add(u)
                
                # Check if we've met the forward search
                if u in visited_forward:
                    total_dist = dist_forward[u] + dist_backward[u]
                    if total_dist < best_distance:
                        best_distance = total_dist
                        meeting_vertex = u
                
                # Relax edges in reverse graph
                for v in graph_reverse.get_neighbors(u):
                    weight = graph_reverse.get_weight(u, v)
                    if dist_backward[u] + weight < dist_backward[v]:
                        dist_backward[v] = dist_backward[u] + weight
                        parent_backward[v] = u
                        heapq.heappush(pq_backward, (dist_backward[v], v))
        
        # Early termination if we can't improve
        if pq_forward and pq_backward:
            if min(pq_forward[0][0], pq_backward[0][0]) >= best_distance:
                break
    
    if meeting_vertex == -1:
        return None, float('inf')
    
    # Reconstruct path through meeting vertex
    path_to_meeting = reconstruct_path(parent_forward, source, meeting_vertex)
    path_from_meeting = reconstruct_path(parent_backward, target, meeting_vertex)
    
    if path_from_meeting:
        path_from_meeting = list(reversed(path_from_meeting[:-1]))  # Exclude meeting vertex
    
    full_path = path_to_meeting + path_from_meeting
    return full_path, best_distance

def a_star(graph, source, target, heuristic):
    """
    A* search with admissible heuristic.
    heuristic(v, target) should estimate distance from v to target.
    Must be admissible: h(v) <= actual distance.
    """
    g_score = [float('inf')] * graph.num_vertices  # Actual distance from source
    f_score = [float('inf')] * graph.num_vertices  # g + heuristic
    parent = [-1] * graph.num_vertices
    
    g_score[source] = 0
    f_score[source] = heuristic(source, target)
    
    pq = [(f_score[source], source)]
    visited = set()
    
    while pq:
        f, u = heapq.heappop(pq)
        
        if u == target:
            path = reconstruct_path(parent, source, target)
            return path, g_score[target]
        
        if u in visited:
            continue
        
        visited.add(u)
        
        for v in graph.get_neighbors(u):
            if v in visited:
                continue
            
            weight = graph.get_weight(u, v)
            tentative_g = g_score[u] + weight
            
            if tentative_g < g_score[v]:
                parent[v] = u
                g_score[v] = tentative_g
                f_score[v] = g_score[v] + heuristic(v, target)
                heapq.heappush(pq, (f_score[v], v))
    
    return None, float('inf')  # Target not reachable

def euclidean_heuristic(graph, v, target, coordinates):
    """
    Euclidean distance heuristic for A*.
    coordinates is a dict mapping vertex -> (x, y).
    """
    x1, y1 = coordinates[v]
    x2, y2 = coordinates[target]
    return ((x2 - x1)**2 + (y2 - y1)**2)**0.5

def manhattan_heuristic(graph, v, target, coordinates):
    """Manhattan distance heuristic for grid graphs"""
    x1, y1 = coordinates[v]
    x2, y2 = coordinates[target]
    return abs(x2 - x1) + abs(y2 - y1)

def find_all_shortest_paths(graph, source, target):
    """Find all shortest paths from source to target"""
    dist, _ = dijkstra(graph, source)
    shortest_dist = dist[target]
    
    if shortest_dist == float('inf'):
        return []
    
    all_paths = []
    
    def dfs(v, path, current_dist):
        if v == target and current_dist == shortest_dist:
            all_paths.append(path[:])
            return
        
        if current_dist > shortest_dist:
            return
        
        for neighbor in graph.get_neighbors(v):
            weight = graph.get_weight(v, neighbor)
            if current_dist + weight <= shortest_dist:
                path.append(neighbor)
                dfs(neighbor, path, current_dist + weight)
                path.pop()
    
    dfs(source, [source], 0)
    return all_paths
```

### Test Cases
- Graph with unique shortest paths
- Graph with multiple shortest paths of same length
- Dense vs. sparse graphs
- Grid graphs (for A*)
- Verify against brute-force for small graphs
- Large random graphs

### Deliverables
- `dijkstra.{py,go,zig}` - Multiple implementations
- `path_reconstruction.{py,go,zig}` - Path recovery utilities
- `bidirectional_dijkstra.{py,go,zig}` - Bidirectional search
- `a_star.{py,go,zig}` - A* implementation
- `heuristics.{py,go,zig}` - Various heuristic functions
- `performance_analysis.md` - Compare implementations
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Correct shortest paths found
- Performance matches theoretical complexity
- Understanding of PQ choice impact
- A* with admissible heuristic never finds suboptimal path
- Can explain when early termination is safe

---

## Lab 7.2: Bellman-Ford Algorithm

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement Bellman-Ford algorithm
- Detect negative weight cycles
- Compare with Dijkstra's
- Understand when negative weights are acceptable

### Requirements

1. **Implement Bellman-Ford**:
   - Relax all edges V-1 times
   - Detect negative cycles
   - Path reconstruction

2. **Optimizations**:
   - Early termination if no updates in a round
   - Queue-based (SPFA variant)

3. **Applications**:
   - Currency arbitrage detection
   - Network routing with costs/penalties

### Implementation Details

```python
def bellman_ford(graph, source):
    """
    Bellman-Ford algorithm for single-source shortest paths.
    Works with negative weights. Returns None if negative cycle detected.
    """
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    dist[source] = 0
    
    # Relax all edges V-1 times
    for i in range(graph.num_vertices - 1):
        updated = False
        
        for u in range(graph.num_vertices):
            if dist[u] == float('inf'):
                continue
            
            for v in graph.get_neighbors(u):
                weight = graph.get_weight(u, v)
                if dist[u] + weight < dist[v]:
                    dist[v] = dist[u] + weight
                    parent[v] = u
                    updated = True
        
        if not updated:
            break  # Early termination
    
    # Check for negative weight cycles
    has_negative_cycle = False
    for u in range(graph.num_vertices):
        if dist[u] == float('inf'):
            continue
        
        for v in graph.get_neighbors(u):
            weight = graph.get_weight(u, v)
            if dist[u] + weight < dist[v]:
                has_negative_cycle = True
                break
        
        if has_negative_cycle:
            break
    
    if has_negative_cycle:
        return None, None, True
    
    return dist, parent, False

def find_negative_cycle(graph):
    """
    Find a negative weight cycle if one exists.
    Returns the cycle as a list of vertices, or None.
    """
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    
    # Use vertex 0 as source (arbitrary choice)
    dist[0] = 0
    
    # Relax all edges V-1 times
    for _ in range(graph.num_vertices - 1):
        for u in range(graph.num_vertices):
            if dist[u] == float('inf'):
                continue
            
            for v in graph.get_neighbors(u):
                weight = graph.get_weight(u, v)
                if dist[u] + weight < dist[v]:
                    dist[v] = dist[u] + weight
                    parent[v] = u
    
    # Find a vertex in a negative cycle
    cycle_vertex = -1
    for u in range(graph.num_vertices):
        if dist[u] == float('inf'):
            continue
        
        for v in graph.get_neighbors(u):
            weight = graph.get_weight(u, v)
            if dist[u] + weight < dist[v]:
                cycle_vertex = v
                break
        
        if cycle_vertex != -1:
            break
    
    if cycle_vertex == -1:
        return None  # No negative cycle
    
    # Walk back V times to ensure we're in the cycle
    for _ in range(graph.num_vertices):
        if parent[cycle_vertex] != -1:
            cycle_vertex = parent[cycle_vertex]
    
    # Reconstruct the cycle
    cycle = []
    current = cycle_vertex
    while True:
        cycle.append(current)
        current = parent[current]
        if current == cycle_vertex:
            cycle.append(current)
            break
    
    return list(reversed(cycle))

def spfa(graph, source):
    """
    Shortest Path Faster Algorithm (queue-based Bellman-Ford).
    Often faster in practice, but worst case is still O(VE).
    """
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    in_queue = [False] * graph.num_vertices
    count = [0] * graph.num_vertices  # Number of times vertex added to queue
    
    dist[source] = 0
    queue = deque([source])
    in_queue[source] = True
    count[source] = 1
    
    while queue:
        u = queue.popleft()
        in_queue[u] = False
        
        for v in graph.get_neighbors(u):
            weight = graph.get_weight(u, v)
            if dist[u] + weight < dist[v]:
                dist[v] = dist[u] + weight
                parent[v] = u
                
                if not in_queue[v]:
                    queue.append(v)
                    in_queue[v] = True
                    count[v] += 1
                    
                    # Negative cycle detection
                    if count[v] > graph.num_vertices:
                        return None, None, True  # Negative cycle
    
    return dist, parent, False

def currency_arbitrage(exchange_rates):
    """
    Detect arbitrage opportunities in currency exchange.
    exchange_rates[i][j] is the rate to convert currency i to currency j.
    
    Uses negative logarithm trick: log(a*b*c) = log(a) + log(b) + log(c).
    Arbitrage exists if there's a cycle where product > 1, i.e., sum of logs > 0,
    i.e., sum of negative logs < 0 (negative cycle).
    """
    import math
    
    n = len(exchange_rates)
    graph = AdjacencyList(n, directed=True)
    
    # Build graph with -log(rate) as weights
    for i in range(n):
        for j in range(n):
            if i != j and exchange_rates[i][j] > 0:
                weight = -math.log(exchange_rates[i][j])
                graph.add_edge(i, j, weight)
    
    # Find negative cycle
    cycle = find_negative_cycle(graph)
    
    if cycle:
        return True, cycle  # Arbitrage exists
    return False, None

def bellman_ford_with_path_tracking(graph, source):
    """
    Bellman-Ford with detailed path tracking.
    Useful for debugging and understanding algorithm.
    """
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    dist[source] = 0
    
    iterations = []
    
    for i in range(graph.num_vertices - 1):
        updated_edges = []
        
        for u in range(graph.num_vertices):
            if dist[u] == float('inf'):
                continue
            
            for v in graph.get_neighbors(u):
                weight = graph.get_weight(u, v)
                if dist[u] + weight < dist[v]:
                    old_dist = dist[v]
                    dist[v] = dist[u] + weight
                    parent[v] = u
                    updated_edges.append((u, v, old_dist, dist[v]))
        
        iterations.append({
            'round': i,
            'updated_edges': updated_edges,
            'distances': dist[:]
        })
        
        if not updated_edges:
            break
    
    return dist, parent, iterations
```

### Test Cases
- Graphs with no negative weights (compare with Dijkstra)
- Graphs with negative weights but no cycles
- Graphs with negative weight cycles
- Currency exchange graphs (arbitrage)
- Large graphs (performance testing)

### Deliverables
- `bellman_ford.{py,go,zig}` - Standard implementation
- `spfa.{py,go,zig}` - Queue-based variant
- `negative_cycle_detection.{py,go,zig}` - Find and extract cycles
- `currency_arbitrage.{py,go,zig}` - Application example
- `comparison.md` - Bellman-Ford vs. Dijkstra
- `benchmarks.{py,go,zig}` - Performance tests

### Success Criteria
- Correct distances even with negative weights
- Detects all negative cycles
- Understanding of when Bellman-Ford is necessary
- SPFA optimization shows practical speedup

---

## Lab 7.3: Floyd-Warshall Algorithm

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Implement all-pairs shortest paths algorithm
- Use dynamic programming approach
- Detect negative cycles globally
- Compute graph properties (diameter, center)

### Requirements

1. **Implement Floyd-Warshall**:
   - DP table for distances
   - Path reconstruction
   - Transitive closure variant

2. **Graph analysis**:
   - Compute graph diameter
   - Find graph center
   - All-pairs distances

3. **Applications**:
   - Network reachability
   - Transitive closure
   - Finding shortest paths between all pairs

### Implementation Details

```python
def floyd_warshall(graph):
    """
    Floyd-Warshall algorithm for all-pairs shortest paths.
    Returns distance matrix and path reconstruction matrix.
    """
    n = graph.num_vertices
    
    # Initialize distance matrix
    dist = [[float('inf')] * n for _ in range(n)]
    next_vertex = [[-1] * n for _ in range(n)]
    
    # Distance from vertex to itself is 0
    for i in range(n):
        dist[i][i] = 0
    
    # Set initial distances from edges
    for u in range(n):
        for v in graph.get_neighbors(u):
            weight = graph.get_weight(u, v)
            dist[u][v] = weight
            next_vertex[u][v] = v
    
    # Floyd-Warshall DP
    for k in range(n):
        for i in range(n):
            for j in range(n):
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
                    next_vertex[i][j] = next_vertex[i][k]
    
    # Check for negative cycles
    negative_cycle = False
    for i in range(n):
        if dist[i][i] < 0:
            negative_cycle = True
            break
    
    return dist, next_vertex, negative_cycle

def reconstruct_path_fw(next_vertex, u, v):
    """Reconstruct path from Floyd-Warshall next_vertex matrix"""
    if next_vertex[u][v] == -1:
        return None  # No path
    
    path = [u]
    while u != v:
        u = next_vertex[u][v]
        path.append(u)
    
    return path

def transitive_closure(graph):
    """
    Compute transitive closure using Floyd-Warshall-like approach.
    Returns boolean matrix where result[i][j] = True if path from i to j exists.
    """
    n = graph.num_vertices
    reach = [[False] * n for _ in range(n)]
    
    # Initialize
    for i in range(n):
        reach[i][i] = True
    
    for u in range(n):
        for v in graph.get_neighbors(u):
            reach[u][v] = True
    
    # Transitive closure
    for k in range(n):
        for i in range(n):
            for j in range(n):
                reach[i][j] = reach[i][j] or (reach[i][k] and reach[k][j])
    
    return reach

def graph_diameter(dist):
    """
    Compute graph diameter (maximum shortest path distance).
    Assumes dist is output from Floyd-Warshall.
    """
    diameter = 0
    for i in range(len(dist)):
        for j in range(len(dist[i])):
            if dist[i][j] != float('inf') and dist[i][j] > diameter:
                diameter = dist[i][j]
    return diameter

def graph_radius_and_center(dist):
    """
    Compute graph radius and center vertices.
    Center vertices are those with minimum eccentricity.
    """
    n = len(dist)
    eccentricity = []
    
    for i in range(n):
        # Eccentricity is max distance from this vertex
        max_dist = 0
        for j in range(n):
            if dist[i][j] != float('inf'):
                max_dist = max(max_dist, dist[i][j])
        eccentricity.append(max_dist if max_dist != 0 else float('inf'))
    
    # Radius is minimum eccentricity
    radius = min(eccentricity)
    
    # Center vertices have eccentricity equal to radius
    center_vertices = [i for i, ecc in enumerate(eccentricity) if ecc == radius]
    
    return radius, center_vertices

def find_all_pairs_bottleneck(graph):
    """
    Find bottleneck (maximum edge weight) on shortest path between all pairs.
    Similar to Floyd-Warshall but tracking maximum edge weight.
    """
    n = graph.num_vertices
    bottleneck = [[float('inf')] * n for _ in range(n)]
    
    # Initialize
    for i in range(n):
        bottleneck[i][i] = 0
    
    for u in range(n):
        for v in graph.get_neighbors(u):
            bottleneck[u][v] = graph.get_weight(u, v)
    
    # Modified Floyd-Warshall
    for k in range(n):
        for i in range(n):
            for j in range(n):
                # Bottleneck on path through k
                path_through_k = max(bottleneck[i][k], bottleneck[k][j])
                bottleneck[i][j] = min(bottleneck[i][j], path_through_k)
    
    return bottleneck

def johnson_algorithm(graph):
    """
    Johnson's algorithm for all-pairs shortest paths.
    More efficient than Floyd-Warshall for sparse graphs: O(V²log V + VE).
    Uses Bellman-Ford once, then Dijkstra's V times with reweighting.
    """
    # Add new vertex s connected to all vertices with weight 0
    n = graph.num_vertices
    augmented = AdjacencyList(n + 1, directed=True)
    
    # Copy original edges
    for u in range(n):
        for v in graph.get_neighbors(u):
            augmented.add_edge(u, v, graph.get_weight(u, v))
    
    # Add edges from new vertex s (= n) to all vertices
    for v in range(n):
        augmented.add_edge(n, v, 0)
    
    # Run Bellman-Ford from s to find h(v) values
    h_values, _, has_neg_cycle = bellman_ford(augmented, n)
    
    if has_neg_cycle:
        return None, True  # Negative cycle exists
    
    # Reweight edges: w'(u,v) = w(u,v) + h(u) - h(v)
    reweighted = AdjacencyList(n, directed=True)
    for u in range(n):
        for v in graph.get_neighbors(u):
            original_weight = graph.get_weight(u, v)
            new_weight = original_weight + h_values[u] - h_values[v]
            reweighted.add_edge(u, v, new_weight)
    
    # Run Dijkstra from each vertex on reweighted graph
    dist = [[float('inf')] * n for _ in range(n)]
    
    for u in range(n):
        dist_from_u, _ = dijkstra(reweighted, u)
        
        # Convert back to original weights
        for v in range(n):
            if dist_from_u[v] != float('inf'):
                dist[u][v] = dist_from_u[v] - h_values[u] + h_values[v]
    
    return dist, False
```

### Test Cases
- Small graphs (verify by hand)
- Dense graphs
- Graphs with negative weights
- Disconnected graphs
- Complete graphs
- Compare Johnson's vs. Floyd-Warshall on sparse graphs

### Deliverables
- `floyd_warshall.{py,go,zig}` - Full implementation
- `transitive_closure.{py,go,zig}` - Reachability matrix
- `graph_metrics.{py,go,zig}` - Diameter, radius, center
- `path_reconstruction.{py,go,zig}` - Reconstruct any shortest path
- `johnson.{py,go,zig}` - Johnson's algorithm
- `complexity_analysis.md` - When to use Floyd-Warshall vs. Dijkstra/Johnson
- `benchmarks.{py,go,zig}` - Performance comparison

### Success Criteria
- Correct all-pairs distances
- O(V³) time complexity
- Understanding of space-time tradeoffs
- Can explain when Johnson's is better

---

## Lab 7.4: Shortest Paths in DAGs

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Exploit DAG structure for linear-time shortest paths
- Implement critical path method
- Apply to scheduling problems

### Requirements

1. **Implement DAG shortest paths**:
   - Topological sort
   - Single-pass relaxation
   - Longest paths (negate weights)

2. **Applications**:
   - Critical Path Method (CPM) for project scheduling
   - PERT charts
   - Longest increasing subsequence (using DAG)

### Implementation Details

```python
from topological_sort import topological_sort_dfs

def dag_shortest_paths(graph, source):
    """
    Shortest paths in DAG - O(V + E) time.
    Works for any edge weights (including negative).
    """
    # Get topological ordering
    topo_order = topological_sort_dfs(graph)
    
    if topo_order is None:
        return None, None  # Not a DAG
    
    dist = [float('inf')] * graph.num_vertices
    parent = [-1] * graph.num_vertices
    dist[source] = 0
    
    # Process vertices in topological order
    for u in topo_order:
        if dist[u] != float('inf'):
            for v in graph.get_neighbors(u):
                weight = graph.get_weight(u, v)
                if dist[u] + weight < dist[v]:
                    dist[v] = dist[u] + weight
                    parent[v] = u
    
    return dist, parent

def dag_longest_paths(graph, source):
    """
    Longest paths in DAG.
    Negate weights and find shortest paths.
    """
    # Create graph with negated weights
    n = graph.num_vertices
    negated = AdjacencyList(n, directed=True)
    
    for u in range(n):
        for v in graph.get_neighbors(u):
            negated.add_edge(u, v, -graph.get_weight(u, v))
    
    dist, parent = dag_shortest_paths(negated, source)
    
    if dist is None:
        return None, None
    
    # Negate distances back
    dist = [-d if d != float('inf') else float('-inf') for d in dist]
    
    return dist, parent

class CriticalPathMethod:
    """
    Critical Path Method for project scheduling.
    Tasks are vertices, dependencies are edges, weights are durations.
    """
    def __init__(self):
        self.tasks = []
        self.graph = None
        self.task_to_id = {}
    
    def add_task(self, task_name, duration, dependencies=None):
        """Add a task with duration and dependencies"""
        if dependencies is None:
            dependencies = []
        self.tasks.append({
            'name': task_name,
            'duration': duration,
            'dependencies': dependencies
        })
    
    def build_graph(self):
        """Build DAG from tasks"""
        # Create mapping
        self.task_to_id = {task['name']: i for i, task in enumerate(self.tasks)}
        n = len(self.tasks)
        
        self.graph = AdjacencyList(n, directed=True)
        
        # Add edges from dependencies to tasks
        for i, task in enumerate(self.tasks):
            for dep in task['dependencies']:
                dep_id = self.task_to_id[dep]
                # Edge weight is the duration of the dependency
                self.graph.add_edge(dep_id, i, self.tasks[dep_id]['duration'])
    
    def find_critical_path(self):
        """Find critical path and project duration"""
        self.build_graph()
        
        # Add virtual start and end nodes
        n = self.graph.num_vertices
        start = n
        end = n + 1
        
        augmented = AdjacencyList(n + 2, directed=True)
        
        # Copy edges
        for u in range(n):
            for v in self.graph.get_neighbors(u):
                augmented.add_edge(u, v, self.graph.get_weight(u, v))
        
        # Connect start to all tasks with no dependencies
        for i, task in enumerate(self.tasks):
            if not task['dependencies']:
                augmented.add_edge(start, i, 0)
        
        # Connect all tasks to end
        for i in range(n):
            augmented.add_edge(i, end, self.tasks[i]['duration'])
        
        # Find longest path from start to end
        longest_dist, parent = dag_longest_paths(augmented, start)
        
        if longest_dist is None:
            return None, "Cycle detected in dependencies!"
        
        project_duration = longest_dist[end]
        
        # Reconstruct critical path
        critical_path = []
        current = end
        while parent[current] != -1:
            if current < n:  # Exclude virtual nodes in output
                critical_path.append(self.tasks[current]['name'])
            current = parent[current]
        
        critical_path = list(reversed(critical_path))
        
        # Calculate earliest and latest start times
        earliest_start = [longest_dist[i] for i in range(n)]
        
        # Latest start times (work backwards)
        latest_finish = [project_duration - self.tasks[i]['duration'] for i in range(n)]
        
        # Calculate slack
        slack = {}
        for i, task in enumerate(self.tasks):
            slack[task['name']] = latest_finish[i] - earliest_start[i]
        
        return {
            'duration': project_duration,
            'critical_path': critical_path,
            'earliest_start': {self.tasks[i]['name']: earliest_start[i] for i in range(n)},
            'slack': slack
        }

def longest_increasing_subsequence_dag(arr):
    """
    Find longest increasing subsequence using DAG approach.
    Build DAG where edge (i,j) exists if arr[i] < arr[j] and i < j.
    """
    n = len(arr)
    graph = AdjacencyList(n, directed=True)
    
    # Build DAG
    for i in range(n):
        for j in range(i + 1, n):
            if arr[i] < arr[j]:
                graph.add_edge(i, j, 1)  # Weight 1 for path length
    
    # Find longest path starting from any vertex
    max_length = 0
    best_path = []
    
    for start in range(n):
        dist, parent = dag_longest_paths(graph, start)
        
        for end in range(n):
            if dist[end] > max_length:
                max_length = dist[end]
                # Reconstruct path
                path = []
                current = end
                while current != -1:
                    path.append(current)
                    current = parent[current]
                best_path = list(reversed(path))
    
    # Convert indices to actual values
    lis = [arr[i] for i in best_path]
    
    return lis, len(lis)
```

### Test Cases
- Simple DAGs with unique longest path
- DAGs with multiple longest paths
- Project scheduling examples with realistic dependencies
- LIS test cases
- Compare DAG algorithm with Bellman-Ford (should be much faster)

### Deliverables
- `dag_shortest_paths.{py,go,zig}` - Linear-time algorithm
- `critical_path.{py,go,zig}` - CPM implementation
- `lis_dag.{py,go,zig}` - LIS using DAG
- `project_scheduler.{py,go,zig}` - Practical scheduling tool
- `performance_comparison.md` - DAG vs. general shortest paths
- `examples/` - Real project scheduling examples

### Success Criteria
- Linear-time complexity for DAGs
- Correct critical path identification
- Understanding of DAG optimizations
- Can apply to real scheduling problems

---

## Lab 7.5: Route Planning Application

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Build practical route planning system
- Implement realistic constraints
- Optimize for real-world usage

### Requirements

1. **Features**:
   - Multi-criteria routing (shortest, fastest, avoid tolls)
   - Turn restrictions
   - Time-dependent edge weights (traffic patterns)
   - Intermediate waypoints

2. **Optimizations**:
   - A* with good heuristics
   - Bidirectional search
   - Caching common routes

3. **Realistic constraints**:
   - One-way streets
   - Road closures
   - Vehicle restrictions

### Implementation Details

```python
class RouteOptions:
    def __init__(self):
        self.optimize_for = "time"  # or "distance", "fuel"
        self.avoid_tolls = False
        self.avoid_highways = False
        self.avoid_ferries = False
        self.waypoints = []
        self.departure_time = None  # For time-dependent routing

class RoadNetwork:
    """Enhanced graph for road networks"""
    def __init__(self, num_vertices):
        self.graph = AdjacencyList(num_vertices, directed=True)
        self.coordinates = {}  # vertex -> (lat, lon)
        self.road_types = {}  # (u,v) -> type (highway, toll, etc.)
        self.speed_limits = {}  # (u,v) -> speed limit
        self.turn_restrictions = set()  # set of (u,v,w) where turn from u->v->w is forbidden
    
    def add_road(self, u, v, distance, road_type="street", speed_limit=30):
        """Add a road segment"""
        self.graph.add_edge(u, v, distance)
        self.road_types[(u, v)] = road_type
        self.speed_limits[(u, v)] = speed_limit
    
    def set_coordinates(self, vertex, lat, lon):
        """Set GPS coordinates for vertex"""
        self.coordinates[vertex] = (lat, lon)
    
    def add_turn_restriction(self, u, v, w):
        """Forbid turn from u->v->w"""
        self.turn_restrictions.add((u, v, w))

def multi_criteria_route(network, source, target, options):
    """
    Find route optimizing for different criteria.
    """
    def get_edge_cost(u, v):
        distance = network.graph.get_weight(u, v)
        road_type = network.road_types.get((u, v), "street")
        speed_limit = network.speed_limits.get((u, v), 30)
        
        # Base cost
        if options.optimize_for == "distance":
            cost = distance
        elif options.optimize_for == "time":
            cost = distance / speed_limit  # Time = distance / speed
        elif options.optimize_for == "fuel":
            # Simplified fuel model
            cost = distance * (1.2 if road_type == "highway" else 1.0)
        else:
            cost = distance
        
        # Apply penalties
        if options.avoid_tolls and road_type == "toll":
            cost *= 10  # Heavy penalty
        
        if options.avoid_highways and road_type == "highway":
            cost *= 5
        
        if options.avoid_ferries and road_type == "ferry":
            cost *= 20
        
        return cost
    
    # Create custom graph with modified costs
    custom_graph = AdjacencyList(network.graph.num_vertices, directed=True)
    for u in range(network.graph.num_vertices):
        for v in network.graph.get_neighbors(u):
            cost = get_edge_cost(u, v)
            custom_graph.add_edge(u, v, cost)
    
    # Use A* with haversine heuristic
    def heuristic(v, target):
        if v not in network.coordinates or target not in network.coordinates:
            return 0
        return haversine_distance(
            network.coordinates[v],
            network.coordinates[target]
        )
    
    path, cost = a_star(custom_graph, source, target, heuristic)
    
    # Check turn restrictions
    if path and len(path) >= 3:
        for i in range(len(path) - 2):
            if (path[i], path[i+1], path[i+2]) in network.turn_restrictions:
                # This path has a forbidden turn, need to find alternative
                # In practice, we'd remove this turn and re-route
                pass
    
    return path, cost

def route_with_waypoints(network, start, waypoints, end, options):
    """
    Find route visiting all waypoints in order.
    """
    full_route = []
    total_cost = 0
    
    current = start
    for waypoint in waypoints:
        segment, cost = multi_criteria_route(network, current, waypoint, options)
        
        if segment is None:
            return None, float('inf')
        
        # Append segment (excluding last vertex to avoid duplicates)
        if full_route:
            full_route.extend(segment[1:])
        else:
            full_route.extend(segment)
        
        total_cost += cost
        current = waypoint
    
    # Final segment to end
    segment, cost = multi_criteria_route(network, current, end, options)
    
    if segment is None:
        return None, float('inf')
    
    full_route.extend(segment[1:])
    total_cost += cost
    
    return full_route, total_cost

def time_dependent_routing(network, source, target, departure_time, traffic_patterns):
    """
    Route planning with time-dependent edge weights (traffic).
    traffic_patterns[(u,v)] is a function: time -> travel_time
    """
    def get_travel_time(u, v, current_time):
        if (u, v) in traffic_patterns:
            return traffic_patterns[(u, v)](current_time)
        
        # Default: use speed limit
        distance = network.graph.get_weight(u, v)
        speed = network.speed_limits.get((u, v), 30)
        return distance / speed
    
    # Modified Dijkstra with time tracking
    dist = [float('inf')] * network.graph.num_vertices
    arrival_time = [float('inf')] * network.graph.num_vertices
    parent = [-1] * network.graph.num_vertices
    
    dist[source] = 0
    arrival_time[source] = departure_time
    
    pq = [(0, source, departure_time)]
    visited = set()
    
    while pq:
        d, u, time = heapq.heappop(pq)
        
        if u == target:
            path = reconstruct_path(parent, source, target)
            return path, dist[target], arrival_time[target]
        
        if u in visited:
            continue
        
        visited.add(u)
        
        for v in network.graph.get_neighbors(u):
            travel_time = get_travel_time(u, v, time)
            new_arrival = time + travel_time
            
            if dist[u] + travel_time < dist[v]:
                dist[v] = dist[u] + travel_time
                arrival_time[v] = new_arrival
                parent[v] = u
                heapq.heappush(pq, (dist[v], v, new_arrival))
    
    return None, float('inf'), None

def haversine_distance(coord1, coord2):
    """
    Calculate great-circle distance between two points on Earth.
    coord1, coord2 are (lat, lon) in degrees.
    Returns distance in kilometers.
    """
    import math
    
    lat1, lon1 = coord1
    lat2, lon2 = coord2
    
    # Convert to radians
    lat1, lon1 = math.radians(lat1), math.radians(lon1)
    lat2, lon2 = math.radians(lat2), math.radians(lon2)
    
    # Haversine formula
    dlat = lat2 - lat1
    dlon = lon2 - lon1
    a = math.sin(dlat/2)**2 + math.cos(lat1) * math.cos(lat2) * math.sin(dlon/2)**2
    c = 2 * math.asin(math.sqrt(a))
    
    # Earth radius in kilometers
    r = 6371
    
    return c * r

class RouteCache:
    """Cache common routes to avoid recomputation"""
    def __init__(self, max_size=1000):
        self.cache = {}
        self.max_size = max_size
    
    def get(self, source, target, options_hash):
        key = (source, target, options_hash)
        return self.cache.get(key)
    
    def put(self, source, target, options_hash, route, cost):
        if len(self.cache) >= self.max_size:
            # Simple LRU: remove oldest
            self.cache.pop(next(iter(self.cache)))
        
        key = (source, target, options_hash)
        self.cache[key] = (route, cost)
    
    def hash_options(self, options):
        """Create hash of options for cache key"""
        return (
            options.optimize_for,
            options.avoid_tolls,
            options.avoid_highways,
            options.avoid_ferries
        )
```

### Test Cases
- Simple point-to-point routing
- Multi-waypoint routes
- Routes with turn restrictions
- Time-dependent routing with traffic
- Compare optimizations (distance vs. time)

### Deliverables
- `route_planner.{py,go,zig}` - Full routing system
- `multi_criteria.{py,go,zig}` - Support multiple optimization criteria
- `waypoint_routing.{py,go,zig}` - Routes with intermediate stops
- `time_dependent.{py,go,zig}` - Traffic-aware routing
- `road_network_loader.{py,go,zig}` - Load real road data
- `route_visualizer.{py,go,zig}` - Visualize routes on map
- `web_interface/` - Simple web UI for route planning (optional)

### Success Criteria
- Handles realistic constraints
- Fast enough for interactive use (< 1 second for typical queries)
- Correct routes for complex queries
- Understanding of practical routing challenges

---

## Module Resources

### CLRS References
- Chapter 24: Single-Source Shortest Paths (Bellman-Ford, Dijkstra)
- Chapter 25: All-Pairs Shortest Paths (Floyd-Warshall, Johnson)

### Additional Reading
- "Algorithm Design" by Kleinberg & Tardos (Network Flow chapter)
- Route planning papers (Google Maps, Waze)
- A* and heuristics research
- Real-world routing challenges

### Key Takeaways

By the end of this module, you should:
1. Know which shortest path algorithm to use when
2. Understand the impact of negative weights
3. Be able to implement efficient shortest path queries
4. Recognize when to use A* and design good heuristics
5. Handle realistic routing constraints
6. Build practical route planning systems

### Next Module

**Module 8: Dynamic Programming** - You'll master the DP paradigm through extensive practice, from classic problems to advanced patterns.

---

## Tips for Success

1. **Draw graphs** - Visualize problems on paper
2. **Trace algorithms** - Step through by hand for small examples
3. **Test with known answers** - Verify against small examples
4. **Use visualization** - See how algorithms explore the graph
5. **Understand relaxation** - Core concept in all shortest path algorithms

## Common Pitfalls

- **Using Dijkstra with negative weights** - Incorrect results
- **Forgetting to initialize distances** - Should be infinity
- **Not checking for negative cycles** - Bellman-Ford requirement
- **Incorrect path reconstruction** - Carefully track parent pointers
- **Inefficient priority queue** - Major performance impact
- **Bad A* heuristics** - Must be admissible (never overestimate)
- **Not handling unreachable vertices** - Check for infinity distances
