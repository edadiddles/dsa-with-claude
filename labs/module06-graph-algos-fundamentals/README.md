# Module 6: Graph Algorithms I - Fundamentals

**Duration**: 3 weeks  
**Difficulty**: Medium

## Overview

Graphs are one of the most fundamental structures in computer science, modeling relationships between entities. This module covers essential graph algorithms - the building blocks you'll use to solve network problems, analyze social networks, route packets, schedule tasks, and much more.

**Why this matters**: Graph algorithms appear everywhere - web crawlers use BFS, compilers use topological sort, network routing uses shortest paths, databases use graph traversals. Mastering these fundamentals opens up an entire class of problems that would otherwise seem intractable.

## Learning Objectives

- Implement multiple graph representations and understand their tradeoffs
- Master graph traversal algorithms (BFS and DFS)
- Apply traversals to solve connectivity and ordering problems
- Implement algorithms for strongly connected components
- Build minimum spanning tree algorithms
- Create visualization and analysis tools for graphs

## Topics Covered

- Graph representations (adjacency matrix, list, edge list)
- Breadth-First Search (BFS) and applications
- Depth-First Search (DFS) and applications
- Topological sorting for DAGs
- Strongly connected components (Kosaraju's and Tarjan's)
- Minimum spanning trees (Kruskal's and Prim's)
- Graph properties and analysis

---

## Lab 6.1: Graph Representations

**Duration**: 4-5 hours  
**Difficulty**: Easy-Medium

### Objectives
- Implement multiple graph representations from scratch
- Understand space-time tradeoffs between representations
- Convert between different representations
- Analyze memory usage

### Requirements

1. **Implement three representations**:
   - **Adjacency Matrix**: 2D array, O(1) edge queries, O(V²) space
   - **Adjacency List**: Array of lists, O(V+E) space, O(degree) edge queries
   - **Edge List**: List of edges, minimal space, linear search

2. **Support both directed and undirected graphs**

3. **Operations for each representation**:
   - `add_vertex()` - Add a new vertex
   - `add_edge(u, v, weight=1)` - Add edge from u to v
   - `remove_vertex(v)` - Remove vertex and incident edges
   - `remove_edge(u, v)` - Remove edge
   - `get_neighbors(v)` - Get all neighbors of v
   - `has_edge(u, v)` - Check if edge exists
   - `get_weight(u, v)` - Get edge weight
   - `degree(v)` - Get vertex degree

4. **Conversion functions between representations**

5. **Memory usage analysis**

### Implementation Details

