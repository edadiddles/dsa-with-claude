# Lab 7.2: Bellman-Ford Algorithm

**Module**: Module 7 - Graph Algorithms II - Shortest Paths  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 7.2: Bellman-Ford Algorithm

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

---

## Back to Module

[← Back to Module 7: Graph Algorithms II - Shortest Paths](../module_07/README.md)

## Navigation

- [Previous Lab](./lab_7_1.md) (if exists)
- [Next Lab](./lab_7_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 7, Lab 2 of 5*
