# Lab 6.1: Graph Representations

**Module**: Module 6 - Graph Algorithms I - Fundamentals  
**Duration**: 4-5 hours  
**Difficulty**: Easy-Medium

---

Lab 6.1: Graph Representations

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

---

## Back to Module

[← Back to Module 6: Graph Algorithms I - Fundamentals](../module_06/README.md)

## Navigation

- [Previous Lab](./lab_6_0.md) (if exists)
- [Next Lab](./lab_6_2.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 6, Lab 1 of 6*