```python
class AdjacencyMatrix:
    def __init__(self, num_vertices, directed=False):
        self.num_vertices = num_vertices
        self.directed = directed
        self.matrix = [[0] * num_vertices for _ in range(num_vertices)]
    
    def add_edge(self, u, v, weight=1):
        self.matrix[u][v] = weight
        if not self.directed:
            self.matrix[v][u] = weight
    
    def remove_edge(self, u, v):
        self.matrix[u][v] = 0
        if not self.directed:
            self.matrix[v][u] = 0
    
    def has_edge(self, u, v):
        return self.matrix[u][v] != 0
    
    def get_weight(self, u, v):
        return self.matrix[u][v]
    
    def get_neighbors(self, u):
        return [v for v in range(self.num_vertices) if self.matrix[u][v] != 0]
    
    def degree(self, u):
        if self.directed:
            # Out-degree for directed graph
            return sum(1 for v in range(self.num_vertices) if self.matrix[u][v] != 0)
        else:
            return len(self.get_neighbors(u))
    
    def get_all_edges(self):
        edges = []
        for u in range(self.num_vertices):
            for v in range(self.num_vertices):
                if self.matrix[u][v] != 0:
                    if self.directed or u <= v:  # Avoid duplicates in undirected
                        edges.append((u, v, self.matrix[u][v]))
        return edges

class AdjacencyList:
    def __init__(self, num_vertices, directed=False):
        self.num_vertices = num_vertices
        self.directed = directed
        self.adj_list = [[] for _ in range(num_vertices)]
    
    def add_edge(self, u, v, weight=1):
        # Check if edge already exists
        for i, (neighbor, _) in enumerate(self.adj_list[u]):
            if neighbor == v:
                self.adj_list[u][i] = (v, weight)  # Update weight
                if not self.directed:
                    for j, (neighbor2, _) in enumerate(self.adj_list[v]):
                        if neighbor2 == u:
                            self.adj_list[v][j] = (u, weight)
                return
        
        self.adj_list[u].append((v, weight))
        if not self.directed:
            self.adj_list[v].append((u, weight))
    
    def remove_edge(self, u, v):
        self.adj_list[u] = [(neighbor, w) for neighbor, w in self.adj_list[u] if neighbor != v]
        if not self.directed:
            self.adj_list[v] = [(neighbor, w) for neighbor, w in self.adj_list[v] if neighbor != u]
    
    def has_edge(self, u, v):
        return any(neighbor == v for neighbor, _ in self.adj_list[u])
    
    def get_weight(self, u, v):
        for neighbor, weight in self.adj_list[u]:
            if neighbor == v:
                return weight
        return 0
    
    def get_neighbors(self, u):
        return [v for v, _ in self.adj_list[u]]
    
    def degree(self, u):
        return len(self.adj_list[u])
    
    def get_all_edges(self):
        edges = []
        for u in range(self.num_vertices):
            for v, weight in self.adj_list[u]:
                if self.directed or u <= v:
                    edges.append((u, v, weight))
        return edges

class EdgeList:
    def __init__(self, num_vertices, directed=False):
        self.num_vertices = num_vertices
        self.directed = directed
        self.edges = []
    
    def add_edge(self, u, v, weight=1):
        # Remove old edge if exists
        self.remove_edge(u, v)
        self.edges.append((u, v, weight))
        if not self.directed:
            self.edges.append((v, u, weight))
    
    def remove_edge(self, u, v):
        self.edges = [(a, b, w) for a, b, w in self.edges if not (a == u and b == v)]
        if not self.directed:
            self.edges = [(a, b, w) for a, b, w in self.edges if not (a == v and b == u)]
    
    def has_edge(self, u, v):
        return any(a == u and b == v for a, b, _ in self.edges)
    
    def get_weight(self, u, v):
        for a, b, weight in self.edges:
            if a == u and b == v:
                return weight
        return 0
    
    def get_neighbors(self, u):
        neighbors = []
        for a, b, _ in self.edges:
            if a == u:
                neighbors.append(b)
        return neighbors
    
    def degree(self, u):
        return len(self.get_neighbors(u))
    
    def get_all_edges(self):
        if self.directed:
            return self.edges[:]
        else:
            # Remove duplicates for undirected
            seen = set()
            unique_edges = []
            for u, v, w in self.edges:
                if (u, v) not in seen and (v, u) not in seen:
                    unique_edges.append((u, v, w))
                    seen.add((u, v))
            return unique_edges

def matrix_to_list(matrix_graph):
    """Convert adjacency matrix to adjacency list"""
    adj_list = AdjacencyList(matrix_graph.num_vertices, matrix_graph.directed)
    for u in range(matrix_graph.num_vertices):
        for v in range(matrix_graph.num_vertices):
            if matrix_graph.matrix[u][v] != 0:
                adj_list.add_edge(u, v, matrix_graph.matrix[u][v])
    return adj_list

def list_to_matrix(list_graph):
    """Convert adjacency list to adjacency matrix"""
    matrix = AdjacencyMatrix(list_graph.num_vertices, list_graph.directed)
    for u in range(list_graph.num_vertices):
        for v, weight in list_graph.adj_list[u]:
            matrix.add_edge(u, v, weight)
    return matrix

def edges_to_list(edge_graph):
    """Convert edge list to adjacency list"""
    adj_list = AdjacencyList(edge_graph.num_vertices, edge_graph.directed)
    for u, v, weight in edge_graph.get_all_edges():
        adj_list.add_edge(u, v, weight)
    return adj_list

def measure_memory_usage(graph):
    """Estimate memory usage of graph representation"""
    import sys
    
    if isinstance(graph, AdjacencyMatrix):
        # Matrix: V² integers
        return graph.num_vertices ** 2 * sys.getsizeof(int())
    elif isinstance(graph, AdjacencyList):
        # List: V lists + E edges
        total = graph.num_vertices * sys.getsizeof([])
        for neighbors in graph.adj_list:
            total += len(neighbors) * sys.getsizeof((0, 0))
        return total
    elif isinstance(graph, EdgeList):
        # Edge list: E edges
        return len(graph.edges) * sys.getsizeof((0, 0, 0))
```

### Test Cases
- Dense graphs (many edges): compare space usage
- Sparse graphs (few edges): compare space usage
- Add/remove edges: verify correctness
- Directed vs. undirected operations
- Conversion between representations preserves structure

### Deliverables
- `adjacency_matrix.{py,go,zig}` - Matrix implementation
- `adjacency_list.{py,go,zig}` - List implementation
- `edge_list.{py,go,zig}` - Edge list implementation
- `graph_converters.{py,go,zig}` - Conversion functions
- `representation_comparison.md` - Detailed tradeoff analysis
- `memory_analysis.{py,go,zig}` - Memory measurement tools
- `benchmarks.{py,go,zig}` - Operation performance comparison

### Success Criteria
- All representations work correctly
- Understanding of when to use each representation
- Conversion functions maintain graph structure
- Can explain space-time tradeoffs

---

## Lab 6.2: BFS and DFS

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

## Lab 6.3: Topological Sort

**Duration**: 3-4 hours  
**Difficulty**: Medium

### Objectives
- Implement topological sorting for DAGs
- Detect cycles in directed graphs
- Apply to dependency resolution problems

### Requirements

1. **Implement two algorithms**:
   - **Kahn's algorithm**: BFS-based using in-degrees
   - **DFS-based**: Using finish times from DFS

