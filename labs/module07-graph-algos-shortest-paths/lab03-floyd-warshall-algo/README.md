# Lab 7.3: Floyd-Warshall Algorithm

**Module**: Module 7 - Graph Algorithms II - Shortest Paths  
**Duration**: 4-5 hours  
**Difficulty**: Medium

---

Lab 7.3: Floyd-Warshall Algorithm

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

---

## Back to Module

[← Back to Module 7: Graph Algorithms II - Shortest Paths](../module_07/README.md)

## Navigation

- [Previous Lab](./lab_7_2.md) (if exists)
- [Next Lab](./lab_7_4.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 7, Lab 3 of 5*
