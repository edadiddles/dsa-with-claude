# Lab 7.1: Dijkstra's Algorithm

**Module**: Module 7 - Graph Algorithms II - Shortest Paths  
**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

---

Lab 7.1: Dijkstra's Algorithm

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

---

## Back to Module

[← Back to Module 7: Graph Algorithms II - Shortest Paths](../module_07/README.md)

## Navigation

- [Previous Lab](./lab_7_0.md) (if exists)
- [Next Lab](./lab_7_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 7, Lab 1 of 5*