2. **Cycle detection** for directed graphs

3. **Applications**:
   - Course prerequisite ordering
   - Build system dependencies
   - Task scheduling
   - Package manager dependency resolution

### Implementation Details

```python
from collections import deque

def topological_sort_kahn(graph):
    """Kahn's algorithm for topological sorting"""
    # Calculate in-degrees
    in_degree = [0] * graph.num_vertices
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            in_degree[v] += 1
    
    # Start with vertices having in-degree 0
    queue = deque([v for v in range(graph.num_vertices) if in_degree[v] == 0])
    topo_order = []
    
    while queue:
        u = queue.popleft()
        topo_order.append(u)
        
        # Remove this vertex from graph (conceptually)
        for v in graph.get_neighbors(u):
            in_degree[v] -= 1
            if in_degree[v] == 0:
                queue.append(v)
    
    # If not all vertices processed, there's a cycle
    if len(topo_order) != graph.num_vertices:
        return None  # Cycle detected
    
    return topo_order

def topological_sort_dfs(graph):
    """DFS-based topological sorting"""
    visited = set()
    rec_stack = set()  # For cycle detection
    topo_order = []
    
    def dfs(v):
        visited.add(v)
        rec_stack.add(v)
        
        for neighbor in graph.get_neighbors(v):
            if neighbor not in visited:
                if not dfs(neighbor):
                    return False  # Cycle detected
            elif neighbor in rec_stack:
                return False  # Back edge (cycle)
        
        rec_stack.remove(v)
        topo_order.append(v)
        return True
    
    for vertex in range(graph.num_vertices):
        if vertex not in visited:
            if not dfs(vertex):
                return None  # Cycle detected
    
    return list(reversed(topo_order))

def all_topological_sorts(graph):
    """Generate all possible topological orderings"""
    in_degree = [0] * graph.num_vertices
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            in_degree[v] += 1
    
    result = []
    
    def backtrack(current_order, remaining_in_degree):
        if len(current_order) == graph.num_vertices:
            result.append(current_order[:])
            return
        
        # Try all vertices with in-degree 0
        for v in range(graph.num_vertices):
            if remaining_in_degree[v] == 0 and v not in current_order:
                # Temporarily use this vertex
                new_in_degree = remaining_in_degree[:]
                new_in_degree[v] = -1  # Mark as used
                
                # Decrease in-degree of neighbors
                for neighbor in graph.get_neighbors(v):
                    new_in_degree[neighbor] -= 1
                
                current_order.append(v)
                backtrack(current_order, new_in_degree)
                current_order.pop()
    
    backtrack([], in_degree[:])
    return result

def has_cycle_directed(graph):
    """Check if directed graph has a cycle"""
    WHITE, GRAY, BLACK = 0, 1, 2
    color = [WHITE] * graph.num_vertices
    
    def dfs(v):
        color[v] = GRAY
        
        for neighbor in graph.get_neighbors(v):
            if color[neighbor] == GRAY:
                return True  # Back edge
            if color[neighbor] == WHITE and dfs(neighbor):
                return True
        
        color[v] = BLACK
        return False
    
    for v in range(graph.num_vertices):
        if color[v] == WHITE:
            if dfs(v):
                return True
    
    return False

class DependencyResolver:
    """Resolve dependencies using topological sort"""
    def __init__(self):
        self.packages = {}
        self.graph = None
        self.pkg_to_id = {}
        self.id_to_pkg = {}
    
    def add_package(self, package, dependencies):
        """Add a package with its dependencies"""
        self.packages[package] = dependencies
    
    def build_graph(self):
        """Build dependency graph"""
        # Create package ID mapping
        all_packages = set(self.packages.keys())
        for deps in self.packages.values():
            all_packages.update(deps)
        
        self.pkg_to_id = {pkg: i for i, pkg in enumerate(sorted(all_packages))}
        self.id_to_pkg = {i: pkg for pkg, i in self.pkg_to_id.items()}
        
        # Build graph
        self.graph = AdjacencyList(len(all_packages), directed=True)
        
        for package, dependencies in self.packages.items():
            pkg_id = self.pkg_to_id[package]
            for dep in dependencies:
                dep_id = self.pkg_to_id[dep]
                # Edge from dependency to package (dep must come first)
                self.graph.add_edge(dep_id, pkg_id)
    
    def get_install_order(self):
        """Get valid installation order"""
        self.build_graph()
        
        order_ids = topological_sort_kahn(self.graph)
        if order_ids is None:
            raise ValueError("Circular dependency detected!")
        
        return [self.id_to_pkg[i] for i in order_ids]

def course_schedule(num_courses, prerequisites):
    """
    Determine if all courses can be finished given prerequisites.
    prerequisites = [(course, prereq), ...]
    """
    graph = AdjacencyList(num_courses, directed=True)
    
    for course, prereq in prerequisites:
        graph.add_edge(prereq, course)
    
    order = topological_sort_kahn(graph)
    return order is not None  # Can finish if no cycle

def find_course_order(num_courses, prerequisites):
    """Find valid order to take all courses"""
    graph = AdjacencyList(num_courses, directed=True)
    
    for course, prereq in prerequisites:
        graph.add_edge(prereq, course)
    
    return topological_sort_kahn(graph)
```

