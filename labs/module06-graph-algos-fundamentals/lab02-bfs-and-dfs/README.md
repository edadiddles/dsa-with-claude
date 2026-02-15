# Lab 6.2: BFS and DFS

**Module**: Module 6 - Graph Algorithms I - Fundamentals  
**Duration**: 5-6 hours  
**Difficulty**: Medium

---

Lab 6.2: BFS and DFS

**Duration**: 5-6 hours  
**Difficulty**: Medium

### Objectives
- Implement breadth-first and depth-first search
- Apply traversals to solve graph problems
- Understand differences between BFS and DFS

### Requirements

1. **Implement Breadth-First Search**:
   - Iterative with queue
   - Track visited vertices
   - Record parent pointers (for path reconstruction)
   - Record distances from source
   - Level-order traversal

2. **Implement Depth-First Search**:
   - Recursive version
   - Iterative with explicit stack
   - Track discovery and finish times
   - Classify edges (tree, back, forward, cross)

3. **Applications**:
   - **BFS Applications**:
     - Shortest path in unweighted graph
     - Test if graph is connected
     - Test if graph is bipartite
     - Level-order properties
   - **DFS Applications**:
     - Detect cycles
     - Find path between vertices
     - Classify edges
     - Count connected components

### Implementation Details

```python
from collections import deque

def bfs(graph, start):
    """Breadth-First Search from start vertex"""
    visited = set()
    queue = deque([start])
    visited.add(start)
    
    order = []
    parent = {start: None}
    distance = {start: 0}
    
    while queue:
        vertex = queue.popleft()
        order.append(vertex)
        
        for neighbor in graph.get_neighbors(vertex):
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append(neighbor)
                parent[neighbor] = vertex
                distance[neighbor] = distance[vertex] + 1
    
    return {
        'order': order,
        'parent': parent,
        'distance': distance,
        'visited': visited
    }

def bfs_shortest_path(graph, start, end):
    """Find shortest path from start to end"""
    result = bfs(graph, start)
    
    if end not in result['visited']:
        return None  # No path exists
    
    # Reconstruct path
    path = []
    current = end
    while current is not None:
        path.append(current)
        current = result['parent'][current]
    
    return list(reversed(path))

def is_connected(graph):
    """Check if undirected graph is connected"""
    if graph.num_vertices == 0:
        return True
    
    result = bfs(graph, 0)
    return len(result['visited']) == graph.num_vertices

def is_bipartite(graph):
    """Check if graph is bipartite (2-colorable)"""
    color = {}
    
    for start in range(graph.num_vertices):
        if start in color:
            continue
        
        queue = deque([start])
        color[start] = 0
        
        while queue:
            vertex = queue.popleft()
            for neighbor in graph.get_neighbors(vertex):
                if neighbor not in color:
                    color[neighbor] = 1 - color[vertex]
                    queue.append(neighbor)
                elif color[neighbor] == color[vertex]:
                    return False  # Odd cycle found
    
    return True

def dfs_recursive(graph, vertex, visited=None, result=None):
    """Recursive DFS"""
    if visited is None:
        visited = set()
        result = {'order': [], 'parent': {vertex: None}}
    
    visited.add(vertex)
    result['order'].append(vertex)
    
    for neighbor in graph.get_neighbors(vertex):
        if neighbor not in visited:
            result['parent'][neighbor] = vertex
            dfs_recursive(graph, neighbor, visited, result)
    
    return result

def dfs_iterative(graph, start):
    """Iterative DFS using explicit stack"""
    visited = set()
    stack = [start]
    order = []
    parent = {start: None}
    
    while stack:
        vertex = stack.pop()
        
        if vertex not in visited:
            visited.add(vertex)
            order.append(vertex)
            
            # Add neighbors in reverse order to match recursive DFS
            for neighbor in reversed(graph.get_neighbors(vertex)):
                if neighbor not in visited:
                    stack.append(neighbor)
                    if neighbor not in parent:
                        parent[neighbor] = vertex
    
    return {
        'order': order,
        'parent': parent,
        'visited': visited
    }

def dfs_with_timestamps(graph):
    """DFS with discovery and finish times"""
    time = [0]  # Use list to allow mutation in nested function
    visited = set()
    discovery = {}
    finish = {}
    parent = {}
    
    def dfs_visit(vertex):
        time[0] += 1
        discovery[vertex] = time[0]
        visited.add(vertex)
        
        for neighbor in graph.get_neighbors(vertex):
            if neighbor not in visited:
                parent[neighbor] = vertex
                dfs_visit(neighbor)
        
        time[0] += 1
        finish[vertex] = time[0]
    
    for v in range(graph.num_vertices):
        if v not in visited:
            parent[v] = None
            dfs_visit(v)
    
    return {
        'discovery': discovery,
        'finish': finish,
        'parent': parent
    }

def classify_edges(graph):
    """Classify edges as tree, back, forward, or cross"""
    result = dfs_with_timestamps(graph)
    discovery = result['discovery']
    finish = result['finish']
    parent = result['parent']
    
    edge_types = {}
    
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            if parent.get(v) == u:
                edge_types[(u, v)] = 'tree'
            elif discovery[u] > discovery[v] and finish[u] < finish[v]:
                edge_types[(u, v)] = 'back'
            elif discovery[u] < discovery[v] and finish[u] > finish[v]:
                edge_types[(u, v)] = 'forward'
            else:
                edge_types[(u, v)] = 'cross'
    
    return edge_types

def detect_cycle_undirected_bfs(graph):
    """Detect cycle in undirected graph using BFS"""
    visited = set()
    
    for start in range(graph.num_vertices):
        if start in visited:
            continue
        
        queue = deque([(start, -1)])  # (vertex, parent)
        visited.add(start)
        
        while queue:
            vertex, parent = queue.popleft()
            
            for neighbor in graph.get_neighbors(vertex):
                if neighbor not in visited:
                    visited.add(neighbor)
                    queue.append((neighbor, vertex))
                elif neighbor != parent:
                    return True  # Cycle found
    
    return False

def detect_cycle_directed_dfs(graph):
    """Detect cycle in directed graph using DFS"""
    WHITE, GRAY, BLACK = 0, 1, 2
    color = [WHITE] * graph.num_vertices
    
    def has_cycle(vertex):
        color[vertex] = GRAY
        
        for neighbor in graph.get_neighbors(vertex):
            if color[neighbor] == GRAY:
                return True  # Back edge (cycle)
            if color[neighbor] == WHITE and has_cycle(neighbor):
                return True
        
        color[vertex] = BLACK
        return False
    
    for v in range(graph.num_vertices):
        if color[v] == WHITE:
            if has_cycle(v):
                return True
    
    return False

def count_connected_components(graph):
    """Count connected components in undirected graph"""
    visited = set()
    count = 0
    
    for v in range(graph.num_vertices):
        if v not in visited:
            # Start new component
            count += 1
            result = bfs(graph, v)
            visited.update(result['visited'])
    
    return count

def find_path_dfs(graph, start, end):
    """Find any path from start to end using DFS"""
    visited = set()
    path = []
    
    def dfs(vertex):
        visited.add(vertex)
        path.append(vertex)
        
        if vertex == end:
            return True
        
        for neighbor in graph.get_neighbors(vertex):
            if neighbor not in visited:
                if dfs(neighbor):
                    return True
        
        path.pop()
        return False
    
    if dfs(start):
        return path
    return None
```

### Test Cases
- Connected and disconnected graphs
- Graphs with cycles
- Bipartite and non-bipartite graphs
- DAGs (Directed Acyclic Graphs)
- Large random graphs
- Edge case: empty graph, single vertex

### Deliverables
- `bfs.{py,go,zig}` - BFS with all variants
- `dfs.{py,go,zig}` - Recursive and iterative DFS
- `graph_applications.{py,go,zig}` - Cycle detection, bipartiteness, shortest paths
- `traversal_visualizer.{py,go,zig}` - Visualize traversal order
- `edge_classification.{py,go,zig}` - Classify edges in directed graphs
- `complexity_analysis.md` - Time and space analysis
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Both traversals visit all reachable vertices
- BFS finds shortest paths in unweighted graphs
- DFS correctly detects cycles
- Understanding of when to use each traversal
- Can explain edge classification in DFS

---

---

## Back to Module

[← Back to Module 6: Graph Algorithms I - Fundamentals](../module_06/README.md)

## Navigation

- [Previous Lab](./lab_6_1.md) (if exists)
- [Next Lab](./lab_6_3.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 6, Lab 2 of 6*
