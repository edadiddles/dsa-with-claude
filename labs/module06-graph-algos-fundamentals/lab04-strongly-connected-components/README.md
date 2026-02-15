# Lab 6.4: Strongly Connected Components

**Module**: Module 6 - Graph Algorithms I - Fundamentals  
**Duration**: 4-5 hours  
**Difficulty**: Hard

---

Lab 6.4: Strongly Connected Components

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

---

## Back to Module

[← Back to Module 6: Graph Algorithms I - Fundamentals](../module_06/README.md)

## Navigation

- [Previous Lab](./lab_6_3.md) (if exists)
- [Next Lab](./lab_6_5.md) (if exists)
- [All Labs](../labs/README.md)

---

*Part of the comprehensive DSA Curriculum*  
*Module 6, Lab 4 of 6*