### Test Cases
- DAG with single valid ordering
- DAG with multiple valid orderings
- Graph with cycle (should return None)
- Complete DAG (many orderings)
- Real dependency graphs (npm packages, courses)
- Empty graph, single vertex

### Deliverables
- `topological_sort.{py,go,zig}` - Both algorithms
- `cycle_detection.{py,go,zig}` - Directed graph cycle detection
- `dependency_resolver.{py,go,zig}` - Practical application
- `all_toposorts.{py,go,zig}` - Find all valid orderings
- `comparison.md` - Kahn's vs. DFS-based
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Both algorithms produce valid topological orderings
- Cycle detection works correctly
- Understanding of when topological sort is applicable
- Can solve real dependency problems

---

## Lab 6.4: Strongly Connected Components

**Duration**: 4-5 hours  
**Difficulty**: Hard

### Objectives
- Implement SCC algorithms (Kosaraju's and/or Tarjan's)
- Understand graph condensation
- Apply to real-world problems

### Requirements

1. **Implement Kosaraju's Algorithm**:
   - Two-pass DFS approach
   - Compute finish times
   - Transpose graph
   - DFS in decreasing finish time order

2. **Implement Tarjan's Algorithm** (optional but recommended):
   - Single-pass DFS
   - Uses low-link values
   - More space-efficient

3. **Applications**:
   - Find SCCs in web graph
   - Analyze package dependencies
   - Social network analysis
   - Build condensation graph

### Implementation Details

```python
def transpose(graph):
    """Reverse all edges in a directed graph"""
    transposed = AdjacencyList(graph.num_vertices, directed=True)
    
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            transposed.add_edge(v, u)
    
    return transposed

def kosaraju_scc(graph):
    """
    Kosaraju's algorithm for finding strongly connected components.
    Returns list of SCCs, where each SCC is a list of vertices.
    """
    # Step 1: First DFS to compute finish times
    visited = set()
    finish_order = []
    
    def dfs1(v):
        visited.add(v)
        for neighbor in graph.get_neighbors(v):
            if neighbor not in visited:
                dfs1(neighbor)
        finish_order.append(v)
    
    for v in range(graph.num_vertices):
        if v not in visited:
            dfs1(v)
    
    # Step 2: Transpose the graph
    graph_t = transpose(graph)
    
    # Step 3: Second DFS on transposed graph in reverse finish order
    visited.clear()
    sccs = []
    
    def dfs2(v, component):
        visited.add(v)
        component.append(v)
        for neighbor in graph_t.get_neighbors(v):
            if neighbor not in visited:
                dfs2(neighbor, component)
    
    for v in reversed(finish_order):
        if v not in visited:
            component = []
            dfs2(v, component)
            sccs.append(component)
    
    return sccs

def tarjan_scc(graph):
    """
    Tarjan's algorithm for finding strongly connected components.
    Single-pass DFS using low-link values.
    """
    index_counter = [0]
    stack = []
    lowlinks = {}
    index = {}
    on_stack = set()
    sccs = []
    
    def strongconnect(v):
        # Set the depth index for v
        index[v] = index_counter[0]
        lowlinks[v] = index_counter[0]
        index_counter[0] += 1
        stack.append(v)
        on_stack.add(v)
        
        # Consider successors of v
        for w in graph.get_neighbors(v):
            if w not in index:
                # Successor w has not yet been visited; recurse on it
                strongconnect(w)
                lowlinks[v] = min(lowlinks[v], lowlinks[w])
            elif w in on_stack:
                # Successor w is in stack and hence in the current SCC
                lowlinks[v] = min(lowlinks[v], index[w])
        
        # If v is a root node, pop the stack and create an SCC
        if lowlinks[v] == index[v]:
            component = []
            while True:
                w = stack.pop()
                on_stack.remove(w)
                component.append(w)
                if w == v:
                    break
            sccs.append(component)
    
    for v in range(graph.num_vertices):
        if v not in index:
            strongconnect(v)
    
    return sccs

def condensation_graph(graph, sccs):
    """
    Build condensation graph (DAG of SCCs).
    Each SCC becomes a single vertex in the condensation graph.
    """
    # Map each vertex to its SCC ID
    scc_id = {}
    for i, scc in enumerate(sccs):
        for v in scc:
            scc_id[v] = i
    
    # Build condensation graph
    condensed = AdjacencyList(len(sccs), directed=True)
    edges_added = set()
    
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            if scc_id[u] != scc_id[v]:
                edge = (scc_id[u], scc_id[v])
                if edge not in edges_added:
                    condensed.add_edge(scc_id[u], scc_id[v])
                    edges_added.add(edge)
    
    return condensed

def is_strongly_connected(graph):
    """Check if entire graph is strongly connected"""
    if graph.num_vertices == 0:
        return True
    
    sccs = kosaraju_scc(graph)
    return len(sccs) == 1

def find_cut_vertices_directed(graph):
    """Find articulation points in directed graph using SCCs"""
    sccs = kosaraju_scc(graph)
    condensed = condensation_graph(graph, sccs)
    
    # Vertices that connect different SCCs are critical
    cut_vertices = set()
    
    for scc in sccs:
        if len(scc) == 1:
            v = scc[0]
            # Check if removing this vertex disconnects the graph
            # (This is a simplified check)
            if len(graph.get_neighbors(v)) > 1:
                cut_vertices.add(v)
    
    return list(cut_vertices)

class SocialNetworkAnalyzer:
    """Analyze social networks using SCCs"""
    def __init__(self, graph):
        self.graph = graph
        self.sccs = kosaraju_scc(graph)
    
    def find_communities(self):
        """SCCs represent tightly-knit communities"""
        return self.sccs
    
    def find_influential_users(self):
        """Users that connect different communities"""
        condensed = condensation_graph(self.graph, self.sccs)
        
        # Find SCCs with high out-degree in condensation
        out_degrees = [len(condensed.get_neighbors(i)) for i in range(len(self.sccs))]
        
        # Return users from highly connected SCCs
        influential = []
        for i, degree in enumerate(out_degrees):
            if degree > 1:
                influential.extend(self.sccs[i])
        
        return influential
    
    def find_sink_components(self):
        """SCCs with no outgoing edges (information sinks)"""
        condensed = condensation_graph(self.graph, self.sccs)
        
        sinks = []
        for i in range(len(self.sccs)):
            if len(condensed.get_neighbors(i)) == 0:
                sinks.append(self.sccs[i])
        
        return sinks
```

### Test Cases
- Graphs with obvious SCCs
- Single SCC (strongly connected)
- DAG (each vertex is its own SCC)
- Complex graphs with nested SCCs
- Web graph example
- Social network graph

### Deliverables
- `kosaraju.{py,go,zig}` - Kosaraju's algorithm
- `tarjan.{py,go,zig}` - Tarjan's algorithm
- `condensation.{py,go,zig}` - Build condensation graph
- `scc_applications.{py,go,zig}` - Practical uses
- `social_network_analyzer.{py,go,zig}` - Application example
- `algorithm_comparison.md` - Kosaraju vs. Tarjan
- `test_suite.{py,go,zig}` - Comprehensive tests

### Success Criteria
- Both algorithms find correct SCCs
- Condensation graph is a DAG
- Understanding of SCC properties and applications
- Can analyze real network data

---

## Lab 6.5: Minimum Spanning Trees

**Duration**: 5-6 hours  
**Difficulty**: Medium-Hard

### Objectives
- Implement Kruskal's and Prim's algorithms for MST
- Use Union-Find for Kruskal's
- Understand cut property and MST uniqueness

### Requirements

1. **Implement Kruskal's Algorithm**:
   - Sort edges by weight
   - Use Union-Find to detect cycles
   - Greedy edge selection

2. **Implement Prim's Algorithm**:
   - Start from arbitrary vertex
   - Use min-heap for edge selection
   - Grow single tree

3. **Optimizations**:
   - Prim's with Fibonacci heap (theoretical)
   - Parallel Kruskal's (Borůvka's algorithm)

4. **Applications**:
   - Network design (minimize cable cost)
   - Clustering (remove heaviest edges)
   - Approximation algorithms

### Implementation Details

```python
import heapq
from union_find import UnionFind  # From Module 5

def kruskal_mst(graph):
    """Kruskal's algorithm for MST"""
    # Get all edges and sort by weight
    edges = graph.get_all_edges()
    edges.sort(key=lambda e: e[2])  # Sort by weight
    
    uf = UnionFind(graph.num_vertices)
    mst = []
    total_weight = 0
    
    for u, v, weight in edges:
        if uf.find(u) != uf.find(v):
            uf.union(u, v)
            mst.append((u, v, weight))
            total_weight += weight
            
            # Early termination: MST has V-1 edges
            if len(mst) == graph.num_vertices - 1:
                break
    
    return mst, total_weight

def prim_mst(graph, start=0):
    """Prim's algorithm for MST"""
    visited = set()
    min_heap = [(0, start, -1)]  # (weight, vertex, parent)
    mst = []
    total_weight = 0
    
    while min_heap and len(visited) < graph.num_vertices:
        weight, u, parent = heapq.heappop(min_heap)
        
        if u in visited:
            continue
        
        visited.add(u)
        
        if parent != -1:
            mst.append((parent, u, weight))
            total_weight += weight
        
        # Add all edges from u to unvisited vertices
        for v in graph.get_neighbors(u):
            if v not in visited:
                edge_weight = graph.get_weight(u, v)
                heapq.heappush(min_heap, (edge_weight, v, u))
    
    return mst, total_weight

def boruvka_mst(graph):
    """
    Borůvka's algorithm - can be parallelized.
    Faster than Kruskal's for dense graphs.
    """
    uf = UnionFind(graph.num_vertices)
    mst = []
    total_weight = 0
    
    while uf.num_sets > 1:
        # Find cheapest edge for each component
        cheapest = [None] * graph.num_vertices
        
        for u in range(graph.num_vertices):
            for v in graph.get_neighbors(u):
                weight = graph.get_weight(u, v)
                root_u = uf.find(u)
                root_v = uf.find(v)
                
                if root_u != root_v:
                    if cheapest[root_u] is None or weight < cheapest[root_u][2]:
                        cheapest[root_u] = (u, v, weight)
                    if cheapest[root_v] is None or weight < cheapest[root_v][2]:
                        cheapest[root_v] = (v, u, weight)
        
        # Add cheapest edges
        for i in range(graph.num_vertices):
            if cheapest[i] is not None:
                u, v, weight = cheapest[i]
                if uf.find(u) != uf.find(v):
                    uf.union(u, v)
                    mst.append((u, v, weight))
                    total_weight += weight
    
    return mst, total_weight

def verify_mst(graph, mst):
    """Verify that MST is valid"""
    if len(mst) != graph.num_vertices - 1:
        return False, "Wrong number of edges"
    
    # Check connectivity using MST edges
    mst_graph = AdjacencyList(graph.num_vertices, directed=False)
    for u, v, _ in mst:
        mst_graph.add_edge(u, v)
    
    # BFS to check if all vertices reachable
    from collections import deque
    visited = set([0])
    queue = deque([0])
    
    while queue:
        u = queue.popleft()
        for v in mst_graph.get_neighbors(u):
            if v not in visited:
                visited.add(v)
                queue.append(v)
    
    if len(visited) != graph.num_vertices:
        return False, "MST not connected"
    
    # Check no cycles (implicitly satisfied if V-1 edges and connected)
    return True, "Valid MST"

def mst_second_best(graph):
    """Find second-best MST"""
    mst, min_weight = kruskal_mst(graph)
    
    second_best_weight = float('inf')
    second_best_mst = None
    
    # Try removing each MST edge and replacing with another
    for i, (u_remove, v_remove, w_remove) in enumerate(mst):
        # Build MST without this edge
        edges = graph.get_all_edges()
        edges = [(u, v, w) for u, v, w in edges if not (u == u_remove and v == v_remove)]
        edges.sort(key=lambda e: e[2])
        
        uf = UnionFind(graph.num_vertices)
        current_mst = []
        current_weight = 0
        
        for u, v, weight in edges:
            if uf.find(u) != uf.find(v):
                uf.union(u, v)
                current_mst.append((u, v, weight))
                current_weight += weight
                
                if len(current_mst) == graph.num_vertices - 1:
                    break
        
        if len(current_mst) == graph.num_vertices - 1:
            if current_weight < second_best_weight and current_weight > min_weight:
                second_best_weight = current_weight
                second_best_mst = current_mst
    
    return second_best_mst, second_best_weight

def clustering_with_mst(graph, k):
    """
    Use MST for clustering: remove k-1 heaviest edges to get k clusters.
    """
    mst, _ = kruskal_mst(graph)
    
    # Sort MST edges by weight (descending)
    mst.sort(key=lambda e: e[2], reverse=True)
    
    # Remove k-1 heaviest edges
    kept_edges = mst[k-1:]
    
    # Find connected components in remaining graph
    uf = UnionFind(graph.num_vertices)
    for u, v, _ in kept_edges:
        uf.union(u, v)
    
    # Build clusters
    clusters = {}
    for v in range(graph.num_vertices):
        root = uf.find(v)
        if root not in clusters:
            clusters[root] = []
        clusters[root].append(v)
    
    return list(clusters.values())

def maximum_spanning_tree(graph):
    """Find maximum spanning tree (negate weights)"""
    # Create graph with negated weights
    edges = graph.get_all_edges()
    negated_graph = EdgeList(graph.num_vertices, directed=False)
    
    for u, v, weight in edges:
        negated_graph.add_edge(u, v, -weight)
    
    mst, neg_weight = kruskal_mst(negated_graph)
    
    # Restore original weights
    mst_with_orig_weights = [(u, v, -w) for u, v, w in mst]
    return mst_with_orig_weights, -neg_weight
```

### Test Cases
- Complete graphs
- Sparse graphs
- Graphs with equal edge weights
- Graphs with unique weights
- Disconnected graphs
- Large random graphs

### Deliverables
- `kruskal.{py,go,zig}` - Kruskal's algorithm
- `prim.{py,go,zig}` - Prim's algorithm
- `boruvka.{py,go,zig}` - Borůvka's algorithm
- `mst_verification.{py,go,zig}` - Verify MST correctness
- `mst_applications.{py,go,zig}` - Clustering, second-best MST
- `algorithm_comparison.md` - Performance comparison
- `benchmarks.{py,go,zig}` - Benchmark all three algorithms

### Success Criteria
- All algorithms find MST of same weight
- Kruskal's and Prim's produce valid MSTs
- Understanding of when each algorithm is preferable
- Can apply to real network design problems

---

## Lab 6.6: Graph Visualization & Analysis Tools

**Duration**: 4-5 hours  
**Difficulty**: Medium

### Objectives
- Build tools to visualize graphs
- Aid in debugging graph algorithms
- Compute graph statistics

### Requirements

1. **Visualization capabilities**:
   - Display graph structure (vertices and edges)
   - Highlight paths, trees, components
   - Animate algorithms (BFS, DFS, MST)
   - Show vertex/edge weights and labels

2. **Output formats**:
   - ASCII art for terminal
   - DOT format for Graphviz
   - Interactive HTML/JavaScript
   - Export to image formats

3. **Analysis tools**:
   - Compute graph statistics (density, diameter, etc.)
   - Find central vertices (degree centrality)
   - Identify bridges and articulation points

### Implementation Details

```python
def graph_to_dot(graph, highlight_edges=None, highlight_vertices=None, labels=None):
    """Convert graph to DOT format for Graphviz"""
    dot = "digraph G {\n" if graph.directed else "graph G {\n"
    edge_symbol = "->" if graph.directed else "--"
    
    # Add vertices
    for v in range(graph.num_vertices):
        color = "red" if highlight_vertices and v in highlight_vertices else "black"
        label = labels.get(v, str(v)) if labels else str(v)
        dot += f'  {v} [label="{label}", color={color}];\n'
    
    # Add edges
    seen_edges = set()
    for u in range(graph.num_vertices):
        for v in graph.get_neighbors(u):
            if not graph.directed and (v, u) in seen_edges:
                continue
            
            weight = graph.get_weight(u, v)
            color = "red" if highlight_edges and (u, v) in highlight_edges else "black"
            dot += f'  {u} {edge_symbol} {v} [label="{weight}", color={color}];\n'
            seen_edges.add((u, v))
    
    dot += "}\n"
    return dot

def ascii_graph(graph, max_width=80):
    """Simple ASCII representation"""
    lines = []
    for u in range(graph.num_vertices):
        neighbors = graph.get_neighbors(u)
        if graph.directed:
            line = f"{u} -> {neighbors}"
        else:
            line = f"{u}: {neighbors}"
        lines.append(line)
    return "\n".join(lines)

def graph_statistics(graph):
    """Compute various graph statistics"""
    stats = {}
    
    # Basic counts
    stats['vertices'] = graph.num_vertices
    stats['edges'] = len(graph.get_all_edges())
    
    # Density
    max_edges = graph.num_vertices * (graph.num_vertices - 1)
    if not graph.directed:
        max_edges //= 2
    stats['density'] = stats['edges'] / max_edges if max_edges > 0 else 0
    
    # Degree distribution
    degrees = [graph.degree(u) for u in range(graph.num_vertices)]
    stats['avg_degree'] = sum(degrees) / len(degrees) if degrees else 0
    stats['max_degree'] = max(degrees) if degrees else 0
    stats['min_degree'] = min(degrees) if degrees else 0
    
    # Components
    if not graph.directed:
        stats['num_components'] = count_connected_components(graph)
    
    # Diameter (for connected graphs)
    if not graph.directed and stats.get('num_components', 1) == 1:
        stats['diameter'] = compute_diameter(graph)
    
    return stats

def compute_diameter(graph):
    """Compute graph diameter (longest shortest path)"""
    diameter = 0
    
    for source in range(graph.num_vertices):
        result = bfs(graph, source)
        max_dist = max(result['distance'].values()) if result['distance'] else 0
        diameter = max(diameter, max_dist)
    
    return diameter

def find_bridges(graph):
    """
    Find bridges (edges whose removal disconnects the graph).
    Uses Tarjan's bridge-finding algorithm.
    """
    visited = set()
    discovery = {}
    low = {}
    parent = {}
    bridges = []
    time = [0]
    
    def dfs(u):
        visited.add(u)
        discovery[u] = low[u] = time[0]
        time[0] += 1
        
        for v in graph.get_neighbors(u):
            if v not in visited:
                parent[v] = u
                dfs(v)
                
                low[u] = min(low[u], low[v])
                
                # If lowest vertex reachable from v is below u, then (u,v) is a bridge
                if low[v] > discovery[u]:
                    bridges.append((u, v))
            
            elif v != parent.get(u):
                low[u] = min(low[u], discovery[v])
    
    for v in range(graph.num_vertices):
        if v not in visited:
            parent[v] = None
            dfs(v)
    
    return bridges

def find_articulation_points(graph):
    """
    Find articulation points (vertices whose removal disconnects the graph).
    """
    visited = set()
    discovery = {}
    low = {}
    parent = {}
    articulation_points = set()
    time = [0]
    
    def dfs(u):
        children = 0
        visited.add(u)
        discovery[u] = low[u] = time[0]
        time[0] += 1
        
        for v in graph.get_neighbors(u):
            if v not in visited:
                children += 1
                parent[v] = u
                dfs(v)
                
                low[u] = min(low[u], low[v])
                
                # u is an articulation point in two cases:
                # 1) u is root of DFS tree and has two or more children
                if parent.get(u) is None and children > 1:
                    articulation_points.add(u)
                
                # 2) u is not root and low value of one of its children is >= discovery of u
                if parent.get(u) is not None and low[v] >= discovery[u]:
                    articulation_points.add(u)
            
            elif v != parent.get(u):
                low[u] = min(low[u], discovery[v])
    
    for v in range(graph.num_vertices):
        if v not in visited:
            parent[v] = None
            dfs(v)
    
    return list(articulation_points)

def animate_bfs(graph, start):
    """Generate steps for BFS animation"""
    steps = []
    visited = set()
    queue = deque([start])
    visited.add(start)
    
    steps.append({
        'type': 'start',
        'vertex': start,
        'queue': list(queue),
        'visited': list(visited)
    })
    
    while queue:
        u = queue.popleft()
        steps.append({
            'type': 'visit',
            'vertex': u,
            'queue': list(queue),
            'visited': list(visited)
        })
        
        for v in graph.get_neighbors(u):
            if v not in visited:
                visited.add(v)
                queue.append(v)
                steps.append({
                    'type': 'discover',
                    'vertex': v,
                    'from': u,
                    'queue': list(queue),
                    'visited': list(visited)
                })
    
    return steps

def export_to_html(graph, algorithm_name="Graph", animation_steps=None):
    """Export interactive graph visualization to HTML"""
    html = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>{algorithm_name}</title>
        <script src="https://d3js.org/d3.v7.min.js"></script>
        <style>
            .node {{ fill: #69b3a2; stroke: #333; stroke-width: 2px; }}
            .node.visited {{ fill: #ff6b6b; }}
            .link {{ stroke: #999; stroke-width: 2px; }}
            .link.highlighted {{ stroke: #ff6b6b; stroke-width: 4px; }}
        </style>
    </head>
    <body>
        <svg width="800" height="600"></svg>
        <script>
            // D3.js visualization code here
            // This is a simplified example
            const nodes = {[{{"id": i}} for i in range(graph.num_vertices)]};
            const edges = {graph.get_all_edges()};
            
            // Visualization logic...
        </script>
    </body>
    </html>
    """
    return html
```

### Test Cases
- Small graphs (easy to verify visually)
- Large graphs (stress test layout)
- Different graph types (trees, DAGs, cyclic)
- Animation of all algorithms

### Deliverables
- `graph_visualizer.{py,go,zig}` - Core visualization library
- `ascii_viz.{py,go,zig}` - Terminal-based visualization
- `dot_export.{py,go,zig}` - Graphviz export
- `graph_stats.{py,go,zig}` - Statistical analysis
- `algorithm_animator.{py,go,zig}` - Animate graph algorithms
- `bridge_articulation.{py,go,zig}` - Find bridges and articulation points
- `interactive_viz.html` - Web-based interactive visualizer (optional)

### Success Criteria
- Clear, readable visualizations
- Supports all graph types implemented
- Useful for debugging and learning
- Statistics are accurate

---

## Module Resources

### CLRS References
- Chapter 22: Elementary Graph Algorithms (BFS, DFS)
- Chapter 23: Minimum Spanning Trees

### Additional Reading
- Graph algorithms textbooks
- Network analysis papers
- Real-world graph datasets (SNAP, NetworkX)

### Key Takeaways

By the end of this module, you should:
1. Understand graph representations deeply
2. Master BFS and DFS and their applications
3. Be able to detect cycles and check connectivity
4. Implement topological sorting correctly
5. Find strongly connected components
6. Build minimum spanning trees
7. Have tools to visualize and analyze graphs

### Next Module

**Module 7: Graph Algorithms II - Shortest Paths** - You'll implement shortest path algorithms including Dijkstra's, Bellman-Ford, Floyd-Warshall, and build a practical route planning application.

---

## Tips for Success

1. **Draw graphs** - Always visualize on paper first
2. **Test small graphs** - Verify by hand
3. **Use your visualizer** - Build it early and use it
4. **Understand edge cases** - Disconnected, single vertex, cycles
5. **Practice on real data** - Social networks, road networks

## Common Pitfalls

- **Directed vs. undirected** - Remember which you're working with
- **Visited tracking** - Essential to avoid infinite loops
- **Off-by-one errors** - Especially in adjacency matrix
- **Not handling disconnected graphs** - Always check all components
- **Cycle detection bugs** - Carefully track parent vertices
- **MST on disconnected graphs** - Returns forest, not tree
